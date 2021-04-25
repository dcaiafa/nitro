package lib

import (
	"testing"
)

func TestDur(t *testing.T) {
	RunSubO(t, "nanosecond", `print(dur(10, "nanosecond"))`, `10ns`)
	RunSubO(t, "microsecond", `print(dur(10, "microsecond"))`, `10µs`)
	RunSubO(t, "millisecond", `print(dur(10, "millisecond"))`, `10ms`)
	RunSubO(t, "second", `print(dur(10, "second"))`, `10s`)
	RunSubO(t, "minute", `print(dur(10, "minute"))`, `10m0s`)
	RunSubO(t, "hour", `print(dur(10, "hour"))`, `10h0m0s`)

	RunSubO(t, "eq1", `print(dur(10, "minute") == dur(600, "second"))`, `true`)
	RunSubO(t, "eq2", `print(dur(10, "minute") == dur(601, "second"))`, `false`)
	RunSubO(t, "ne1", `print(dur(10, "minute") != dur(600, "second"))`, `false`)
	RunSubO(t, "ne2", `print(dur(10, "minute") != dur(601, "second"))`, `true`)
	RunSubO(t, "add", `print(dur(1, "second") + dur(100, "millisecond"))`, `1.1s`)
	RunSubErr(t, "add_err", `print(dur(1, "second") + 1)`, errInvalidOp)
	RunSubO(t, "sub", `print(dur(1, "second") - dur(100, "millisecond"))`, `900ms`)
	RunSubO(t, "mul", `print(dur(1, "second") * 2)`, `2s`)
	RunSubErr(t, "mul_err", `print(dur(1, "second") * dur(2, "second"))`, errInvalidOp)
	RunSubO(t, "div_int", `print(dur(1, "second") / 2)`, `500ms`)
	RunSubO(t, "div_dur", `print(dur(1, "second") / dur(200, "millisecond"))`, `5`)
	RunSubErr(t, "div_dur0_err", `print(dur(1, "second") / dur(0, "millisecond"))`, errDivByZero)
	RunSubErr(t, "div_int0_err", `print(dur(1, "second") / 0)`, errDivByZero)
	RunSubErr(t, "div_err", `print(dur(1, "second") / "hi")`, errInvalidOp)
	RunSubO(t, "mod", `print(dur(250, "millisecond") % dur(200, "millisecond"))`, `50ms`)
	RunSubErr(t, "mod_div0_err", `print(dur(250, "millisecond") % dur(0, "millisecond"))`, errDivByZero)
	RunSubO(t, "lt1", `print(dur(250, "millisecond") < dur(200, "millisecond"))`, `false`)
	RunSubO(t, "lt2", `print(dur(200, "millisecond") < dur(250, "millisecond"))`, `true`)
	RunSubO(t, "lt3", `print(dur(200, "millisecond") < dur(200, "millisecond"))`, `false`)
	RunSubO(t, "le1", `print(dur(250, "millisecond") <= dur(200, "millisecond"))`, `false`)
	RunSubO(t, "le2", `print(dur(200, "millisecond") <= dur(250, "millisecond"))`, `true`)
	RunSubO(t, "le3", `print(dur(200, "millisecond") <= dur(200, "millisecond"))`, `true`)
	RunSubO(t, "gt1", `print(dur(250, "millisecond") > dur(200, "millisecond"))`, `true`)
	RunSubO(t, "gt2", `print(dur(200, "millisecond") > dur(250, "millisecond"))`, `false`)
	RunSubO(t, "gt3", `print(dur(200, "millisecond") > dur(200, "millisecond"))`, `false`)
	RunSubO(t, "ge1", `print(dur(250, "millisecond") >= dur(200, "millisecond"))`, `true`)
	RunSubO(t, "ge2", `print(dur(200, "millisecond") >= dur(250, "millisecond"))`, `false`)
	RunSubO(t, "ge3", `print(dur(200, "millisecond") >= dur(200, "millisecond"))`, `true`)
	RunSubO(t, "uminus", `print(-dur(200, "millisecond"))`, `-200ms`)

	RunSubErr(t, "errUsage1", `print(dur("hour", "10"))`, errDurUsage)
	RunSubErr(t, "errUsage2", `print(dur("10"))`, errDurUsage)
	RunSubErr(t, "errUsage3", `print(dur())`, errDurUsage)
	RunSubErr(t, "errUsage4", `print(dur(10, 20))`, errDurUsage)
	RunSubErr(t, "errUsage5", `print(dur(10, "days"))`, errDurUsage)
}

func TestTruncDur(t *testing.T) {
	RunSubO(t, "val", `print(truncdur(dur(1999, "millisecond"), dur(10, "millisecond")))`, `1.99s`)
	RunSubO(t, "val", `print(truncdur(dur(1999, "millisecond"), "second"))`, `1s`)
	RunSubErr(t, "val_usage_err1", `truncdur()`, errTruncDurUsage)
	RunSubErr(t, "val_usage_err2", `truncdur(dur(1999, "millisecond"), 1000)`, errTruncDurUsage)
	RunSubErr(t, "val_usage_err3", `truncdur(1000, "second")`, errTruncDurUsage)
	RunSubErr(t, "val_usage_err4", `truncdur(dur(1999, "millisecond"), "years")`, errTruncDurUsage)
}
