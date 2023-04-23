package nfadfa

import (
	"encoding/binary"
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/dcaiafa/nitro/internal/stack"
)

type DFAState struct {
	ID          uint32
	Transitions map[any]*DFAState
	Accept      bool
	NFAStates   []*NFAState
	Data        any
}

func (s *DFAState) addTransition(toState *DFAState, input any) {
	if s.Transitions == nil {
		s.Transitions = make(map[any]*DFAState)
	}
	s.Transitions[input] = toState
}

func (s *DFAState) signature() string {
	sig := make([]byte, len(s.NFAStates)*4)
	for i, nfaState := range s.NFAStates {
		binary.BigEndian.PutUint32(sig[i*4:], nfaState.ID)
	}
	return string(sig)
}

type DFA struct {
	Start  *DFAState
	States []*DFAState
}

func (d *DFA) Print(out io.Writer) {
	fmt.Fprintf(out, "digraph G {\n")
	fmt.Fprintf(out, "  rankdir=\"LR\";\n")

	type Edge struct {
		from  *DFAState
		to    *DFAState
		input any
	}

	var edges []Edge
	var stack stack.Stack[*DFAState]
	stack.Push(d.Start)
	visited := make(map[*DFAState]bool)
	for !stack.Empty() {
		state := stack.Pop()
		if _, ok := visited[state]; ok {
			continue
		}
		visited[state] = true
		for input, destState := range state.Transitions {
			edges = append(edges, Edge{
				from:  state,
				to:    destState,
				input: input,
			})
			stack.Push(destState)
		}
	}

	sort.SliceStable(edges, func(i, j int) bool {
		if edges[i].from.ID == edges[j].from.ID {
			return edges[i].to.ID < edges[j].to.ID
		} else {
			return edges[i].from.ID < edges[j].from.ID
		}
	})

	for _, e := range edges {
		inputStr := fmt.Sprintf("%v", e.input)
		fmt.Fprintf(out, "  %v -> %v [label=%q];\n",
			e.from.ID, e.to.ID, inputStr)
	}

	for _, state := range d.States {
		shape := "circle"
		if state.Accept {
			shape = "doublecircle"
		}
		nfaIDs := &strings.Builder{}
		for i, nfaState := range state.NFAStates {
			if i > 0 {
				nfaIDs.WriteString(" ")
			}
			fmt.Fprintf(nfaIDs, "%v", nfaState.ID)
		}
		fmt.Fprintf(out, "  %v [label=\"%v\\n%v\", shape=%q];\n",
			state.ID, state.ID, nfaIDs.String(), shape)
	}

	fmt.Fprintf(out, "}\n")
}
