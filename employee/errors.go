package employee

import "errors"

var (
	ErrEmployeeNotFound    = errors.New("employee not found")
	ErrUnionMemberNotFound = errors.New("union member not found")
	ErrNoTimeCard          = errors.New("timecard not found")
)
