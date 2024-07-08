package lib

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"

	"github.com/dcaiafa/nitro"
	"github.com/dcaiafa/nitro/internal/vm"
	"github.com/dcaiafa/nitro/lib/core"
)

type parseCSVOptions struct {
	Columns []int64 `nitro:"columns"`
}

var parseCSVConv core.Value2Structer

var errParseCSVUsage = errors.New(
	`invalid usage. Expected parse_csv(reader, map?)`)

func parseCSV(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	if len(args) != 1 && len(args) != 2 {
		return nil, errParseCSVUsage
	}

	reader, err := nitro.MakeReader(m, args[0])
	if err != nil {
		return nil, errParseCSVUsage
	}

	var opts parseCSVOptions
	if len(args) == 2 {
		err = parseCSVConv.Convert(args[1], &opts)
		if err != nil {
			return nil, err
		}
	}

	var columns []int
	if opts.Columns != nil {
		columns = make([]int, len(opts.Columns))
		for i, colInt64 := range opts.Columns {
			columns[i] = int(colInt64)
		}
	}

	i := csvIter{
		origReader: reader,
		csvReader:  csv.NewReader(reader),
		columns:    columns,
	}
	i.csvReader.ReuseRecord = true

	return []nitro.Value{nitro.NewIterator(i.Next, i.Close, 1)}, nil
}

type csvIter struct {
	origReader io.Reader
	csvReader  *csv.Reader
	columns    []int
}

func (i *csvIter) Next(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	record, err := i.csvReader.Read()
	if err != nil {
		core.CloseReader(i.origReader)
		if err == io.EOF {
			return nil, nil
		} else {
			return nil, err
		}
	}

	var res []nitro.Value
	if i.columns != nil {
		res = make([]nitro.Value, len(i.columns))
		for i, col := range i.columns {
			if col < 0 || col >= len(record) {
				return nil, fmt.Errorf(
					"column %v requested, but csv record has %v columns",
					col, len(record))
			}
			res[i] = nitro.NewString(record[col])
		}
	} else {
		res = make([]nitro.Value, len(record))
		for i, entry := range record {
			res[i] = nitro.NewString(entry)
		}
	}

	return []nitro.Value{nitro.NewArrayFromSlice(res)}, nil
}

func (i *csvIter) Close(vm *nitro.VM) error {
	core.CloseReader(i.origReader)
	return nil
}

func writeCSV(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	iter, err := getIterArg(m, args, 0)
	if err != nil {
		return nil, err
	}

	out, err := getWriterArg(args, 1)
	if err != nil {
		return nil, err
	}

	var records []string

	csvOut := csv.NewWriter(out)

	for {
		vals, err := m.IterNext(iter, 1)
		if err != nil {
			return nil, err
		}
		if vals == nil {
			break
		}

		recordValues, ok := vals[0].(*vm.Array)
		if !ok {
			return nil, fmt.Errorf(
				"iterator must return lists, but returned %v",
				vm.TypeName(vals[0]))
		}

		if records == nil {
			records = make([]string, recordValues.Len())
		} else if len(records) != recordValues.Len() {
			return nil, fmt.Errorf("all csv rows must have the same number of records")
		}

		for i := 0; i < recordValues.Len(); i++ {
			records[i] = recordValues.Get(i).String()
		}

		err = csvOut.Write(records)
		if err != nil {
			return nil, err
		}
	}

	csvOut.Flush()

	if csvOut.Error() != nil {
		return nil, csvOut.Error()
	}

	return nil, nil
}
