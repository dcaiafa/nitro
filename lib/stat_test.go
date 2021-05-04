package lib

import "testing"

func TestAvg(t *testing.T) {
	RunSubO(t, `accum`, `
		var a = avg(nil, 1)
		a = avg(a, 2)
		a = avg(a, 3)
		a = avg(a, 4.5)
		a = avg(a)
		print(a)
	`, `2.625`)
	RunSubO(t, `accum_nil`, `
		var a = avg(nil, 1)
		a = avg(a, 2)
		a = avg(a, 3)
		a = avg(a, 4.5)
		a = avg(a, nil)
		print(a)
	`, `2.625`)
	RunSubO(t, `err_accum_nil`, `avg(nil) | print()`, `<nil>`)
	RunSubErr(t, `err_accum_not_num`, `
		var a = avg(nil, 1)
		a = avg(a, 2)
		a = avg(a, "foo")
		a = avg(a, 4.5)
		a = avg(a, nil)
		print(a)
	`, errAvgUsage)

	RunSubO(t, `args`, `print(avg(1, 2, 3, 4.5))`, `2.625`)
	RunSubErr(t, `err_args_not_num`, `print(avg(1, 2, "foo", 4.5))`, errAvgUsage)
	RunSubErr(t, `err_args_usage`, `print(avg())`, errAvgUsage)

	RunSubO(t, `iter`, `print(avg([1, 2, 3, 4.5]))`, `2.625`)
	RunSubO(t, `iter_empty`, `print(avg([]))`, `<nil>`)
	RunSubErr(t, `err_iter_not_num`, `print(avg([1, 2, "foo", 4.5]))`, errAvgUsage)

	RunSubO(t, `reduce`, `reduce([1, 2, 3, 4.5], avg) | print()`, `2.625`)
	RunSubO(t, `reduce`, `reduce([], avg) | print()`, `<nil>`)
}
