package lib

import (
	"io"
	"io/ioutil"

	"github.com/dcaiafa/nitro"
	"github.com/dcaiafa/nitro/lib/core"
)

func read(m *nitro.VM, args []nitro.Value, nRet int) ([]nitro.Value, error) {
	var err error

	if len(args) > 2 {
		return nil, errTooManyArgs
	}

	reader, err := getReaderArg(m, args, 0)
	if err != nil {
		return nil, err
	}

	count := -1

	if len(args) == 2 {
		countArg, err := getIntArg(args, 1)
		if err != nil {
			return nil, err
		}
		count = int(countArg)
	}

	var data []byte
	if count == -1 {
		defer core.CloseReader(reader)
		data, err = ioutil.ReadAll(reader)
		if err != nil {
			return nil, err
		}
	} else {
		buf := make([]byte, count)
		n, err := io.ReadAtLeast(reader, buf, count)
		if err != nil && err != io.ErrUnexpectedEOF {
			return nil, err
		}
		data = buf[:n]
	}

	return []nitro.Value{nitro.NewString(string(data))}, nil
}
