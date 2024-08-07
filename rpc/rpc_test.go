package rpc_test

import (
	"educationalsp/rpc"
	"testing"
)

// EncodeExample is a test structure for encoding
type EncodeExample struct {
	Testing bool `json:"Testing"`
}

// TestEncode verifies the EncodeMessage function
func TestEncode(t *testing.T) {
	expected := "Content-Length: 16\r\n\r\n{\"Testing\":true}"
	actual := rpc.EncodeMessage(EncodeExample{Testing: true})
	if expected != actual {
		t.Fatalf("Expected %s, actual %s", expected, actual)
	}
}

// TestDecode verifies the DecodeMessage function
func TestDecode(t *testing.T) {
	incomingMsg := "Content-Length: 15\r\n\r\n{\"Method\":\"hi\"}"
	method, content, err := rpc.DecodeMessage([]byte(incomingMsg))
	contentLength := len(content)

	if err != nil {
		t.Fatal(err)
	}

	if contentLength != 15 {
		t.Fatalf("Expected 15 and got %d", contentLength)
	}

	if method != "hi" {
		t.Fatalf("Expected 'hi' and got %s", method)
	}
}
