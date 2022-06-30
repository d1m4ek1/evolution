package newerror

import (
	"fmt"
	"log"
	"runtime"
	"strings"
)

// Struct for structuring an error
type WrapError struct {
	WhereError string
	ErrorStr   error
	File       string
	IsLine     int
}

func (ne *WrapError) wrapError() string {
	ne.File = strings.ReplaceAll(ne.File, "/home/dmitrydvp/go/src/iNote/www/backend/", "")

	return fmt.Sprintf("ERROR:\nWhere error: %s;\nLine: %d;\nFile: %s;\nError: %s;\n<------>",
		ne.WhereError, ne.IsLine, ne.File, ne.ErrorStr)
}

func Wrap(whereError string, errorStr error) {
	_, file, isLine, _ := runtime.Caller(1)
	wrapError := &WrapError{
		WhereError: whereError,
		ErrorStr:   errorStr,
		File:       file,
		IsLine:     isLine,
	}

	log.Println(wrapError.wrapError())
}
