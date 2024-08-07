package lsp

// HoverRequest represents a request for a hover action
type HoverRequest struct {
	Request
	Params HoverParams `json:"params"`
}

// HoverParams contains parameters for a hover request
type HoverParams struct {
	TextDocumentPositionParams
}

// HoverResponse represents a response to a hover request
type HoverResponse struct {
	Response
	Result HoverResult `json:"result"`
}

// HoverResult represents the result of a hover action
type HoverResult struct {
	Contents string `json:"contents"`
}
