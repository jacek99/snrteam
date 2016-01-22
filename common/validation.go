package common
import (
	"github.com/asaskevich/govalidator"
	"strings"
)

// common validation logic
func Validate(entity interface{}, entityType string) error {
	if _, err := govalidator.ValidateStruct(entity); err != nil {

		var errs BadRequestErrors

		switch err.(type){
		case govalidator.Errors:

			errors := err.(govalidator.Errors)
			for _, single := range errors {
				splits := strings.Split(single.Error(),":")
				errs = append(errs, BadRequestError{splits[1],entityType,splits[0], nil})
			}

			default:
			splits := strings.Split(err.Error(),":")
			errs = append(errs, BadRequestError{splits[1],entityType,splits[0],nil})
		}

		return errs

	} else {
		// all gooc
		return nil
	}
}



