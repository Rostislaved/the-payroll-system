package employee

import "errors"

var (
	ErrNoEmployeeFound = errors.New("employee not found")
	ErrNoTimeCard      = errors.New("timecard not found")
)
