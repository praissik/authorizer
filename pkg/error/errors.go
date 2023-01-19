package errors

const (
	BusyEmail              = "e-mail address \"%v\" is busy"
	EmptyEmail             = "e-mail cannot be empty"
	InvalidEmail           = "invalid e-mail address \"%v\""
	ErrorSendingEmail      = "e-mail has not been sent"
	EmptyPassword          = "password cannot be empty"
	InvalidPassword        = "password cannot contain space sign"
	ToShortPassword        = "minimum password length is 8 characters"
	DifferentPasswords     = "passwords are different"
	InvalidEmailOrPassword = "invalid email or password"
	RegisterTokenNotExist  = "the registration token does not exist"
	RegisterTokenExpired   = "the registration token has expired"
	RegisterTokenUsedUp    = "the registration token used up"
	SomethingWentWrong     = "oops... something went wrong, let's try again"
)

type FieldErrors map[string]string

type UnauthorizedError struct {
	Message string
}

func (e *UnauthorizedError) Error() string {
	if e.Message != "" {
		return e.Message
	}
	return "unauthorized"
}

type PermissionError struct {
	Message string
}

func (e *PermissionError) Error() string {
	if e.Message != "" {
		return e.Message
	}
	return "permission denied"
}

type IgnoreError struct {
	Message string
}

func (e *IgnoreError) Error() string {
	if e.Message != "" {
		return e.Message
	}
	return "ignored"
}

func HandleCustomErrors(formErrors map[string]string) FieldErrors {
	var fe FieldErrors = make(map[string]string)
	for field, msg := range formErrors {
		fe[field] = msg
	}
	return fe
}
