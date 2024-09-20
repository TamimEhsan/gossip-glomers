package main

import (
	"encoding/json"
	"log"

	"github.com/google/uuid"
	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type server struct {
	node *maelstrom.Node
}

func (s *server) handleUniqueIds(msg maelstrom.Message) error {
	var body map[string]any
	if err := json.Unmarshal(msg.Body, &body); err != nil {
		return err
	}

	// Update the message type to return back.
	body["type"] = "generate_ok"
	body["id"] = uuid.New().String()

	// Echo the original message back with the updated message type.
	return s.node.Reply(msg, body)
}

func main() {
	s := &server{}
	s.node = maelstrom.NewNode()

	s.node.Handle("generate", s.handleUniqueIds)
	if err := s.node.Run(); err != nil {
		log.Fatal(err)
	}
} 
