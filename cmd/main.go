package main

// import (
// 	"fmt"
// 	"time"

// 	"github.com/abhaybhu10/kv-store/storage"
// 	"github.com/google/uuid"
// )

// func main1() {
// 	kv := storage.NewKV()
// 	for {
// 		go func() {
// 			key := uuid.NewString()
// 			value := uuid.NewString()
// 			kv.Set(key, []byte(value))
// 			val, err := kv.Get(key)

// 			if err != nil {
// 				fmt.Printf("Error occured %s\n", err.Error())
// 			} else {
// 				fmt.Printf("Key %s : value %s\n", "abhay", string(val))
// 			}
// 		}()
// 	}
// 	time.Sleep(10 * time.Second)

// }
