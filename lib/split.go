package lib

func fnSplit(ctx context.Context, caps []nitro.ValueRef, args []nitro.Value, retN int) ([]nitro.Value, error) {
	if len(args) < 2 {
		return nil, errNotEnoughArgs
	}

	str, err := getStringArg(args, 0)
	if err != nil {
		return nil, err
	}

	n := -1
	if len(args) >= 3 {
		n, err = getIntArg(args, 2)
		if err != nil {
			return nil, err
		}
	}

	var parts []string
	switch sep := args[1].(type) {
	case nitro.String:
		parts = strings.SplitN(str, sep, n)

	case *nitro.Regex:
		parts = sep.Split(str, n)

	default:
		return nil, fmt.Errorf(
			"Invalid arg#2 type %v. Expected String or Regex.",
			nitro.TypeName(sep))
	}

	a := nitro.NewArrayWithCapacity(len(parts))
	for _, :w


}
