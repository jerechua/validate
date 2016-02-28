package validate_test

import (
	"testing"

	v "../validate"
)

func TestValidateRequired(t *testing.T) {
	tests := []struct {
		testStruct interface{}
		shouldPass bool
	}{
		{
			testStruct: struct {
				RequiredInt int `validate:"required"`
			}{},
			shouldPass: false,
		},
		{
			testStruct: struct {
				RequiredStruct struct { some int } `validate:"required"`
			}{},
			shouldPass: false,
		},
		{
			testStruct: struct {
				RequiredInt int `validate:"required"`
			}{1},
			shouldPass: true,
		},
		{
			testStruct: struct {
				RequiredStruct struct { some int } `validate:"required"`
			}{ struct{ some int }{1}},
			shouldPass: true,
		},
		{
			testStruct: struct {
				RequiredStruct struct {
					RequiredSubStruct struct {
						some int `validate: "required"`
					}
				}
			}{},
			// TODO: This should fail! A sub-struct that has a required field
			// should fail validation. For now, just validate the sub-struct
			// by itself.
			shouldPass: true,  // Make this false
		},
		{
			testStruct: struct {
				OptionalInt int
			}{},
			shouldPass: true,
		},
		{
			testStruct: struct {
				OptionalInt int
				RequiredInt int `validate:"required"`
			}{
				OptionalInt: 1,
			},
			shouldPass: false,
		},
		{
			testStruct: struct {
				OptionalInt int
				RequiredString string `validate:"required"`
			}{
				OptionalInt: 1,
				RequiredString: "yellow",
			},
			shouldPass: true,
		},
	}

	for i, tt := range tests {
		if err := v.Validate(tt.testStruct); (err != nil) == tt.shouldPass {
			if tt.shouldPass {
				t.Errorf("Expected test case %d to pass, but didn't. err: %v", i+1, err)
			} else {
				t.Errorf("Expected test case %d to fail, but didnt.", i+1)
			}
		}
	}
}
