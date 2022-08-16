package lib

import "testing"

func TestMap(t *testing.T) {
	RunSubO(t, "clone", `
    var m = { a: 1, b: 2 }
    var c = map.clone(m)
    m.a = 10
    c.c = 30
    print(m)
    print(c)
  `, `
{a: 10, b: 2}
{a: 1, b: 2, c: 30}
`)
	RunSubO(t, "update", `
    var m = { a: 1, b: 2 }
    var r = map.update(m, { b: 20, c: 30 })
    print(m == r)
    print(m)
`, `
true
{a: 1, b: 20, c: 30}
`)
	RunSubO(t, "update_func", `
    var m = { a: 1, b: 2 }
    var r = map.update(m, &x -> { a: x.a + 1, c: x.a + x.b })
    print(m == r)
    print(m)
`, `
true
{a: 2, b: 2, c: 3}
`)
	RunSubO(t, "delete", `
    var m = { a: 1, b: 2, c: 3, d: 4 }
    var r = map.delete(m, "b", "d")
    print(m == r)
    print(m)
`, `
true
{a: 1, c: 3}
`)
}
