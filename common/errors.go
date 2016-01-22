package common

// standard error, will result in 500
type GenericError struct {
	Message string
	EntityType string
	EntityField string
	EntityId interface{}
}

func (e GenericError) Error() string {
	return e.Message
}

type GenericErrors []GenericError

// standard error, will result in 404
type NotFoundError GenericError
func (e NotFoundError) Error() string {
	return e.Message
}

// standard error, will result in 409
type ConflictError GenericError
func (e ConflictError) Error() string {
	return e.Message
}

// standard error will result in 400
type BadRequestError GenericError
func (e BadRequestError) Error() string {
	return e.Message
}

// allows to return multiple bad request errors as error
type BadRequestErrors []BadRequestError
func (e BadRequestErrors) Error() string {
	var err string
	for _, e := range e {
		err += e.Error() + ";"
	}
	return err
}

// standard wrapper that allows returning multiple errors
type RequestError struct {
	Errors []GenericError
}
func (e RequestError) Error() string {
	var err string
	for _, e := range e.Errors {
		err += e.Error() + ";"
	}
	return err
}