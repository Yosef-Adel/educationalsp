package rpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

func EncodeMessage(msg any) string {
	content, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("Content-Length: %d\r\n\r\n%s", len(content), content)
}

type BaseMessage struct {
	Method string `json:"method"`
}

func DecodeMessage(msg []byte) (string, []byte, error) {
	header, content, found := bytes.Cut(msg, []byte{'\r', '\n', '\r', '\n'})

	if !found {
		return "", nil, errors.New("Did not found the seperator")
	}

	// Content-Length: <number>
	contentLenghtBytes := header[len("Content-Length: "):]
	contentLenght, err := strconv.Atoi(string(contentLenghtBytes))
	if err != nil {
		return "", nil, err
	}

	var baseMessage BaseMessage
	if err := json.Unmarshal(content[:contentLenght], &baseMessage); err != nil {
		return "", nil, err
	}

	return baseMessage.Method, content[:contentLenght], nil
}

// IMPLEMENT Something like this to give to the scanner
// type SplitFunc func(data []byte, atEOF bool) (advance int, token []byte, err error)
func Split(data []byte, _ bool) (advance int, token []byte, err error) {

	header, content, found := bytes.Cut(data, []byte{'\r', '\n', '\r', '\n'})

	if !found {
		return 0, nil, nil
	}
	contentLenghtBytes := header[len("Content-Length: "):]
	contentLenght, err := strconv.Atoi(string(contentLenghtBytes))
	if err != nil {
		return 0, nil, err
	}
	if len(content) < contentLenght {
		return 0, nil, nil
	}

	totalLenght := len(header) + 4 + contentLenght
	return totalLenght, data[:totalLenght], nil
}
