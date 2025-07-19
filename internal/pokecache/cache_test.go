package pokecache

import (
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	cache := NewCache(5 * time.Second)

	cache.Add("test-key", []byte("test-data"))
	result, exists := cache.Get("test-key")

	if !exists {
		t.Error("нема ключа")
	}
	if string(result) != "test-data" {
		t.Error("дані не вірні")
	}
}
