package storage

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"sync"
)

type Storage interface {
	Write(file string, data []byte)
	Read(file string, offset, length int) ([]byte, error)
}

type DiskStorage struct {
	fileHandles map[string]*os.File
	mu          *sync.Mutex
}

func NewDiskStorage() *DiskStorage {
	return &DiskStorage{
		fileHandles: map[string]*os.File{},
		mu:          &sync.Mutex{},
	}
}

func (st *DiskStorage) Write(fileName string, data []byte) {
	fileHandle, ok := st.fileHandles[fileName]
	if !ok {
		newHandle, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0600)
		if err != nil {
			fmt.Printf("Error while opening file %s %s\n", fileName, err.Error())
		}
		st.fileHandles[fileName] = fileHandle
		fileHandle = newHandle
	}
	fileHandle.Write(data)
	fmt.Printf("Data write to file %s\n", fileName)
}

func (st *DiskStorage) Read(file string, start, length int) ([]byte, error) {
	handle, err := os.Open(file)
	if err != nil {
		fmt.Printf("File %s does not exits\n", file)
	}
	buffer := bytes.NewBuffer(make([]byte, length))
	for length > 0 {
		data := make([]byte, length)
		n, err := handle.ReadAt(data, int64(start))
		if err != nil && err != io.EOF {
			return nil, errors.New(fmt.Sprintf("Read failed from file %s", err.Error()))
		}
		if err == io.EOF {
			return nil, errors.New("Corrrupted data")
		}

		buffer.Write(data[:n])
		start += n
		length -= n
	}
	return buffer.Bytes(), nil
}
