package errors

const (
	BusyEmail              = "e-mail address \"%v\" is busy"
	EmptyEmail             = "e-mail cannot be empty"
	InvalidEmail           = "invalid e-mail address \"%v\""
	ErrorSendingEmail      = "e-mail has not been sent"
	EmptyPassword          = "password cannot be empty"
	InvalidPassword        = "password cannot contain space sign"
	ToShortPassword        = "minimum password length is 8 characters"
	InvalidEmailOrPassword = "invalid email or password"
	RegisterTokenNotExist  = "the registration token does not exist"
	RegisterTokenExpired   = "the registration token has expired"
	RegisterTokenUsedUp    = "the registration token used up"
	SomethingWentWrong     = "oops... something went wrong, let's try again"
)
