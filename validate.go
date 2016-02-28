package validate

import (
	"fmt"
	"reflect"
	"strings"
)

const (
	tagSep      = ";"
	validateTag = "validate"
)

// Validate validates the json finding the "validate" tag.
func Validate(in interface{}) error {
	// TODO(Issue #1): Add nested struct validation
	// TODO(Issue #2): Add validation for emails
	value := reflect.Indirect(reflect.ValueOf(in))
	typeOfT := value.Type()

	for i := 0; i < value.NumField(); i++ {
		field := typeOfT.Field(i)
		tags := strings.Split(field.Tag.Get(validateTag), tagSep)
		for _, tag := range tags {
			switch tag {
			case "":
				// Because strings.Split likes to return "" when there is nothing to
				// split. It will just end up going to default and erroring :(
				break
			case "required":
				if err := required(value.Field(i), field); err != nil {
					return err
				}
				break
			default:
				return fmt.Errorf("unknown tag found: %s", tag)
			}
		}
	}
	return nil
}

// required checks the Zero state of the input interface. For example if the
// input is an int, the zero state is 0, and if the input is a string, zero
// state is ""
func required(field reflect.Value, structField reflect.StructField) error {
	if field.Interface() == reflect.Zero(field.Type()).Interface() {
		return fmt.Errorf("%s must be set", structField.Name)
	}
	return nil
}
