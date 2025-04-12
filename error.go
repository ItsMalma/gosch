package gosch

import "fmt"

type TypeError struct {
	Expected string
	Actual   string
}

func (typeError TypeError) Error() string {
	return "expected " + typeError.Expected + ", got " + typeError.Actual
}

type RuleName uint

const (
	RuleNotEmpty RuleName = iota
	RuleLength
	RuleMinLength
	RuleMaxLength
	RuleMinValue
	RuleMaxValue
	RuleField
)

type RuleError struct {
	Name   RuleName
	Value  any
	Params []any
}

func (ruleError RuleError) Error() string {
	switch ruleError.Name {
	case RuleNotEmpty:
		return "value must not be empty"
	case RuleLength:
		return fmt.Sprintf("value must be exactly %d in length", ruleError.Params[0])
	case RuleMinLength:
		return fmt.Sprintf("value must be at least %d in length", ruleError.Params[0])
	case RuleMaxLength:
		return fmt.Sprintf("value must be at most %d in length", ruleError.Params[0])
	case RuleMinValue:
		return fmt.Sprintf("value must be at least %d", ruleError.Params[0])
	case RuleMaxValue:
		return fmt.Sprintf("value must be at most %d", ruleError.Params[0])
	case RuleField:
		return fmt.Sprintf("value must contain field %s", ruleError.Params[0])
	default:
		return "unknown error"
	}
}

type FieldError struct {
	Name  string
	Value any
	Err   error
}

func (fieldError FieldError) Error() string {
	return fmt.Sprintf("field %s: %s", fieldError.Name, fieldError.Err)
}

type ElementError struct {
	Index any
	Value any
	Err   error
}

func (elementError ElementError) Error() string {
	return fmt.Sprintf("index %v: %v", elementError.Index, elementError.Err)
}

type KeyError struct {
	Key any
	Err error
}

func (keyError KeyError) Error() string {
	return fmt.Sprintf("key %v: %v", keyError.Key, keyError.Err)
}
