package lib

import (
	"testing"

	"github.com/dcaiafa/nitro/internal/vm"
)

func TestTimeDuration(t *testing.T) {
	RunSubO(t, "nanosecond", `10 * time.NANOSECOND | print`, `10ns`)
	RunSubO(t, "microsecond", `10 * time.MICROSECOND | print`, `10µs`)
	RunSubO(t, "millisecond", `10 * time.MILLISECOND | print`, `10ms`)
	RunSubO(t, "second", `10 * time.SECOND | print`, `10s`)
	RunSubO(t, "minute", `10 * time.MINUTE | print`, `10m0s`)
	RunSubO(t, "hour", `10 * time.HOUR | print`, `10h0m0s`)

	RunSubO(t, "eq1", `print(10 * time.MINUTE == 600 * time.SECOND)`, `true`)
	RunSubO(t, "eq2", `print(10 * time.MINUTE == 601 * time.SECOND)`, `false`)
	RunSubO(t, "ne1", `print(10 * time.MINUTE != 600 * time.SECOND)`, `false`)
	RunSubO(t, "ne2", `print(10 * time.MINUTE != 601 * time.SECOND)`, `true`)

	RunSubO(t, "add", `print(1 * time.SECOND + 100 * time.MILLISECOND)`, `1.1s`)
	RunSubErr(t, "add_err", `print(1 * time.SECOND + 1)`, vm.ErrOperationNotSupported)

	RunSubO(t, "sub", `print(1 * time.SECOND - 100 * time.MILLISECOND)`, `900ms`)
	RunSubO(t, "mul", `print(1 * time.SECOND * 2)`, `2s`)
	RunSubErr(t, "mul_err", `print(1 * time.SECOND * 2 * time.SECOND)`, vm.ErrOperationNotSupported)
	RunSubO(t, "div_int", `print(1 * time.SECOND / 2)`, `500ms`)
	RunSubO(t, "div_dur", `print(1 * time.SECOND / (200 * time.MILLISECOND))`, `5`)
	RunSubErr(t, "div_dur0_err", `print(1 * time.SECOND / (0 * time.MILLISECOND))`, vm.ErrDivByZero)
	RunSubErr(t, "div_int0_err", `print(1 * time.SECOND / 0)`, vm.ErrDivByZero)
	RunSubErr(t, "div_err", `print(1 * time.SECOND / "hi")`, vm.ErrOperationNotSupported)
	RunSubO(t, "mod", `print(250 * time.MILLISECOND % (200 * time.MILLISECOND))`, `50ms`)
	RunSubErr(t, "mod_div0_err", `print(250 * time.MILLISECOND % (0 * time.MILLISECOND))`, nil)

	RunSubO(t, "lt1", `print(250 * time.MILLISECOND < 200 * time.MILLISECOND)`, `false`)
	RunSubO(t, "lt2", `print(200 * time.MILLISECOND < 250 * time.MILLISECOND)`, `true`)
	RunSubO(t, "lt3", `print(200 * time.MILLISECOND < 200 * time.MILLISECOND)`, `false`)
	RunSubO(t, "le1", `print(250 * time.MILLISECOND <= 200 * time.MILLISECOND)`, `false`)
	RunSubO(t, "le2", `print(200 * time.MILLISECOND <= 250 * time.MILLISECOND)`, `true`)
	RunSubO(t, "le3", `print(200 * time.MILLISECOND <= 200 * time.MILLISECOND)`, `true`)
	RunSubO(t, "gt1", `print(250 * time.MILLISECOND > 200 * time.MILLISECOND)`, `true`)
	RunSubO(t, "gt2", `print(200 * time.MILLISECOND > 250 * time.MILLISECOND)`, `false`)
	RunSubO(t, "gt3", `print(200 * time.MILLISECOND > 200 * time.MILLISECOND)`, `false`)
	RunSubO(t, "ge1", `print(250 * time.MILLISECOND >= 200 * time.MILLISECOND)`, `true`)
	RunSubO(t, "ge2", `print(200 * time.MILLISECOND >= 250 * time.MILLISECOND)`, `false`)
	RunSubO(t, "ge3", `print(200 * time.MILLISECOND >= 200 * time.MILLISECOND)`, `true`)
	RunSubO(t, "uminus", `print(-(200 * time.MILLISECOND))`, `-200ms`)
}

func TestTimeTruncate(t *testing.T) {
	RunSubO(t, "val", `1999 * time.MILLISECOND | time.truncate(10 * time.MILLISECOND) | print`, `1.99s`)
	RunSubO(t, "val2", `2999 * time.MILLISECOND | time.truncate(1 * time.SECOND) | print`, `2s`)
	RunSubErr(t, "val_usage_err2", `1999 * time.MILLISECOND | time.truncate(1000) | print`, nil)
}
