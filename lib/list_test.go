package lib

import "testing"

func TestList(t *testing.T) {
	RunSubO(t, "append", `[1,2,3] | list.append(4) | print`, `[1 2 3 4]`)
	RunSubO(t, "append_iter", `[10] | list.append_iter(range(3)) | print`, `[10 0 1 2]`)
	RunSubO(t, "find_pos", `["a", 3, "c", 4] | list.find("c") | print`, `2`)
	RunSubO(t, "find_neg", `["a", 3, "c", 4] | list.find(5) | print`, `<nil>`)
}
