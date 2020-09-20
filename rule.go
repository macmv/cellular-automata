package main

type Rule struct {

}

func (r *Rule) Eval(alive, nearby int) bool {
  return true
}
