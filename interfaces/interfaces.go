package interfaces

import (
	"fmt"
	"sort"
)

type Sequence []int

// Methods required by sort.Interface.
func (s Sequence) Len() int {
	return len(s)
}
func (s Sequence) Less(i, j int) bool {
	return s[i] < s[j]
}
func (s Sequence) Swap(i, j int) {
	if i == j {
		panic("Same index selected")
	}
	//s[i], s[j] = s[j], s[i]
	k := s[i]
	s[i] = s[j]
	s[j] = k
}

// Copy returns a copy of the Sequence.
func (s Sequence) Copy() Sequence {
	copy := make(Sequence, 0, len(s))
	return append(copy, s...)
}

// Method for printing - sorts the elements before printing.
func (s Sequence) String() string {
	s = s.Copy() // Make a copy; don't overwrite argument.
	sort.Sort(s)
	str := "["
	for i, elem := range s { // Loop is O(N²); will fix that in next example.
		if i > 0 {
			str += " "
		}
		str += fmt.Sprint(elem)
	}
	return str + "]"
}

func RunSeq() {
	seq := Sequence{
		1, 5, 6, 7, 8, 9,
	}
	fmt.Printf("%v \n", seq)
	seq.Swap(0, 1)
	seq.Swap(3, 5)
	fmt.Printf("%v \n", seq.String())
}
