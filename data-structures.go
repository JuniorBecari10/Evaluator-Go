package main

type Stack struct {
  value []string
}

type Queue struct {
  value []string
}

// --- //

func NewStack(values []string) *Stack {
  return &Stack { values }
}

func (s *Stack) Push(v string) {
  s.value = append(s.value, v)
}

func (s *Stack) Pop() string {
  ret := s.value[len(s.value) - 1]
  s.value = s.value[:len(s.value) - 1]
  
  return ret
}

func (s *Stack) Peek() string {
  ret := s.value[len(s.value) - 1]
  
  return ret
}

func (s *Stack) Get(index int) string {
  return s.value[index]
}

func (s *Stack) Len() int {
  return len(s.value)
}

// --- //

func NewQueue(values []string) *Queue {
  return &Queue { values }
}

func (q *Queue) Enqueue(v string) {
  q.value = append(q.value, v)
}

func (q *Queue) Dequeue() string {
  ret := q.value[0]
  q.value = q.value[1:]
  
  return ret
}

func (q *Queue) Poll() string {
  ret := q.value[0]
  
  return ret
}

func (q *Queue) Get(index int) string {
  return q.value[index]
}

func (q *Queue) Len() int {
  return len(q.value)
}