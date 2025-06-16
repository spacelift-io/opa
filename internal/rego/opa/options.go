package opa

import (
	"io"
	"time"

	"github.com/spacelift-io/opa/ast"
	"github.com/spacelift-io/opa/metrics"
	"github.com/spacelift-io/opa/topdown/builtins"
	"github.com/spacelift-io/opa/topdown/cache"
	"github.com/spacelift-io/opa/topdown/print"
)

// Result holds the evaluation result.
type Result struct {
	Result []byte
}

// EvalOpts define options for performing an evaluation.
type EvalOpts struct {
	Input                       *interface{}
	Metrics                     metrics.Metrics
	Entrypoint                  int32
	Time                        time.Time
	Seed                        io.Reader
	InterQueryBuiltinCache      cache.InterQueryCache
	InterQueryBuiltinValueCache cache.InterQueryValueCache
	NDBuiltinCache              builtins.NDBCache
	PrintHook                   print.Hook
	Capabilities                *ast.Capabilities
}
