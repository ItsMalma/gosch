package gosch

import "reflect"

type Uint16Rule func(value uint16) error

type Uint16Schema struct {
	nilable bool
	rules   []Uint16Rule
}

// Uint16 validate data type of the input.
// If the input is not an uint16, it will return an error.
func Uint16() Uint16Schema {
	return Uint16Schema{
		nilable: false,
		rules:   []Uint16Rule{},
	}
}

// Nil will pass nil input.
func (uint16Schema Uint16Schema) Nil() Uint16Schema {
	uint16Schema.nilable = true
	return uint16Schema
}

// MinValue validate the minimum value of an uint16.
// If the input is less than the minimum value, it will return an error.
func (uint16Schema Uint16Schema) MinValue(min uint16) Uint16Schema {
	uint16Schema.rules = append(uint16Schema.rules, func(value uint16) error {
		if value < min {
			return RuleError{
				Name:   RuleMinValue,
				Value:  value,
				Params: []any{min},
			}
		}
		return nil
	})

	return uint16Schema
}

// MaxValue validate the maximum value of an uint16.
// If the input is greater than the maximum value, it will return an error.
func (uint16Schema Uint16Schema) MaxValue(max uint16) Uint16Schema {
	uint16Schema.rules = append(uint16Schema.rules, func(value uint16) error {
		if value > max {
			return RuleError{
				Name:   RuleMaxValue,
				Value:  value,
				Params: []any{max},
			}
		}
		return nil
	})

	return uint16Schema
}

func (uint16Schema Uint16Schema) Validate(value any) error {
	reflectedValue := reflect.ValueOf(value)
	reflectedType := reflect.TypeOf(value)

	if reflectedType == nil {
		if uint16Schema.nilable {
			return nil
		}

		return TypeError{
			Expected: "uint16",
			Actual:   "nil",
		}
	}

	if reflectedValue.Kind() == reflect.Ptr {
		reflectedValue = reflectedValue.Elem()
		reflectedType = reflectedType.Elem()
	}

	if reflectedType.Kind() != reflect.Uint16 {
		return TypeError{
			Expected: "uint16",
			Actual:   reflectedType.Kind().String(),
		}
	}

	uint16Value := uint16(reflectedValue.Uint())

	for _, rule := range uint16Schema.rules {
		if err := rule(uint16Value); err != nil {
			return err
		}
	}

	return nil
}
