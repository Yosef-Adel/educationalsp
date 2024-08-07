package lsp

// DidOpenTextDocumentNotification represents a notification for opening a text document
type DidOpenTextDocumentNotification struct {
	Notification
	Params DidOpenTextDocumentParams `json:"params"`
}

// DidOpenTextDocumentParams contains parameters for a text document open notification
type DidOpenTextDocumentParams struct {
	TextDocument TextDocumentItem `json:"textDocument"`
}
