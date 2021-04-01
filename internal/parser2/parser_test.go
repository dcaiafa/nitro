package parser2

import (
	"testing"

	"github.com/dcaiafa/nitro/internal/errlogger"
	"github.com/stretchr/testify/require"
)

func TestHelloWorld(t *testing.T) {
	prog := `
lines() | split("\t")

a | b() | c()

1 + 2 | b() | c()

print(in() | fromjson() | &d -> d.id)

c(b(a))

try
	json_lines() | 
		join(json_lines("nodes.json"), ".metadata.node", ".node_name") |
		select(&e->e.metadata.labels.team=="t2")
	
	throw "boom"
catch e
	log("shit's on fire yo " + e)
end

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
		`

	module, err := Parse("test.nitro", prog, true, &errlogger.ConsoleErrLogger{})
	require.NoError(t, err)
	require.NotNil(t, module)
}
