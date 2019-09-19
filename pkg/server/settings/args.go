package settings

import (
	"bytes"
	"fmt"
)



// Args contains the startup arguments to instantiate Sygen.
type Args struct { 
	ExecutionTime int
	RequestRate int
}



// String produces a stringified version of the arguments for debugging.
func (a *Args) String() string {
	buf := &bytes.Buffer{}

	_, _ = fmt.Fprintf(buf, "ExecutionTime: %d\n", a.ExecutionTime)

	return buf.String()
}