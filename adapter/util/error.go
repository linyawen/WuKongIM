package util

import (
	"fmt"
	"github.com/pkg/errors"
)

func PanicIfErrorf(err error, format string, args ...interface{}) {
	if err != nil {
		panic(errors.Wrap(err, fmt.Sprintf(format, args...)))
	}
}

func PanicIfError(err error) {
	if err != nil {
		panic(errors.Wrap(err, "panic when error"))
	}
}
