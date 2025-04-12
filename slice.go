package gosch

import "reflect"

type SliceRule func(value []any) error

type SliceSchema struct {
	nilable bool
	element Schema
	rules   []SliceRule
}

// Slice validate data type of the input.
// If the input is not a slice, it will return an error.
func Slice() SliceSchema {
	return SliceSchema{
		nilable: false,
		element: nil,
		rules:   []SliceRule{},
	}
}

// Nil will pass nil input.
func (sliceSchema SliceSchema) Nil() SliceSchema {
	sliceSchema.nilable = true
	return sliceSchema
}

// Element validate the element of a slice.
// If the element is not match the schema, it will return an error.
func (sliceSchema SliceSchema) Element(schema Schema) SliceSchema {
	sliceSchema.element = schema

	return sliceSchema
}

// MinLength validate the minimum length of a slice.
// If the input is less than the minimum length, it will return an error.
func (sliceSchema SliceSchema) MinLength(length uint) SliceSchema {
	sliceSchema.rules = append(sliceSchema.rules, func(value []any) error {
		if len(value) < int(length) {
			return RuleError{
				Name:   RuleMinLength,
				Value:  value,
				Params: []any{length},
			}
		}
		return nil
	})

	return sliceSchema
}

// MaxLength validate the maximum length of a slice.
// If the input is greater than the maximum length, it will return an error.
func (sliceSchema SliceSchema) MaxLength(length uint) SliceSchema {
	sliceSchema.rules = append(sliceSchema.rules, func(value []any) error {
		if len(value) > int(length) {
			return RuleError{
				Name:   RuleMaxLength,
				Value:  value,
				Params: []any{length},
			}
		}
		return nil
	})

	return sliceSchema
}

func (sliceSchema SliceSchema) Validate(value any) error {
	reflectedValue := reflect.ValueOf(value)
	reflectedType := reflect.TypeOf(value)

	if reflectedType == nil {
		if sliceSchema.nilable {
			return nil
		}

		return TypeError{
			Expected: "slice",
			Actual:   "nil",
		}
	}

	if reflectedValue.Kind() == reflect.Ptr {
		reflectedValue = reflectedValue.Elem()
		reflectedType = reflectedType.Elem()
	}

	if reflectedType.Kind() != reflect.Slice {
		return TypeError{
			Expected: "slice",
			Actual:   reflectedType.Kind().String(),
		}
	}

	sliceValue := make([]any, reflectedValue.Len())

	i := 0
	for _, element := range reflectedValue.Seq2() {
		elementValue := element.Interface()

		if err := sliceSchema.element.Validate(elementValue); err != nil {
			return ElementError{
				Index: i,
				Value: element,
				Err:   err,
			}
		}

		sliceValue[i] = elementValue

		i++
	}

	for _, rule := range sliceSchema.rules {
		if err := rule(sliceValue); err != nil {
			return err
		}
	}

	return nil
}
