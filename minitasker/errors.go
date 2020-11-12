package minitasker

import (
	"errors"
	"fmt"
)

var (
	errAddDupNameTask = "duplicated task name"
	errNoTask = "no such task"
)

func taskError(errStr, taskName string) error {
	return errors.New(fmt.Sprintf("%v: %v", errStr, taskName))
}
