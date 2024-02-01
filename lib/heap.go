package lib

/* Implementing this binary heap
https://runestone.academy/ns/books/published/pythonds/Trees/BinaryHeapImplementation.html
*/

// Returns true if a is before b
type HeapComparisonFunc[T any] func(a, b T) bool

type Heap[T any] struct {
	items  []T
	before HeapComparisonFunc[T]
}

func NewHeap[T any](before HeapComparisonFunc[T]) Heap[T] {
	var zero T
	return Heap[T]{
		items:  []T{zero},
		before: before,
	}
}

func (heap *Heap[T]) Insert(item T) Heap[T] {
	heap.items = append(heap.items, item)

	// Compare the new item with its parent, and swap if necessary
	i := heap.Size()
	for i/2 > 0 {
		if heap.before(heap.items[i], heap.items[i/2]) {
			heap.items[i], heap.items[i/2] = heap.items[i/2], heap.items[i]
			i = i / 2
		} else {
			break
		}
	}

	return *heap
}

func (heap *Heap[T]) Pop() T {
	result := heap.items[1]

	// Promote last item to the top
	heap.items[1] = heap.items[heap.Size()]
	heap.items = heap.items[:heap.Size()]

	// Compare the top item with its smallest child, and swap if necessary
	i := 1
	for i*2 <= heap.Size() {
		leftChildIndex := i * 2
		rightChildIndex := i*2 + 1
		var smallestChildIndex int
		if rightChildIndex <= heap.Size() && heap.before(heap.items[rightChildIndex], heap.items[leftChildIndex]) {
			smallestChildIndex = rightChildIndex
		} else {
			smallestChildIndex = leftChildIndex
		}

		if heap.before(heap.items[smallestChildIndex], heap.items[i]) {
			heap.items[i], heap.items[smallestChildIndex] = heap.items[smallestChildIndex], heap.items[i]
			i = smallestChildIndex
		} else {
			break
		}
	}

	return result
}

func (heap Heap[T]) Size() int {
	return len(heap.items) - 1
}
