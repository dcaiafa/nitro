package parser

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParser1(t *testing.T) {
	_, err := Parse(`
	{
		foo: "bar",
		x: 1 + 2,
	if 1 then
		y: "some stuff",
		z: "moar",
		["if"] : "then",
		[some_var.x] : "then",
	end
	if a + 123 != "that-makes-no-sense" then
		y: {
			take: "that",
			if dont_feel_sad then
				have: 1,
				cookie: "yes",
			elif is_there_plan_b then
				execute_order: "66",
			elif is_there_plan_c and soon > later then
				run: "now",
			else
				bwaa: 123
			end
		}
	end
		w: "one last thing"
	}
`)
	require.NoError(t, err)
}

func TestParser2(t *testing.T) {
	_, err := Parse(`
	a = {
		foo: "bar"
		bar: "foo"
		if 1 then
		end
		w: "one last thing"
	}
`)
	require.NoError(t, err)
}

func TestParserLvalue(t *testing.T) {
	_, err := Parse(`
	a.b.c = 123
	(a - 123).b = 999
	("hi").blah = 1
`)
	require.NoError(t, err)
}
