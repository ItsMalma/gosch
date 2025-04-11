package gosch

import (
	"reflect"
)

type StringRule func(value string) error

type StringSchema struct {
	nilable bool
	rules   []StringRule
}

// String validate data type of the input.
// If the input is not a string, it will return an error.
func String() StringSchema {
	return StringSchema{
		nilable: false,
		rules:   []StringRule{},
	}
}

// Nil will pass nil input.
func (stringSchema StringSchema) Nil() StringSchema {
	stringSchema.nilable = true
	return stringSchema
}

// NotEmpty validate that a string is not empty.
// If the input is empty, it will return an error.
func (stringSchema StringSchema) NotEmpty() StringSchema {
	stringSchema.rules = append(stringSchema.rules, func(value string) error {
		if value == "" {
			return RuleError{
				Name:  RuleNotEmpty,
				Value: value,
			}
		}
		return nil
	})

	return stringSchema
}

// MinLength validate the minimum length of a string.
// If the input is less than the minimum length, it will return an error.
func (stringSchema StringSchema) MinLength(length uint) StringSchema {
	stringSchema.rules = append(stringSchema.rules, func(value string) error {
		if len(value) < int(length) {
			return RuleError{
				Name:   RuleMinLength,
				Value:  value,
				Params: []any{length},
			}
		}
		return nil
	})

	return stringSchema
}

// MaxLength validate the maximum length of a string.
// If the input is greater than the maximum length, it will return an error.
func (stringSchema StringSchema) MaxLength(length uint) StringSchema {
	stringSchema.rules = append(stringSchema.rules, func(value string) error {
		if len(value) > int(length) {
			return RuleError{
				Name:   RuleMaxLength,
				Value:  value,
				Params: []any{length},
			}
		}
		return nil
	})

	return stringSchema
}

func (stringSchema StringSchema) Validate(value any) error {
	reflectedValue := reflect.ValueOf(value)
	reflectedType := reflect.TypeOf(value)

	if reflectedType == nil {
		if stringSchema.nilable {
			return nil
		}

		return TypeError{
			Expected: "string",
			Actual:   "nil",
		}
	}

	if reflectedValue.Kind() == reflect.Ptr {
		reflectedValue = reflectedValue.Elem()
		reflectedType = reflectedType.Elem()
	}

	if reflectedType.Kind() != reflect.String {
		return TypeError{
			Expected: "string",
			Actual:   reflectedType.Kind().String(),
		}
	}

	stringValue := reflectedValue.String()

	for _, rule := range stringSchema.rules {
		if err := rule(stringValue); err != nil {
			return err
		}
	}

	return nil
}
