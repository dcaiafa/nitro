package main

type Flag struct {
	Name     string
	System   bool
	Provided bool
	Value    interface{}
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
	var err error
	for len(args) > 0 {
		cur := args[0]
		if len(cur) >= 2 && cur[0] == '-' {
			isUserFlag := cur[1] == '-'
			if isUserFlag {
				args, err = f.parseUserFlag(args)
			} else {
				args, err = f.parseSysFlag(args)
			}
			if err != nil {
				return err
			}
		} else {
			f.args = append(f.args, cur)
			args = args[1:]
		}
	}
	return nil
}

func (f *Flags) parseUserFlag(args []string) ([]string, error) {
	// Skip the '--'.
	flagName := args[0][2:]

	flagValue := ""
	idxEq := strings.IndexByte(flagName, '=')
	if idxEq != -1 {
		flagName = flagName[:idxEq]
		flagValue = flagName[idxEq+1:]
	}
	args = args[1:]

	for _, flag := range f.userFlags {
		if flag.Name == flagName {
			if idxEq == -1 {
				if _, ok := flag.Value.(*bool); !ok {
					if len(args) < 1 {
						return nil, fmt.Errorf("flag --%v requires a value", flag.Name)
					}
					flagValue = args[0]
					args = args[1:]
				}
			}
		}
	}

}
*/
