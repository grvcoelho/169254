package kv

import (
	"reflect"
	"strings"
)

type Node struct {
	Value    interface{}
	Children map[string]*Node
}

func NewNode(value interface{}) (*Node, error) {
	return &Node{
		Value:    value,
		Children: make(map[string]*Node),
	}, nil
}

func (n *Node) Put(key string, value interface{}) error {
	keys := strings.Split(key, "/")

	if len(keys) == 1 {
		child, _ := NewNode(value)
		n.Children[keys[0]] = child

		return nil
	}

	newKey := strings.Join(keys[1:], "/")
	child, _ := NewNode(nil)
	n.Children[keys[0]] = child

	return child.Put(newKey, value)
}

func (n *Node) Get(key string) (*Node, bool) {
	keys := strings.Split(key, "/")

	child, ok := n.Children[keys[0]]

	if len(keys) == 1 {
		return child, ok
	}

	if !ok {
		return nil, false
	}

	newKey := strings.Join(keys[1:], "/")
	return child.Get(newKey)
}

func (n *Node) ListKeys() []string {
	keys := make([]string, 0, len(n.Children))

	for _, key := range reflect.ValueOf(n.Children).MapKeys() {
		keys = append(keys, key.Interface().(string))
	}

	return keys
}
