package gosch

import "reflect"

type MapRule func(value map[any]any) error

type MapSchema struct {
	nilable bool
	key     Schema
	element Schema
	rules   []MapRule
}

// Map validate data type of the input.
// If the input is not a map key to element, it will return an error.
func Map(key, element Schema) MapSchema {
	return MapSchema{
		nilable: false,
		key:     key,
		element: element,
		rules:   []MapRule{},
	}
}

// Key validate the key of a map.
// If the input key is not match the key schema, it will return an error.
func (mapSchema MapSchema) Key(keySchema Schema) MapSchema {
	mapSchema.key = keySchema
	return mapSchema
}

// Element validate the element of a map.
// If the input element is not match the element schema, it will return an error.
func (mapSchema MapSchema) Element(elementSchema Schema) MapSchema {
	mapSchema.element = elementSchema
	return mapSchema
}

// Nil will pass nil input.
func (mapSchema MapSchema) Nil() MapSchema {
	mapSchema.nilable = true
	return mapSchema
}

// MinLength validate the minimum length of a map.
// If the input is less than the minimum length, it will return an error.
func (mapSchema MapSchema) MinLength(length uint) MapSchema {
	mapSchema.rules = append(mapSchema.rules, func(value map[any]any) error {
		if len(value) < int(length) {
			return RuleError{
				Name:   RuleMinLength,
				Value:  value,
				Params: []any{length},
			}
		}
		return nil
	})

	return mapSchema
}

// MaxLength validate the maximum length of a map.
// If the input is greater than the maximum length, it will return an error.
func (mapSchema MapSchema) MaxLength(length uint) MapSchema {
	mapSchema.rules = append(mapSchema.rules, func(value map[any]any) error {
		if len(value) > int(length) {
			return RuleError{
				Name:   RuleMaxLength,
				Value:  value,
				Params: []any{length},
			}
		}
		return nil
	})

	return mapSchema
}

func (mapSchema MapSchema) Validate(value any) error {
	reflectedValue := reflect.ValueOf(value)
	reflectedType := reflect.TypeOf(value)

	if reflectedType == nil {
		if mapSchema.nilable {
			return nil
		}

		return TypeError{
			Expected: "map",
			Actual:   "nil",
		}
	}

	if reflectedValue.Kind() == reflect.Ptr {
		reflectedValue = reflectedValue.Elem()
		reflectedType = reflectedType.Elem()
	}

	if reflectedType.Kind() != reflect.Map {
		return TypeError{
			Expected: "map",
			Actual:   reflectedType.Kind().String(),
		}
	}

	mapValue := make(map[any]any, reflectedValue.Len())

	for key, element := range reflectedValue.Seq2() {
		keyValue := key.Interface()
		elementValue := element.Interface()

		if err := mapSchema.key.Validate(keyValue); err != nil {
			return KeyError{
				Key: keyValue,
				Err: err,
			}
		}

		if err := mapSchema.element.Validate(elementValue); err != nil {
			return ElementError{
				Index: keyValue,
				Value: element,
				Err:   err,
			}
		}

		mapValue[keyValue] = elementValue
	}

	for _, rule := range mapSchema.rules {
		if err := rule(mapValue); err != nil {
			return err
		}
	}

	return nil
}
