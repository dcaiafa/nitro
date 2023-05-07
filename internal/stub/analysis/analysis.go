package analysis

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"regexp"
	"sort"
	"strings"

	"github.com/dcaiafa/nitro/internal/nfadfa"
	"github.com/dcaiafa/nitro/internal/util/set"
)

const vmPackage = "github.com/dcaiafa/nitro/internal/vm"
const stubPackage = "github.com/dcaiafa/nitro/internal/stub"
const exportPackage = "github.com/dcaiafa/nitro/internal/export"

var valueGoType = GoType{Package: vmPackage, Name: "Value"}
var readerGoType = GoType{Package: vmPackage, Name: "Reader"}
var readableGoType = GoType{Package: vmPackage, Name: "Readable"}
var iterableGoType = GoType{Package: vmPackage, Name: "Iterable"}
var iteratorGoType = GoType{Package: vmPackage, Name: "Iterator"}
var strGoType = GoType{Package: vmPackage, Name: "String"}
var mapGoType = GoType{Package: vmPackage, Name: "Object", Ref: true}

type Analysis struct {
	pkg          string
	goPkg        string
	types        map[string]Type
	funcs        map[string]*Func
	structs      map[string]*Struct
	addedImports map[ /*alias*/ string] /*import*/ string
	imports      []string
	importMap    map[ /*package*/ string] /*alias*/ string
}

func NewAnalysis() *Analysis {
	a := &Analysis{
		types:        make(map[string]Type),
		funcs:        make(map[string]*Func),
		structs:      make(map[string]*Struct),
		addedImports: make(map[string]string),
	}

	types := []struct {
		Name   string
		GoName string
		Ref    bool
	}{
		{"Any", "Value", false},
		{"Bool", "Bool", false},
		{"Callable", "Callable", false},
		{"Float", "Float", false},
		{"Int", "Int", false},
		{"Iter", "Iterator", false},
		{"List", "Array", true},
		{"Map", "Object", true},
		{"Reader", "Reader", false},
		{"Regex", "Regex", true},
		{"Str", "String", false},
		{"Writer", "Writer", false},
	}

	for _, typ := range types {
		err := a.AddGoType(Type{
			Name: typ.Name,
			GoType: GoType{
				Ref:     typ.Ref,
				Package: vmPackage,
				Name:    typ.GoName,
			},
		})
		if err != nil {
			panic(err)
		}
	}

	return a
}

func (a *Analysis) SetPackage(name string) {
	a.pkg = name
	a.goPkg = path.Base(a.pkg)
}

func (a *Analysis) AddImport(alias string, imp string) error {
	if _, ok := a.addedImports[alias]; ok {
		return fmt.Errorf("there is already an import with alias %q", alias)
	}
	a.addedImports[alias] = imp
	return nil
}

func (a *Analysis) AddGoType(typ Type) error {
	if _, ok := a.types[typ.Name]; ok {
		return fmt.Errorf("there is already a type %q", typ.Name)
	}
	a.types[typ.Name] = typ
	return nil
}

func (a *Analysis) GetGoType(name string) (Type, bool) {
	typ, ok := a.types[name]
	return typ, ok
}

func (a *Analysis) AddFunc(fun *Func) error {
	existing := a.funcs[fun.Name]
	if existing != nil {
		existing.Signatures = append(existing.Signatures, fun.Signatures...)
		return nil
	}
	a.funcs[fun.Name] = fun
	return nil
}

func (a *Analysis) AddStruct(st *Struct) error {
	goType := Type{
		Name: st.Name,
		GoType: GoType{
			Name:   st.Name,
			Ref:    true,
			Struct: true,
		},
	}
	err := a.AddGoType(goType)
	if err != nil {
		return err
	}

	a.structs[st.Name] = st
	return nil
}

func (a *Analysis) Emit(w *bytes.Buffer) {
	fmt.Fprintf(w, "package %v\n", a.goPkg)

	importSet := set.Set[ /*package*/ string]{}
	importSet.Add(vmPackage)
	importSet.Add(stubPackage)
	importSet.Add(exportPackage)

	for _, typ := range a.types {
		if typ.GoType.Package != "" {
			importSet.Add(typ.GoType.Package)
		}
	}
	for _, imp := range a.addedImports {
		importSet.Add(imp)
	}

	a.imports = importSet.ToSlice()
	sort.Strings(a.imports)
	a.importMap = make(map[ /*package*/ string] /*alias*/ string, len(a.imports))
	for i, imp := range a.imports {
		alias := fmt.Sprintf("_p%d", i)
		a.importMap[imp] = alias
		fmt.Fprintf(w, "import %v %q\n", alias, imp)
	}

	structNames := make([]string, 0, len(a.structs))
	for structName := range a.structs {
		structNames = append(structNames, structName)
	}
	sort.Strings(structNames)

	for _, structName := range structNames {
		st := a.structs[structName]
		a.emitStruct(w, st)
	}

	funcNames := make([]string, 0, len(a.funcs))
	for funcName := range a.funcs {
		funcNames = append(funcNames, funcName)
	}
	sort.Strings(funcNames)

	for _, funcName := range funcNames {
		fn := a.funcs[funcName]
		a.emitFunc(w, fn)
	}

	exportAlias := a.importMap[exportPackage]
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "var Exports = %s.Exports{\n", exportAlias)
	for _, funcName := range funcNames {
		fmt.Fprintf(w, "  {N: %q, T: %s.Func, F: _%s},\n", funcName, exportAlias, funcName)
	}
	fmt.Fprintf(w, "}\n")
}

func (a *Analysis) emitStruct(w *bytes.Buffer, st *Struct) {
	stubAlias := a.importMap[stubPackage]

	stType := a.types[st.Name]
	fmt.Fprintf(w, "type %s struct {\n", stType.GoType.Name)
	for _, field := range st.Fields {
		fmt.Fprintf(w, "  %s %s\n",
			a.fieldGoName(field.Name),
			a.goType(a.getConvenienceType(field.Type)))
	}
	fmt.Fprintf(w, "}\n")
	fmt.Fprintf(w, "func (m *%s) FromMap(v %s) error {\n",
		stType.GoType.Name, a.goType(
			GoType{
				Package: vmPackage,
				Name:    "Object",
				Ref:     true,
			},
		))
	fmt.Fprintf(w, "var err error\n")
	fmt.Fprintf(w, "_ = err\n")
	fmt.Fprintf(w, "v.ForEach(func(k, v %s) bool {\n", a.goType(valueGoType))
	fmt.Fprintf(w, "n, ok := k.(%v)\n", a.goType(strGoType))
	fmt.Fprintf(w, "if !ok {\n")
	fmt.Fprintf(w, "  err = %s.ErrMapKeyMustBeStr\n", stubAlias)
	fmt.Fprintf(w, "  return false\n")
	fmt.Fprintf(w, "}\n")
	fmt.Fprintf(w, "switch n.String() {\n")
	for _, field := range st.Fields {
		fmt.Fprintf(w, "case %q:\n", field.Name)
		fmt.Fprintf(w, "cv, ok := v.(%v)\n", a.goType(field.Type.GoType))
		fmt.Fprintf(w, "if !ok {\n")
		fmt.Fprintf(w, "  err = %s.ErrInvalidFieldType\n", stubAlias)
		fmt.Fprintf(w, "  return false\n")
		fmt.Fprintf(w, "}\n")
		if field.Type.GoType.Struct {
			panic("notimpl")
		} else {
			fmt.Fprintf(w, "tv := %v\n",
				a.toConvenienceType("cv", field.Type))
		}
		fmt.Fprintf(w, "m.%s = tv\n", a.fieldGoName(field.Name))
	}
	fmt.Fprintf(w, "default:\n")
	fmt.Fprintf(w, "  err = %s.StructDoesNotHaveField(%q, n.String())\n",
		stubAlias, st.Name)
	fmt.Fprintf(w, "  return false\n")
	fmt.Fprintf(w, "}\n")
	fmt.Fprintf(w, "return true\n")
	fmt.Fprintf(w, "})\n")
	fmt.Fprintf(w, "if err != nil {\n")
	fmt.Fprintf(w, " return err\n")
	fmt.Fprintf(w, "}\n")
	fmt.Fprintf(w, "return nil\n")
	fmt.Fprintf(w, "}\n")
}

func (a *Analysis) Dump() {
	fmt.Println("Types")
	fmt.Println("=====")
	for name, typ := range a.types {
		fmt.Println("  ", name, typ.GoType)
	}
	fmt.Println("Funcs")
	fmt.Println("=====")
	for _, fun := range a.funcs {
		fmt.Println(fun.String())
		dfa := fun.DFA()
		dfa.Print(os.Stdout)
	}
}

func (a *Analysis) emitFunc(w *bytes.Buffer, fn *Func) {
	vmAlias := a.importMap[vmPackage]
	stubAlias := a.importMap[stubPackage]

	dfa := fn.DFA()
	fmt.Fprintf(w, "func _%s(vm *%s.VM, args []%s.Value, nret int)([]%s.Value, error) {\n",
		fn.Name, vmAlias, vmAlias, vmAlias)

	fmt.Fprintf(w, "var err error\n")
	fmt.Fprintf(w, "_ = err\n")

	if dfa.MinArgs > 0 {
		fmt.Fprintf(w, "if len(args) < %d {\n", dfa.MinArgs)
		fmt.Fprintf(w, "  return nil, %v.ErrInsufficientArgs\n", stubAlias)
		fmt.Fprintf(w, "}\n")
	}

	var emitDFA func(argIdx int, fromState *nfadfa.DFAState)
	emitDFA = func(argIdx int, fromState *nfadfa.DFAState) {
		if fromState.Accept {
			isCond := len(fromState.Transitions) > 0
			isVarArg := fromState.NFAStates[0].Data != nil && fromState.NFAStates[0].Data.(*FuncNFAData).Param.VarArg
			if isCond {
				fmt.Fprintf(w, "if len(args) == %d {\n", argIdx)
			} else if !isVarArg {
				fmt.Fprintf(w, "if len(args) > %d {\n", argIdx)
				fmt.Fprintf(w, "  return nil, %s.ErrTooManyArgs\n", stubAlias)
				fmt.Fprintf(w, "}\n")
			}

			// N.B. if the DFA has a single state, it will have multiple NFA states.
			var sig *Signature
			for _, nfaState := range fromState.NFAStates {
				if nfaState.Data != nil {
					sig = nfaState.Data.(*FuncNFAData).Signature
					break
				}
			}

			for i := argIdx; i < len(sig.Params); i++ {
				param := &sig.Params[i]
				if !param.HasDefault {
					panic("remaining parameters must have defaults")
				}
				def, err := a.expr(param.DefaultExpr)
				if err != nil {
					// We should check expressions before emit.
					panic(err)
				}
				fmt.Fprintf(w, "var _a%d %s = %v\n",
					i, a.goTypeVM(param.Type.GoType),
					a.fromConvenienceType(def, param.Type))
			}

			// Function invocation.

			fmt.Fprintf(w, "{\n")
			for i, param := range sig.Params {
				if st := a.structs[param.Type.Name]; st != nil {
					fmt.Fprintf(w, "var _ta%d %s\n", i, a.goType(param.Type.GoType))
					fmt.Fprintf(w, "if _a%d != nil {\n", i)
					fmt.Fprintf(w, "_ta%d = new(%v)\n", i, a.goType(param.Type.GoType.NoRef()))
					fmt.Fprintf(w, "err = _ta%d.FromMap(_a%d)\n", i, i)
					fmt.Fprintf(w, "if err != nil {\n")
					fmt.Fprintf(w, "  return nil, %s.InvalidArgErr(args, %d, err)\n",
						stubAlias, argIdx)
					fmt.Fprintf(w, "}\n")
					fmt.Fprintf(w, "}\n")
				} else if param.VarArg && param.Type.Name != "Any" {
					fmt.Fprintf(w, "_ta%d := make([]%s, len(_a%d))\n", i, a.goType(a.getConvenienceType(param.Type)), i)
					fmt.Fprintf(w, "for i := range _a%d {\n", i)
					fmt.Fprintf(w, "  switch _a := _a%d[i].(type) {\n", i)
					a.emitTypeCase(w, param.Type)
					fmt.Fprintf(w, "  _ta%d[i] = %s\n", i, a.toConvenienceType("_a", param.Type))
					fmt.Fprintf(w, "default:\n")
					fmt.Fprintf(w, "  return nil, %s.InvalidArg(args, %d)\n", stubAlias, i)
					fmt.Fprintf(w, "}\n")
					fmt.Fprintf(w, "}\n")
				} else {
					fmt.Fprintf(w, "_ta%d := %v\n",
						i, a.toConvenienceType(fmt.Sprintf("_a%d", i), param.Type))
				}
			}
			for i := range sig.Rets {
				if i != 0 {
					fmt.Fprintf(w, ", ")
				}
				retVar := fmt.Sprintf("_r%d", i)
				fmt.Fprint(w, retVar)
			}
			if len(sig.Rets) > 0 {
				fmt.Fprintf(w, ", ")
			}
			fmt.Fprintf(w, "err := ")

			fmt.Fprintf(w, "%s%d(vm", fn.Name, fn.GetSignatureIndex(sig))
			for i := range sig.Params {
				fmt.Fprintf(w, ", _ta%d", i)
			}
			fmt.Fprintf(w, ")\n")
			fmt.Fprintf(w, "if err != nil {\n")
			fmt.Fprintf(w, "  return nil, err\n")
			fmt.Fprintf(w, "}\n")
			fmt.Fprintf(w, "return []%s.Value{", vmAlias)
			for i, ret := range sig.Rets {
				if i != 0 {
					fmt.Fprintf(w, ", ")
				}
				retVar := fmt.Sprintf("_r%d", i)
				fmt.Fprint(w, a.fromConvenienceType(retVar, ret))
			}
			fmt.Fprintf(w, "}, nil\n")
			fmt.Fprintf(w, "}\n")

			if isCond {
				fmt.Fprintf(w, "}\n")
			}
		}
		if len(fromState.Transitions) != 0 {
			type transition struct {
				Type    Type
				ToState *nfadfa.DFAState
			}
			transitions := make([]transition, 0, len(fromState.Transitions))
			for typ, toState := range fromState.Transitions {
				t := transition{
					Type:    typ.(Type),
					ToState: toState,
				}
				transitions = append(transitions, t)
			}
			sort.Slice(transitions, func(i, j int) bool {
				return transitions[i].Type.Name < transitions[j].Type.Name
			})

			varArgParam := transitions[0].ToState.NFAStates[0].Data.(*FuncNFAData).Param
			isVarArg := varArgParam.VarArg
			if isVarArg {
				if len(transitions) != 1 {
					panic("there should be only one transition for vararg param")
				}
				fmt.Fprintf(w, "var _a%d []%s.Value = args[%d:]\n", argIdx, vmAlias, argIdx)
				emitDFA(argIdx+1, transitions[0].ToState)
			} else {
				fmt.Fprintf(w, "switch _a%d := args[%d].(type) {\n", argIdx, argIdx)
				for _, t := range transitions {
					a.emitTypeCase(w, t.Type)
					emitDFA(argIdx+1, t.ToState)
				}
				fmt.Fprintf(w, "default:\n")
				fmt.Fprintf(w, "  return nil, %s.InvalidArg(args, %d)\n",
					stubAlias, argIdx)
				fmt.Fprintf(w, "}\n")
			}
		}
	}

	emitDFA(0, dfa.Start)
	fmt.Fprintf(w, "}\n")
}

func (a *Analysis) emitTypeCase(w *bytes.Buffer, typ Type) {
	switch typ.Name {
	case "Reader":
		fmt.Fprintf(w, "case %s,%s:\n",
			a.goType(readerGoType),
			a.goType(readableGoType),
		)
	case "Iter":
		fmt.Fprintf(w, "case %s,%s:\n",
			a.goType(iterableGoType),
			a.goType(iteratorGoType),
		)
	default:
		fmt.Fprintf(w, "case %s:\n", a.goTypeVM(typ.GoType))
	}
}

func (a *Analysis) getConvenienceType(typ Type) GoType {
	switch typ.Name {
	case "Str":
		return GoType{Name: "string"}
	case "Int":
		return GoType{Name: "int64"}
	case "Float":
		return GoType{Name: "float64"}
	case "Bool":
		return GoType{Name: "bool"}
	default:
		return typ.GoType
	}
}

func (a *Analysis) toConvenienceType(v string, typ Type) string {
	switch typ.Name {
	case "Str":
		return fmt.Sprintf("(%v).String()", v)
	case "Int":
		return fmt.Sprintf("(%v).Int64()", v)
	case "Float":
		return fmt.Sprintf("(%v).Float64()", v)
	case "Bool":
		return fmt.Sprintf("(%v).Bool()", v)
	case "Reader":
		return fmt.Sprintf("%s.MustMakeReader(vm, %v)", a.importMap[stubPackage], v)
	case "Iter":
		return fmt.Sprintf("%s.MustMakeIter(vm, %v)", a.importMap[stubPackage], v)
	default:
		return v
	}
}

func (a *Analysis) fromConvenienceType(v string, typ Type) string {
	alias := a.importMap["github.com/dcaiafa/nitro/internal/vm"]
	switch typ.Name {
	case "Str":
		return fmt.Sprintf("%s.NewString(%v)", alias, v)
	case "Int":
		return fmt.Sprintf("%s.NewInt(%v)", alias, v)
	case "Float":
		return fmt.Sprintf("%s.NewFloat(%v)", alias, v)
	case "Bool":
		return fmt.Sprintf("%s.NewBool(%v)", alias, v)
	default:
		return v
	}
}

func (a *Analysis) goType(goType GoType) string {
	ref := ""
	if goType.Ref {
		ref = "*"
	}
	importPrefix := ""
	if goType.Package != "" {
		importPrefix = a.importMap[goType.Package] + "."
	}

	return ref + importPrefix + goType.Name
}

func (a *Analysis) goTypeVM(goType GoType) string {
	if goType.Struct {
		return a.goType(mapGoType)
	}
	return a.goType(goType)
}

func (a *Analysis) vmAlias() string {
	return a.importMap[vmPackage]
}

func (a *Analysis) fieldGoName(n string) string {
	// TODO: handle snake.
	return strings.Title(n)
}

var aliasRegex = regexp.MustCompile(`#([a-z]\w*)\.`)

func (a *Analysis) expr(expr string) (string, error) {
	var err error
	expr = aliasRegex.ReplaceAllStringFunc(expr, func(s string) string {
		alias := s[1 : len(s)-1]
		imp, ok := a.addedImports[alias]
		if !ok {
			err = fmt.Errorf("there is no import with alias %q", alias)
			return s
		}
		actualAlias := a.importMap[imp]
		if actualAlias == "" {
			panic("added import not in importMap")
		}
		return actualAlias + "."
	})
	if err != nil {
		return "", err
	}
	return expr, nil
}
