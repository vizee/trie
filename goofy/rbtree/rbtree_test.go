package rbtree

import (
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestInsert(t *testing.T) {
	tr := &Tree{}
	for i := 0; i < 100; i++ {
		n := rand.Int() % 100
		t.Log("> insert", n)
		t.Log(tr.Insert(n))
	}
}

func TestDelete(t *testing.T) {
	nums := make([]int, 100)
	for i := range nums {
		nums[i] = rand.Int() % 100
	}
	tr := &Tree{}
	for _, n := range nums {
		t.Log("> insert", n, tr.Insert(n))
	}
	for _, n := range nums {
		t.Log("> delete", n)
		t.Log(tr.Delete(n))
	}
}
