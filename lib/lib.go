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
	{P: "", N: "close", T: exportFunc, F: closep},
	{P: "", N: "color", T: exportFunc, F: color},
	{P: "", N: "count", T: exportFunc, F: count},
	{P: "", N: "discard", T: exportFunc, F: discard},
	{P: "", N: "do", T: exportFunc, F: dop},
	{P: "", N: "env", T: exportFunc, F: env},
	{P: "", N: "filter", T: exportFunc, F: filter},
	{P: "", N: "first", T: exportFunc, F: first},
	{P: "", N: "flatten", T: exportFunc, F: flatten},
	{P: "", N: "from_crlf", T: exportFunc, F: fromCRLF},
	{P: "", N: "get", T: exportFunc, F: get},
	{P: "", N: "group", T: exportFunc, F: group},
	{P: "", N: "has", T: exportFunc, F: has},
	{P: "", N: "hist", T: exportFunc, F: hist},
	{P: "", N: "join", T: exportFunc, F: join},
	{P: "", N: "len", T: exportFunc, F: Len},
	{P: "", N: "lines", T: exportFunc, F: lines},
	{P: "", N: "log", T: exportFunc, F: log},
	{P: "", N: "logf", T: exportFunc, F: logf},
	{P: "", N: "ls", T: exportFunc, F: filepathLs},
	{P: "", N: "map", T: exportFunc, F: imap},
	{P: "", N: "mapreduce", T: exportFunc, F: mapreduce},
	{P: "", N: "max", T: exportFunc, F: max},
	{P: "", N: "min", T: exportFunc, F: min},
	{P: "", N: "narg", T: exportFunc, F: narg},
	{P: "", N: "nonl", T: exportFunc, F: nonl},
	{P: "", N: "open", T: exportFunc, F: fileOpen},
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
	{P: "", N: "reduce", T: exportFunc, F: reduce},
	{P: "", N: "regex", T: exportFunc, F: regex},
	{P: "", N: "sha1", T: exportFunc, F: sha1},
	{P: "", N: "shuffle", T: exportFunc, F: shuffle},
	{P: "", N: "skip", T: exportFunc, F: skip},
	{P: "", N: "sleep", T: exportFunc, F: sleep},
	{P: "", N: "sort", T: exportFunc, F: sort},
	{P: "", N: "sprintf", T: exportFunc, F: sprintf},
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
	{P: "", N: "to_string", T: exportFunc, F: toString},
	{P: "", N: "type", T: exportFunc, F: typep},
	{P: "", N: "unique", T: exportFunc, F: unique},
	{P: "buf", N: "cap", T: exportFunc, F: bufCap},
	{P: "buf", N: "len", T: exportFunc, F: bufLen},
	{P: "buf", N: "new", T: exportFunc, F: bufNew},
	{P: "buf", N: "read", T: exportFunc, F: bufRead},
	{P: "buf", N: "read_byte", T: exportFunc, F: bufReadByte},
	{P: "buf", N: "read_from", T: exportFunc, F: bufReadFrom},
	{P: "buf", N: "read_rune", T: exportFunc, F: bufReadRune},
	{P: "buf", N: "unread_byte", T: exportFunc, F: bufUnreadByte},
	{P: "co", N: "run_with_timeout", T: exportFunc, F: runWithTimeout},
	{P: "co", N: "start", T: exportFunc, F: start},
	{P: "exec", N: "exec", T: exportFunc, F: execExec},
	{P: "exec", N: "with_stderr", T: exportFunc, F: execWithStderr},
	{P: "file", N: "close", T: exportFunc, F: closep},
	{P: "file", N: "create", T: exportFunc, F: fileCreate},
	{P: "file", N: "create_temp", T: exportFunc, F: fileCreateTemp},
	{P: "file", N: "open", T: exportFunc, F: fileOpen},
	{P: "file", N: "read", T: exportFunc, F: fileRead},
	{P: "file", N: "seek", T: exportFunc, F: fileSeek},
	{P: "file", N: "stat", T: exportFunc, F: fileStat},
	{P: "file", N: "write_to", T: exportFunc, F: fileWriteTo},
	{P: "filepath", N: "abs", T: exportFunc, F: filepathAbs},
	{P: "filepath", N: "base", T: exportFunc, F: filepathBase},
	{P: "filepath", N: "clean", T: exportFunc, F: filepathClean},
	{P: "filepath", N: "dir", T: exportFunc, F: filepathDir},
	{P: "filepath", N: "eval_symlinks", T: exportFunc, F: filepathEvalSymlinks},
	{P: "filepath", N: "exists", T: exportFunc, F: filepathExists},
	{P: "filepath", N: "ext", T: exportFunc, F: filepathExt},
	{P: "filepath", N: "from_slash", T: exportFunc, F: filepathFromSlash},
	{P: "filepath", N: "is_abs", T: exportFunc, F: filepathIsAbs},
	{P: "filepath", N: "is_dir", T: exportFunc, F: filepathIsDir},
	{P: "filepath", N: "join", T: exportFunc, F: filepathJoin},
	{P: "filepath", N: "ls", T: exportFunc, F: filepathLs},
	{P: "filepath", N: "match", T: exportFunc, F: filepathMatch},
	{P: "filepath", N: "mkdir", T: exportFunc, F: filepathMkdir},
	{P: "filepath", N: "mkdir_all", T: exportFunc, F: filepathMkdirAll},
	{P: "filepath", N: "mkdir_temp", T: exportFunc, F: filepathMkdirTemp},
	{P: "filepath", N: "rel", T: exportFunc, F: filepathRel},
	{P: "filepath", N: "remove", T: exportFunc, F: filepathRemove},
	{P: "filepath", N: "remove_all", T: exportFunc, F: filepathRemoveAll},
	{P: "filepath", N: "rename", T: exportFunc, F: filepathRename},
	{P: "filepath", N: "split", T: exportFunc, F: filepathSplit},
	{P: "filepath", N: "to_slash", T: exportFunc, F: filepathToSlash},
	{P: "filepath", N: "volume_name", T: exportFunc, F: filepathVolumeName},
	{P: "iter", N: "into", T: exportFunc, F: iter_into},
	{P: "iter", N: "next", T: exportFunc, F: iter_next},
	{P: "list", N: "append", T: exportFunc, F: listAppend},
	{P: "list", N: "append_iter", T: exportFunc, F: listAppendIter},
	{P: "list", N: "find", T: exportFunc, F: listFind},
	{P: "maps", N: "clone", T: exportFunc, F: mapsClone},
	{P: "maps", N: "delete", T: exportFunc, F: mapsDelete},
	{P: "maps", N: "into", T: exportFunc, F: mapsInto},
	{P: "maps", N: "update", T: exportFunc, F: mapsUpdate},
	{P: "math", N: "trunc", T: exportFunc, F: mathTrunc},
	{P: "os", N: "user_home_dir", T: exportFunc, F: osUserHomeDir},
	{P: "os", N: "wd", T: exportFunc, F: osWD},
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
	{P: "time", N: "truncate", T: exportFunc, F: timeTruncate},
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
