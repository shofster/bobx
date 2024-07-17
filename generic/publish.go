package generic

import "sync"

/*

  File:    publish.go
  Author:  Bob Shofner

  MIT License - https://opensource.org/license/mit/

  This permission notice shall be included in all copies
    or substantial portions of the Software.

*/
/*

    Description: Publish / Subscribe.
  Allows many subscriber channels for a single publisher channel.

*/

type publisher[T any] struct {
	lock        sync.Mutex
	input       chan T
	subscribe   chan chan<- T
	unsubscribe chan chan<- T
	outputs     map[chan<- T]bool
}

// The Publisher interface describes the main entry points to Publisher(s).
type Publisher[T any] interface {
	// Register a new channel to receive broadcasts
	Register(chan<- T)
	// Unregister a channel so that it no longer receives broadcasts.
	Unregister(chan<- T)
	// Close this publisher
	Close()
	// Submit a new object to all subscribers return false if input chan is full
	Submit(T) bool
}

// NewPublisher creates a publisher with the given channel buffer length.
func NewPublisher[T any](size int) Publisher[T] {
	p := &publisher[T]{
		input:       make(chan T, size),      // bi-directional
		subscribe:   make(chan chan<- T),     // send only
		unsubscribe: make(chan chan<- T),     // send only
		outputs:     make(map[chan<- T]bool), // send only
	}
	go p.run()
	return p
}

// run - forever handling requests.
func (p *publisher[T]) run() {
	for {
		select {
		case m := <-p.input: // publish (from a p.Submit(m T))
			for ch := range p.outputs { // to all the registered receivers
				ch <- m
			}
		case ch, ok := <-p.subscribe: // new subscriber
			if ok {
				p.outputs[ch] = true
			} else {
				return
			}
		case ch := <-p.unsubscribe: // quit subscribing
			delete(p.outputs, ch)
		}
	}
}

// Submit attempts to submit an item to be published.
// Returns true if successful, else false.
func (p *publisher[T]) Submit(m T) bool {
	select { // block until a case can run and buffer has room
	case p.input <- m:
		return true
	default: // input channel is not ready. ignore
		return false
	}
}

// Register a receiver channel for T messages.
func (p *publisher[T]) Register(ch chan<- T) {
	defer p.lock.Unlock()
	p.lock.Lock()
	p.subscribe <- ch
}

// Unregister a receiver for T messages
func (p *publisher[T]) Unregister(ch chan<- T) {
	defer p.lock.Unlock()
	p.lock.Lock()
	p.unsubscribe <- ch
}

// Close the publish channel
func (p *publisher[T]) Close() {
	defer p.lock.Unlock()
	p.lock.Lock()
	close(p.subscribe)
	close(p.unsubscribe)
	close(p.input)
}
