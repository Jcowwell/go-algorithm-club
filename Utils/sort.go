package util

func isNaN[N Numeric](n N) bool {
	return n != n
}

func less[N Numeric](seq []N, i, j int) bool {
	size := len(seq)
	if size < 2 {
		panic("slice: Index out of bounds")
	}
	if i >= size || j >= size {
		panic("slice: Index out of bounds")
	}
	return (seq[i] < seq[j]) || (isNaN(seq[i]) && !isNaN(seq[j]))
}

func swap[N Numeric](seq []N, i, j int) {
	size := len(seq)
	if size < 2 {
		return
	}
	if i >= size || j >= size {
		panic("slice: Index out of bounds")
	}
	seq[i], seq[j] = seq[j], seq[i]
}

func siftDown[N Numeric](seq []N, lo, hi, first int) {
	root := lo
	for {
		child := 2*root + 1
		if child >= hi {
			break
		}
		if child+1 < hi && less(seq, first+child, first+child+1) {
			child++
		}
		if !less(seq, first+root, first+child) {
			return
		}
		swap(seq, first+root, first+child)
		root = child
	}
}

func heapSort[N Numeric](seq []N, a, b int) {
	first := a
	lo := 0
	hi := b - a

	// Build heap with greatest element at top.
	for i := (hi - 1) / 2; i >= 0; i-- {
		siftDown(seq, i, hi, first)
	}

	// Pop elements, largest first, into end of seq.
	for i := hi - 1; i >= 0; i-- {
		swap(seq, first, first+i)
		siftDown(seq, lo, i, first)
	}
}

// medianOfThree moves the median of the three values seq[m0], seq[m1], seq[m2] into seq[m1].
func medianOfThree[N Numeric](seq []N, m1, m0, m2 int) {
	// sort 3 elements
	if less(seq, m1, m0) {
		swap(seq, m1, m0)
	}
	// seq[m0] <= seq[m1]
	if less(seq, m2, m1) {
		swap(seq, m2, m1)
		// seq[m0] <= seq[m2] && seq[m1] < seq[m2]
		if less(seq, m1, m0) {
			swap(seq, m1, m0)
		}
	}
	// now seq[m0] <= seq[m1] <= seq[m2]
}

func doPivot[N Numeric](seq []N, lo, hi int) (midlo, midhi int) {
	m := int(uint(lo+hi) >> 1) // Written like this to avoid integer overflow.
	if hi-lo > 40 {
		// Tukey's ``Ninther,'' median of three medians of three.
		s := (hi - lo) / 8
		medianOfThree(seq, lo, lo+s, lo+2*s)
		medianOfThree(seq, m, m-s, m+s)
		medianOfThree(seq, hi-1, hi-1-s, hi-1-2*s)
	}
	medianOfThree(seq, lo, m, hi-1)

	// Invariants are:
	//	seq[lo] = pivot (set up by ChoosePivot)
	//	seq[lo < i < a] < pivot
	//	seq[a <= i < b] <= pivot
	//	seq[b <= i < c] unexamined
	//	seq[c <= i < hi-1] > pivot
	//	seq[hi-1] >= pivot
	pivot := lo
	a, c := lo+1, hi-1

	for ; a < c && less(seq, a, pivot); a++ {
	}
	b := a
	for {
		for ; b < c && !less(seq, pivot, b); b++ { // seq[b] <= pivot
		}
		for ; b < c && less(seq, pivot, c-1); c-- { // seq[c-1] > pivot
		}
		if b >= c {
			break
		}
		// seq[b] > pivot; seq[c-1] <= pivot
		swap(seq, b, c-1)
		b++
		c--
	}
	// If hi-c<3 then there are duplicates (by property of median of nine).
	// Let's be a bit more conservative, and set border to 5.
	protect := hi-c < 5
	if !protect && hi-c < (hi-lo)/4 {
		// Lets test some points for equality to pivot
		dups := 0
		if !less(seq, pivot, hi-1) { // seq[hi-1] = pivot
			swap(seq, c, hi-1)
			c++
			dups++
		}
		if !less(seq, b-1, pivot) { // seq[b-1] = pivot
			b--
			dups++
		}
		// m-lo = (hi-lo)/2 > 6
		// b-lo > (hi-lo)*3/4-1 > 8
		// ==> m < b ==> seq[m] <= pivot
		if !less(seq, m, pivot) { // seq[m] = pivot
			swap(seq, m, b-1)
			b--
			dups++
		}
		// if at least 2 points are equal to pivot, assume skewed distribution
		protect = dups > 1
	}
	if protect {
		// Protect against a lot of duplicates
		// Add invariant:
		//	seq[a <= i < b] unexamined
		//	seq[b <= i < c] = pivot
		for {
			for ; a < b && !less(seq, b-1, pivot); b-- { // seq[b] == pivot
			}
			for ; a < b && less(seq, a, pivot); a++ { // seq[a] < pivot
			}
			if a >= b {
				break
			}
			// seq[a] == pivot; seq[b-1] < pivot
			swap(seq, a, b-1)
			a++
			b--
		}
	}
	// Swap pivot into middle
	swap(seq, pivot, b-1)
	return b - 1, c
}

func insertionSort[N Numeric](seq []N, a, b int) {
	for i := a + 1; i < b; i++ {
		for j := i; j > a && less(seq, j, j-1); j-- {
			swap(seq, j, j-1)
		}
	}
}

func quickSort[N Numeric](seq []N, a, b, maxDepth int) {
	for b-a > 12 { // Use ShellSort for slices <= 12 elements
		if maxDepth == 0 {
			heapSort(seq, a, b)
			return
		}
		maxDepth--
		mlo, mhi := doPivot(seq, a, b)
		// Avoiding recursion on the larger subproblem guarantees
		// a stack depth of at most lg(b-a).
		if mlo-a < b-mhi {
			quickSort(seq, a, mlo, maxDepth)
			a = mhi // i.e., quickSort(seq, mhi, b)
		} else {
			quickSort(seq, mhi, b, maxDepth)
			b = mlo // i.e., quickSort(seq, a, mlo)
		}
	}
	if b-a > 1 {
		// Do ShellSort pass with gap 6
		// It could be written in this simplified form cause b-a <= 12
		for i := a + 6; i < b; i++ {
			if less(seq, i, i-6) {
				swap(seq, i, i-6)
			}
		}
		insertionSort(seq, a, b)
	}
}

// maxDepth returns a threshold at which quicksort should switch
// to heapsort. It returns 2*ceil(lg(n+1)).
func maxDepth(n int) int {
	var depth int
	for i := n; i > 0; i >>= 1 {
		depth++
	}
	return depth * 2
}

func SortNumerics[N Numeric](numbers []N) {
	length := len(numbers)
	quickSort(numbers, 0, length, maxDepth(length))
}
