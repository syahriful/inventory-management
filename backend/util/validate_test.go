package util

import (
	"github.com/stretchr/testify/assert"
	"inventory-management/backend/internal/http/response"
	"testing"
)

func TestValidateStruct(t *testing.T) {
	type args struct {
		s interface{}
	}
	tests := []struct {
		name           string
		args           args
		expectedErrors []*response.ErrorResponse
	}{
		{
			name: "Test validate with one field error",
			args: args{
				s: struct {
					Name string `validate:"required"`
				}{Name: ""},
			},
			expectedErrors: []*response.ErrorResponse{
				{
					FailedField: "Name",
					Tag:         "required",
					Value:       "Error validation 'required' for 'Name' field",
				},
			},
		},
		{
			name: "Test validate with two fields error",
			args: args{
				s: struct {
					Name  string `validate:"required"`
					Email string `validate:"required,email"`
				}{Name: "", Email: "widdy"},
			},
			expectedErrors: []*response.ErrorResponse{
				{
					FailedField: "Name",
					Tag:         "required",
					Value:       "Error validation 'required' for 'Name' field",
				},
				{
					FailedField: "Email",
					Tag:         "email",
					Value:       "Error validation 'email' for 'Email' field",
				},
			},
		},
		{
			name: "Test validate with no error",
			args: args{
				s: struct {
					Name  string `validate:"required"`
					Email string `validate:"required,email"`
				}{Name: "widdy", Email: "widdyarfian@gmail.com"},
			},
			expectedErrors: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errValidate := ValidateStruct(tt.args.s)
			assert.Equal(t, tt.expectedErrors, errValidate)
		})
	}
}
