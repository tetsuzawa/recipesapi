package core

import "github.com/go-playground/validator/v10"

var Validate *validator.Validate

func init() {
	Validate = validator.New()
}

func ValidateRecipe(recipe *Recipe) []string {
	var errorMessages []string

	err := Validate.Struct(recipe)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {

			var errorMessage string
			fieldName := err.Field()

			switch fieldName {
			case "Title":
				errorMessage = "field Title is required"
			case "MakingTime":
				errorMessage = "field MakingTime is required"
			case "Serves":
				errorMessage = "field Serve is required"
			case "Ingredients":
				errorMessage = "field Ingredients is required"
			case "Cost":
				errorMessage = "field Cost is required"
			default:
				errorMessage = "unprecedented error"
			}
			errorMessages = append(errorMessages, errorMessage)
		}
	}
	return errorMessages
	
}
