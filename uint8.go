package gosch

import "reflect"

type Uint8Rule func(value uint8) error

type Uint8Schema struct {
	nilable bool
	rules   []Uint8Rule
}

// Uint8 validate data type of the input.
// If the input is not an uint8, it will return an error.
func Uint8() Uint8Schema {
	return Uint8Schema{
		nilable: false,
		rules:   []Uint8Rule{},
	}
}

// Nil will pass nil input.
func (uint8Schema Uint8Schema) Nil() Uint8Schema {
	uint8Schema.nilable = true
	return uint8Schema
}

// MinValue validate the minimum value of an uint8.
// If the input is less than the minimum value, it will return an error.
func (uint8Schema Uint8Schema) MinValue(min uint8) Uint8Schema {
	uint8Schema.rules = append(uint8Schema.rules, func(value uint8) error {
		if value < min {
			return RuleError{
				Name:   RuleMinValue,
				Value:  value,
				Params: []any{min},
			}
		}
		return nil
	})

	return uint8Schema
}

// MaxValue validate the maximum value of an uint8.
// If the input is greater than the maximum value, it will return an error.
func (uint8Schema Uint8Schema) MaxValue(max uint8) Uint8Schema {
	uint8Schema.rules = append(uint8Schema.rules, func(value uint8) error {
		if value > max {
			return RuleError{
				Name:   RuleMaxValue,
				Value:  value,
				Params: []any{max},
			}
		}
		return nil
	})

	return uint8Schema
}

func (uint8Schema Uint8Schema) Validate(value any) error {
	reflectedValue := reflect.ValueOf(value)
	reflectedType := reflect.TypeOf(value)

	if reflectedType == nil {
		if uint8Schema.nilable {
			return nil
		}

		return TypeError{
			Expected: "uint8",
			Actual:   "nil",
		}
	}

	if reflectedValue.Kind() == reflect.Ptr {
		reflectedValue = reflectedValue.Elem()
		reflectedType = reflectedType.Elem()
	}

	if reflectedType.Kind() != reflect.Uint8 {
		return TypeError{
			Expected: "uint8",
			Actual:   reflectedType.Kind().String(),
		}
	}

	uint8Value := uint8(reflectedValue.Uint())

	for _, rule := range uint8Schema.rules {
		if err := rule(uint8Value); err != nil {
			return err
		}
	}

	return nil
}
