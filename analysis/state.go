package analysis

import (
	"educationalsp/lsp"
	"fmt"
	"strings"
)

// State holds the document content mapped by their URI
type State struct {
	Document map[string]string
}

// NewState initializes and returns a new State
func NewState() State {
	return State{Document: make(map[string]string)}
}

// getDiagnosticsForFile generates diagnostics for a given text
func getDiagnosticsForFile(text string) []lsp.Diagnostic {
	var diagnostics []lsp.Diagnostic
	for row, line := range strings.Split(text, "\n") {
		if strings.Contains(line, "VS Code") {
			idx := strings.Index(line, "VS Code")
			diagnostics = append(diagnostics, lsp.Diagnostic{
				Range:    LineRange(row, idx, idx+len("VS Code")),
				Severity: 1,
				Source:   "Common Sense",
				Message:  "Please make sure we use good language in this Editor",
			})
		}

		if strings.Contains(line, "Neovim") {
			idx := strings.Index(line, "Neovim")
			diagnostics = append(diagnostics, lsp.Diagnostic{
				Range:    LineRange(row, idx, idx+len("Neovim")),
				Severity: 2,
				Source:   "Common Sense",
				Message:  "Great choice :)",
			})
		}
	}

	return diagnostics
}

// OpenDocument handles the opening of a document, storing its content and returning diagnostics
func (s *State) OpenDocument(uri, text string) []lsp.Diagnostic {
	s.Document[uri] = text
	return getDiagnosticsForFile(text)
}

// UpdateDocument handles the updating of a document, storing its content and returning diagnostics
func (s *State) UpdateDocument(uri, text string) []lsp.Diagnostic {
	s.Document[uri] = text
	return getDiagnosticsForFile(text)
}

// Hover provides a hover response for a given position in the document
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

// Definition provides a definition response for a given position in the document
func (s *State) Definition(id int, uri string, position lsp.Position) lsp.DefinitionResponse {
	return lsp.DefinitionResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: lsp.Location{
			URI: uri,
			Range: lsp.Range{
				Start: lsp.Position{Line: position.Line - 1, Character: 0},
				End:   lsp.Position{Line: position.Line - 1, Character: 0},
			},
		},
	}
}

// TextDocumentCodeAction provides code actions for a given document
func (s *State) TextDocumentCodeAction(id int, uri string) lsp.TextDocumentCodeActionResponse {
	text := s.Document[uri]
	var actions []lsp.CodeAction
	for row, line := range strings.Split(text, "\n") {
		idx := strings.Index(line, "VS Code")
		if idx >= 0 {
			replaceChange := map[string][]lsp.TextEdit{
				uri: {
					{
						Range:   LineRange(row, idx, idx+len("VS Code")),
						NewText: "Neovim",
					},
				},
			}

			actions = append(actions, lsp.CodeAction{
				Title: "Replace VS C*de with a superior editor",
				Edit:  &lsp.WorkspaceEdit{Changes: replaceChange},
			})

			censorChange := map[string][]lsp.TextEdit{
				uri: {
					{
						Range:   LineRange(row, idx, idx+len("VS Code")),
						NewText: "VS C*de",
					},
				},
			}

			actions = append(actions, lsp.CodeAction{
				Title: "Censor to VS C*de",
				Edit:  &lsp.WorkspaceEdit{Changes: censorChange},
			})
		}
	}

	return lsp.TextDocumentCodeActionResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: actions,
	}
}

// TextDocumentCompletion provides completion items for a given document
func (s *State) TextDocumentCompletion(id int, uri string) lsp.CompletionResponse {
	// Placeholder for static analysis tools to provide good completions
	items := []lsp.CompletionItem{
		{
			Label:         "Neovim (BTW)",
			Detail:        "Very cool editor",
			Documentation: "An Editor That Can Entertain You",
		},
	}

	return lsp.CompletionResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: items,
	}
}

// LineRange creates a range object for a specific line and character positions
func LineRange(line, start, end int) lsp.Range {
	return lsp.Range{
		Start: lsp.Position{
			Line:      line,
			Character: start,
		},
		End: lsp.Position{
			Line:      line,
			Character: end,
		},
	}
}
