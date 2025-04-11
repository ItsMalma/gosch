package gosch

import "reflect"

type Float64Rule func(value float64) error

type Float64Schema struct {
	nilable bool
	rules   []Float64Rule
}

// Float64 validate data type of the input.
// If the input is not an float, it will return an error.
func Float64() Float64Schema {
	return Float64Schema{
		nilable: false,
		rules:   []Float64Rule{},
	}
}

// Nil will pass nil input.
func (float64Schema Float64Schema) Nil() Float64Schema {
	float64Schema.nilable = true
	return float64Schema
}

// MinValue validate the minimum value of an float.
// If the input is less than the minimum value, it will return an error.
func (float64Schema Float64Schema) MinValue(min float64) Float64Schema {
	float64Schema.rules = append(float64Schema.rules, func(value float64) error {
		if value < min {
			return RuleError{
				Name:   RuleMinValue,
				Value:  value,
				Params: []any{min},
			}
		}
		return nil
	})

	return float64Schema
}

// MaxValue validate the maximum value of an float.
// If the input is greater than the maximum value, it will return an error.
func (float64Schema Float64Schema) MaxValue(max float64) Float64Schema {
	float64Schema.rules = append(float64Schema.rules, func(value float64) error {
		if value > max {
			return RuleError{
				Name:   RuleMaxValue,
				Value:  value,
				Params: []any{max},
			}
		}
		return nil
	})

	return float64Schema
}

func (float64Schema Float64Schema) Validate(value any) error {
	reflectedValue := reflect.ValueOf(value)
	reflectedType := reflect.TypeOf(value)

	if reflectedType == nil {
		if float64Schema.nilable {
			return nil
		}

		return TypeError{
			Expected: "float64",
			Actual:   "nil",
		}
	}

	if reflectedValue.Kind() == reflect.Ptr {
		reflectedValue = reflectedValue.Elem()
		reflectedType = reflectedType.Elem()
	}

	if reflectedType.Kind() != reflect.Float64 {
		return TypeError{
			Expected: "float64",
			Actual:   reflectedType.Kind().String(),
		}
	}

	float64Value := float64(reflectedValue.Float())

	for _, rule := range float64Schema.rules {
		if err := rule(float64Value); err != nil {
			return err
		}
	}

	return nil
}
