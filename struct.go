package gosch

import "reflect"

type StructSchema struct {
	nilable bool
	fields  map[string]Schema
}

// Struct validate data type of the input.
// If the input is not a struct, it will return an error.
func Struct() StructSchema {
	return StructSchema{
		fields: map[string]Schema{},
	}
}

// Nil will pass nil input.
func (structSchema StructSchema) Nil() StructSchema {
	structSchema.nilable = true
	return structSchema
}

// Field validate the field of a struct.
// If the field is not in the struct, it will return an error.
// If the field is not match the schema, it will return an error.
func (structSchema StructSchema) Field(name string, schema Schema) StructSchema {
	structSchema.fields[name] = schema

	return structSchema
}

func (structSchema StructSchema) Validate(value any) error {
	reflectedValue := reflect.ValueOf(value)
	reflectedType := reflect.TypeOf(value)

	if reflectedType == nil {
		if structSchema.nilable {
			return nil
		}

		return TypeError{
			Expected: "struct",
			Actual:   "nil",
		}
	}

	if reflectedValue.Kind() == reflect.Ptr {
		reflectedValue = reflectedValue.Elem()
		reflectedType = reflectedType.Elem()
	}

	if reflectedType.Kind() != reflect.Struct {
		return TypeError{
			Expected: "struct",
			Actual:   reflectedType.Kind().String(),
		}
	}

	for fieldName, fieldSchema := range structSchema.fields {
		fieldValue := reflectedValue.FieldByName(fieldName)

		if !fieldValue.IsValid() {
			return RuleError{
				Name:   RuleField,
				Value:  fieldValue,
				Params: []any{fieldName},
			}
		}

		if err := fieldSchema.Validate(fieldValue.Interface()); err != nil {
			return FieldError{
				Name:  fieldName,
				Value: fieldValue,
				Err:   err,
			}
		}
	}

	return nil
}
