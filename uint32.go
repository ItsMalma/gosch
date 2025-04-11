package gosch

import "reflect"

type Uint32Rule func(value uint32) error

type Uint32Schema struct {
	nilable bool
	rules   []Uint32Rule
}

// Uint32 validate data type of the input.
// If the input is not an uint32, it will return an error.
func Uint32() Uint32Schema {
	return Uint32Schema{
		nilable: false,
		rules:   []Uint32Rule{},
	}
}

// Nil will pass nil input.
func (uint32Schema Uint32Schema) Nil() Uint32Schema {
	uint32Schema.nilable = true
	return uint32Schema
}

// MinValue validate the minimum value of an uint32.
// If the input is less than the minimum value, it will return an error.
func (uint32Schema Uint32Schema) MinValue(min uint32) Uint32Schema {
	uint32Schema.rules = append(uint32Schema.rules, func(value uint32) error {
		if value < min {
			return RuleError{
				Name:   RuleMinValue,
				Value:  value,
				Params: []any{min},
			}
		}
		return nil
	})

	return uint32Schema
}

// MaxValue validate the maximum value of an uint32.
// If the input is greater than the maximum value, it will return an error.
func (uint32Schema Uint32Schema) MaxValue(max uint32) Uint32Schema {
	uint32Schema.rules = append(uint32Schema.rules, func(value uint32) error {
		if value > max {
			return RuleError{
				Name:   RuleMaxValue,
				Value:  value,
				Params: []any{max},
			}
		}
		return nil
	})

	return uint32Schema
}

func (uint32Schema Uint32Schema) Validate(value any) error {
	reflectedValue := reflect.ValueOf(value)
	reflectedType := reflect.TypeOf(value)

	if reflectedType == nil {
		if uint32Schema.nilable {
			return nil
		}

		return TypeError{
			Expected: "uint32",
			Actual:   "nil",
		}
	}

	if reflectedValue.Kind() == reflect.Ptr {
		reflectedValue = reflectedValue.Elem()
		reflectedType = reflectedType.Elem()
	}

	if reflectedType.Kind() != reflect.Uint32 {
		return TypeError{
			Expected: "uint32",
			Actual:   reflectedType.Kind().String(),
		}
	}

	uint32Value := uint32(reflectedValue.Uint())

	for _, rule := range uint32Schema.rules {
		if err := rule(uint32Value); err != nil {
			return err
		}
	}

	return nil
}
