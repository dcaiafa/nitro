package context

import (
	"fmt"
	"io"
)

type Printer interface {
	Print(p *GraphPrinter)
}

type GraphPrinter struct {
	out   io.Writer
	next  int
	stack []int
}

func NewGraphPrinter(out io.Writer) *GraphPrinter {
	return &GraphPrinter{
		out: out,
	}
}

func (p *GraphPrinter) PrintGraph(root Printer) {
	p.push()
	fmt.Fprintf(p.out, "digraph G {\n")
	root.Print(p)
	fmt.Fprintf(p.out, "}\n")
}

func (p *GraphPrinter) PrintNode(label string, children ...Printer) {
	currentID := p.current()
	fmt.Fprintf(p.out, "  %d [label=%q];\n", currentID, label)
	for _, child := range children {
		childID := p.push()
		fmt.Fprintf(p.out, "  %d -> %d;\n", currentID, childID)
		child.Print(p)
		p.pop()
	}
}

func (p *GraphPrinter) push() int {
	p.next++
	id := p.next
	p.stack = append(p.stack, id)
	return id
}

func (p *GraphPrinter) pop() {
	p.stack = p.stack[:len(p.stack)-1]
}

func (p *GraphPrinter) current() int {
	return p.stack[len(p.stack)-1]
}
