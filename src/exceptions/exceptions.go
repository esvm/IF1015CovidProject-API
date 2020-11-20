package exceptions

import (
	"encoding/json"
	"errors"
	"fmt"
)

type FormattedErr struct {
	Method      string
	Description string
	PreviousErr error
}

func (e *FormattedErr) Error() string {
	errJSONArray := e.toJSONArray()
	json, marshalErr := json.Marshal(errJSONArray)

	if marshalErr != nil {
		return fmt.Sprintf("[{\"method\":\"src/exceptions/FormattedErr.Error\",\"description\":\"unable to build stack trace due to error: %s\"}]", marshalErr)
	}

	return string(json)
}

func (e *FormattedErr) Equals(e2 *FormattedErr) bool {
	return e.Description == e2.Description
}

type ErrJSON struct {
	Method      string `json:"method"`
	Description string `json:"description"`
}

func (e *FormattedErr) toJSONArray() []*ErrJSON {

	errJSONArray := []*ErrJSON{}

	if e == nil {
		return errJSONArray
	}

	for {
		errJSONArray = append(errJSONArray, e.toJSONObject())

		if e.PreviousErr == nil {
			return errJSONArray
		}

		if _, ok := e.PreviousErr.(*FormattedErr); !ok {
			err := &FormattedErr{
				Description: e.PreviousErr.Error(),
			}

			errJSONArray = append(errJSONArray, err.toJSONObject())
			return errJSONArray
		}

		e = e.PreviousErr.(*FormattedErr)
	}
}

func (e *FormattedErr) toJSONObject() *ErrJSON {
	errJSON := &ErrJSON{
		Method:      e.Method,
		Description: e.Description,
	}

	return errJSON
}

var (
	ErrGeneral        = errors.New("something bad happened")
	ErrNotFound       = NotFoundError{}
	ErrAlreadyExists  = errors.New("already exists")
	ErrCreationFailed = errors.New("failed to create")
	ErrInvalidParam   = errors.New("invalid parameter")

	/* authentication and authorization */

	ErrAuthUsernameNotProvided  = errors.New("username has not been provided")
	ErrAuthUnsupportedAlgorithm = errors.New("unsupported signing method")
	ErrAuthMissingJWT           = errors.New("missing or invalid jwt in the request header")
	ErrAuthUnauthorized         = UnauthorizedError{}
	ErrInvalidCredentials       = InvalidCredentialsError{}
	ErrAuthMissingHeader        = errors.New("missing or invalid authorization header")
	ErrAuthUnsupportedType      = errors.New("unsupported type for authorization header")
	ErrAuthMissingAPIKey        = errors.New("missing api key")
)

// API error messages
const (
	MsgInvalidApplicationID = "validation against application service failed"
)

type NotFoundError struct{}

func (e NotFoundError) Error() string {
	return "resource not found"
}

type UnauthorizedError struct{}

func (e UnauthorizedError) Error() string {
	return "user does not have permissions to modify or view this resource"
}

type InvalidCredentialsError struct{}

func (e InvalidCredentialsError) Error() string {
	return "credentials are not valid for this service"
}

type ValidationError struct {
	Message string
	Details map[string]string
}

func (e ValidationError) Error() string {
	return e.Message
}

func (e ValidationError) Add(key, message string) {
	e.Details[key] = message
}

func (e ValidationError) AddAll(m map[string]string) {
	if m == nil {
		return
	}

	for k, v := range m {
		e.Details[k] = v
	}
}
