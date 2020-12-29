package ast

type Stack struct {
	stack []AST
}

func (s *Stack) Push(ast AST) {
	s.stack = append(s.stack, ast)
}

func (s *Stack) Pop() {
	s.stack = s.stack[:len(s.stack)-1]
}
