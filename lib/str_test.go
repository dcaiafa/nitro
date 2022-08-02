package lib

import "testing"

func TestStrFind(t *testing.T) {
	RunSubO(t, "found", `"banana" | str.find("na") | print `, `2`)
	RunSubO(t, "not_found", `"banana" | str.find("no") | print `, `<nil>`)
}

func TestStrFindLast(t *testing.T) {
	RunSubO(t, "found", `"banana" | str.find_last("na") | print `, `4`)
	RunSubO(t, "not_found", `"banana" | str.find_last("no") | print `, `<nil>`)
}

func TestStrMatch(t *testing.T) {
	RunSubO(t, "match",
		"var re = r`(\\d+)/.*$`"+`
    "foo/123/bar" | str.match(re) | print`, `
[123/bar 123]
  `)
	RunSubO(t, "no_match",
		"var re = r`(\\d+)/.*$`"+`
    "foo/12a/bar" | str.match(re) | print`, `
<nil>
  `)
}

func TestStrMatchAll(t *testing.T) {
	RunSubO(t, "match",
		"var re = r`foo(.?)`"+`
    "seafood fool" | str.match_all(re) | print`, `
[[food d] [fool l]]
  `)
	RunSubO(t, "no_match",
		"var re = r`foo(.?)`"+`
    "seafeod fiol" | str.match_all(re) | print`, `
<nil>
  `)
}
