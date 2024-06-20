package starlarktest

type PanicReporter struct {
}

// Error implements Reporter.
func (p *PanicReporter) Error(args ...interface{}) {
	panic(args)
}

var (
	_ Reporter = (*PanicReporter)(nil)
)
