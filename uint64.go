package gosch

import "reflect"

type Uint64Rule func(value uint64) error

type Uint64Schema struct {
	nilable bool
	rules   []Uint64Rule
}

// Uint64 validate data type of the input.
// If the input is not an uint64, it will return an error.
func Uint64() Uint64Schema {
	return Uint64Schema{
		nilable: false,
		rules:   []Uint64Rule{},
	}
}

// Nil will pass nil input.
func (uint64Schema Uint64Schema) Nil() Uint64Schema {
	uint64Schema.nilable = true
	return uint64Schema
}

// MinValue validate the minimum value of an uint64.
// If the input is less than the minimum value, it will return an error.
func (uint64Schema Uint64Schema) MinValue(min uint64) Uint64Schema {
	uint64Schema.rules = append(uint64Schema.rules, func(value uint64) error {
		if value < min {
			return RuleError{
				Name:   RuleMinValue,
				Value:  value,
				Params: []any{min},
			}
		}
		return nil
	})

	return uint64Schema
}

// MaxValue validate the maximum value of an uint64.
// If the input is greater than the maximum value, it will return an error.
func (uint64Schema Uint64Schema) MaxValue(max uint64) Uint64Schema {
	uint64Schema.rules = append(uint64Schema.rules, func(value uint64) error {
		if value > max {
			return RuleError{
				Name:   RuleMaxValue,
				Value:  value,
				Params: []any{max},
			}
		}
		return nil
	})

	return uint64Schema
}

func (uint64Schema Uint64Schema) Validate(value any) error {
	reflectedValue := reflect.ValueOf(value)
	reflectedType := reflect.TypeOf(value)

	if reflectedType == nil {
		if uint64Schema.nilable {
			return nil
		}

		return TypeError{
			Expected: "uint64",
			Actual:   "nil",
		}
	}

	if reflectedValue.Kind() == reflect.Ptr {
		reflectedValue = reflectedValue.Elem()
		reflectedType = reflectedType.Elem()
	}

	if reflectedType.Kind() != reflect.Uint64 {
		return TypeError{
			Expected: "uint64",
			Actual:   reflectedType.Kind().String(),
		}
	}

	uint64Value := uint64(reflectedValue.Uint())

	for _, rule := range uint64Schema.rules {
		if err := rule(uint64Value); err != nil {
			return err
		}
	}

	return nil
}
