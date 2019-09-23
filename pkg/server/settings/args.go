package settings

import (
	"bytes"
	"fmt"
)

// Args contains the startup arguments to instantiate Sygen.
type Args struct {
	ExecutionTime int
	RequestRate   int
	Endpoint      string
	Payload       string
	Headers       map[string]string
}

// String produces a stringified version of the arguments for debugging.
func (a *Args) String() string {
	buf := &bytes.Buffer{}

	_, _ = fmt.Fprintf(buf, "ExecutionTime: %d\n", a.ExecutionTime)
	_, _ = fmt.Fprintf(buf, "RequestRate: %d\n", a.RequestRate)
	_, _ = fmt.Fprintf(buf, "Endpoint: %s\n", a.Endpoint)
	_, _ = fmt.Fprintf(buf, "Payload: %s\n", a.Payload)
	_, _ = fmt.Fprintf(buf, "Headers: %v\n", a.Headers)

	return buf.String()
}
