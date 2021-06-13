package storage

import (
	"fmt"
	"sync"
)

const (
	cacheTheshold = 1024
)

type KV struct {
	sstable *SSTable
	cache   *Cache
	offsets map[string]int
}

func NewKV() *KV {
	return &KV{
		sstable: NewSSTable(),
		cache: &Cache{
			mapping: map[string][]byte{},
			mu:      &sync.Mutex{},
		},
		offsets: map[string]int{},
	}
}

func (kv *KV) Set(key string, value []byte) {
	kv.cache.Set(key, value)
	if kv.cache.Size() >= cacheTheshold {
		data := kv.cache.Reset()
		kv.sstable.Insert(data)
	}
}

func (kv *KV) Get(key string) ([]byte, error) {
	cached, ok := kv.cache.Get(key)
	if !ok {
		fmt.Printf("Cache miss for key %s\n", key)
		return kv.sstable.Find(key)
	}
	return cached, nil
}
