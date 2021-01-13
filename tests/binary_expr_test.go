package tests

import "testing"

func TestBinaryExpr(t *testing.T) {
	RunSubO(t, "int/add", `print(1+2)`, `3`)
	RunSubO(t, "int/sub", `print(1-2)`, `-1`)
	RunSubO(t, "int/mul", `print(2*3)`, `6`)
	RunSubO(t, "int/div", `print(6/2)`, `3`)
	RunSubO(t, "int/gt", `print(2>3, 3>2, 3>3)`, `false true false`)
	RunSubO(t, "int/ge", `print(2>=3, 3>=2, 3>=3)`, `false true true`)
	RunSubO(t, "int/lt", `print(2<3, 3<2, 3<3)`, `true false false`)
	RunSubO(t, "int/le", `print(2<=3, 3<=2, 3<=3)`, `true false true`)

	RunSubO(t, "float/add", `printf("%.2f", 1.1+2.2)`, `3.30`)
	RunSubO(t, "float/sub", `print(1.1-2.2)`, `-1.1`)
	RunSubO(t, "float/mul", `print(2.2*3.3)`, `7.26`)
	RunSubO(t, "float/div", `print(7.26/2.2)`, `3.3`)
	RunSubO(t, "float/gt", `print(2>3, 3>2, 3>3)`, `false true false`)
	RunSubO(t, "float/ge", `print(2>=3, 3>=2, 3>=3)`, `false true true`)
	RunSubO(t, "float/lt", `print(2<3, 3<2, 3<3)`, `true false false`)
	RunSubO(t, "float/le", `print(2<=3, 3<=2, 3<=3)`, `true false true`)

	RunSubO(t, "int2float_coercion", `print(1+1.1, 1.1+1, 3/2, 3.0/2)`, `2.1 2.1 1 1.5`)
	RunSubO(t, "precedence", `print(6/2 + 2*3, 6/(2+2)*3) `, `9 3`)
}
