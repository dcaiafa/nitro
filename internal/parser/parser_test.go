package parser

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParserLvalue(t *testing.T) {
	_, err := Parse(`
	a.b.c = 123
	(a - 123).b = 999
	("hi").blah = 1
`)
	require.NoError(t, err)
}

func TestObject(t *testing.T) {
	_, err := Parse(`
	{
		foo: "bar"
		x: 1 + 2

	if 1 then
		y: "some stuff"
		z: "moar"
		["if"] : "then"
		[some_var.x] : "then"
	end

	if a + 123 != "that-makes-no-sense" then
		y: {
			take: "that"
			if dont_feel_sad then
				have: 1
				cookie: "yes"
			elif is_there_plan_b then
				execute_order: "66"
			elif is_there_plan_c and soon > later then
				run: "now"
			else
				bwaa: 123
			end
		}
	end

		w: "one last thing"
	}

	{}
	{ a: 1 }
	{ a: 1, }
	{ a: 1 b: 2 }
	{ a: 1, b: 2 }
	{ a: 1, b: 2, }
`)
	require.NoError(t, err)
}

func TestParserArray(t *testing.T) {
	_, err := Parse(`
	a = [ if foo then end ]
	a = [ if foo then 1, end ]
	a = [ if foo then 1,2 end ]
	a = [ if foo then 1,2,3 end ]
	a = [ 1, if foo then 2,3 end ]
	a = [ 1, if foo then 2,3 end 4 ]
	a = [ 1, if foo then 2,3 end 4, ]
	b = [
  	1,
	if foo then
		2, 3,
	elif bar then
		4,
	elif plan_b or plan_c then
		4,
	else
		5,
	end
		999,
	]
`)
	require.NoError(t, err)
}
