package gosch

import "reflect"

type Int8Rule func(value int8) error

type Int8Schema struct {
	nilable bool
	rules   []Int8Rule
}

// Int8 validate data type of the input.
// If the input is not an int8, it will return an error.
func Int8() Int8Schema {
	return Int8Schema{
		nilable: false,
		rules:   []Int8Rule{},
	}
}

// Nil will pass nil input.
func (int8Schema Int8Schema) Nil() Int8Schema {
	int8Schema.nilable = true
	return int8Schema
}

// MinValue validate the minimum value of an int8.
// If the input is less than the minimum value, it will return an error.
func (int8Schema Int8Schema) MinValue(min int8) Int8Schema {
	int8Schema.rules = append(int8Schema.rules, func(value int8) error {
		if value < min {
			return RuleError{
				Name:   RuleMinValue,
				Value:  value,
				Params: []any{min},
			}
		}
		return nil
	})

	return int8Schema
}

// MaxValue validate the maximum value of an int8.
// If the input is greater than the maximum value, it will return an error.
func (int8Schema Int8Schema) MaxValue(max int8) Int8Schema {
	int8Schema.rules = append(int8Schema.rules, func(value int8) error {
		if value > max {
			return RuleError{
				Name:   RuleMaxValue,
				Value:  value,
				Params: []any{max},
			}
		}
		return nil
	})

	return int8Schema
}

func (int8Schema Int8Schema) Validate(value any) error {
	reflectedValue := reflect.ValueOf(value)
	reflectedType := reflect.TypeOf(value)

	if reflectedType == nil {
		if int8Schema.nilable {
			return nil
		}

		return TypeError{
			Expected: "int8",
			Actual:   "nil",
		}
	}

	if reflectedValue.Kind() == reflect.Ptr {
		reflectedValue = reflectedValue.Elem()
		reflectedType = reflectedType.Elem()
	}

	if reflectedType.Kind() != reflect.Int8 {
		return TypeError{
			Expected: "int8",
			Actual:   reflectedType.Kind().String(),
		}
	}

	int8Value := int8(reflectedValue.Int())

	for _, rule := range int8Schema.rules {
		if err := rule(int8Value); err != nil {
			return err
		}
	}

	return nil
}
