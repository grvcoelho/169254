package kv

type Store struct {
	Root *Node
}

func NewStore() *Store {
	root := NewNode(nil)

	return &Store{
		Root: root,
	}
}

func (s *Store) Put(key, value string) error {
	return s.Root.Put(key, value)
}

func (s *Store) Get(key string) (interface{}, bool) {
	child, ok := s.Root.Get(key)

	if !ok {
		return nil, false
	}

	if len(child.Children) == 0 {
		return child.Value, true
	}

	if len(child.Children) > 0 {
		return child.ListKeys(), true
	}

	return nil, false
}
