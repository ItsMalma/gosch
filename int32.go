package gosch

import "reflect"

type Int32Rule func(value int32) error

type Int32Schema struct {
	nilable bool
	rules   []Int32Rule
}

// Int32 validate data type of the input.
// If the input is not an int32, it will return an error.
func Int32() Int32Schema {
	return Int32Schema{
		nilable: false,
		rules:   []Int32Rule{},
	}
}

// Nil will pass nil input.
func (int32Schema Int32Schema) Nil() Int32Schema {
	int32Schema.nilable = true
	return int32Schema
}

// MinValue validate the minimum value of an int32.
// If the input is less than the minimum value, it will return an error.
func (int32Schema Int32Schema) MinValue(min int32) Int32Schema {
	int32Schema.rules = append(int32Schema.rules, func(value int32) error {
		if value < min {
			return RuleError{
				Name:   RuleMinValue,
				Value:  value,
				Params: []any{min},
			}
		}
		return nil
	})

	return int32Schema
}

// MaxValue validate the maximum value of an int32.
// If the input is greater than the maximum value, it will return an error.
func (int32Schema Int32Schema) MaxValue(max int32) Int32Schema {
	int32Schema.rules = append(int32Schema.rules, func(value int32) error {
		if value > max {
			return RuleError{
				Name:   RuleMaxValue,
				Value:  value,
				Params: []any{max},
			}
		}
		return nil
	})

	return int32Schema
}

func (int32Schema Int32Schema) Validate(value any) error {
	reflectedValue := reflect.ValueOf(value)
	reflectedType := reflect.TypeOf(value)

	if reflectedType == nil {
		if int32Schema.nilable {
			return nil
		}

		return TypeError{
			Expected: "int32",
			Actual:   "nil",
		}
	}

	if reflectedValue.Kind() == reflect.Ptr {
		reflectedValue = reflectedValue.Elem()
		reflectedType = reflectedType.Elem()
	}

	if reflectedType.Kind() != reflect.Int32 {
		return TypeError{
			Expected: "int32",
			Actual:   reflectedType.Kind().String(),
		}
	}

	int32Value := int32(reflectedValue.Int())

	for _, rule := range int32Schema.rules {
		if err := rule(int32Value); err != nil {
			return err
		}
	}

	return nil
}
