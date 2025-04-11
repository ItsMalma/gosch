package gosch

import "reflect"

type IntRule func(value int) error

type IntSchema struct {
	nilable bool
	rules   []IntRule
}

// Int validate data type of the input.
// If the input is not an int, it will return an error.
func Int() IntSchema {
	return IntSchema{
		nilable: false,
		rules:   []IntRule{},
	}
}

// Nil will pass nil input.
func (intSchema IntSchema) Nil() IntSchema {
	intSchema.nilable = true
	return intSchema
}

// MinValue validate the minimum value of an int.
// If the input is less than the minimum value, it will return an error.
func (intSchema IntSchema) MinValue(min int) IntSchema {
	intSchema.rules = append(intSchema.rules, func(value int) error {
		if value < min {
			return RuleError{
				Name:   RuleMinValue,
				Value:  value,
				Params: []any{min},
			}
		}
		return nil
	})

	return intSchema
}

// MaxValue validate the maximum value of an int.
// If the input is greater than the maximum value, it will return an error.
func (intSchema IntSchema) MaxValue(max int) IntSchema {
	intSchema.rules = append(intSchema.rules, func(value int) error {
		if value > max {
			return RuleError{
				Name:   RuleMaxValue,
				Value:  value,
				Params: []any{max},
			}
		}
		return nil
	})

	return intSchema
}

func (intSchema IntSchema) Validate(value any) error {
	reflectedValue := reflect.ValueOf(value)
	reflectedType := reflect.TypeOf(value)

	if reflectedType == nil {
		if intSchema.nilable {
			return nil
		}

		return TypeError{
			Expected: "int",
			Actual:   "nil",
		}
	}

	if reflectedValue.Kind() == reflect.Ptr {
		reflectedValue = reflectedValue.Elem()
		reflectedType = reflectedType.Elem()
	}

	if reflectedType.Kind() != reflect.Int {
		return TypeError{
			Expected: "int",
			Actual:   reflectedType.Kind().String(),
		}
	}

	intValue := int(reflectedValue.Int())

	for _, rule := range intSchema.rules {
		if err := rule(intValue); err != nil {
			return err
		}
	}

	return nil
}
