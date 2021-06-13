package storage

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/google/uuid"
)

//storage offsets for each segment. writes segment to the disk.
//maintains files names,

const (
	filePrefix = "/tmp/uncompact/"
)

type SSTable struct {
	ssids   []node
	disk    *DiskStorage
	counter int
	mu      sync.Mutex
}

func NewSSTable() *SSTable {
	return &SSTable{
		ssids:   []node{},
		disk:    NewDiskStorage(),
		counter: 0,
		mu:      sync.Mutex{},
	}
}

type position struct {
	start  int
	length int
}
type node struct {
	id       int
	offset   map[string]position
	location string
}

func (ss *SSTable) Insert(cache map[string][]byte) {
	ss.mu.Lock()
	defer ss.mu.Unlock()
	data, offsets := seralizeMap(cache)
	ssNode := node{
		offset:   offsets,
		location: ss.getFileLocation(),
		id:       ss.counter,
	}
	ss.ssids = append(ss.ssids, ssNode)
	ss.disk.Write(ssNode.location, data)
}

func (ss *SSTable) Find(key string) ([]byte, error) {
	for i := range ss.ssids {
		node := ss.ssids[len(ss.ssids)-i-1]
		if offset, ok := node.offset[key]; ok {
			data, err := ss.disk.Read(node.location, offset.start, offset.length)
			if err != nil {
				return nil, err
			}
			_, value := deserelize(data)
			return value, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("key %s does not exit", key))
}

func (ss *SSTable) getFileLocation() string {
	return filePrefix + uuid.NewString()
}

func seralizeMap(data map[string][]byte) ([]byte, map[string]position) {
	positions := map[string]position{}
	var b bytes.Buffer
	for k, v := range data {
		start := 0
		sb := strings.Builder{}
		sb.WriteString(k)
		sb.WriteString(":")
		sb.Write(v)
		sb.WriteString(",")
		s := sb.String()
		position := position{start: start, length: len(s) - 1}
		positions[k] = position
		start += len(s)
		b.WriteString(sb.String())
	}
	return b.Bytes(), positions
}

func deserelize(data []byte) (string, []byte) {
	kv := bytes.Split(data, []byte(":"))
	return string(kv[0]), kv[1]
}
