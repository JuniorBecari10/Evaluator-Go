package main

import "fmt"

type Any interface {}

type DataStruct interface {
  func Get(index int) Any
  func Len() int
  
}

type Stack struct {
  value []Any
}

type Queue struct {
  value []Any
}

// --- //

func NewStack(values []Any) *Stack {
  return &Stack { values }
}

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

func (s *Stack) Len() int {
  return len(s.value)
}

// --- //

func NewQueue(values []Any) *Queue {
  return &Queue { values }
}

func (q *Queue) Enqueue(v Any) {
  q.value = append(q.value, v)
}

func (q *Queue) Dequeue() Any {
  ret := q.value[0]
  q.value = q.value[1:]
  
  return ret
}

func (q *Queue) Poll() Any {
  ret := q.value[0]
  
  return ret
}

func (q *Queue) Get(index int) Any {
  return q.value[index]
}

func (q *Queue) Len() int {
  return len(q.value)
}