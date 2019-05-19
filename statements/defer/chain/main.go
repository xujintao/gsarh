package main

type Slice []int

func NewSlice() *Slice {
	return new(Slice)
}

func (s *Slice) Add(i int) *Slice {
	*s = append(*s, i)
	println(i)
	return s
}

func main() {
	s := NewSlice()
	defer s.Add(1).Add(2).Add(3)
	// defer s.Add(1).Add(2)
	// defer func() {
	// 	s.Add(1).Add(2).Add(3)
	// }()
	s.Add(4)
}
