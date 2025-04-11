package gosch

import "reflect"

type Int16Rule func(value int16) error

type Int16Schema struct {
	nilable bool
	rules   []Int16Rule
}

// Int16 validate data type of the input.
// If the input is not an int16, it will return an error.
func Int16() Int16Schema {
	return Int16Schema{
		nilable: false,
		rules:   []Int16Rule{},
	}
}

// Nil will pass nil input.
func (int16Schema Int16Schema) Nil() Int16Schema {
	int16Schema.nilable = true
	return int16Schema
}

// MinValue validate the minimum value of an int16.
// If the input is less than the minimum value, it will return an error.
func (int16Schema Int16Schema) MinValue(min int16) Int16Schema {
	int16Schema.rules = append(int16Schema.rules, func(value int16) error {
		if value < min {
			return RuleError{
				Name:   RuleMinValue,
				Value:  value,
				Params: []any{min},
			}
		}
		return nil
	})

	return int16Schema
}

// MaxValue validate the maximum value of an int16.
// If the input is greater than the maximum value, it will return an error.
func (int16Schema Int16Schema) MaxValue(max int16) Int16Schema {
	int16Schema.rules = append(int16Schema.rules, func(value int16) error {
		if value > max {
			return RuleError{
				Name:   RuleMaxValue,
				Value:  value,
				Params: []any{max},
			}
		}
		return nil
	})

	return int16Schema
}

func (int16Schema Int16Schema) Validate(value any) error {
	reflectedValue := reflect.ValueOf(value)
	reflectedType := reflect.TypeOf(value)

	if reflectedType == nil {
		if int16Schema.nilable {
			return nil
		}

		return TypeError{
			Expected: "int16",
			Actual:   "nil",
		}
	}

	if reflectedValue.Kind() == reflect.Ptr {
		reflectedValue = reflectedValue.Elem()
		reflectedType = reflectedType.Elem()
	}

	if reflectedType.Kind() != reflect.Int16 {
		return TypeError{
			Expected: "int16",
			Actual:   reflectedType.Kind().String(),
		}
	}

	int16Value := int16(reflectedValue.Int())

	for _, rule := range int16Schema.rules {
		if err := rule(int16Value); err != nil {
			return err
		}
	}

	return nil
}
