package lib

import (
	"errors"
	gosort "sort"

	"github.com/dcaiafa/nitro"
	"github.com/dcaiafa/nitro/internal/vm"
)

var (
	errNotEnoughArgs = errors.New("not enough arguments")
	errTooManyArgs   = errors.New("too many arguments")
	errTakesNoArgs   = errors.New("function does not take arguments")
)

type fnInfo struct {
	Package string
	Name    string
	Func    func(vm *nitro.VM, args []nitro.Value, nret int) ([]nitro.Value, error)
}

var allFuncs = []fnInfo{
	{"", "$exec", exec},
	{"", "args", args},
	{"", "avg", avg},
	{"", "batch", batch},
	{"", "buf", buf},
	{"", "close", closep},
	{"", "color", color},
	{"", "copy_file", copyFile},
	{"", "count", count},
	{"", "create", create},
	{"", "create_temp", createTemp},
	{"", "delete", delete_},
	{"", "discard", discard},
	{"", "do", dop},
	{"", "dur", dur},
	{"", "env", env},
	{"", "file_exists", fileExists},
	{"", "filter", filter},
	{"", "first", first},
	{"", "flatten", flatten},
	{"", "from_crlf", fromCRLF},
	{"", "group", group},
	{"", "has", has},
	{"", "hist", hist},
	{"", "home_dir", homeDir},
	{"", "is_dir", is_dir},
	{"", "join", join},
	{"", "len", Len},
	{"", "lines", lines},
	{"", "log", log},
	{"", "logf", logf},
	{"", "ls", ls},
	{"", "map", mapp},
	{"", "mapreduce", mapreduce},
	{"", "max", max},
	{"", "min", min},
	{"", "mkdir_all", mkdirAll},
	{"", "move_file", moveFile},
	{"", "narg", narg},
	{"", "nonl", nonl},
	{"", "now", now},
	{"", "open", open},
	{"", "parse_base64", parseBase64},
	{"", "parse_csv", parseCSV},
	{"", "parse_float", parseFloat},
	{"", "parse_int", parseInt},
	{"", "parse_json", parseJSON},
	{"", "parse_time", parseTime},
	{"", "pop_stdout", popStdout},
	{"", "print", print},
	{"", "print_table", printTable},
	{"", "printf", printf},
	{"", "probe", probe},
	{"", "prompt", prompt},
	{"", "push_stdout", pushStdout},
	{"", "range", range_},
	{"", "read", read},
	{"", "read_file", readFile},
	{"", "reduce", reduce},
	{"", "regex", regex},
	{"", "remove_all", removeAll},
	{"", "remove_file", removeFile},
	{"", "seek", seek},
	{"", "sha1", sha1},
	{"", "shuffle", shuffle},
	{"", "skip", skip},
	{"", "sleep", sleep},
	{"", "sort", sort},
	{"", "sprintf", sprintf},
	{"", "stat", stat},
	{"", "stderr", stderr},
	{"", "stdin", stdin},
	{"", "stdout", stdout},
	{"", "sum", sum},
	{"", "take", take},
	{"", "time_from_unix", timeFromUnix},
	{"", "to_array", toArray},
	{"", "to_base64", toBase64},
	{"", "to_bool", toBool},
	{"", "to_crlf", toCRLF},
	{"", "to_hex", toHex},
	{"", "to_int", toInt},
	{"", "to_json", toJSON},
	{"", "to_map", toMap},
	{"", "to_string", toString},
	{"", "type", typep},
	{"", "unique", unique},
	{"", "write_file", writeFile},
	{"co", "run_with_timeout", runWithTimeout},
	{"co", "start", start},
	{"exec", "exec", exec},
	{"math", "trunc", mathTrunc},
	{"path", "base", pathBase},
	{"path", "clean", pathClean},
	{"path", "dir", pathDir},
	{"path", "ext", pathExt},
	{"path", "from_slash", pathFromSlash},
	{"path", "join", pathJoin},
	{"path", "match", pathMatch},
	{"path", "to_slash", pathToSlash},
}

type funcRegistry struct {
}

func NewFuncRegistry() vm.NativeFnRegistry {
	return &funcRegistry{}
}

func (r *funcRegistry) IsValidPackage(pkg string) bool {
	i := gosort.Search(len(allFuncs), func(i int) bool {
		return allFuncs[i].Package >= pkg
	})
	return i < len(allFuncs) && allFuncs[i].Package == pkg
}

func (r *funcRegistry) GetNativeFn(pkg, name string) *vm.NativeFn {
	i := gosort.Search(len(allFuncs), func(i int) bool {
		if allFuncs[i].Package > pkg {
			return true
		} else if allFuncs[i].Package == pkg && allFuncs[i].Name >= name {
			return true
		}
		return false
	})
	if i >= len(allFuncs) ||
		allFuncs[i].Package != pkg ||
		allFuncs[i].Name != name {
		return nil
	}

	return vm.NewNativeFn(allFuncs[i].Func)
}
