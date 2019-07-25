package event

// Emitter wraps emit and listen function
type Emitter interface {
	Emit(data interface{}) error
}
