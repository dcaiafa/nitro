package lib

import (
	"github.com/dcaiafa/nitro/internal/export"
	"github.com/dcaiafa/nitro/lib/encoding/json"
	"github.com/dcaiafa/nitro/lib/file"
	"github.com/dcaiafa/nitro/lib/io"
	"github.com/dcaiafa/nitro/lib/maps"
	"github.com/dcaiafa/nitro/lib/path/filepath"
	"github.com/dcaiafa/nitro/lib/str"
	libtime "github.com/dcaiafa/nitro/lib/time"
)

var GlobalPackage = export.Exports{
	{N: "$exec", T: export.Func, F: execExec},
	{N: "$home", T: export.Func, F: osHome},
	{N: "args", T: export.Func, F: args},
	{N: "avg", T: export.Func, F: avg},
	{N: "batch", T: export.Func, F: batch},
	{N: "close", T: export.Func, F: closep},
	{N: "color", T: export.Func, F: color},
	{N: "count", T: export.Func, F: count},
	{N: "discard", T: export.Func, F: discard},
	{N: "do", T: export.Func, F: dop},
	{N: "env", T: export.Func, F: env},
	{N: "filter", T: export.Func, F: filter},
	{N: "first", T: export.Func, F: first},
	{N: "flatten", T: export.Func, F: flatten},
	{N: "from_crlf", T: export.Func, F: fromCRLF},
	{N: "get", T: export.Func, F: get},
	{N: "group", T: export.Func, F: group},
	{N: "has", T: export.Func, F: has},
	{N: "hist", T: export.Func, F: hist},
	{N: "join", T: export.Func, F: join},
	{N: "len", T: export.Func, F: Len},
	{N: "lines", T: export.Func, F: lines},
	{N: "log", T: export.Func, F: log},
	{N: "logf", T: export.Func, F: logf},
	{N: "map", T: export.Func, F: imap},
	{N: "mapreduce", T: export.Func, F: mapreduce},
	{N: "max", T: export.Func, F: max},
	{N: "min", T: export.Func, F: min},
	{N: "narg", T: export.Func, F: narg},
	{N: "nonl", T: export.Func, F: nonl},
	{N: "parse_base64", T: export.Func, F: parseBase64},
	{N: "parse_csv", T: export.Func, F: parseCSV},
	{N: "parse_float", T: export.Func, F: parseFloat},
	{N: "parse_int", T: export.Func, F: parseInt},
	{N: "print", T: export.Func, F: print},
	{N: "print_table", T: export.Func, F: printTable},
	{N: "printf", T: export.Func, F: printf},
	{N: "probe", T: export.Func, F: probe},
	{N: "prompt", T: export.Func, F: prompt},
	{N: "range", T: export.Func, F: range_},
	{N: "read", T: export.Func, F: read},
	{N: "reduce", T: export.Func, F: reduce},
	{N: "regex", T: export.Func, F: regex},
	{N: "sha1", T: export.Func, F: sha1},
	{N: "shuffle", T: export.Func, F: shuffle},
	{N: "skip", T: export.Func, F: skip},
	{N: "sleep", T: export.Func, F: sleep},
	{N: "sort", T: export.Func, F: sort},
	{N: "sprintf", T: export.Func, F: sprintf},
	{N: "sum", T: export.Func, F: sum},
	{N: "take", T: export.Func, F: take},
	{N: "to_array", T: export.Func, F: toArray},
	{N: "to_base64", T: export.Func, F: toBase64},
	{N: "to_bool", T: export.Func, F: toBool},
	{N: "to_crlf", T: export.Func, F: toCRLF},
	{N: "to_hex", T: export.Func, F: toHex},
	{N: "to_int", T: export.Func, F: toInt},
	{N: "to_string", T: export.Func, F: toString},
	{N: "type", T: export.Func, F: typep},
	{N: "unique", T: export.Func, F: unique},
}

var BufPackage = export.Exports{
	{N: "cap", T: export.Func, F: bufCap},
	{N: "len", T: export.Func, F: bufLen},
	{N: "new", T: export.Func, F: bufNew},
	{N: "read", T: export.Func, F: bufRead},
	{N: "read_byte", T: export.Func, F: bufReadByte},
	{N: "read_from", T: export.Func, F: bufReadFrom},
	{N: "read_rune", T: export.Func, F: bufReadRune},
	{N: "unread_byte", T: export.Func, F: bufUnreadByte},
}

var CoPackage = export.Exports{
	{N: "run_with_timeout", T: export.Func, F: runWithTimeout},
	{N: "start", T: export.Func, F: start},
}

var ExecPackage = export.Exports{
	{N: "exec", T: export.Func, F: execExec},
	{N: "with_stderr", T: export.Func, F: execWithStderr},
}

var IterPackage = export.Exports{
	{N: "into", T: export.Func, F: iter_into},
	{N: "next", T: export.Func, F: iter_next},
}

var ListPackage = export.Exports{
	{N: "append", T: export.Func, F: listAppend},
	{N: "append_iter", T: export.Func, F: listAppendIter},
	{N: "find", T: export.Func, F: listFind},
}

var MathPackage = export.Exports{
	{N: "trunc", T: export.Func, F: mathTrunc},
}

var OSPackage = export.Exports{
	{N: "home", T: export.Func, F: osHome},
	{N: "wd", T: export.Func, F: osWD},
}

type BuiltinRegistry interface {
	RegisterBuiltins(pkgName string, exports export.Exports)
}

func RegisterAll(registry BuiltinRegistry) {
	registry.RegisterBuiltins("$global", GlobalPackage)
	registry.RegisterBuiltins("buf", BufPackage)
	registry.RegisterBuiltins("co", CoPackage)
	registry.RegisterBuiltins("encoding/json", json.Exports)
	registry.RegisterBuiltins("exec", ExecPackage)
	registry.RegisterBuiltins("file", file.Exports)
	registry.RegisterBuiltins("path/filepath", filepath.Exports)
	registry.RegisterBuiltins("io", io.Exports)
	registry.RegisterBuiltins("iter", IterPackage)
	registry.RegisterBuiltins("list", ListPackage)
	registry.RegisterBuiltins("maps", maps.Exports)
	registry.RegisterBuiltins("math", MathPackage)
	registry.RegisterBuiltins("os", OSPackage)
	registry.RegisterBuiltins("str", str.Exports)
	registry.RegisterBuiltins("time", libtime.Exports)
}
