package main

import (
	"fmt"
	"log"
	"runtime"
)

type sqrtError struct {
	value    float64
	function string
	file     string
	line     int
	msg      string
}

func newSqrtError(v float64, m string) sqrtError {
	pc, fileName, lineNmbr, ok := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	name := f.Name()
	if !ok {
		return sqrtError{value: v, msg: m}
	}
	return sqrtError{value: v,
		function: name,
		file:     fileName,
		line:     lineNmbr,
		msg:      m}
}

func (err sqrtError) Error() string {
	return fmt.Sprintf("file=%s, func=%s, line=%d, trying to use '%g', msg=%s\n",
		err.file, err.function, err.line, err.value, err.msg)
}

func main() {
	if _, err := sqrt(-12); err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, newSqrtError(f, "sqrt for negative number")
	}
	return 2, nil
}
