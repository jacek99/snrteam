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

// standard error, will result in 404
type NotFoundError struct {
	Message string
	EntityType string
	EntityField string
	EntityId interface{}
}

func (e NotFoundError) Error() string {
	return e.Message
}

// standard error, will result in 409
type ConflictError struct {
	Message string
	EntityType string
	EntityField string
	EntityId interface{}
}

func (e ConflictError) Error() string {
	return e.Message
}

