package main

import (
)

type Rule struct {
  values []int
}

func NewRule(values... int) *Rule {
  r := Rule{}
  r.values = values
  return &r
}

func (r *Rule) True(nearby int) bool {
  for _, v := range r.values {
    if v == nearby {
      return true
    }
  }
  return false
}
