package heap

import (
	"math/rand"
	"testing"

	. "github.com/Jcowwell/go-algorithm-club/Utils"
	. "golang.org/x/exp/slices"
)

func verifyMaxHeap(h Heap[int]) bool {
	for i := 0; i < h.count(); i++ {
		left := h.leftChildIndex(i)
		right := h.rightChildIndex(i)
		parent := h.parentIndex(i)
		if left < h.count() && h.nodes[i] < h.nodes[left] {
			return false
		}
		if right < h.count() && h.nodes[i] < h.nodes[right] {
			return false
		}
		if i > 0 && h.nodes[parent] < h.nodes[i] {
			return false
		}
	}
	return true
}

func verifyMinHeap(h Heap[int]) bool {
	for i := 0; i < h.count(); i++ {
		left := h.leftChildIndex(i)
		right := h.rightChildIndex(i)
		parent := h.parentIndex(i)
		if left < h.count() && h.nodes[i] > h.nodes[left] {
			return false
		}
		if right < h.count() && h.nodes[i] > h.nodes[right] {
			return false
		}
		if i > 0 && h.nodes[parent] > h.nodes[i] {
			return false
		}
	}
	return true
}

func isPermutation(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for len(a) > 0 {
		if i := Index(b, a[0]); i != -1 {
			a = Delete(a, 0, 1)
			b = Delete(b, i, i+1)
		} else {
			return false
		}
	}
	return len(b) == 0
}

func TestEmptyHeap(t *testing.T) {
	heap := Heap[int]{nodes: []int{}, orderCriteria: LessThan[int]}
	AssertTrue(heap.isEmpty(), t)
	AssertEqual(heap.count(), 0, t)
	_, validPeek := heap.pop()
	AssertFalse(validPeek, t)
	_, validPop := heap.pop()
	AssertFalse(validPop, t)
}

func TestIsEmpty(t *testing.T) {
	heap := Heap[int]{nodes: []int{}, orderCriteria: GreaterThan[int]}
	AssertTrue(heap.isEmpty(), t)
	heap.insert(1)
	AssertFalse(heap.isEmpty(), t)
	_, _ = heap.pop()
	AssertTrue(heap.isEmpty(), t)
}

func TestCount(t *testing.T) {
	heap := Heap[int]{nodes: []int{}, orderCriteria: GreaterThan[int]}
	AssertEqual(heap.count(), 0, t)
	heap.insert(1)
	AssertEqual(heap.count(), 1, t)
}

func TestMaxHeapOneElement(t *testing.T) {
	heap := Heap[int]{nodes: []int{10}, orderCriteria: GreaterThan[int]}
	AssertTrue(verifyMaxHeap(heap), t)
	AssertTrue(verifyMinHeap(heap), t)
	AssertFalse(heap.isEmpty(), t)
	AssertEqual(heap.count(), 1, t)
	valuePeek, _ := heap.peek()
	AssertEqual(valuePeek, 10, t)
}

func TestCreateMaxHeap(t *testing.T) {
	h1 := *HeapSliceInit([]int{1, 2, 3, 4, 5, 6, 7}, GreaterThan[int])
	AssertTrue(verifyMaxHeap(h1), t)
	AssertFalse(verifyMinHeap(h1), t)
	AssertEqualSlice(h1.nodes, []int{7, 5, 6, 4, 2, 1, 3}, t)
	AssertFalse(h1.isEmpty(), t)
	AssertEqual(h1.count(), 7, t)
	valuePeek1, _ := h1.peek()
	AssertEqual(valuePeek1, 7, t)

	h2 := *HeapSliceInit([]int{7, 6, 5, 4, 3, 2, 1}, GreaterThan[int])
	AssertTrue(verifyMaxHeap(h2), t)
	AssertFalse(verifyMinHeap(h2), t)
	AssertEqualSlice(h2.nodes, []int{7, 6, 5, 4, 3, 2, 1}, t)
	AssertFalse(h2.isEmpty(), t)
	AssertEqual(h2.count(), 7, t)
	valuePeek2, _ := h2.peek()
	AssertEqual(valuePeek2, 7, t)

	h3 := *HeapSliceInit([]int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}, GreaterThan[int])
	AssertTrue(verifyMaxHeap(h3), t)
	AssertFalse(verifyMinHeap(h3), t)
	AssertEqualSlice(h3.nodes, []int{16, 14, 10, 8, 7, 9, 3, 2, 4, 1}, t)
	AssertFalse(h3.isEmpty(), t)
	AssertEqual(h3.count(), 10, t)
	valuePeek3, _ := h3.peek()
	AssertEqual(valuePeek3, 16, t)

	h4 := *HeapSliceInit([]int{27, 17, 3, 16, 13, 10, 1, 5, 7, 12, 4, 8, 9, 0}, GreaterThan[int])
	AssertTrue(verifyMaxHeap(h4), t)
	AssertFalse(verifyMinHeap(h4), t)
	AssertEqualSlice(h4.nodes, []int{27, 17, 10, 16, 13, 9, 1, 5, 7, 12, 4, 8, 3, 0}, t)
	AssertFalse(h4.isEmpty(), t)
	AssertEqual(h4.count(), 14, t)
	valuePeek5, _ := h4.peek()
	AssertEqual(valuePeek5, 27, t)
}

func TestCreateMinHeap(t *testing.T) {
	h1 := *HeapSliceInit([]int{1, 2, 3, 4, 5, 6, 7}, LessThan[int])
	AssertTrue(verifyMinHeap(h1), t)
	AssertFalse(verifyMaxHeap(h1), t)
	AssertEqualSlice(h1.nodes, []int{1, 2, 3, 4, 5, 6, 7}, t)
	AssertFalse(h1.isEmpty(), t)
	AssertEqual(h1.count(), 7, t)
	valuePeek1, _ := h1.peek()
	AssertEqual(valuePeek1, 1, t)

	h2 := *HeapSliceInit([]int{7, 6, 5, 4, 3, 2, 1}, LessThan[int])
	AssertTrue(verifyMinHeap(h2), t)
	AssertFalse(verifyMaxHeap(h2), t)
	AssertEqualSlice(h2.nodes, []int{1, 3, 2, 4, 6, 7, 5}, t)
	AssertFalse(h2.isEmpty(), t)
	AssertEqual(h2.count(), 7, t)
	valuePeek2, _ := h2.peek()
	AssertEqual(valuePeek2, 1, t)

	h3 := *HeapSliceInit([]int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}, LessThan[int])
	AssertTrue(verifyMinHeap(h3), t)
	AssertFalse(verifyMaxHeap(h3), t)
	AssertEqualSlice(h3.nodes, []int{1, 2, 3, 4, 7, 9, 10, 14, 8, 16}, t)
	AssertFalse(h3.isEmpty(), t)
	AssertEqual(h3.count(), 10, t)
	valuePeek3, _ := h3.peek()
	AssertEqual(valuePeek3, 1, t)

	h4 := *HeapSliceInit([]int{27, 17, 3, 16, 13, 10, 1, 5, 7, 12, 4, 8, 9, 0}, LessThan[int])
	AssertTrue(verifyMinHeap(h4), t)
	AssertFalse(verifyMaxHeap(h4), t)
	AssertEqualSlice(h4.nodes, []int{0, 4, 1, 5, 12, 8, 3, 16, 7, 17, 13, 10, 9, 27}, t)
	AssertFalse(h4.isEmpty(), t)
	AssertEqual(h4.count(), 14, t)
	valuePeek4, _ := h4.peek()
	AssertEqual(valuePeek4, 0, t)
}

func TestCreateMaxHeapEqualnodes(t *testing.T) {
	heap := *HeapSliceInit([]int{1, 1, 1, 1, 1}, GreaterThan[int])
	AssertTrue(verifyMaxHeap(heap), t)
	AssertTrue(verifyMinHeap(heap), t)
	AssertEqualSlice(heap.nodes, []int{1, 1, 1, 1, 1}, t)
}

func TestCreateMinHeapEqualnodes(t *testing.T) {
	heap := *HeapSliceInit([]int{1, 1, 1, 1, 1}, LessThan[int])
	AssertTrue(verifyMinHeap(heap), t)
	AssertTrue(verifyMaxHeap(heap), t)
	AssertEqualSlice(heap.nodes, []int{1, 1, 1, 1, 1}, t)
}

func randomArray(n int) []int {
	a := []int{}
	for i := 0; i < n; i++ {
		a = append(a, rand.Int())
	}
	return a
}

func TestCreateRandomMaxHeap(t *testing.T) {
	for n := 1; n < 40; n++ {
		a := randomArray(n)
		h := *HeapSliceInit(a, GreaterThan[int])
		AssertTrue(verifyMaxHeap(h), t)
		AssertFalse(h.isEmpty(), t)
		AssertEqual(h.count(), n, t)
		AssertTrue(isPermutation(a, h.nodes), t)
	}
}

func TestCreateRandomMinHeap(t *testing.T) {
	for n := 1; n < 40; n++ {
		a := randomArray(n)
		h := *HeapSliceInit(a, LessThan[int])
		AssertTrue(verifyMinHeap(h), t)
		AssertFalse(h.isEmpty(), t)
		AssertEqual(h.count(), n, t)
		AssertTrue(isPermutation(a, h.nodes), t)
	}
}

func TestPoping(t *testing.T) {
	h := *HeapSliceInit([]int{100, 50, 70, 10, 20, 60, 65}, GreaterThan[int])
	AssertTrue(verifyMaxHeap(h), t)
	AssertEqualSlice(h.nodes, []int{100, 50, 70, 10, 20, 60, 65}, t)

	//test index out of bounds
	_, validPop := h.popAt(10)
	AssertFalse(validPop, t)
	AssertTrue(verifyMaxHeap(h), t)
	AssertEqualSlice(h.nodes, []int{100, 50, 70, 10, 20, 60, 65}, t)

	valuePop2, _ := h.popAt(5)
	AssertEqual(valuePop2, 60, t)
	AssertTrue(verifyMaxHeap(h), t)
	AssertEqualSlice(h.nodes, []int{100, 50, 70, 10, 20, 65}, t)

	valuePop3, _ := h.popAt(4)
	AssertEqual(valuePop3, 20, t)
	AssertTrue(verifyMaxHeap(h), t)
	AssertEqualSlice(h.nodes, []int{100, 65, 70, 10, 50}, t)

	valuePop4, _ := h.popAt(4)
	AssertEqual(valuePop4, 50, t)
	AssertTrue(verifyMaxHeap(h), t)
	AssertEqualSlice(h.nodes, []int{100, 65, 70, 10}, t)

	valuePop5, _ := h.popAt(0)
	AssertEqual(valuePop5, 100, t)
	AssertTrue(verifyMaxHeap(h), t)
	AssertEqualSlice(h.nodes, []int{70, 65, 10}, t)

	valuePeek1, _ := h.peek()
	AssertEqual(valuePeek1, 70, t)
	valuePop6, _ := h.pop()
	AssertEqual(valuePop6, 70, t)
	AssertTrue(verifyMaxHeap(h), t)
	AssertEqualSlice(h.nodes, []int{65, 10}, t)

	valuePeek2, _ := h.peek()
	AssertEqual(valuePeek2, 65, t)
	valuePop7, _ := h.pop()
	AssertEqual(valuePop7, 65, t)
	AssertTrue(verifyMaxHeap(h), t)
	AssertEqualSlice(h.nodes, []int{10}, t)

	valuePeek3, _ := h.peek()
	AssertEqual(valuePeek3, 10, t)
	valuePop8, _ := h.pop()
	AssertEqual(valuePop8, 10, t)
	AssertTrue(verifyMaxHeap(h), t)
	AssertEqualSlice(h.nodes, []int{}, t)
	_, validPeek4 := h.peek()
	AssertFalse(validPeek4, t)
}

func TestRemoveEmpty(t *testing.T) {
	heap := Heap[int]{orderCriteria: GreaterThan[int]}
	_, validPop := heap.pop()
	AssertFalse(validPop, t)
}

func TestRemoveRoot(t *testing.T) {
	h := Heap[int]{nodes: []int{15, 13, 9, 5, 12, 8, 7, 4, 0, 6, 2, 1}, orderCriteria: GreaterThan[int]}
	AssertTrue(verifyMaxHeap(h), t)
	AssertEqualSlice(h.nodes, []int{15, 13, 9, 5, 12, 8, 7, 4, 0, 6, 2, 1}, t)
	valuePeek, _ := h.peek()
	AssertEqual(valuePeek, 15, t)
	valuePop, _ := h.pop()
	AssertEqual(valuePop, 15, t)
	AssertTrue(verifyMaxHeap(h), t)
	AssertEqualSlice(h.nodes, []int{13, 12, 9, 5, 6, 8, 7, 4, 0, 1, 2}, t)
}

func TestRemoveRandomItems(t *testing.T) {
	for n := 1; n < 40; n++ {
		a := randomArray(n)
		h := *HeapSliceInit(a, GreaterThan[int])
		AssertTrue(verifyMaxHeap(h), t)
		AssertTrue(isPermutation(a, h.nodes), t)

		m := (n + 1) / 2
		for k := 1; k < m; k++ {
			i := int(rand.Int31n(int32(n - k + 1)))
			valuePop, _ := h.popAt(i)
			j := Index(a, valuePop)
			a = Delete(a, j, j+1)

			AssertTrue(verifyMaxHeap(h), t)
			AssertEqual(h.count(), len(a), t)
			AssertEqual(h.count(), n-k, t)
			AssertTrue(isPermutation(a, h.nodes), t)
		}
	}
}

func TestInsert(t *testing.T) {
	h := Heap[int]{nodes: []int{15, 13, 9, 5, 12, 8, 7, 4, 0, 6, 2, 1}, orderCriteria: GreaterThan[int]}
	AssertTrue(verifyMaxHeap(h), t)
	AssertEqualSlice(h.nodes, []int{15, 13, 9, 5, 12, 8, 7, 4, 0, 6, 2, 1}, t)

	h.insert(10)
	AssertTrue(verifyMaxHeap(h), t)
	AssertEqualSlice(h.nodes, []int{15, 13, 10, 5, 12, 9, 7, 4, 0, 6, 2, 1, 8}, t)
}

func TestInsertArrayAndRemove(t *testing.T) {
	heap := Heap[int]{orderCriteria: GreaterThan[int]}
	heap.insertSequence([]int{1, 3, 2, 7, 5, 9}...)
	AssertEqualSlice(heap.nodes, []int{9, 5, 7, 1, 3, 2}, t)

	valuePop, _ := heap.pop()
	AssertEqual(valuePop, 9, t)
	valuePop2, _ := heap.pop()
	AssertEqual(valuePop2, 7, t)
	valuePop3, _ := heap.pop()
	AssertEqual(valuePop3, 5, t)
	valuePop4, _ := heap.pop()
	AssertEqual(valuePop4, 3, t)
	valuePop5, _ := heap.pop()
	AssertEqual(valuePop5, 2, t)
	valuePop6, _ := heap.pop()
	AssertEqual(valuePop6, 1, t)
	_, validPop7 := heap.pop()
	AssertFalse(validPop7, t)
}

func TestReplace(t *testing.T) {
	h := Heap[int]{nodes: []int{16, 14, 10, 8, 7, 9, 3, 2, 4, 1}, orderCriteria: GreaterThan[int]}
	AssertTrue(verifyMaxHeap(h), t)

	h.replace(5, 13)
	AssertTrue(verifyMaxHeap(h), t)

	//test index out of bounds
	h.replace(20, 2)
	AssertTrue(verifyMaxHeap(h), t)
}
