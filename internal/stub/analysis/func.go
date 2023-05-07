package analysis

import (
	"fmt"
	"strings"

	"github.com/dcaiafa/nitro/internal/nfadfa"
)

type Func struct {
	Name       string
	Signatures []*Signature
}

func (s *Func) String() string {
	str := &strings.Builder{}
	for i, sig := range s.Signatures {
		if i != 0 {
			str.WriteString("\n")
		}
		str.WriteString(s.Name)
		str.WriteString(sig.String())
	}
	return str.String()
}

func (s *Func) GetSignatureIndex(v *Signature) int {
	for i, sig := range s.Signatures {
		if sig == v {
			return i
		}
	}
	panic("signature not found")
}

type FuncNFAData struct {
	Signature *Signature
	Param     Param
}

type FuncDFA struct {
	*nfadfa.DFA
	MinArgs int
}

func (s *Func) DFA() *FuncDFA {
	nfa := nfadfa.NewNFA()
	for _, sig := range s.Signatures {
		from := nfa.NewState()
		from.Data = &FuncNFAData{
			Signature: sig,
			Param:     Param{},
		}
		nfa.AddTransition(nfa.Start, from, nfadfa.Epsilon)
		for _, param := range sig.Params {
			if param.HasDefault {
				from.Accept = true
			}
			to := nfa.NewState()
			to.Data = &FuncNFAData{
				Signature: sig,
				Param:     param,
			}
			nfa.AddTransition(from, to, param.Type)
			from = to
		}
		from.Accept = true
	}

	var calcMinArgs func(min, cur int, s *nfadfa.DFAState) int
	calcMinArgs = func(min, cur int, s *nfadfa.DFAState) int {
		if cur > 0 {
			hasDefault := s.NFAStates[0].Data.(*FuncNFAData).Param.HasDefault
			for _, nfaState := range s.NFAStates {
				if nfaState.Data.(*FuncNFAData).Param.HasDefault != hasDefault {
					panic("default should be consistent in all nfa states")
				}
			}
			if hasDefault {
				return min
			}
		}
		if cur > min {
			min = cur
		}
		for _, toState := range s.Transitions {
			min = calcMinArgs(min, cur+1, toState)
		}
		return min
	}

	dfa := &FuncDFA{
		DFA: nfa.DFA(),
	}
	dfa.MinArgs = calcMinArgs(0, 0, dfa.Start)

	return dfa
}

type Signature struct {
	Params []Param
	Rets   []Type
}

func (s *Signature) String() string {
	paramStrs := make([]string, len(s.Params))
	for i, param := range s.Params {
		paramStrs[i] = fmt.Sprintf("%v %v", param.Name, param.Type.String())
	}
	sigStr := "(" + strings.Join(paramStrs, ", ") + ")"
	if len(s.Rets) == 1 {
		sigStr += " " + s.Rets[0].String()
	} else if len(s.Rets) > 1 {
		retStrs := make([]string, len(s.Rets))
		for i, ret := range s.Rets {
			retStrs[i] = ret.String()
		}
		sigStr += " (" + strings.Join(retStrs, ", ") + ")"
	}
	return sigStr
}

type Param struct {
	Name        string
	Type        Type
	HasDefault  bool
	DefaultExpr string
	VarArg      bool
}
