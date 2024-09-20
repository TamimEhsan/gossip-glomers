package main

import (
	"encoding/json"
	"log"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type server struct {
	node *maelstrom.Node
}

func (s *server) echoHandler(msg maelstrom.Message) error {
	var body map[string]any
	if err := json.Unmarshal(msg.Body, &body); err != nil {
		return err
	}

	// Update the message type to return back.
	body["type"] = "echo_ok"

	// Echo the original message back with the updated message type.
	return s.node.Reply(msg, body)
}

func main() {
	s := &server{}
	s.node = maelstrom.NewNode()

	s.node.Handle("echo", s.echoHandler)
	if err := s.node.Run(); err != nil {
		log.Fatal(err)
	}
}
