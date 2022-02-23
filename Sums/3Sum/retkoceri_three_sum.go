package sum

import . "github.com/Jcowwell/go-algorithm-club/Utils"

func insert[E any](seq []E, index int, value E) []E {
	if len(seq) == index {
		return append(seq, value)
	}
	seq = append(seq[:index+1], seq[index:]...)
	seq[index] = value
	return seq
}

func prep[N Numeric](nums []N) ([]N, []N, map[N]int, map[N]int) {
	positives, negatives, HTP, HTN := []N{}, []N{}, make(map[N]int), make(map[N]int)
	for index, num := range nums {
		if num > 0 {
			positives = append(positives, num)
			HTP[num] = index
		} else {
			negatives = insert(negatives, 0, num)
			HTN[num] = index

		}
	}
	return positives, negatives, HTP, HTN
}

// Based off Muhamed Retkoceri's Efficient Algorithm for solving 3SUM problem
// https://knowledgecenter.ubt-uni.net/cgi/viewcontent.cgi?article=1137&context=conference
func RetkoceriThreeSum[N Numeric](nums []N) (result [][]N) {
	SortNumerics(nums)
	P, NE, HTP, HTN := prep(nums)
	n := len(nums) // size of nums
	p := len(P)    // number of positives
	for i := 0; i < n-p; i++ {
		a := NE[i]
		for j := 0; j < p; j++ {
			b := P[j]
			c := -(a + b)
			if c < b {
				break
			}
			if _, ok := HTP[c]; ok {
				if b == c {
					break
				} else {
					result = append(result, []N{a, b, c})
				}
			}
		}
	}
	for i := 0; i < p; i++ {
		a := P[i]
		for j := 0; j < n-p; j++ {
			b := NE[j]
			c := -(a + b)
			if c > b {
				break
			}
			if val, ok := HTN[c]; ok {
				if b == c && j == val {
					break
				} else {
					result = append(result, []N{a, b, c})
				}
			}
		}
	}
	return
}
