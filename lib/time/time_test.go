package time_test

import (
	"testing"

	"github.com/dcaiafa/nitro/internal/btesting"
	"github.com/dcaiafa/nitro/internal/vm"
)

func TestTimeDuration(t *testing.T) {
	btesting.RunSubO(t, "nanosecond", `10 * time.NANOSECOND | print`, `10ns`)
	btesting.RunSubO(t, "microsecond", `10 * time.MICROSECOND | print`, `10µs`)
	btesting.RunSubO(t, "millisecond", `10 * time.MILLISECOND | print`, `10ms`)
	btesting.RunSubO(t, "second", `10 * time.SECOND | print`, `10s`)
	btesting.RunSubO(t, "minute", `10 * time.MINUTE | print`, `10m0s`)
	btesting.RunSubO(t, "hour", `10 * time.HOUR | print`, `10h0m0s`)

	btesting.RunSubO(t, "eq1", `print(10 * time.MINUTE == 600 * time.SECOND)`, `true`)
	btesting.RunSubO(t, "eq2", `print(10 * time.MINUTE == 601 * time.SECOND)`, `false`)
	btesting.RunSubO(t, "ne1", `print(10 * time.MINUTE != 600 * time.SECOND)`, `false`)
	btesting.RunSubO(t, "ne2", `print(10 * time.MINUTE != 601 * time.SECOND)`, `true`)

	btesting.RunSubO(t, "add", `print(1 * time.SECOND + 100 * time.MILLISECOND)`, `1.1s`)
	btesting.RunSubErr(t, "add_err", `print(1 * time.SECOND + 1)`, vm.ErrOperationNotSupported)

	btesting.RunSubO(t, "sub", `print(1 * time.SECOND - 100 * time.MILLISECOND)`, `900ms`)
	btesting.RunSubO(t, "mul", `print(1 * time.SECOND * 2)`, `2s`)
	btesting.RunSubErr(t, "mul_err", `print(1 * time.SECOND * 2 * time.SECOND)`, vm.ErrOperationNotSupported)
	btesting.RunSubO(t, "div_int", `print(1 * time.SECOND / 2)`, `500ms`)
	btesting.RunSubO(t, "div_dur", `print(1 * time.SECOND / (200 * time.MILLISECOND))`, `5`)
	btesting.RunSubErr(t, "div_dur0_err", `print(1 * time.SECOND / (0 * time.MILLISECOND))`, vm.ErrDivByZero)
	btesting.RunSubErr(t, "div_int0_err", `print(1 * time.SECOND / 0)`, vm.ErrDivByZero)
	btesting.RunSubErr(t, "div_err", `print(1 * time.SECOND / "hi")`, vm.ErrOperationNotSupported)
	btesting.RunSubO(t, "mod", `print(250 * time.MILLISECOND % (200 * time.MILLISECOND))`, `50ms`)
	btesting.RunSubErr(t, "mod_div0_err", `print(250 * time.MILLISECOND % (0 * time.MILLISECOND))`, nil)

	btesting.RunSubO(t, "lt1", `print(250 * time.MILLISECOND < 200 * time.MILLISECOND)`, `false`)
	btesting.RunSubO(t, "lt2", `print(200 * time.MILLISECOND < 250 * time.MILLISECOND)`, `true`)
	btesting.RunSubO(t, "lt3", `print(200 * time.MILLISECOND < 200 * time.MILLISECOND)`, `false`)
	btesting.RunSubO(t, "le1", `print(250 * time.MILLISECOND <= 200 * time.MILLISECOND)`, `false`)
	btesting.RunSubO(t, "le2", `print(200 * time.MILLISECOND <= 250 * time.MILLISECOND)`, `true`)
	btesting.RunSubO(t, "le3", `print(200 * time.MILLISECOND <= 200 * time.MILLISECOND)`, `true`)
	btesting.RunSubO(t, "gt1", `print(250 * time.MILLISECOND > 200 * time.MILLISECOND)`, `true`)
	btesting.RunSubO(t, "gt2", `print(200 * time.MILLISECOND > 250 * time.MILLISECOND)`, `false`)
	btesting.RunSubO(t, "gt3", `print(200 * time.MILLISECOND > 200 * time.MILLISECOND)`, `false`)
	btesting.RunSubO(t, "ge1", `print(250 * time.MILLISECOND >= 200 * time.MILLISECOND)`, `true`)
	btesting.RunSubO(t, "ge2", `print(200 * time.MILLISECOND >= 250 * time.MILLISECOND)`, `false`)
	btesting.RunSubO(t, "ge3", `print(200 * time.MILLISECOND >= 200 * time.MILLISECOND)`, `true`)
	btesting.RunSubO(t, "uminus", `print(-(200 * time.MILLISECOND))`, `-200ms`)
}

func TestTimeTruncate(t *testing.T) {
	btesting.RunSubO(t, "val", `1999 * time.MILLISECOND | time.truncate(10 * time.MILLISECOND) | print`, `1.99s`)
	btesting.RunSubO(t, "val2", `2999 * time.MILLISECOND | time.truncate(1 * time.SECOND) | print`, `2s`)
	btesting.RunSubErr(t, "val_usage_err2", `1999 * time.MILLISECOND | time.truncate(1000) | print`, nil)
}
