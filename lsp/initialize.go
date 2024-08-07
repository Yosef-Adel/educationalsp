package lsp

// InitializeRequest represents an initialization request from the client
type InitializeRequest struct {
	Request
	Params InitializeRequestParams `json:"params"`
}

// InitializeRequestParams contains parameters for the initialization request
type InitializeRequestParams struct {
	ClientInfo *ClientInfo `json:"clientInfo"`
}

// ClientInfo represents information about the client
type ClientInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// InitializeResponse represents a response to an initialization request
type InitializeResponse struct {
	Response
	Result InitializeResult `json:"result"`
}

// InitializeResult contains the result of the initialization response
type InitializeResult struct {
	Capabilities ServerCapabilities `json:"capabilities"`
	ServerInfo   ServerInfo         `json:"serverInfo"`
}

// ServerCapabilities represents the capabilities of the server
type ServerCapabilities struct {
	TextDocumentSync   int            `json:"textDocumentSync"`
	HoverProvider      bool           `json:"hoverProvider"`
	DefinitionProvider bool           `json:"definitionProvider"`
	CodeActionProvider bool           `json:"codeActionProvider"`
	CompletionProvider map[string]any `json:"completionProvider"`
}

// ServerInfo represents information about the server
type ServerInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// NewInitializeResponse creates a new InitializeResponse with predefined capabilities and server info
func NewInitializeResponse(id int) InitializeResponse {
	return InitializeResponse{
		Response: Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: InitializeResult{
			Capabilities: ServerCapabilities{
				TextDocumentSync:   1,
				HoverProvider:      true,
				DefinitionProvider: true,
				CodeActionProvider: true,
				CompletionProvider: map[string]any{},
			},
			ServerInfo: ServerInfo{
				Name:    "educationalsp",
				Version: "0.0.0-beta1",
			},
		},
	}
}
