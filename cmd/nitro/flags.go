package main

type Flag struct {
	Name   string
	System bool
	Value  interface{}
}

type Flags struct {
	systemFlags []*Flag
	userFlags   []*Flag
	args        []string
}

func (f *Flags) AddFlag(flag *Flag) {
	if flag.System {
		f.systemFlags = append(f.userFlags, flag)
	} else {
		f.userFlags = append(f.userFlags, flag)
	}
}

/*
func (f *Flags) Parse(args []string) error {
	for len(args) > 0 {
		cur := args[0]
		if len(cur) >= 2 && cur[0] == '-' {
			flagExpr := cur[1:]
			isUser := cur[1] == '-'
			if isUser {
				flagExpr = cur[2:]
			}
		} else {
			f.args = append(f.args, cur)
			args = args[1:]
		}
	}
}
*/
