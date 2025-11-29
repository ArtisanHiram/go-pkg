package __obf_68a92d5117d8d921

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

// Error implements the error interface and can represents multiple
// errors that occur in the course of a single decode.
type Error struct {
	Errors []string
}

func (__obf_d04d2a8d186ca580 *Error) Error() string {
	__obf_aa64cd73b06aefce := make([]string, len(__obf_d04d2a8d186ca580.Errors))
	for __obf_a6629621259bc2bc, __obf_ef25cdf0d96b47a8 := range __obf_d04d2a8d186ca580.Errors {
		__obf_aa64cd73b06aefce[__obf_a6629621259bc2bc] = fmt.Sprintf("* %s", __obf_ef25cdf0d96b47a8)
	}

	sort.Strings(__obf_aa64cd73b06aefce)
	return fmt.Sprintf(
		"%d error(s) decoding:\n\n%s",
		len(__obf_d04d2a8d186ca580.Errors), strings.Join(__obf_aa64cd73b06aefce, "\n"))
}

// WrappedErrors implements the errwrap.Wrapper interface to make this
// return value more useful with the errwrap and go-multierror libraries.
func (__obf_d04d2a8d186ca580 *Error) WrappedErrors() []error {
	if __obf_d04d2a8d186ca580 == nil {
		return nil
	}
	__obf_84b0717583bd55a4 := make([]error, len(__obf_d04d2a8d186ca580.Errors))
	for __obf_a6629621259bc2bc, __obf_d04d2a8d186ca580 := range __obf_d04d2a8d186ca580.Errors {
		__obf_84b0717583bd55a4[__obf_a6629621259bc2bc] = errors.New(__obf_d04d2a8d186ca580)
	}

	return __obf_84b0717583bd55a4
}

func __obf_39ea766e30fe264e(errors []string, __obf_ef25cdf0d96b47a8 error) []string {
	switch __obf_d04d2a8d186ca580 := __obf_ef25cdf0d96b47a8.(type) {
	case *Error:
		return append(errors, __obf_d04d2a8d186ca580.Errors...)
	default:
		return append(errors, __obf_d04d2a8d186ca580.Error())
	}
}
