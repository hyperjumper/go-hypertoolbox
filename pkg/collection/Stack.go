package collection


import "sync"

type Stack struct {
	lock  sync.Mutex
	stack []interface{}
}

func (s *Stack) PeekStack() []interface{} {
	s.lock.Lock()
	defer s.lock.Unlock()
	ret := make([]interface{}, 0)
	for _, v := range s.stack {
		ret = append(ret, v)
	}
	return ret
}

func (s *Stack) Clear() {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.stack = s.stack[0:0]
}

func (s *Stack) Size() int {
	return len(s.stack)
}

func (s *Stack) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Stack) Push(v interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.stack = append(s.stack, v)
}

func (s *Stack) Pop() interface{} {
	s.lock.Lock()
	defer s.lock.Unlock()
	if len(s.stack) == 0 {
		return nil
	}
	res := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return res
}

func (s *Stack) Peek() interface{} {
	s.lock.Lock()
	defer s.lock.Unlock()
	if len(s.stack) == 0 {
		return nil
	}
	return s.stack[len(s.stack)-1]
}