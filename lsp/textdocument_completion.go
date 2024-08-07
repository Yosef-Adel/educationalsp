package lsp

// CompletionRequest represents a request for code completions
type CompletionRequest struct {
	Request
	Params CompletionParams `json:"params"`
}

// CompletionParams contains parameters for a completion request
type CompletionParams struct {
	TextDocumentPositionParams
}

// CompletionResponse represents a response for a completion request
type CompletionResponse struct {
	Response
	Result []CompletionItem `json:"result"`
}

// CompletionItem represents a single completion item
type CompletionItem struct {
	Label         string `json:"label"`
	Detail        string `json:"detail"`
	Documentation string `json:"documentation"`
}
