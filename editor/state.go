package editor

type State struct {
	Err error
}

func NewState() *State {
	return &State{}
}
