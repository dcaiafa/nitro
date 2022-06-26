package mod

type Version struct {
	Major int
	Minor int
	Patch int
}

func (v Version) Compare(v2 Version) int {
	if v.Major == v2.Major {
		if v.Minor == v2.Minor {
			if v.Patch == v2.Patch {
				return 0
			} else if v.Patch < v2.Patch {
				return -1
			} else {
				return 1
			}
		} else if v.Minor < v2.Minor {
			return -1
		} else {
			return 1
		}

	} else if v.Major < v2.Major {
		return -1
	} else {
		return 1
	}
}
