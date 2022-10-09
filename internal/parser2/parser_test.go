package parser2

import (
	"testing"

	"github.com/dcaiafa/nitro/internal/errlogger"
	"github.com/stretchr/testify/require"
)

const prog = `
!param disable_drm = false {type:"any", required, desc:"Disable DRM"}
!flag has_cc = true
!flag multi_line {
	name: "multi_line"
	type: "bool";
	desc: "Parameter attributes in multiple lines",
	required
}

import "foo/bar"
import yo "github.com/dcaiafa/nitro/yay"

lines() | split("\t")

a | b() | c()

a = a + 'c'

a++
b--

1 + 2 | b() | c()

if foo {
	print("stuff")
} else if other+x < 2 {
	print("yay")
} else if not nope {
} else {
	print("last resort")
}

for i in range(10) {
	print(i)
}
for i in range(10) { print(i) }
for i in range(10) { 
	;
}

print(1); print2;;

# This is cool.
print(in() | fromjson() | (&d -> d.id))

c(b(a)) # and so is this

try {
	json_lines() | 
		join(json_lines("nodes.json"), ".metadata.node", ".node_name") |
		select(&e->e.metadata.labels.team=="t2")
	
	throw "boom"
} catch e {
	log("shit's on fire yo " + e)
}

return

if cancel {
	return
}

return (
	foo), bar

select(
  join(
		json_lines(),
		json_lines("nodes.json"), ".metadata.node", ".node_name"),
	&e->e.metadata.labels.team=="t2")

	var preset = {
		video_encoders: [
			{
				id: "vid-60"
				width: 1920
				height: 1080
				frame_rate: "FR_60"
				codec: "H264"
			}
			{
				id: "vid-30"
				width: 1920
				height: 1080
				frame_rate: "FR_30"
				codec: "H264"
			}
		]
	}

	var config = {
  	transcoder: {
			preset()...

			audio_encoders: [
				{
        	id: "foo"
				}
			]
		}
	}

emit(
{
}
)

{
	foo: x ? y : bar
	bar: 1
}
`

func TestHelloWorld(t *testing.T) {
	unit, err := ParseUnit("test.nitro", prog, true, &errlogger.ConsoleErrLogger{})
	require.NoError(t, err)
	require.NotNil(t, unit)

	prologue, err := ParsePrologue("test.nitro", prog, true, &errlogger.ConsoleErrLogger{})
	require.NoError(t, err)
	require.NotNil(t, prologue)
	require.Len(t, prologue.Imports, 2)
}

func BenchmarkUnit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		unit, err := ParseUnit("test.nitro", prog, true, &errlogger.ConsoleErrLogger{})
		require.NoError(b, err)
		require.NotNil(b, unit)
	}
}

func BenchmarkPrologue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prologue, err := ParsePrologue("test.nitro", prog, true, &errlogger.ConsoleErrLogger{})
		require.NoError(b, err)
		require.NotNil(b, prologue)
	}
}
