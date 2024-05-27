package repository

import (
	errors2 "errors"
	"github.com/pkg/errors"
)

// ErrUserAlreadyExists is used when a user already exists.
var ErrUserAlreadyExists = errors2.New("user already exists")

// ErrUserInvalid is used when the user is invalid.
var ErrUserInvalid = errors2.New("user is invalid")

// ErrUserNotCreated is used when the user can't be created.
var ErrUserNotCreated = errors2.New("user not created")

// ErrUserNotUpdated is used when the user can't be updated.
var ErrUserNotUpdated = errors2.New("user not updated")

// ErrUserNotDeleted is used when the user can't be deleted.
var ErrUserNotDeleted = errors2.New("user not deleted")

// ErrUserNotListed is used when the users can't be listed.
var ErrUserNotListed = errors2.New("users not listed")

// ErrUserNotFound is used when the user can't be found.
var ErrUserNotFound = errors2.New("user not found")

// ErrNotFound not found
var ErrNotFound = errors.New("not found")

// ErrInvalidOrder invalid order
var ErrInvalidOrder = errors.New("invalid order")

// ErrUnmarshalBinary unmarshal binary
var ErrUnmarshalBinary = errors.New("unmarshal binary")

// ErrOrderNotFound order not found
var ErrOrderNotFound = errors.New("order not found")

// ErrBucketNotFound bucket not found
var ErrBucketNotFound = errors.New("bucket not found")

// ErrInvalidEntity invalid entity
var ErrInvalidEntity = errors.New("invalid entity")
