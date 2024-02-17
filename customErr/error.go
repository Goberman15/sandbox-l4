package customErr

type CustomError struct {
	status  int
	message string
}

func (e CustomError) Error() string {
	return e.message
}

func (e CustomError) StatusCode() int {
	return e.status
}

func NewBadRequestError(message string) CustomError {
	return CustomError{
		status:  400,
		message: message,
	}
}

func NewUnauthorizedError(message string) CustomError {
	return CustomError{
		status:  401,
		message: message,
	}
}

func NewNotFoundError(message string) CustomError {
	return CustomError{
		status:  404,
		message: message,
	}
}

func NewInternalServerError(message string) CustomError {
	return CustomError{
		status:  500,
		message: message,
	}
}




