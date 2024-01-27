package lib

import (
	"testing"
)

func TestNewHeap(t *testing.T) {
	result := NewHeap[int](func(a, b int) bool {
		return a < b
	})

	if len(result.items) != 1 {
		t.Errorf("Expected empty heap, got %v", result.items)
	}
}

func TestInsert(t *testing.T) {
	heap := NewHeap[int](func(a, b int) bool {
		return a < b
	})

	heap.Insert(1)
	heap.Insert(2)

	if len(heap.items) != 3 {
		t.Errorf("Expected heap to have 3 items, got %v", heap.items)
	}
}

func TestPop(t *testing.T) {
	heap := NewHeap[int](func(a, b int) bool {
		return a < b
	})

	heap.Insert(1)

	result := heap.Pop()

	if result != 1 {
		t.Errorf("Expected heap to pop 1, got %v", result)
	}
}

func TestPopReducesLength(t *testing.T) {
	heap := NewHeap[int](func(a, b int) bool {
		return a < b
	})

	heap.Insert(1)

	length := len(heap.items)

	heap.Pop()

	if len(heap.items) != length-1 {
		t.Errorf("Expected heap to have length %v, got %v", length-1, len(heap.items))
	}
}

func TestOuputOrder(t *testing.T) {
	heap := NewHeap[int](func(a, b int) bool {
		return a < b
	})

	heap.Insert(4)
	heap.Insert(5)
	heap.Insert(6)
	heap.Insert(7)
	heap.Insert(3)
	heap.Insert(2)
	heap.Insert(1)

	result := heap.Pop()

	if result != 1 {
		t.Errorf("Expected heap to pop 1, got %v", result)
	}

	result = heap.Pop()

	if result != 2 {
		t.Errorf("Expected heap to pop 2, got %v", result)
	}

	result = heap.Pop()

	if result != 3 {
		t.Errorf("Expected heap to pop 3, got %v", result)
	}
}
