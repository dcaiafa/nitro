package parser

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParser(t *testing.T) {
	parse(t, `
package "foo/bar"

type Time "github.com/dcaiafa/nitro/lib/time".Time

const LAYOUT Str = "01/02 03:04:05PM '06 -0700"

func utc(time Time) Time

func split(path Str) (Str, Str)

func min(a Any, b Any) Any

func replace(s Str, old Str, new Str) Str
func replace(s Str, old Regex, new Str) Str
func replace(s Str, old Regex, new Callable) Str

func foo(o any?) any

struct ParseOptions {
	Indent Str;
	Prefix Str;
}

`)

}

func parse(t *testing.T, input string) {
	unit, err := Parse("test.stub", []byte(input))
	require.NoError(t, err)
	fmt.Printf("%+v\n", unit)
}
