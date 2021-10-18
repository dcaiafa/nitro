package lib

import "testing"

func TestParseBase64(t *testing.T) {
	RunSubO(t, "valid", `"aGVsbG8="|parse_base64()|print()`, "hello")
	RunSubO(t, "valid_empty", `""|parse_base64()|print()`, "")
	RunSubErr(t, "err_invalid", `"aGVsbG8"|parse_base64()|print()`, nil)
	RunSubErr(t, "err_usage1", `parse_base64(123)`, errParseBase64Usage)
	RunSubErr(t, "err_usage2", `parse_base64()`, errParseBase64Usage)
	RunSubErr(t, "err_usage3", `parse_base64("aGVsbG8=", "x")`, errParseBase64Usage)
}

func TestToBase64(t *testing.T) {
	RunSubO(t, "valid", `"hello"|to_base64()|print()`, "aGVsbG8=")
	RunSubO(t, "valid_empty", `""|to_base64()|print()`, "")
	RunSubErr(t, "err_usage1", `to_base64(1)`, nil)
	RunSubErr(t, "err_usage2", `to_base64()`, nil)
	RunSubErr(t, "err_usage3", `to_base64("hello", 1)`, nil)
}
