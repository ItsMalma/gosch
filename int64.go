package gosch

import "reflect"

type Int64Rule func(value int64) error

type Int64Schema struct {
	nilable bool
	rules   []Int64Rule
}

// Int64 validate data type of the input.
// If the input is not an int64, it will return an error.
func Int64() Int64Schema {
	return Int64Schema{
		nilable: false,
		rules:   []Int64Rule{},
	}
}

// Nil will pass nil input.
func (int64Schema Int64Schema) Nil() Int64Schema {
	int64Schema.nilable = true
	return int64Schema
}

// MinValue validate the minimum value of an int64.
// If the input is less than the minimum value, it will return an error.
func (int64Schema Int64Schema) MinValue(min int64) Int64Schema {
	int64Schema.rules = append(int64Schema.rules, func(value int64) error {
		if value < min {
			return RuleError{
				Name:   RuleMinValue,
				Value:  value,
				Params: []any{min},
			}
		}
		return nil
	})

	return int64Schema
}

// MaxValue validate the maximum value of an int64.
// If the input is greater than the maximum value, it will return an error.
func (int64Schema Int64Schema) MaxValue(max int64) Int64Schema {
	int64Schema.rules = append(int64Schema.rules, func(value int64) error {
		if value > max {
			return RuleError{
				Name:   RuleMaxValue,
				Value:  value,
				Params: []any{max},
			}
		}
		return nil
	})

	return int64Schema
}

func (int64Schema Int64Schema) Validate(value any) error {
	reflectedValue := reflect.ValueOf(value)
	reflectedType := reflect.TypeOf(value)

	if reflectedType == nil {
		if int64Schema.nilable {
			return nil
		}

		return TypeError{
			Expected: "int64",
			Actual:   "nil",
		}
	}

	if reflectedValue.Kind() == reflect.Ptr {
		reflectedValue = reflectedValue.Elem()
		reflectedType = reflectedType.Elem()
	}

	if reflectedType.Kind() != reflect.Int64 {
		return TypeError{
			Expected: "int64",
			Actual:   reflectedType.Kind().String(),
		}
	}

	int64Value := int64(reflectedValue.Int())

	for _, rule := range int64Schema.rules {
		if err := rule(int64Value); err != nil {
			return err
		}
	}

	return nil
}
