package editor

type Editor struct {
	s       *State
	in      chan Command
	refresh chan bool
}

func NewEditor(s *State) *Editor {
	e := &Editor{
		s:       s,
		in:      make(chan Command),
		refresh: make(chan bool),
	}
	go e.Run()

	return e
}

func (e *Editor) In() chan Command {
	return e.in
}

func (e *Editor) Out() chan bool {
	return e.refresh
}

func (e *Editor) Run() {
	for cmd := range e.in {
		cmd.Do(e.s)
		e.refresh <- true
	}
}
