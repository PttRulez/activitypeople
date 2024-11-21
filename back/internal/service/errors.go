package service

import "errors"

var ErrDomainNotFound = errors.New("not found")
var ErrAlreadyExists = errors.New("already exists")
