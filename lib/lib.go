package lib

import (
	"time"

	"github.com/dcaiafa/nitro"
	"github.com/dcaiafa/nitro/internal/export"
)

var Exports = export.Exports{
	{P: "", N: "$exec", T: export.Func, F: execExec},
	{P: "", N: "$home", T: export.Func, F: osUserHomeDir},
	{P: "", N: "args", T: export.Func, F: args},
	{P: "", N: "avg", T: export.Func, F: avg},
	{P: "", N: "batch", T: export.Func, F: batch},
	{P: "", N: "close", T: export.Func, F: closep},
	{P: "", N: "color", T: export.Func, F: color},
	{P: "", N: "count", T: export.Func, F: count},
	{P: "", N: "discard", T: export.Func, F: discard},
	{P: "", N: "do", T: export.Func, F: dop},
	{P: "", N: "env", T: export.Func, F: env},
	{P: "", N: "filter", T: export.Func, F: filter},
	{P: "", N: "first", T: export.Func, F: first},
	{P: "", N: "flatten", T: export.Func, F: flatten},
	{P: "", N: "from_crlf", T: export.Func, F: fromCRLF},
	{P: "", N: "get", T: export.Func, F: get},
	{P: "", N: "group", T: export.Func, F: group},
	{P: "", N: "has", T: export.Func, F: has},
	{P: "", N: "hist", T: export.Func, F: hist},
	{P: "", N: "join", T: export.Func, F: join},
	{P: "", N: "len", T: export.Func, F: Len},
	{P: "", N: "lines", T: export.Func, F: lines},
	{P: "", N: "log", T: export.Func, F: log},
	{P: "", N: "logf", T: export.Func, F: logf},
	{P: "", N: "ls", T: export.Func, F: filepathLs},
	{P: "", N: "map", T: export.Func, F: imap},
	{P: "", N: "mapreduce", T: export.Func, F: mapreduce},
	{P: "", N: "max", T: export.Func, F: max},
	{P: "", N: "min", T: export.Func, F: min},
	{P: "", N: "narg", T: export.Func, F: narg},
	{P: "", N: "nonl", T: export.Func, F: nonl},
	{P: "", N: "open", T: export.Func, F: fileOpen},
	{P: "", N: "parse_base64", T: export.Func, F: parseBase64},
	{P: "", N: "parse_csv", T: export.Func, F: parseCSV},
	{P: "", N: "parse_float", T: export.Func, F: parseFloat},
	{P: "", N: "parse_int", T: export.Func, F: parseInt},
	{P: "", N: "parse_json", T: export.Func, F: parseJSON},
	{P: "", N: "pop_stdout", T: export.Func, F: popStdout},
	{P: "", N: "print", T: export.Func, F: print},
	{P: "", N: "print_table", T: export.Func, F: printTable},
	{P: "", N: "printf", T: export.Func, F: printf},
	{P: "", N: "probe", T: export.Func, F: probe},
	{P: "", N: "prompt", T: export.Func, F: prompt},
	{P: "", N: "push_stdout", T: export.Func, F: pushStdout},
	{P: "", N: "range", T: export.Func, F: range_},
	{P: "", N: "read", T: export.Func, F: read},
	{P: "", N: "reduce", T: export.Func, F: reduce},
	{P: "", N: "regex", T: export.Func, F: regex},
	{P: "", N: "sha1", T: export.Func, F: sha1},
	{P: "", N: "shuffle", T: export.Func, F: shuffle},
	{P: "", N: "skip", T: export.Func, F: skip},
	{P: "", N: "sleep", T: export.Func, F: sleep},
	{P: "", N: "sort", T: export.Func, F: sort},
	{P: "", N: "sprintf", T: export.Func, F: sprintf},
	{P: "", N: "stderr", T: export.Func, F: stderr},
	{P: "", N: "stdin", T: export.Func, F: stdin},
	{P: "", N: "stdout", T: export.Func, F: stdout},
	{P: "", N: "sum", T: export.Func, F: sum},
	{P: "", N: "take", T: export.Func, F: take},
	{P: "", N: "time_from_unix", T: export.Func, F: timeFromUnix},
	{P: "", N: "to_array", T: export.Func, F: toArray},
	{P: "", N: "to_base64", T: export.Func, F: toBase64},
	{P: "", N: "to_bool", T: export.Func, F: toBool},
	{P: "", N: "to_crlf", T: export.Func, F: toCRLF},
	{P: "", N: "to_hex", T: export.Func, F: toHex},
	{P: "", N: "to_int", T: export.Func, F: toInt},
	{P: "", N: "to_json", T: export.Func, F: toJSON},
	{P: "", N: "to_string", T: export.Func, F: toString},
	{P: "", N: "type", T: export.Func, F: typep},
	{P: "", N: "unique", T: export.Func, F: unique},
	{P: "buf", N: "cap", T: export.Func, F: bufCap},
	{P: "buf", N: "len", T: export.Func, F: bufLen},
	{P: "buf", N: "new", T: export.Func, F: bufNew},
	{P: "buf", N: "read", T: export.Func, F: bufRead},
	{P: "buf", N: "read_byte", T: export.Func, F: bufReadByte},
	{P: "buf", N: "read_from", T: export.Func, F: bufReadFrom},
	{P: "buf", N: "read_rune", T: export.Func, F: bufReadRune},
	{P: "buf", N: "unread_byte", T: export.Func, F: bufUnreadByte},
	{P: "co", N: "run_with_timeout", T: export.Func, F: runWithTimeout},
	{P: "co", N: "start", T: export.Func, F: start},
	{P: "exec", N: "exec", T: export.Func, F: execExec},
	{P: "exec", N: "with_stderr", T: export.Func, F: execWithStderr},
	{P: "file", N: "close", T: export.Func, F: closep},
	{P: "file", N: "create", T: export.Func, F: fileCreate},
	{P: "file", N: "create_temp", T: export.Func, F: fileCreateTemp},
	{P: "file", N: "open", T: export.Func, F: fileOpen},
	{P: "file", N: "read", T: export.Func, F: fileRead},
	{P: "file", N: "seek", T: export.Func, F: fileSeek},
	{P: "file", N: "stat", T: export.Func, F: fileStat},
	{P: "file", N: "write_to", T: export.Func, F: fileWriteTo},
	{P: "filepath", N: "abs", T: export.Func, F: filepathAbs},
	{P: "filepath", N: "base", T: export.Func, F: filepathBase},
	{P: "filepath", N: "clean", T: export.Func, F: filepathClean},
	{P: "filepath", N: "dir", T: export.Func, F: filepathDir},
	{P: "filepath", N: "eval_symlinks", T: export.Func, F: filepathEvalSymlinks},
	{P: "filepath", N: "exists", T: export.Func, F: filepathExists},
	{P: "filepath", N: "ext", T: export.Func, F: filepathExt},
	{P: "filepath", N: "from_slash", T: export.Func, F: filepathFromSlash},
	{P: "filepath", N: "is_abs", T: export.Func, F: filepathIsAbs},
	{P: "filepath", N: "is_dir", T: export.Func, F: filepathIsDir},
	{P: "filepath", N: "join", T: export.Func, F: filepathJoin},
	{P: "filepath", N: "ls", T: export.Func, F: filepathLs},
	{P: "filepath", N: "match", T: export.Func, F: filepathMatch},
	{P: "filepath", N: "mkdir", T: export.Func, F: filepathMkdir},
	{P: "filepath", N: "mkdir_all", T: export.Func, F: filepathMkdirAll},
	{P: "filepath", N: "mkdir_temp", T: export.Func, F: filepathMkdirTemp},
	{P: "filepath", N: "rel", T: export.Func, F: filepathRel},
	{P: "filepath", N: "remove", T: export.Func, F: filepathRemove},
	{P: "filepath", N: "remove_all", T: export.Func, F: filepathRemoveAll},
	{P: "filepath", N: "rename", T: export.Func, F: filepathRename},
	{P: "filepath", N: "split", T: export.Func, F: filepathSplit},
	{P: "filepath", N: "to_slash", T: export.Func, F: filepathToSlash},
	{P: "filepath", N: "volume_name", T: export.Func, F: filepathVolumeName},
	{P: "iter", N: "into", T: export.Func, F: iter_into},
	{P: "iter", N: "next", T: export.Func, F: iter_next},
	{P: "list", N: "append", T: export.Func, F: listAppend},
	{P: "list", N: "append_iter", T: export.Func, F: listAppendIter},
	{P: "list", N: "find", T: export.Func, F: listFind},
	{P: "maps", N: "clone", T: export.Func, F: mapsClone},
	{P: "maps", N: "delete", T: export.Func, F: mapsDelete},
	{P: "maps", N: "make", T: export.Func, F: mapsMake},
	{P: "maps", N: "update", T: export.Func, F: mapsUpdate},
	{P: "math", N: "trunc", T: export.Func, F: mathTrunc},
	{P: "os", N: "user_home_dir", T: export.Func, F: osUserHomeDir},
	{P: "os", N: "wd", T: export.Func, F: osWD},
	{P: "str", N: "fields", T: export.Func, F: strFields},
	{P: "str", N: "find", T: export.Func, F: strFind},
	{P: "str", N: "find_last", T: export.Func, F: strFindLast},
	{P: "str", N: "has_prefix", T: export.Func, F: strHasPrefix},
	{P: "str", N: "has_suffix", T: export.Func, F: strHasSuffix},
	{P: "str", N: "match", T: export.Func, F: strMatch},
	{P: "str", N: "match_all", T: export.Func, F: strMatchAll},
	{P: "str", N: "repeat", T: export.Func, F: strRepeat},
	{P: "str", N: "replace", T: export.Func, F: strReplace},
	{P: "str", N: "split", T: export.Func, F: strSplit},
	{P: "str", N: "to_lower", T: export.Func, F: strToLower},
	{P: "str", N: "to_upper", T: export.Func, F: strToUpper},
	{P: "str", N: "trim", T: export.Func, F: strTrim},
	{P: "str", N: "trim_left", T: export.Func, F: strTrimLeft},
	{P: "str", N: "trim_prefix", T: export.Func, F: strTrimPrefix},
	{P: "str", N: "trim_right", T: export.Func, F: strTrimRight},
	{P: "str", N: "trim_suffix", T: export.Func, F: strTrimSuffix},
	{P: "time", N: "format", T: export.Func, F: timeFormat},
	{P: "time", N: "from_unix", T: export.Func, F: timeFromUnix},
	{P: "time", N: "hour", T: export.Custom, I: int64(time.Hour), C: exportDurationConst},
	{P: "time", N: "local", T: export.Func, F: timeLocal},
	{P: "time", N: "microsecond", T: export.Custom, I: int64(time.Microsecond), C: exportDurationConst},
	{P: "time", N: "millisecond", T: export.Custom, I: int64(time.Millisecond), C: exportDurationConst},
	{P: "time", N: "minute", T: export.Custom, I: int64(time.Minute), C: exportDurationConst},
	{P: "time", N: "nanosecond", T: export.Custom, I: int64(time.Nanosecond), C: exportDurationConst},
	{P: "time", N: "now", T: export.Func, F: timeNow},
	{P: "time", N: "parse", T: export.Func, F: timeParse},
	{P: "time", N: "second", T: export.Custom, I: int64(time.Second), C: exportDurationConst},
	{P: "time", N: "to_map", T: export.Func, F: timeToMap},
	{P: "time", N: "truncate", T: export.Func, F: timeTruncate},
	{P: "time", N: "unix", T: export.Func, F: timeUnix},
	{P: "time", N: "unix_nano", T: export.Func, F: timeUnixNano},
	{P: "time", N: "utc", T: export.Func, F: timeUTC},
}

func exportDurationConst(e *export.Export) nitro.Value {
	return NewDuration(time.Duration(e.I))
}
