package helpers

import (
	"testSagara/utils"

	"github.com/go-playground/validator/v10"
)

func GenerateValidationResponse(err error) (response utils.ValidationResponse) {
	var a struct{}
	response.Success = false
	response.Data = a
	var validations []utils.Validation

	// Set validation error
	fieldErr := err.(validator.ValidationErrors)

	for _, value := range fieldErr {

		// Get validation field & rule
		field, rule := value.Field(), value.Tag()

		// Create validation object
		validation := utils.Validation{Field: field, Message: utils.GenerateValidationMessage(field, rule)}

		// Add validation object to validations
		validations = append(validations, validation)
	}

	// Set validation response
	response.Validations = validations

	return response

}
