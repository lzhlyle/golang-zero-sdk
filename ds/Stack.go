package ds

type IntStack struct {
	data []int
	top  int
}

func NewIntStack() *IntStack {
	return &IntStack{
		data: make([]int, 0),
		top:  0,
	}
}

func (s *IntStack) push(v int) {
	if s.top == len(s.data) {
		s.data = append(s.data, v)
	} else {
		s.data[s.top] = v
	}
	s.top++
}

func (s *IntStack) pop() int {
	res := s.data[s.top-1]
	s.top--
	return res
}

func (s *IntStack) peek() int {
	return s.data[s.top-1]
}

func (s *IntStack) size() int {
	return s.top
}
