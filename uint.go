package gosch

import "reflect"

type UintRule func(value uint) error

type UintSchema struct {
	nilable bool
	rules   []UintRule
}

// Uint validate data type of the input.
// If the input is not an uint, it will return an error.
func Uint() UintSchema {
	return UintSchema{
		nilable: false,
		rules:   []UintRule{},
	}
}

// Nil will pass nil input.
func (uintSchema UintSchema) Nil() UintSchema {
	uintSchema.nilable = true
	return uintSchema
}

// MinValue validate the minimum value of an uint.
// If the input is less than the minimum value, it will return an error.
func (uintSchema UintSchema) MinValue(min uint) UintSchema {
	uintSchema.rules = append(uintSchema.rules, func(value uint) error {
		if value < min {
			return RuleError{
				Name:   RuleMinValue,
				Value:  value,
				Params: []any{min},
			}
		}
		return nil
	})

	return uintSchema
}

// MaxValue validate the maximum value of an uint.
// If the input is greater than the maximum value, it will return an error.
func (uintSchema UintSchema) MaxValue(max uint) UintSchema {
	uintSchema.rules = append(uintSchema.rules, func(value uint) error {
		if value > max {
			return RuleError{
				Name:   RuleMaxValue,
				Value:  value,
				Params: []any{max},
			}
		}
		return nil
	})

	return uintSchema
}

func (uintSchema UintSchema) Validate(value any) error {
	reflectedValue := reflect.ValueOf(value)
	reflectedType := reflect.TypeOf(value)

	if reflectedType == nil {
		if uintSchema.nilable {
			return nil
		}

		return TypeError{
			Expected: "uint",
			Actual:   "nil",
		}
	}

	if reflectedValue.Kind() == reflect.Ptr {
		reflectedValue = reflectedValue.Elem()
		reflectedType = reflectedType.Elem()
	}

	if reflectedType.Kind() != reflect.Uint {
		return TypeError{
			Expected: "uint",
			Actual:   reflectedType.Kind().String(),
		}
	}

	uintValue := uint(reflectedValue.Uint())

	for _, rule := range uintSchema.rules {
		if err := rule(uintValue); err != nil {
			return err
		}
	}

	return nil
}
