package gosch

import "reflect"

type Float32Rule func(value float32) error

type Float32Schema struct {
	nilable bool
	rules   []Float32Rule
}

// Float32 validate data type of the input.
// If the input is not an float, it will return an error.
func Float32() Float32Schema {
	return Float32Schema{
		nilable: false,
		rules:   []Float32Rule{},
	}
}

// Nil will pass nil input.
func (float32Schema Float32Schema) Nil() Float32Schema {
	float32Schema.nilable = true
	return float32Schema
}

// MinValue validate the minimum value of an float.
// If the input is less than the minimum value, it will return an error.
func (float32Schema Float32Schema) MinValue(min float32) Float32Schema {
	float32Schema.rules = append(float32Schema.rules, func(value float32) error {
		if value < min {
			return RuleError{
				Name:   RuleMinValue,
				Value:  value,
				Params: []any{min},
			}
		}
		return nil
	})

	return float32Schema
}

// MaxValue validate the maximum value of an float.
// If the input is greater than the maximum value, it will return an error.
func (float32Schema Float32Schema) MaxValue(max float32) Float32Schema {
	float32Schema.rules = append(float32Schema.rules, func(value float32) error {
		if value > max {
			return RuleError{
				Name:   RuleMaxValue,
				Value:  value,
				Params: []any{max},
			}
		}
		return nil
	})

	return float32Schema
}

func (float32Schema Float32Schema) Validate(value any) error {
	reflectedValue := reflect.ValueOf(value)
	reflectedType := reflect.TypeOf(value)

	if reflectedType == nil {
		if float32Schema.nilable {
			return nil
		}

		return TypeError{
			Expected: "float32",
			Actual:   "nil",
		}
	}

	if reflectedValue.Kind() == reflect.Ptr {
		reflectedValue = reflectedValue.Elem()
		reflectedType = reflectedType.Elem()
	}

	if reflectedType.Kind() != reflect.Float32 {
		return TypeError{
			Expected: "float32",
			Actual:   reflectedType.Kind().String(),
		}
	}

	float32Value := float32(reflectedValue.Float())

	for _, rule := range float32Schema.rules {
		if err := rule(float32Value); err != nil {
			return err
		}
	}

	return nil
}
