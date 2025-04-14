package shardkv

// SkipList is a placeholder implementation. Replace with your actual implementation.
type SkipList struct {
	Data map[string]string
}

func NewSkipList() *SkipList {
	return &SkipList{Data: make(map[string]string)}
}

func (sl *SkipList) Get(key string) (string, bool) {
	value, ok := sl.Data[key]
	return value, ok
}

func (sl *SkipList) Put(key, value string) {
	sl.Data[key] = value
}

func (sl *SkipList) GetAll() map[string]string {
	copy := make(map[string]string)
	for k, v := range sl.Data {
		copy[k] = v
	}
	return copy
}

type MemoryKVStateMachine struct {
	KV     *SkipList
	Status ShardStatus
}

func NewMemoryKVStateMachine() *MemoryKVStateMachine {
	return &MemoryKVStateMachine{
		KV:     NewSkipList(),
		Status: Normal,
	}
}

func (mkv *MemoryKVStateMachine) Get(key string) (string, Err) {
	value, ok := mkv.KV.Get(key)
	if !ok {
		return "", ErrNoKey
	}
	return value, OK
}

func (mkv *MemoryKVStateMachine) Put(key, value string) Err {
	mkv.KV.Put(key, value)
	return OK
}

func (mkv *MemoryKVStateMachine) Append(key, value string) Err {
	oldValue, ok := mkv.KV.Get(key)
	if ok {
		mkv.KV.Put(key, oldValue+value)
	} else {
		mkv.KV.Put(key, value)
	}
	return OK
}

func (mkv *MemoryKVStateMachine) copyData() map[string]string {
	return mkv.KV.GetAll()
}
