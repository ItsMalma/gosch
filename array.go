package gosch

import "reflect"

type ArraySchema struct {
	nilable bool
	element Schema
	length  int
}

// Array validate data type of the input.
// If the input is not an array, it will return an error.
func Array() ArraySchema {
	return ArraySchema{
		element: nil,
		length:  0,
	}
}

// Nil will pass nil input.
func (arraySchema ArraySchema) Nil() ArraySchema {
	arraySchema.nilable = true
	return arraySchema
}

// Element validate the element of an array.
// If the element is not match the schema, it will return an error.
func (arraySchema ArraySchema) Element(schema Schema) ArraySchema {
	arraySchema.element = schema

	return arraySchema
}

// Length validate the length of an array.
// If the input is not match the length, it will return an error.
func (arraySchema ArraySchema) Length(length int) ArraySchema {
	if length < 0 {
		panic("array length must be greater than or equal to 0")
	}

	arraySchema.length = length

	return arraySchema
}

func (arraySchema ArraySchema) Validate(value any) error {
	reflectedValue := reflect.ValueOf(value)
	reflectedType := reflect.TypeOf(value)

	if reflectedType == nil {
		if arraySchema.nilable {
			return nil
		}

		return TypeError{
			Expected: "array",
			Actual:   "nil",
		}
	}

	if reflectedValue.Kind() == reflect.Ptr {
		reflectedValue = reflectedValue.Elem()
		reflectedType = reflectedType.Elem()
	}

	if reflectedType.Kind() != reflect.Array {
		return TypeError{
			Expected: "array",
			Actual:   reflectedType.Kind().String(),
		}
	}

	if reflectedValue.Len() != arraySchema.length {
		return RuleError{
			Name:   RuleLength,
			Value:  reflectedValue,
			Params: []any{arraySchema.length},
		}
	}

	i := 0
	for _, element := range reflectedValue.Seq2() {
		if err := arraySchema.element.Validate(element.Interface()); err != nil {
			return ElementError{
				Index: i,
				Value: element,
				Err:   err,
			}
		}
		i++
	}

	return nil
}
