package lib

import (
	gosort "sort"
	"time"

	"github.com/dcaiafa/nitro"
	"github.com/dcaiafa/nitro/internal/vm"
)

type exportType int

const (
	exportFunc exportType = iota
	exportString
	exportInt
	exportFloat
	exportCustom
)

type export struct {
	P  string
	N  string
	T  exportType
	F  func(vm *nitro.VM, args []nitro.Value, nret int) ([]nitro.Value, error)
	S  string
	I  int64
	Fl float64
	C  func(e *export) nitro.Value
}

var allExports = []export{
	{P: "", N: "$exec", T: exportFunc, F: execExec},
	{P: "", N: "args", T: exportFunc, F: args},
	{P: "", N: "avg", T: exportFunc, F: avg},
	{P: "", N: "batch", T: exportFunc, F: batch},
	{P: "", N: "buf", T: exportFunc, F: buf},
	{P: "", N: "close", T: exportFunc, F: closep},
	{P: "", N: "color", T: exportFunc, F: color},
	{P: "", N: "copy_file", T: exportFunc, F: copyFile},
	{P: "", N: "count", T: exportFunc, F: count},
	{P: "", N: "create", T: exportFunc, F: create},
	{P: "", N: "create_temp", T: exportFunc, F: createTemp},
	{P: "", N: "delete", T: exportFunc, F: delete_},
	{P: "", N: "discard", T: exportFunc, F: discard},
	{P: "", N: "do", T: exportFunc, F: dop},
	{P: "", N: "dur", T: exportFunc, F: dur},
	{P: "", N: "env", T: exportFunc, F: env},
	{P: "", N: "file_exists", T: exportFunc, F: fileExists},
	{P: "", N: "filter", T: exportFunc, F: filter},
	{P: "", N: "first", T: exportFunc, F: first},
	{P: "", N: "flatten", T: exportFunc, F: flatten},
	{P: "", N: "from_crlf", T: exportFunc, F: fromCRLF},
	{P: "", N: "get", T: exportFunc, F: get},
	{P: "", N: "group", T: exportFunc, F: group},
	{P: "", N: "has", T: exportFunc, F: has},
	{P: "", N: "hist", T: exportFunc, F: hist},
	{P: "", N: "home_dir", T: exportFunc, F: homeDir},
	{P: "", N: "is_dir", T: exportFunc, F: is_dir},
	{P: "", N: "join", T: exportFunc, F: join},
	{P: "", N: "len", T: exportFunc, F: Len},
	{P: "", N: "lines", T: exportFunc, F: lines},
	{P: "", N: "log", T: exportFunc, F: log},
	{P: "", N: "logf", T: exportFunc, F: logf},
	{P: "", N: "ls", T: exportFunc, F: ls},
	{P: "", N: "map", T: exportFunc, F: mapp},
	{P: "", N: "mapreduce", T: exportFunc, F: mapreduce},
	{P: "", N: "max", T: exportFunc, F: max},
	{P: "", N: "min", T: exportFunc, F: min},
	{P: "", N: "mkdir_all", T: exportFunc, F: mkdirAll},
	{P: "", N: "move_file", T: exportFunc, F: moveFile},
	{P: "", N: "narg", T: exportFunc, F: narg},
	{P: "", N: "nonl", T: exportFunc, F: nonl},
	{P: "", N: "open", T: exportFunc, F: open},
	{P: "", N: "parse_base64", T: exportFunc, F: parseBase64},
	{P: "", N: "parse_csv", T: exportFunc, F: parseCSV},
	{P: "", N: "parse_float", T: exportFunc, F: parseFloat},
	{P: "", N: "parse_int", T: exportFunc, F: parseInt},
	{P: "", N: "parse_json", T: exportFunc, F: parseJSON},
	{P: "", N: "pop_stdout", T: exportFunc, F: popStdout},
	{P: "", N: "print", T: exportFunc, F: print},
	{P: "", N: "print_table", T: exportFunc, F: printTable},
	{P: "", N: "printf", T: exportFunc, F: printf},
	{P: "", N: "probe", T: exportFunc, F: probe},
	{P: "", N: "prompt", T: exportFunc, F: prompt},
	{P: "", N: "push_stdout", T: exportFunc, F: pushStdout},
	{P: "", N: "range", T: exportFunc, F: range_},
	{P: "", N: "read", T: exportFunc, F: read},
	{P: "", N: "read_file", T: exportFunc, F: readFile},
	{P: "", N: "reduce", T: exportFunc, F: reduce},
	{P: "", N: "regex", T: exportFunc, F: regex},
	{P: "", N: "remove_all", T: exportFunc, F: removeAll},
	{P: "", N: "remove_file", T: exportFunc, F: removeFile},
	{P: "", N: "seek", T: exportFunc, F: seek},
	{P: "", N: "sha1", T: exportFunc, F: sha1},
	{P: "", N: "shuffle", T: exportFunc, F: shuffle},
	{P: "", N: "skip", T: exportFunc, F: skip},
	{P: "", N: "sleep", T: exportFunc, F: sleep},
	{P: "", N: "sort", T: exportFunc, F: sort},
	{P: "", N: "sprintf", T: exportFunc, F: sprintf},
	{P: "", N: "stat", T: exportFunc, F: stat},
	{P: "", N: "stderr", T: exportFunc, F: stderr},
	{P: "", N: "stdin", T: exportFunc, F: stdin},
	{P: "", N: "stdout", T: exportFunc, F: stdout},
	{P: "", N: "sum", T: exportFunc, F: sum},
	{P: "", N: "take", T: exportFunc, F: take},
	{P: "", N: "time_from_unix", T: exportFunc, F: timeFromUnix},
	{P: "", N: "to_array", T: exportFunc, F: toArray},
	{P: "", N: "to_base64", T: exportFunc, F: toBase64},
	{P: "", N: "to_bool", T: exportFunc, F: toBool},
	{P: "", N: "to_crlf", T: exportFunc, F: toCRLF},
	{P: "", N: "to_hex", T: exportFunc, F: toHex},
	{P: "", N: "to_int", T: exportFunc, F: toInt},
	{P: "", N: "to_json", T: exportFunc, F: toJSON},
	{P: "", N: "to_map", T: exportFunc, F: toMap},
	{P: "", N: "to_string", T: exportFunc, F: toString},
	{P: "", N: "type", T: exportFunc, F: typep},
	{P: "", N: "unique", T: exportFunc, F: unique},
	{P: "", N: "write_file", T: exportFunc, F: writeFile},
	{P: "co", N: "run_with_timeout", T: exportFunc, F: runWithTimeout},
	{P: "co", N: "start", T: exportFunc, F: start},
	{P: "exec", N: "exec", T: exportFunc, F: execExec},
	{P: "exec", N: "with_stderr", T: exportFunc, F: execWithStderr},
	{P: "list", N: "append", T: exportFunc, F: listAppend},
	{P: "list", N: "append_iter", T: exportFunc, F: listAppendIter},
	{P: "list", N: "find", T: exportFunc, F: listFind},
	{P: "math", N: "trunc", T: exportFunc, F: mathTrunc},
	{P: "path", N: "base", T: exportFunc, F: pathBase},
	{P: "path", N: "clean", T: exportFunc, F: pathClean},
	{P: "path", N: "dir", T: exportFunc, F: pathDir},
	{P: "path", N: "ext", T: exportFunc, F: pathExt},
	{P: "path", N: "from_slash", T: exportFunc, F: pathFromSlash},
	{P: "path", N: "join", T: exportFunc, F: pathJoin},
	{P: "path", N: "match", T: exportFunc, F: pathMatch},
	{P: "path", N: "to_slash", T: exportFunc, F: pathToSlash},
	{P: "str", N: "fields", T: exportFunc, F: strFields},
	{P: "str", N: "find", T: exportFunc, F: strFind},
	{P: "str", N: "find_last", T: exportFunc, F: strFindLast},
	{P: "str", N: "has_prefix", T: exportFunc, F: strHasPrefix},
	{P: "str", N: "has_suffix", T: exportFunc, F: strHasSuffix},
	{P: "str", N: "match", T: exportFunc, F: strMatch},
	{P: "str", N: "match_all", T: exportFunc, F: strMatchAll},
	{P: "str", N: "repeat", T: exportFunc, F: strRepeat},
	{P: "str", N: "replace", T: exportFunc, F: strReplace},
	{P: "str", N: "split", T: exportFunc, F: strSplit},
	{P: "str", N: "to_lower", T: exportFunc, F: strToLower},
	{P: "str", N: "to_upper", T: exportFunc, F: strToUpper},
	{P: "str", N: "trim", T: exportFunc, F: strTrim},
	{P: "str", N: "trim_left", T: exportFunc, F: strTrimLeft},
	{P: "str", N: "trim_prefix", T: exportFunc, F: strTrimPrefix},
	{P: "str", N: "trim_right", T: exportFunc, F: strTrimRight},
	{P: "str", N: "trim_suffix", T: exportFunc, F: strTrimSuffix},
	{P: "time", N: "format", T: exportFunc, F: timeFormat},
	{P: "time", N: "from_unix", T: exportFunc, F: timeFromUnix},
	{P: "time", N: "hour", T: exportCustom, I: int64(time.Hour), C: exportDurationConst},
	{P: "time", N: "local", T: exportFunc, F: timeLocal},
	{P: "time", N: "microsecond", T: exportCustom, I: int64(time.Microsecond), C: exportDurationConst},
	{P: "time", N: "millisecond", T: exportCustom, I: int64(time.Millisecond), C: exportDurationConst},
	{P: "time", N: "minute", T: exportCustom, I: int64(time.Minute), C: exportDurationConst},
	{P: "time", N: "nanosecond", T: exportCustom, I: int64(time.Nanosecond), C: exportDurationConst},
	{P: "time", N: "now", T: exportFunc, F: timeNow},
	{P: "time", N: "parse", T: exportFunc, F: timeParse},
	{P: "time", N: "second", T: exportCustom, I: int64(time.Second), C: exportDurationConst},
	{P: "time", N: "to_map", T: exportFunc, F: timeToMap},
	{P: "time", N: "unix", T: exportFunc, F: timeUnix},
	{P: "time", N: "unix_nano", T: exportFunc, F: timeUnixNano},
	{P: "time", N: "utc", T: exportFunc, F: timeUTC},
}

type exportRegistry struct {
}

func NewExportRegistry() vm.ExportRegistry {
	return &exportRegistry{}
}

func (r *exportRegistry) IsValidPackage(pkg string) bool {
	i := gosort.Search(len(allExports), func(i int) bool {
		return allExports[i].P >= pkg
	})
	return i < len(allExports) && allExports[i].P == pkg
}

func (r *exportRegistry) GetExport(pkg, name string) vm.Value {
	i := gosort.Search(len(allExports), func(i int) bool {
		if allExports[i].P > pkg {
			return true
		} else if allExports[i].P == pkg && allExports[i].N >= name {
			return true
		}
		return false
	})
	if i >= len(allExports) ||
		allExports[i].P != pkg ||
		allExports[i].N != name {
		return nil
	}

	switch allExports[i].T {
	case exportFunc:
		return vm.NewNativeFn(allExports[i].F)
	case exportString:
		return vm.NewString(allExports[i].S)
	case exportInt:
		return vm.NewInt(allExports[i].I)
	case exportFloat:
		return vm.NewFloat(allExports[i].Fl)
	case exportCustom:
		return allExports[i].C(&allExports[i])
	default:
		panic("not reached")
	}
}

func exportDurationConst(e *export) nitro.Value {
	return NewDuration(time.Duration(e.I))
}
