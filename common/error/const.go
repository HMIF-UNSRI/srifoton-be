package error

var (
	domainErrorTranslatorDirectories = map[string]*ClientError{
		"GET_USER_PERMISSION.NOT_AUTHORIZED": NewUnauthorizedError("user not authorized"),
	}
)
