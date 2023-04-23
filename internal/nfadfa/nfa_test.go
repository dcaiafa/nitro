package nfadfa

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func simpleNFA() *NFA {
	n := NewNFA()

	s1 := n.NewState()
	s2 := n.NewState()
	s3 := n.NewState()
	s4 := n.NewState()
	s5 := n.NewState()
	s6 := n.NewState()
	s7 := n.NewState()
	s8 := n.NewState()
	s9 := n.NewState()
	s10 := n.NewState()

	n.AddTransition(n.Start, s1, Epsilon)
	n.AddTransition(s1, s2, Epsilon)
	n.AddTransition(s2, s3, "a")
	n.AddTransition(s3, s6, Epsilon)
	n.AddTransition(s1, s4, Epsilon)
	n.AddTransition(s4, s5, "b")
	n.AddTransition(s5, s6, Epsilon)
	n.AddTransition(s6, s7, Epsilon)
	n.AddTransition(s7, s8, "a")
	n.AddTransition(s8, s9, "b")
	n.AddTransition(s9, s10, "b")
	n.AddTransition(n.Start, s7, Epsilon)
	s10.Accept = true

	return n
}

func TestNFA_Print(t *testing.T) {
	n := simpleNFA()

	str := &strings.Builder{}
	n.Print(str)
	require.Equal(t, strings.TrimSpace(str.String()), strings.TrimSpace(`
digraph G {
  rankdir="LR";
  0 -> 1 [label="ε"];
  0 -> 7 [label="ε"];
  1 -> 2 [label="ε"];
  1 -> 4 [label="ε"];
  2 -> 3 [label="a"];
  3 -> 6 [label="ε"];
  4 -> 5 [label="b"];
  5 -> 6 [label="ε"];
  6 -> 7 [label="ε"];
  7 -> 8 [label="a"];
  8 -> 9 [label="b"];
  9 -> 10 [label="b"];
  0 [label="0", shape="circle"];
  1 [label="1", shape="circle"];
  2 [label="2", shape="circle"];
  3 [label="3", shape="circle"];
  4 [label="4", shape="circle"];
  5 [label="5", shape="circle"];
  6 [label="6", shape="circle"];
  7 [label="7", shape="circle"];
  8 [label="8", shape="circle"];
  9 [label="9", shape="circle"];
  10 [label="10", shape="doublecircle"];
}
`))
}

func TestEClosure(t *testing.T) {
	n := simpleNFA()

	ids := func(nfaStates []*NFAState) []uint32 {
		ids := make([]uint32, len(nfaStates))
		for i, s := range nfaStates {
			ids[i] = s.ID
		}
		return ids
	}

	c := eClosure(map[*NFAState]bool{n.Start: true})
	require.Equal(t, []uint32{0, 1, 2, 4, 7}, ids(c.NFAStates))
}

func TestDFA(t *testing.T) {
	n := simpleNFA()

	d := n.DFA()
	d.Print(os.Stdout)
}
