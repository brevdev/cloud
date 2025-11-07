package errors

import (
	stderrors "errors"
	"fmt"
	"runtime"
	"strconv"
	"strings"

	pkgerrors "github.com/pkg/errors"
)

type ValidationError struct {
	Message string
}

var _ error = ValidationError{}

func NewValidationError(message string) ValidationError {
	return ValidationError{Message: message}
}

func (v ValidationError) Error() string {
	return v.Message
}

var New = stderrors.New

var Errorf = fmt.Errorf

func Wrap(err error, msg string) error {
	if err == nil {
		return nil
	}
	return Errorf("%s: %w", msg, err)
}

var As = stderrors.As

var Unwrap = stderrors.Unwrap

func Unwraps(err error) []error {
	u, ok := err.(interface {
		Unwrap() []error
	})
	if !ok {
		return nil
	}
	return u.Unwrap()
}

func Root(err error) error {
	for Unwrap(err) != nil {
		err = Unwrap(err)
	}
	joinedErrs := Unwraps(err)
	if len(joinedErrs) == 0 {
		return err
	}
	return Roots(joinedErrs)
}

func Roots(errs []error) error {
	if len(errs) == 0 {
		return nil
	}
	rootedErrs := make([]error, len(errs))
	for i, e := range errs {
		rootedErrs[i] = Root(e)
	}
	return Join(rootedErrs...)
}

// flattens error tree
func Flatten(err error) []error {
	if err == nil {
		return nil
	}
	joinedErrs := Unwraps(err)
	if joinedErrs == nil {
		return []error{err}
	}
	flatErrs := []error{}
	for _, e := range joinedErrs {
		flatErrs = append(flatErrs, Flatten(e)...)
	}
	return flatErrs
}

// var ReturnTrace = errtrace.Wrap

func Join(errs ...error) error {
	noNilErrs := make([]error, 0, len(errs))
	for _, err := range errs {
		if err != nil {
			noNilErrs = append(noNilErrs, err)
		}
	}
	if len(noNilErrs) == 0 {
		return nil
	}
	if len(noNilErrs) == 1 {
		return noNilErrs[0]
	}
	return stderrors.Join(errs...) //nolint:wrapcheck // this is a wrapper
}

// if multi err, combine similar errors
func CombineByString(err error) error {
	if err == nil {
		return nil
	}
	errs := Flatten(err)
	mapE := make(map[string]error)
	mapEList := []error{}
	for _, e := range errs {
		_, ok := mapE[e.Error()]
		if !ok {
			mapE[e.Error()] = e
			mapEList = append(mapEList, e)
		}
	}
	return Join(mapEList...)
}

var Is = stderrors.Is

var WrapAndTrace = WrapAndTraceInMsg

func WrapAndTraceInMsg(err error) error {
	if err == nil {
		return nil
	}
	return pkgerrors.Wrap(err, makeErrorMessage("", 0)) // this wrap also adds a stacktrace which can be nice
}

func WrapAndTrace2[T any](t T, err error) (T, error) {
	if err == nil {
		return t, nil
	}
	return t, pkgerrors.Wrap(err, makeErrorMessage("", 0))
}

func makeErrorMessage(message string, skip int) string {
	skip += 2
	pc, file, line, _ := runtime.Caller(skip)

	funcName := "unknown"
	fn := runtime.FuncForPC(pc)
	if fn != nil {
		funcName = fn.Name()
	}

	lineNum := strconv.Itoa(line)
	return fmt.Sprintf("[error] %s\n%s\n%s:%s\n", message, funcName, file, lineNum)
}

func HandleErrDefer(f func() error) {
	_ = f()
	// logger.L().Error("", zap.Error(err))
}

func ErrorContainsAny(err error, substrs ...string) bool {
	for _, substr := range substrs {
		if ErrorContains(err, substr) {
			return true
		}
	}
	return false
}

func ErrorContains(err error, substr string) bool {
	return err != nil && strings.Contains(err.Error(), substr)
}

func IsErrorExcept(err error, errs ...error) bool {
	return err != nil && !IsAny(err, errs...)
}

func IsErrorExceptSubstr(err error, substr ...string) bool {
	return err != nil && !ErrorContainsAny(err, substr...)
}

func IsAny(err error, errs ...error) bool {
	for _, e := range errs {
		if Is(err, e) {
			return true
		}
	}
	return false
}

// TruncateErrorForLogging truncates a long error message to a more manageable size
// while preserving the most important parts of the error message
func TruncateErrorForLogging(err error, maxLength int) error {
	if err == nil {
		return nil
	}

	errStr := err.Error()
	if len(errStr) <= maxLength {
		return err
	}

	// Otherwise truncate with indication
	return New(fmt.Sprintf("ERROR (truncated): %s... (truncated)", errStr[:maxLength-20]))
}
