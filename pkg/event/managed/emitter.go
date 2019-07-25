package managed

import (
	"context"
	"log"
	"time"
)

//NewManagedEmitter return new managed event listener
func NewManagedEmitter(stream Stream, store Store) *ManagedEmitter {
	return &ManagedEmitter{
		stream: stream,
		store:  store,
	}
}

//ManagedEmitter is managed event emitter
type ManagedEmitter struct {
	stream  Stream
	store   Store
	success int
	failed  int
}

//Emit emits data
func (e *ManagedEmitter) Emit(data interface{}) error {
	err := e.stream.Push(data)
	if err != nil {
		e.store.Push(data)
		e.failed++
		return err
	}
	e.success++
	return nil
}

//Watch is a routine ensures data is emited
func (e *ManagedEmitter) Watch(ctx context.Context) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				data := e.store.Pop()
				if data != nil {
					err := e.Emit(data)
					if err != nil {
						log.Println(err)
					}
				}
				time.Sleep(1 * time.Millisecond)
			}
		}
	}()
}

//Dispose release resources used by emitter
func (e *ManagedEmitter) Dispose() {
	e.stream.Dispose()
	e.store.Dispose()
}

// Success returns count for success emit
func (e ManagedEmitter) Success() int {
	return e.success
}

// Failed returns count for failed emit
func (e ManagedEmitter) Failed() int {
	return e.failed
}
