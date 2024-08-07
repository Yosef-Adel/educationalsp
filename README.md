# Educational LSP

This project is an educational implementation of the Language Server Protocol (LSP). It provides various LSP features such as text document synchronization, hover information, code actions, completion, and diagnostics. The goal of this project is to help developers understand the implementation details and workings of an LSP server.

## Features

- **Initialize**: Establishes the connection between the client and server.
- **Text Document Did Open**: Handles opening of text documents.
- **Text Document Did Change**: Handles changes in text documents.
- **Hover**: Provides hover information for symbols.
- **Definition**: Provides the definition of symbols.
- **Code Action**: Suggests code actions for the document.
- **Completion**: Provides code completions.
- **Diagnostics**: Publishes diagnostics for documents.

## Project Structure

The project is organized into the following packages:

- `analysis`: Contains the logic for handling document state and providing diagnostics.
- `lsp`: Defines the structures and types for LSP requests, responses, and notifications.
- `rpc`: Handles encoding and decoding of messages between the client and server.

## Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/yosef-adel/educationalsp.git
   ```
2. Navigate to the project directory:
   ```sh
   cd educationalsp
   ```
3. Build the project:

   ```sh
   go build -o main

   ```

## Usage

Run the LSP server:

```sh
./main
```

The server will start and listen for incoming LSP messages from the client.

## Neovim Configuration

To connect Neovim to the LSP server, add the following configuration to your Neovim setup:

```lua
local client = vim.lsp.start_client({
    name = "educationalsp",
    cmd = { "PATH_TO_BIN" },
})

if not client then
    vim.notify("LSP not installed")
    return
end

vim.api.nvim_create_autocmd("FileType", {
    pattern = "markdown",
    callback = function()
        vim.lsp.buf_attach_client(0, client)
    end,
})
```

Replace `"PATH_TO_BIN"` with the path to your compiled LSP server binary.

## Example Messages

### Initialize Request

```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "method": "initialize",
  "params": {
    "clientInfo": {
      "name": "ExampleClient",
      "version": "1.0.0"
    }
  }
}
```

### Hover Request

```json
{
  "jsonrpc": "2.0",
  "id": 2,
  "method": "textDocument/hover",
  "params": {
    "textDocument": {
      "uri": "file://path/to/document"
    },
    "position": {
      "line": 10,
      "character": 5
    }
  }
}
```

## Logging

The server logs its activities to a file named `log.txt`. You can view the logs to understand the server's operations and debug any issues.

## Testing

The project includes unit tests for the `rpc` package. To run the tests, use the following command:

```sh
go test ./rpc
```
