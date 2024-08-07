package analysis

import (
	"educationalsp/lsp"
	"fmt"
)

type State struct {
	Document map[string]string
}

func NewState() State {
	return State{Document: make(map[string]string)}
}

func (s *State) OpenDocument(uri, text string) {
	s.Document[uri] = text
}

func (s *State) UpdateDocument(uri, text string) {
	s.Document[uri] = text

}

func (s *State) Hover(id int, uri string, position lsp.Position) lsp.HoverResponse {

	document := s.Document[uri]
	return lsp.HoverResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: lsp.HoverResult{
			Contents: fmt.Sprintf("File: %s, Line: %d, Col: %d Character Length: %d", uri, position.Line, position.Character, len(document)),
		},
	}

}
