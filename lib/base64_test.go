package lib

import "testing"

func TestParseBase64(t *testing.T) {
	RunSubO(t, "valid", `"aGVsbG8="|parsebase64()|print()`, "hello")
	RunSubO(t, "valid_empty", `""|parsebase64()|print()`, "")
	RunSubErr(t, "err_invalid", `"aGVsbG8"|parsebase64()|print()`, nil)
	RunSubErr(t, "err_usage1", `parsebase64(123)`, errParseBase64Usage)
	RunSubErr(t, "err_usage2", `parsebase64()`, errParseBase64Usage)
	RunSubErr(t, "err_usage3", `parsebase64("aGVsbG8=", "x")`, errParseBase64Usage)
}

func TestToBase64(t *testing.T) {
	RunSubO(t, "valid", `"hello"|tobase64()|print()`, "aGVsbG8=")
	RunSubO(t, "valid_empty", `""|tobase64()|print()`, "")
	RunSubErr(t, "err_usage1", `tobase64(1)`, nil)
	RunSubErr(t, "err_usage2", `tobase64()`, nil)
	RunSubErr(t, "err_usage3", `tobase64("hello", 1)`, nil)
}
