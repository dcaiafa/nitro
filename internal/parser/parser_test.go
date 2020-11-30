package parser

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParserProg(t *testing.T) {
	_, err := Parse(`
	fn counter(n)
		var i = 0 
  	fn() i = i + 1; i; end
	end
`)
	require.NoError(t, err)
}

func TestParserLvalue(t *testing.T) {
	_, err := Parse(`
	a.b.c = 123
	(a - 123).b = 999;;
	("hi").blah = 1
	;
`)
	require.NoError(t, err)
}

func TestObject(t *testing.T) {
	_, err := Parse(`
		{ 
			a: "foo", b: "bar";

		if ans == 42 then
			fade_away: "let it go"
			moar: like + "stone"
		elif plan_b or plan_c then
			next_best_thing: 123
		else
			last_recourse: false
		end

		if x then x: 1, y: 2 end,
		if x then x: 1, y: 2 end
		if x then x: 1, y: 2, end;
		if x then x: 1 end
		if x then x: 1, end
		if x then x: 1; end

			b: "nope"
		} | map("foo", 1234 + youtube(false))
`)
	require.NoError(t, err)
}

/*
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

	a = [ { name: "alice" }, { name: "bob" } ]
`)
	require.NoError(t, err)
}
*/
