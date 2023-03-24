package lib

import (
	"time"

	"github.com/dcaiafa/nitro"
	"github.com/dcaiafa/nitro/internal/export"
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
	{N: "ls", T: export.Func, F: filepathLs},
	{N: "map", T: export.Func, F: imap},
	{N: "mapreduce", T: export.Func, F: mapreduce},
	{N: "max", T: export.Func, F: max},
	{N: "min", T: export.Func, F: min},
	{N: "narg", T: export.Func, F: narg},
	{N: "nonl", T: export.Func, F: nonl},
	{N: "open", T: export.Func, F: fileOpen},
	{N: "parse_base64", T: export.Func, F: parseBase64},
	{N: "parse_csv", T: export.Func, F: parseCSV},
	{N: "parse_float", T: export.Func, F: parseFloat},
	{N: "parse_int", T: export.Func, F: parseInt},
	{N: "parse_json", T: export.Func, F: parseJSON},
	{N: "pop_stdout", T: export.Func, F: popStdout},
	{N: "print", T: export.Func, F: print},
	{N: "print_table", T: export.Func, F: printTable},
	{N: "printf", T: export.Func, F: printf},
	{N: "probe", T: export.Func, F: probe},
	{N: "prompt", T: export.Func, F: prompt},
	{N: "push_stdout", T: export.Func, F: pushStdout},
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
	{N: "stderr", T: export.Func, F: stderr},
	{N: "stdin", T: export.Func, F: stdin},
	{N: "stdout", T: export.Func, F: stdout},
	{N: "sum", T: export.Func, F: sum},
	{N: "take", T: export.Func, F: take},
	{N: "time_from_unix", T: export.Func, F: timeFromUnix},
	{N: "to_array", T: export.Func, F: toArray},
	{N: "to_base64", T: export.Func, F: toBase64},
	{N: "to_bool", T: export.Func, F: toBool},
	{N: "to_crlf", T: export.Func, F: toCRLF},
	{N: "to_hex", T: export.Func, F: toHex},
	{N: "to_int", T: export.Func, F: toInt},
	{N: "to_json", T: export.Func, F: toJSON},
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

var FilePackage = export.Exports{
	{N: "close", T: export.Func, F: closep},
	{N: "create", T: export.Func, F: fileCreate},
	{N: "create_temp", T: export.Func, F: fileCreateTemp},
	{N: "open", T: export.Func, F: fileOpen},
	{N: "read", T: export.Func, F: fileRead},
	{N: "seek", T: export.Func, F: fileSeek},
	{N: "stat", T: export.Func, F: fileStat},
	{N: "write_to", T: export.Func, F: fileWriteTo},
}

var FilepathPackage = export.Exports{
	{N: "abs", T: export.Func, F: filepathAbs},
	{N: "base", T: export.Func, F: filepathBase},
	{N: "clean", T: export.Func, F: filepathClean},
	{N: "dir", T: export.Func, F: filepathDir},
	{N: "eval_symlinks", T: export.Func, F: filepathEvalSymlinks},
	{N: "exists", T: export.Func, F: filepathExists},
	{N: "ext", T: export.Func, F: filepathExt},
	{N: "from_slash", T: export.Func, F: filepathFromSlash},
	{N: "is_abs", T: export.Func, F: filepathIsAbs},
	{N: "is_dir", T: export.Func, F: filepathIsDir},
	{N: "join", T: export.Func, F: filepathJoin},
	{N: "ls", T: export.Func, F: filepathLs},
	{N: "match", T: export.Func, F: filepathMatch},
	{N: "mkdir", T: export.Func, F: filepathMkdir},
	{N: "mkdir_all", T: export.Func, F: filepathMkdirAll},
	{N: "mkdir_temp", T: export.Func, F: filepathMkdirTemp},
	{N: "rel", T: export.Func, F: filepathRel},
	{N: "remove", T: export.Func, F: filepathRemove},
	{N: "remove_all", T: export.Func, F: filepathRemoveAll},
	{N: "rename", T: export.Func, F: filepathRename},
	{N: "split", T: export.Func, F: filepathSplit},
	{N: "to_slash", T: export.Func, F: filepathToSlash},
	{N: "volume_name", T: export.Func, F: filepathVolumeName},
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

var MapsPackage = export.Exports{
	{N: "clone", T: export.Func, F: mapsClone},
	{N: "delete", T: export.Func, F: mapsDelete},
	{N: "make", T: export.Func, F: mapsMake},
	{N: "update", T: export.Func, F: mapsUpdate},
}

var MathPackage = export.Exports{
	{N: "trunc", T: export.Func, F: mathTrunc},
}

var OSPackage = export.Exports{
	{N: "home", T: export.Func, F: osHome},
	{N: "wd", T: export.Func, F: osWD},
}

var StrPackage = export.Exports{
	{N: "fields", T: export.Func, F: strFields},
	{N: "find", T: export.Func, F: strFind},
	{N: "find_last", T: export.Func, F: strFindLast},
	{N: "has_prefix", T: export.Func, F: strHasPrefix},
	{N: "has_suffix", T: export.Func, F: strHasSuffix},
	{N: "match", T: export.Func, F: strMatch},
	{N: "match_all", T: export.Func, F: strMatchAll},
	{N: "repeat", T: export.Func, F: strRepeat},
	{N: "replace", T: export.Func, F: strReplace},
	{N: "split", T: export.Func, F: strSplit},
	{N: "to_lower", T: export.Func, F: strToLower},
	{N: "to_upper", T: export.Func, F: strToUpper},
	{N: "trim", T: export.Func, F: strTrim},
	{N: "trim_left", T: export.Func, F: strTrimLeft},
	{N: "trim_prefix", T: export.Func, F: strTrimPrefix},
	{N: "trim_right", T: export.Func, F: strTrimRight},
	{N: "trim_suffix", T: export.Func, F: strTrimSuffix},
}

var TimePackage = export.Exports{
	{N: "format", T: export.Func, F: timeFormat},
	{N: "from_unix", T: export.Func, F: timeFromUnix},
	{N: "hour", T: export.Custom, I: int64(time.Hour), C: exportDurationConst},
	{N: "local", T: export.Func, F: timeLocal},
	{N: "microsecond", T: export.Custom, I: int64(time.Microsecond), C: exportDurationConst},
	{N: "millisecond", T: export.Custom, I: int64(time.Millisecond), C: exportDurationConst},
	{N: "minute", T: export.Custom, I: int64(time.Minute), C: exportDurationConst},
	{N: "nanosecond", T: export.Custom, I: int64(time.Nanosecond), C: exportDurationConst},
	{N: "now", T: export.Func, F: timeNow},
	{N: "parse", T: export.Func, F: timeParse},
	{N: "second", T: export.Custom, I: int64(time.Second), C: exportDurationConst},
	{N: "to_map", T: export.Func, F: timeToMap},
	{N: "truncate", T: export.Func, F: timeTruncate},
	{N: "unix", T: export.Func, F: timeUnix},
	{N: "unix_nano", T: export.Func, F: timeUnixNano},
	{N: "utc", T: export.Func, F: timeUTC},
}

type BuiltinRegistry interface {
	RegisterBuiltins(pkgName string, exports export.Exports)
}

func RegisterAll(registry BuiltinRegistry) {
	registry.RegisterBuiltins("$global", GlobalPackage)
	registry.RegisterBuiltins("buf", BufPackage)
	registry.RegisterBuiltins("co", CoPackage)
	registry.RegisterBuiltins("exec", ExecPackage)
	registry.RegisterBuiltins("file", FilePackage)
	registry.RegisterBuiltins("filepath", FilepathPackage)
	registry.RegisterBuiltins("iter", IterPackage)
	registry.RegisterBuiltins("list", ListPackage)
	registry.RegisterBuiltins("maps", MapsPackage)
	registry.RegisterBuiltins("math", MathPackage)
	registry.RegisterBuiltins("os", OSPackage)
	registry.RegisterBuiltins("str", StrPackage)
	registry.RegisterBuiltins("time", TimePackage)
}

func exportDurationConst(e *export.Export) nitro.Value {
	return NewDuration(time.Duration(e.I))
}
