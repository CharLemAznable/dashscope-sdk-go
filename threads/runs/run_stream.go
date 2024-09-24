package runs

import (
	"github.com/gogf/gf/v2/container/gtype"
	"github.com/gogf/gf/v2/os/gmutex"
)

type RunStream interface {
	Event() <-chan RunStreamEvent
	Err() error
	Drain()
}

func newRunStream() *runStream {
	return &runStream{
		ch:     make(chan RunStreamEvent),
		mutex:  &gmutex.Mutex{},
		closed: gtype.NewBool(),
	}
}

type runStream struct {
	ch     chan RunStreamEvent
	mutex  *gmutex.Mutex
	err    error
	closed *gtype.Bool
}

func (s *runStream) Push(event RunStreamEvent) {
	s.ch <- event
}

func (s *runStream) Close(err error) {
	if !s.closed.Cas(false, true) {
		return // already closed
	}
	s.mutex.LockFunc(func() {
		s.err = err
		close(s.ch)
	})
}

func (s *runStream) Event() <-chan RunStreamEvent {
	return s.ch
}

func (s *runStream) Err() (err error) {
	s.mutex.LockFunc(func() {
		err = s.err
	})
	return
}

func (s *runStream) Drain() {
	for range s.ch {
		// drain channel
	}
}
