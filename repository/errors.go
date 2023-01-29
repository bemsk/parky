package repository

import "errors"

var (
	ErrRepositoryTicketNotRegistered = errors.New("[repository] ticket not registered")
	ErrRepositoryTicketAlreadyRegistered = errors.New("[repository] ticket already registered")
)