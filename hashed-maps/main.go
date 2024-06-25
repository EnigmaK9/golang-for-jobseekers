package main

import (
    "fmt"
)

type HashedMaps struct {
    items [100]string
}

func NewHashedMaps() HashedMaps {
    return HashedMaps{items: [100]string{}}
}

func (h HashedMaps) GetHash(key string) int {
    totalSum := 0
    for _, v := range key {
        totalSum += int(v)
    }
    hashKey := totalSum % 100
    return hashKey
}

func (h *HashedMaps) Set(key, val string) {
    hashedKey := h.GetHash(key)
    h.items[hashedKey] = val
}

func (h *HashedMaps) Get(key string) string {
    hashedKey := h.GetHash(key)
    return h.items[hashedKey]
}

func main() {
    hashMap := NewHashedMaps()
    hashMap.Set("foo", "bar")
    hashMap.Set("baz", "qux")

    fmt.Println("Value for 'foo':", hashMap.Get("foo"))
    fmt.Println("Value for 'baz':", hashMap.Get("baz"))
}

