package error

func TranslateDomainError(err error) *ClientError {
	return domainErrorTranslatorDirectories[err.Error()]
}
