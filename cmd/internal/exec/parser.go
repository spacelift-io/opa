package exec

import (
	"io"

	"github.com/spacelift-io/opa/util"
)

type parser interface {
	Parse(io.Reader) (interface{}, error)
}

type utilParser struct {
}

func (utilParser) Parse(r io.Reader) (interface{}, error) {
	bs, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	var x interface{}
	return x, util.Unmarshal(bs, &x)
}
