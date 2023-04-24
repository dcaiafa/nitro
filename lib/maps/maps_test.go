package maps_test

import (
	"testing"

	"github.com/dcaiafa/nitro/internal/btesting"
)

func TestMaps(t *testing.T) {
	btesting.RunSubO(t, "clone", `
    var m = { a: 1, b: 2 }
    var c = maps.clone(m)
    m.a = 10
    c.c = 30
    print(m)
    print(c)
  `, `
{a: 10, b: 2}
{a: 1, b: 2, c: 30}
`)
	btesting.RunSubO(t, "update", `
    var m = { a: 1, b: 2 }
    var r = maps.update(m, { b: 20, c: 30 })
    print(m == r)
    print(m)
`, `
true
{a: 1, b: 20, c: 30}
`)
	btesting.RunSubO(t, "update_func", `
    var m = { a: 1, b: 2 }
    var r = maps.update(m, &x -> { a: x.a + 1, c: x.a + x.b })
    print(m == r)
    print(m)
`, `
true
{a: 2, b: 2, c: 3}
`)
	btesting.RunSubO(t, "delete", `
    var m = { a: 1, b: 2, c: 3, d: 4 }
    var r = maps.delete(m, "b")
    print(m == r)
    print(m)
`, `
true
{a: 1, c: 3, d: 4}
`)
}
