package lsp

// DefinitionRequest represents a request for a definition
type DefinitionRequest struct {
	Request
	Params DefinitionParams `json:"params"`
}

// DefinitionParams contains parameters for a definition request
type DefinitionParams struct {
	TextDocumentPositionParams
}

// DefinitionResponse represents a response for a definition request
type DefinitionResponse struct {
	Response
	Result Location `json:"result"`
}
