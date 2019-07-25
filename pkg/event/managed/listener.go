package managed

import (
	"context"
	"sync"
	"time"

	"github.com/payfazz/payfazz-notification/pkg/event"
)

//NewManagedListener return new managed event listener
func NewManagedListener(stream Stream, store Store) *ManagedListener {
	return &ManagedListener{
		stream: stream,
		store:  store,
		ch:     make(chan interface{}, 100),
		mutex:  &sync.Mutex{},
	}
}

// ManagedListener is managed event listener
type ManagedListener struct {
	stream  Stream
	store   Store
	ch      chan interface{}
	success int
	failed  int
	mutex   *sync.Mutex
}

//f is a routine listening for events
func (e *ManagedListener) count(i *int) {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	*i++
}

//Listen is a routine listening for events
func (e *ManagedListener) Listen(ctx context.Context, handler event.ListenerHandler) {
	// read stream
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				data, err := e.stream.Pull()
				if err != nil {
					e.count(&e.failed)
				}
				if data != nil {
					e.ch <- data
				}
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()

	// go func() {
	for {
		select {
		case <-ctx.Done():
			return
		case data := <-e.ch:
			err := handler(ctx, data)
			if err != nil {
				e.count(&e.failed)
				e.store.Push(data)
			} else {
				e.count(&e.success)
			}
		}
		time.Sleep(100 * time.Millisecond)
	}
	// }()
}

// Watch watches listener
func (e *ManagedListener) Watch(ctx context.Context) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				data := e.store.Pop()
				if data != nil {
					e.ch <- data
				}
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()
}

//Dispose release resources used by listener
func (e *ManagedListener) Dispose() {
	e.stream.Dispose()
	e.store.Dispose()
	close(e.ch)
}

// Success returns count for success emit
func (e ManagedListener) Success() int {
	return e.success
}

// Failed returns count for failed emit
func (e ManagedListener) Failed() int {
	return e.failed
}
