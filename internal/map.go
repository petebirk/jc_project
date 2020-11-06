/*
 * JumpCloud Project API
 *
 * This is a simple API for JumpCloud Project
 *
 * API version: 1.0.0
 * Contact: birk.pete@gmail.com
 */
package main

import (
	"sync"
	"time"
	"fmt"
)

/* TODO: This information needs to be persisted.
   This is fairly easy to convert to a DB client
   as it's isolated from the handler. */
var (
	counter int = 1
	pending int = 0
	allData = make(map[int]string)
	rwmutex     sync.RWMutex
)

func GetHash(key int) string {
	rwmutex.RLock()
	defer rwmutex.RUnlock()
	return allData[key]

}

func RequestsPending() bool {
	fmt.Printf("pending count: %d\n", pending);
	return pending != 0;
}
func SetHashWithDelay(key int, value string) {
	defer decrementPendingCount();
	time.Sleep(5 * time.Second)
	encodedPassword := HashPassword(value)
	allData[key] = encodedPassword
}

func decrementPendingCount() {
	pending--;
}

func incrementPendingCount() {
	pending++;
}

func SetHash(value string) int {
	rwmutex.Lock()
	defer rwmutex.Unlock()
	incrementPendingCount();
	var key int
	key = counter
	counter++
	go SetHashWithDelay(key, value)
	return key
}

/**
func main() {
	key1 := SetHash("angryMonkey1")
	result1 := GetHash(key1)
	fmt.Println("key = ", key1, "value = ", result1)

	key2 := SetHash("angryMonkey2")
	result2 := GetHash(key2)
	fmt.Println("key = ", key2, "value = ", result2)

	key3 := SetHash("angryMonkey3")
	key4 := SetHash("angryMonkey4")
	key5 := SetHash("angryMonkey5")
	result3 := GetHash(key3)
	result4 := GetHash(key4)
	result5 := GetHash(key5)
	fmt.Println("key = ", key3, "value = ", result3)
	fmt.Println("key = ", key4, "value = ", result4)
	fmt.Println("key = ", key5, "value = ", result5)

	result1 = GetHash(key1)
	fmt.Println("key = ", key1, "value = ", result1)
	result2 = GetHash(key2)
	fmt.Println("key = ", key2, "value = ", result2)
	
}
**/