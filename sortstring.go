package sortstring

type SortString []byte

func NewSortString(s string) SortString {
	return SortString([]byte(s))
}

func (s SortString) Len () int { return len(s)}

func (s SortString) Swap (i, j int) { s[i],s[j] = s[j],s[i]}

func (s SortString) Less (i, j int) bool { return s[i] < s[j]}

func (s SortString) String () string { return string([]byte(s)) }

// Nearly STL reverse
func (s SortString) Reverse (beg, end int) {
	for beg != end {
		end--
		if beg == end {
			return
		}
		s.Swap(beg, end)
		beg++
	}
}

// Nearly STL next_permutation, but I don't have to play all the
// bidirectional iterator games due to my use of indexes
func (s SortString) NextPermutation (beg, end int) bool {
	if beg == end { return false }
	i := beg + 1
	if i == end { return false }
	i = end - 1

	for {
		ii := i
		i--
		if s[i] < s[ii] {
			j := end
			j--
			for !(s[i] < s[j]) { j-- }
			s.Swap(i, j)
			s.Reverse(ii, end)
			return true
		}
		if i == beg {
			s.Reverse(beg, end)
			return false
		}
	}
	return true // never happens, but 6g complains :-(
}

