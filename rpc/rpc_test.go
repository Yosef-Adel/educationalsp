package rpc_test

import (
	"educationalsp/rpc"
	"testing"
)

type EncodeingExample struct {
	Testing bool
}

func TestEncode(t *testing.T) {
	expected := "Content-Length: 16\r\n\r\n{\"Testing\":true}"
	actual := rpc.EncodeMessage(EncodeingExample{Testing: true})
	if expected != actual {
		t.Fatalf("Expected %s, actual %s", expected, actual)
	}
}

func TestDecode(t *testing.T) {
	incomingMsg := "Content-Length: 15\r\n\r\n{\"Method\":\"hi\"}"
	method, content, err := rpc.DecodeMessage([]byte(incomingMsg))
	contentLenght := len(content)

	if err != nil {
		t.Fatal(err)
	}

	if contentLenght != 15 {
		t.Fatalf("Expected 16 and Got %d", contentLenght)
	}

	if method != "hi" {
		t.Fatalf("Expected hi and Got %s", method)
	}

}
