package editor

type Command interface {
	Do(s *State)
	//Undo(s *State) error
}
