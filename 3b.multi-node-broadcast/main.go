package main

import (
	"encoding/json"
	"log"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type server struct {
	node  *maelstrom.Node
	store []int
}

func (s *server) handleBroadcast(msg maelstrom.Message) error {
	var body map[string]any
	if err := json.Unmarshal(msg.Body, &body); err != nil {
		return err
	}
	var ret map[string]any = make(map[string]any)
	// Update the message type to return back.
	ret["type"] = "broadcast_ok"
	message := int(body["message"].(float64))
	s.store = append(s.store, message)

	// Echo the original message back with the updated message type.
	return s.node.Reply(msg, ret)
}

func (s *server) handleRead(msg maelstrom.Message) error {
	var body map[string]any
	if err := json.Unmarshal(msg.Body, &body); err != nil {
		return err
	}
	var ret map[string]any = make(map[string]any)

	// Update the message type to return back.
	ret["type"] = "read_ok"
	ret["messages"] = s.store

	// Echo the original message back with the updated message type.
	return s.node.Reply(msg, ret)
}

func (s *server) handleTopology(msg maelstrom.Message) error {
	var body map[string]any
	if err := json.Unmarshal(msg.Body, &body); err != nil {
		return err
	}
	var ret map[string]any = make(map[string]any)

	// Update the message type to return back.
	ret["type"] = "topology_ok"

	// Echo the original message back with the updated message type.
	return s.node.Reply(msg, ret)
}

func main() {
	s := &server{}
	s.node = maelstrom.NewNode()

	s.node.Handle("broadcast", s.handleBroadcast)
	s.node.Handle("read", s.handleRead)
	s.node.Handle("topology", s.handleTopology)

	if err := s.node.Run(); err != nil {
		log.Fatal(err)
	}
}
