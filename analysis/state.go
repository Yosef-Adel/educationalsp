package analysis

type State struct {
	Document map[string]string
}

func NewState() State {
	return State{Document: make(map[string]string)}
}

func (s *State) OpenDocument(document, text string) {
	s.Document[document] = text
}
