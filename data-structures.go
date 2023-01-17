package main

type Stack struct {
  value []Any
}

type Queue struct {
  value []Any
}

// --- //

func (s *Stack) Push(v Any) {
  s.value = append(s.value, v)
}

func (s *Stack) Pop() Any {
  ret := s.value[len(s.value) - 1]
  s.value = s.value[:len(s.value) - 1]
  
  return ret
}

func (s *Stack) Peek() Any {
  ret := s.value[len(s.value) - 1]
  
  return ret
}

func (s *Stack) Get(index int) Any {
  return s.value[index]
}

// --- //

func (q *Queue) Enqueue(v Any) {
  s.value = append(s.value, v)
}

func (q *Queue) Dequeue() Any {
  ret := s.value[len(s.value) - 1]
  s.value = s.value[:len(s.value) - 1]
  
  return ret
}

func (q *Queue) Poll() Any {
  ret := s.value[0]
  
  return ret
}

func (q *Queue) Get(index int) Any {
  return s.value[index]
}