package parser

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParserProg(t *testing.T) {
	err := parse(`
	fn counter(n)
		var i = 0 
  	return fn() 
			if i < n then
				var ret = i
				i = i + 1
				return i, true
			else
				return 0, false
			end
		end
	end

	for x in counter(5) do
		print(x)
	end

	var less = fn(x,y)
		x < y
	end

  var someMult2 = [ 1, 2, 3 ].map(fn(x) x*2 end)

	{
		id: {
			customer: customer
			name: id
			region: region
		}

		video_encoders: [
			for enc in videoEncoders do
				{
					id: enc.id
					frame_rate: fixFrameRate(enc.frameRate)
					width: enc.width
					height: enc.height
					bit_rate: enc.bitRate
					codec: "H264"
					h264: {
						nal_hrd: "NAL_HRD_NONE",
						profile: "HIGH",
					}
				}
			end
		]

		publishPoints: {
			for pp in publishPoints do
				[pp.id] : {
					publishBaseUrl: pp.publishURL
					playbackBaseUrl: pp.playbackURL
					doNotMonitor: false
				}
			end
		}
	}
`)
	require.NoError(t, err)
}

func TestParserProg2(t *testing.T) {
	err := parse(`
		var a = 1
		fn Dwit(x)
			var a = 0
			fn()
				var t = a
				a = a + 1
				return t
			end
		end
`)
	require.NoError(t, err)
}

func TestParserLvalue(t *testing.T) {
	err := parse(`
	a.b.c = 123
	(a - 123).b = 999;;
	("hi").blah = 1
	;
`)
	require.NoError(t, err)
}

func TestObject(t *testing.T) {
	err := parse(`
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
		}.map("foo", 1234 + youtube(false))
`)
	require.NoError(t, err)
}

func TestParserArray(t *testing.T) {
	err := parse(`
	a = [ if foo then end ]
	a = [ if foo then 1, end ]
	a = [ if foo then 1,2 end ]
	a = [ if foo then 1,2,3 end ]
	a = [ 1, if foo then 2,3 end ]
	a = [ 1, if foo then 2,3 end, 4 ]
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

func parse(input string) error {
	_, err := Parse("test.nit", []byte(input))
	return err
}
