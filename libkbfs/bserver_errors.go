package libkbfs

import (
	"errors"

	"github.com/keybase/client/go/libkb"
	keybase1 "github.com/keybase/client/go/protocol"
	"github.com/keybase/go-framed-msgpack-rpc"
)

const (
	// StatusCodeBServerError is the error code for a generic block server error.
	StatusCodeBServerError = 2700
	// StatusCodeBServerErrorBadRequest is the error code for a generic client error.
	StatusCodeBServerErrorBadRequest = 2701
	// StatusCodeBServerErrorUnauthorized is the error code for when the session has not been validated
	StatusCodeBServerErrorUnauthorized = 2702
	// StatusCodeBServerErrorOverQuota is the error code for when the user has exceeded his quota
	StatusCodeBServerErrorOverQuota = 2703
	// StatusCodeBServerErrorBlockNonExistent is the error code for when bserver cannot find a block
	StatusCodeBServerErrorBlockNonExistent = 2704
	// StatusCodeBServerErrorBlockArchived is the error code for a block has been archived
	StatusCodeBServerErrorBlockArchived = 2705
	// StatusCodeBServerErrorNoPermission is the error code for when there's no permission
	StatusCodeBServerErrorNoPermission = 2706
	// StatusCodeBServerErrorBlockDeleted is the error code for a block has been deleted
	StatusCodeBServerErrorBlockDeleted = 2707
	// StatusCodeBServerErrorNonceNonExistent is the error code when a nonce cannot be found
	StatusCodeBServerErrorNonceNonExistent = 2708
	// StatusCodeBServerErrorThrottle is the error code to indicate the client should initiate backoff.
	StatusCodeBServerErrorThrottle = 2799
)

// BServerError is a generic bserver-side error.
type BServerError struct {
	Msg string
}

// ToStatus implements the ExportableError interface for BServerError.
func (e BServerError) ToStatus() (s keybase1.Status) {
	s.Code = StatusCodeBServerError
	s.Name = "SERVER_ERROR"
	s.Desc = e.Msg
	return
}

// Error implements the Error interface for BServerError.
func (e BServerError) Error() string {
	return e.Msg
}

// BServerErrorBadRequest is a generic client-side error.
type BServerErrorBadRequest struct {
	Msg string
}

// ToStatus implements the ExportableError interface for BServerError.
func (e BServerErrorBadRequest) ToStatus() (s keybase1.Status) {
	s.Code = StatusCodeBServerErrorBadRequest
	s.Name = "BAD_REQUEST"
	s.Desc = e.Msg
	return
}

// Error implements the Error interface for BServerError.
func (e BServerErrorBadRequest) Error() string {
	if e.Msg == "" {
		return "BServer: bad client request"
	}
	return e.Msg
}

// BServerErrorUnauthorized is a generic client-side error.
type BServerErrorUnauthorized struct {
	Msg string
}

// ToStatus implements the ExportableError interface for BServerErrorUnauthorized.
func (e BServerErrorUnauthorized) ToStatus() (s keybase1.Status) {
	s.Code = StatusCodeBServerErrorUnauthorized
	s.Name = "SESSION_UNAUTHORIZED"
	s.Desc = e.Msg
	return
}

// Error implements the Error interface for BServerErrorUnauthorized.
func (e BServerErrorUnauthorized) Error() string {
	if e.Msg == "" {
		return "BServer: session not validated"
	}
	return e.Msg
}

// BServerErrorOverQuota is a generic client-side error.
type BServerErrorOverQuota struct {
	Msg string
}

// ToStatus implements the ExportableError interface for BServerErrorOverQuota.
func (e BServerErrorOverQuota) ToStatus() (s keybase1.Status) {
	s.Code = StatusCodeBServerErrorOverQuota
	s.Name = "QUOTA_EXCEEDED"
	s.Desc = e.Msg
	return
}

// Error implements the Error interface for BServerErrorOverQuota.
func (e BServerErrorOverQuota) Error() string {
	if e.Msg == "" {
		return "BServer: user has exceeded quota"
	}
	return e.Msg
}

//BServerErrorBlockNonExistent is an exportable error from bserver
type BServerErrorBlockNonExistent struct {
	Msg string
}

// ToStatus implements the ExportableError interface for BServerErrorBlockNonExistent
func (e BServerErrorBlockNonExistent) ToStatus() (s keybase1.Status) {
	s.Code = StatusCodeBServerErrorBlockNonExistent
	s.Name = "BLOCK_NONEXISTENT"
	s.Desc = e.Msg
	return
}

// Error implements the Error interface for BServerErrorBlockNonExistent.
func (e BServerErrorBlockNonExistent) Error() string {
	if e.Msg == "" {
		return "BServer: block does not exist"
	}
	return e.Msg
}

//BServerErrorBlockArchived is an exportable error from bserver
type BServerErrorBlockArchived struct {
	Msg string
}

// ToStatus implements the ExportableError interface for BServerErrorBlockArchived
func (e BServerErrorBlockArchived) ToStatus() (s keybase1.Status) {
	s.Code = StatusCodeBServerErrorBlockArchived
	s.Name = "BLOCK_ARCHIVED"
	s.Desc = e.Msg
	return
}

// Error implements the Error interface for BServerErrorBlockArchived.
func (e BServerErrorBlockArchived) Error() string {
	if e.Msg == "" {
		return "BServer: block is archived"
	}
	return e.Msg
}

//BServerErrorBlockDeleted is an exportable error from bserver
type BServerErrorBlockDeleted struct {
	Msg string
}

// ToStatus implements the ExportableError interface for BServerErrorBlockDeleted
func (e BServerErrorBlockDeleted) ToStatus() (s keybase1.Status) {
	s.Code = StatusCodeBServerErrorBlockDeleted
	s.Name = "BLOCK_DELETED"
	s.Desc = e.Msg
	return
}

// Error implements the Error interface for BServerErrorBlockDeleted
func (e BServerErrorBlockDeleted) Error() string {
	if e.Msg == "" {
		return "BServer: block is deleted"
	}
	return e.Msg
}

//BServerErrorNoPermission is an exportable error from bserver
type BServerErrorNoPermission struct {
	Msg string
}

// ToStatus implements the ExportableError interface for BServerErrorBlockArchived
func (e BServerErrorNoPermission) ToStatus() (s keybase1.Status) {
	s.Code = StatusCodeBServerErrorNoPermission
	s.Name = "NO_PERMISSION"
	s.Desc = e.Msg
	return
}

// Error implements the Error interface for BServerErrorNoPermission.
func (e BServerErrorNoPermission) Error() string {
	if e.Msg == "" {
		return "BServer: permission denied"
	}
	return e.Msg
}

//BServerErrorNonceNonExistent is an exportable error from bserver
type BServerErrorNonceNonExistent struct {
	Msg string
}

// ToStatus implements the ExportableError interface for BServerErrorNonceNonExistent
func (e BServerErrorNonceNonExistent) ToStatus() (s keybase1.Status) {
	s.Code = StatusCodeBServerErrorNonceNonExistent
	s.Name = "BLOCK_NONCENONEXISTENT"
	s.Desc = e.Msg
	return
}

// Error implements the Error interface for BServerErrornonceNonExistent.
func (e BServerErrorNonceNonExistent) Error() string {
	if e.Msg == "" {
		return "BServer: reference nonce does not exist"
	}
	return e.Msg
}

// BServerErrorThrottle is returned when the server wants the client to backoff.
type BServerErrorThrottle struct {
	Msg string
}

// Error implements the Error interface for BServerErrorThrottle.
func (e BServerErrorThrottle) Error() string {
	return e.Msg
}

// ToStatus implements the ExportableError interface for BServerErrorThrottle.
func (e BServerErrorThrottle) ToStatus() (s keybase1.Status) {
	s.Code = StatusCodeBServerErrorThrottle
	s.Name = "THROTTLE"
	s.Desc = e.Msg
	return
}

type bServerErrorUnwrapper struct{}

var _ rpc.ErrorUnwrapper = bServerErrorUnwrapper{}

func (eu bServerErrorUnwrapper) MakeArg() interface{} {
	return &keybase1.Status{}
}

func (eu bServerErrorUnwrapper) UnwrapError(arg interface{}) (appError error, dispatchError error) {
	s, ok := arg.(*keybase1.Status)
	if !ok {
		return nil, errors.New("Error converting arg to keybase1.Status object in bServerErrorUnwrapper.UnwrapError")
	}

	if s == nil || s.Code == 0 {
		return nil, nil
	}

	switch s.Code {
	case StatusCodeBServerError:
		appError = BServerError{Msg: s.Desc}
		break
	case StatusCodeBServerErrorBadRequest:
		appError = BServerErrorBadRequest{Msg: s.Desc}
		break
	case StatusCodeBServerErrorUnauthorized:
		appError = BServerErrorUnauthorized{Msg: s.Desc}
		break
	case StatusCodeBServerErrorOverQuota:
		appError = BServerErrorOverQuota{Msg: s.Desc}
		break
	case StatusCodeBServerErrorBlockNonExistent:
		appError = BServerErrorBlockNonExistent{Msg: s.Desc}
		break
	case StatusCodeBServerErrorBlockArchived:
		appError = BServerErrorBlockArchived{Msg: s.Desc}
		break
	case StatusCodeBServerErrorNoPermission:
		appError = BServerErrorNoPermission{Msg: s.Desc}
		break
	case StatusCodeBServerErrorThrottle:
		appError = BServerErrorThrottle{Msg: s.Desc}
		break
	default:
		ase := libkb.AppStatusError{
			Code:   s.Code,
			Name:   s.Name,
			Desc:   s.Desc,
			Fields: make(map[string]string),
		}
		for _, f := range s.Fields {
			ase.Fields[f.Key] = f.Value
		}
		appError = ase
	}

	return appError, nil
}
