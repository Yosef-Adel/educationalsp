package lsp

// DidChangeTextDocumentNotification represents a notification for changes in a text document
type DidChangeTextDocumentNotification struct {
	Notification
	Params DidChangeTextDocumentParams `json:"params"`
}

// DidChangeTextDocumentParams contains parameters for a text document change notification
type DidChangeTextDocumentParams struct {
	TextDocument   VersionTextDocumentIdentifier    `json:"textDocument"`
	ContentChanges []TextDocumentContentChangeEvent `json:"contentChanges"`
}

// TextDocumentContentChangeEvent represents a content change event in a text document
type TextDocumentContentChangeEvent struct {
	Text string `json:"text"`
}
