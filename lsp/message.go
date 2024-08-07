package lsp

// Request represents a generic LSP request
type Request struct {
	RPC    string `json:"jsonrpc"`
	ID     int    `json:"id"`
	Method string `json:"method"`
	// Parameters for specific requests will be added in request types
}

// Response represents a generic LSP response
type Response struct {
	RPC string `json:"jsonrpc"`
	ID  *int   `json:"id,omitempty"`
	// Result and Error will be specified in response types
}

// Notification represents a generic LSP notification
type Notification struct {
	RPC    string `json:"jsonrpc"`
	Method string `json:"method"`
}
