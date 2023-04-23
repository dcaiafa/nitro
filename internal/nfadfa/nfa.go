package nfadfa

import (
	"fmt"
	"io"
	"sort"

	"github.com/dcaiafa/nitro/internal/stack"
)

type epsilon struct {
}

func (e epsilon) String() string {
	return "ε"
}

// Epsilon is the ε (empty string) input.
var Epsilon = epsilon{}

// NFAState is a state in a NFA.
// NFAStates should only created using NFA.NewState().
type NFAState struct {
	// ID of the state.
	// It is assigned by the NFA and should be read-only.
	ID uint32

	// Transitions of the state.
	// Transitions should only be added using NFA.AddTransitions.
	Transitions map[any][]*NFAState

	// Accept indicates that the state machine should accept/recognize the input.
	// You set this yourself.
	Accept bool

	// Data is some user-data associated with this state.
	// Your set this yourself (or don't, I don't care).
	Data any
}

// NFA represents a Nondeterministic Finite Automaton.
// https://en.wikipedia.org/wiki/Nondeterministic_finite_automaton
// The NFA is state machine where:
// - A state is allowed to have multiple transitions for the same input.
// - The input ε (empty string) is allowed.
type NFA struct {
	Start  *NFAState
	States []*NFAState
}

// NewNFA creates a new NFA.
func NewNFA() *NFA {
	n := &NFA{}
	n.Start = n.NewState()
	return n
}

// AddTransition adds a transition between two states on a given input.
// The input can be any comparable data, including Epsilon.
func (n *NFA) AddTransition(from, to *NFAState, input any) *NFAState {
	if from.Transitions == nil {
		from.Transitions = make(map[any][]*NFAState)
	}
	from.Transitions[input] = append(from.Transitions[input], to)
	return to
}

// NewState creates a new state in the NFA.
func (n *NFA) NewState() *NFAState {
	s := &NFAState{
		ID: uint32(len(n.States)),
	}
	n.States = append(n.States, s)
	return s
}

// DFA converts the NFA into a DFA.
func (n *NFA) DFA() *DFA {
	// Create the DFA using the Subset Construction algorigthm.
	dfa := new(DFA)

	// DFA states already created, indexed by their signature:
	// the unique combination of NFA states that define a DFA state.
	states := make(map[string]*DFAState)

	// Create the start state.
	start := eClosure(map[*NFAState]bool{n.Start: true})
	dfa.Start = start
	dfa.States = append(dfa.States, start)
	states[start.signature()] = start

	// Starting from the newly create 'start' state, build all DFA states, one by
	// one.
	var stack stack.Stack[*DFAState]
	stack.Push(start)
	for !stack.Empty() {
		from := stack.Pop()

		for input := range getInputs(from.NFAStates) {
			// Create a subset of all NFA states reachable from the the NFA states
			// composing 'from', using 'input'.
			subset := map[*NFAState]bool{}
			for _, fromNFA := range from.NFAStates {
				for _, toNFA := range fromNFA.Transitions[input] {
					subset[toNFA] = true
				}
			}

			// Expand the subset to include states reachable via input ε.
			to := eClosure(subset)
			toSig := to.signature()

			// Look for an existing DFA state with the same set of NFA states
			// (uniquely represented by its signature).
			if existing := states[toSig]; existing != nil {
				// Reuse state.
				from.addTransition(existing, input)
			} else {
				// Create a new state.
				to.ID = uint32(len(dfa.States))
				dfa.States = append(dfa.States, to)
				states[toSig] = to
				from.addTransition(to, input)

				// Add it to the stack for processing.
				stack.Push(to)
			}
		}
	}

	return dfa
}

func (n *NFA) Print(out io.Writer) {
	fmt.Fprintf(out, "digraph G {\n")
	fmt.Fprintf(out, "  rankdir=\"LR\";\n")

	type Edge struct {
		from  *NFAState
		to    *NFAState
		input any
	}
	var edges []Edge
	var stack stack.Stack[*NFAState]
	stack.Push(n.Start)
	visited := make(map[*NFAState]bool)
	for !stack.Empty() {
		state := stack.Pop()
		if _, ok := visited[state]; ok {
			continue
		}
		visited[state] = true
		for input, destStates := range state.Transitions {
			for _, destState := range destStates {
				edges = append(edges, Edge{
					from:  state,
					to:    destState,
					input: input,
				})
				stack.Push(destState)
			}
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
		fmt.Fprintf(out, "  %v -> %v [label=%q];\n", e.from.ID, e.to.ID, inputStr)
	}

	for _, state := range n.States {
		shape := "circle"
		if state.Accept {
			shape = "doublecircle"
		}
		fmt.Fprintf(out, "  %v [label=\"%v\", shape=%q];\n", state.ID, state.ID, shape)
	}
	fmt.Fprintf(out, "}\n")
}

func eClosure(nfaStates map[*NFAState]bool) *DFAState {
	dfaState := new(DFAState)

	closure := make(map[uint32]*NFAState)
	var stack stack.Stack[*NFAState]
	for s := range nfaStates {
		closure[s.ID] = s
		stack.Push(s)
		dfaState.Accept = dfaState.Accept || s.Accept
	}
	for !stack.Empty() {
		state := stack.Pop()
		eTransitions := state.Transitions[Epsilon]
		for _, to := range eTransitions {
			if _, ok := closure[to.ID]; !ok {
				closure[to.ID] = to
				stack.Push(to)
				dfaState.Accept = dfaState.Accept || to.Accept
			}
		}
	}

	dfaState.NFAStates = make([]*NFAState, 0, len(closure))
	for _, nfaState := range closure {
		dfaState.NFAStates = append(dfaState.NFAStates, nfaState)
	}
	sort.Slice(dfaState.NFAStates, func(i, j int) bool {
		return dfaState.NFAStates[i].ID < dfaState.NFAStates[j].ID
	})

	return dfaState
}

func getInputs(nfaStates []*NFAState) map[any]bool {
	inputs := map[any]bool{}
	for _, nfaState := range nfaStates {
		for input := range nfaState.Transitions {
			if input != Epsilon {
				inputs[input] = true
			}
		}
	}
	return inputs
}
