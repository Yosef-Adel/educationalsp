package lsp

// CodeActionRequest represents a request for code actions
type CodeActionRequest struct {
	Request
	Params TextDocumentCodeActionParams `json:"params"`
}

// TextDocumentCodeActionParams contains parameters for a code action request
type TextDocumentCodeActionParams struct {
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	Range        Range                  `json:"range"`
	Context      CodeActionContext      `json:"context"`
}

// TextDocumentCodeActionResponse represents a response for a code action request
type TextDocumentCodeActionResponse struct {
	Response
	Result []CodeAction `json:"result"`
}

// CodeActionContext provides context for a code action
type CodeActionContext struct {
	// Add fields for CodeActionContext as needed
}

// CodeAction represents a code action
type CodeAction struct {
	Title   string         `json:"title"`
	Edit    *WorkspaceEdit `json:"edit,omitempty"`
	Command *Command       `json:"command,omitempty"`
}

// Command represents a command to be executed
type Command struct {
	Title     string        `json:"title"`
	Command   string        `json:"command"`
	Arguments []interface{} `json:"arguments,omitempty"`
}
