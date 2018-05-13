package kv

type KVStore struct {
	Root *Node
}

func NewKVStore() *KVStore {
	root := NewNode(nil)

	return &KVStore{
		Root: root,
	}
}

func (s *KVStore) Put(key, value string) error {
	return s.Root.Put(key, value)
}

func (s *KVStore) Get(key string) (interface{}, bool) {
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
