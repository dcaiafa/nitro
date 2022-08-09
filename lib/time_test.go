package lib

import (
	"testing"

	"github.com/dcaiafa/nitro/internal/vm"
)

func TestTimeDuration(t *testing.T) {
	RunSubO(t, "nanosecond", `10 * time.nanosecond | print`, `10ns`)
	RunSubO(t, "microsecond", `10 * time.microsecond | print`, `10µs`)
	RunSubO(t, "millisecond", `10 * time.millisecond | print`, `10ms`)
	RunSubO(t, "second", `10 * time.second | print`, `10s`)
	RunSubO(t, "minute", `10 * time.minute | print`, `10m0s`)
	RunSubO(t, "hour", `10 * time.hour | print`, `10h0m0s`)

	RunSubO(t, "eq1", `print(10 * time.minute == 600 * time.second)`, `true`)
	RunSubO(t, "eq2", `print(10 * time.minute == 601 * time.second)`, `false`)
	RunSubO(t, "ne1", `print(10 * time.minute != 600 * time.second)`, `false`)
	RunSubO(t, "ne2", `print(10 * time.minute != 601 * time.second)`, `true`)

	RunSubO(t, "add", `print(1 * time.second + 100 * time.millisecond)`, `1.1s`)
	RunSubErr(t, "add_err", `print(1 * time.second + 1)`, vm.ErrOperationNotSupported)

	RunSubO(t, "sub", `print(1 * time.second - 100 * time.millisecond)`, `900ms`)
	RunSubO(t, "mul", `print(1 * time.second * 2)`, `2s`)
	RunSubErr(t, "mul_err", `print(1 * time.second * 2 * time.second)`, vm.ErrOperationNotSupported)
	RunSubO(t, "div_int", `print(1 * time.second / 2)`, `500ms`)
	RunSubO(t, "div_dur", `print(1 * time.second / (200 * time.millisecond))`, `5`)
	RunSubErr(t, "div_dur0_err", `print(1 * time.second / (0 * time.millisecond))`, vm.ErrDivByZero)
	RunSubErr(t, "div_int0_err", `print(1 * time.second / 0)`, vm.ErrDivByZero)
	RunSubErr(t, "div_err", `print(1 * time.second / "hi")`, vm.ErrOperationNotSupported)
	RunSubO(t, "mod", `print(250 * time.millisecond % (200 * time.millisecond))`, `50ms`)
	RunSubErr(t, "mod_div0_err", `print(250 * time.millisecond % (0 * time.millisecond))`, nil)

	RunSubO(t, "lt1", `print(250 * time.millisecond < 200 * time.millisecond)`, `false`)
	RunSubO(t, "lt2", `print(200 * time.millisecond < 250 * time.millisecond)`, `true`)
	RunSubO(t, "lt3", `print(200 * time.millisecond < 200 * time.millisecond)`, `false`)
	RunSubO(t, "le1", `print(250 * time.millisecond <= 200 * time.millisecond)`, `false`)
	RunSubO(t, "le2", `print(200 * time.millisecond <= 250 * time.millisecond)`, `true`)
	RunSubO(t, "le3", `print(200 * time.millisecond <= 200 * time.millisecond)`, `true`)
	RunSubO(t, "gt1", `print(250 * time.millisecond > 200 * time.millisecond)`, `true`)
	RunSubO(t, "gt2", `print(200 * time.millisecond > 250 * time.millisecond)`, `false`)
	RunSubO(t, "gt3", `print(200 * time.millisecond > 200 * time.millisecond)`, `false`)
	RunSubO(t, "ge1", `print(250 * time.millisecond >= 200 * time.millisecond)`, `true`)
	RunSubO(t, "ge2", `print(200 * time.millisecond >= 250 * time.millisecond)`, `false`)
	RunSubO(t, "ge3", `print(200 * time.millisecond >= 200 * time.millisecond)`, `true`)
	RunSubO(t, "uminus", `print(-(200 * time.millisecond))`, `-200ms`)
}

func TestTimeTruncate(t *testing.T) {
	RunSubO(t, "val", `1999 * time.millisecond | time.truncate(10 * time.millisecond) | print`, `1.99s`)
	RunSubO(t, "val2", `2999 * time.millisecond | time.truncate(1 * time.second) | print`, `2s`)
	RunSubErr(t, "val_usage_err2", `1999 * time.millisecond | time.truncate(1000) | print`, nil)
}
