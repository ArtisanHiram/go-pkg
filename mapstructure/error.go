package __obf_4dc3483102e0d35a

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

func (__obf_823e6570785aac52 *Error) Error() string {
	__obf_66e0b5566f320af7 := make([]string, len(__obf_823e6570785aac52.Errors))
	for __obf_99f0e7230007b81b, __obf_6585aef4313e6005 := range __obf_823e6570785aac52.Errors {
		__obf_66e0b5566f320af7[__obf_99f0e7230007b81b] = fmt.Sprintf("* %s", __obf_6585aef4313e6005)
	}

	sort.Strings(__obf_66e0b5566f320af7)
	return fmt.Sprintf(
		"%d error(s) decoding:\n\n%s",
		len(__obf_823e6570785aac52.Errors), strings.Join(__obf_66e0b5566f320af7, "\n"))
}

// WrappedErrors implements the errwrap.Wrapper interface to make this
// return value more useful with the errwrap and go-multierror libraries.
func (__obf_823e6570785aac52 *Error) WrappedErrors() []error {
	if __obf_823e6570785aac52 == nil {
		return nil
	}
	__obf_382efc6537d96d52 := make([]error, len(__obf_823e6570785aac52.Errors))
	for __obf_99f0e7230007b81b, __obf_823e6570785aac52 := range __obf_823e6570785aac52.Errors {
		__obf_382efc6537d96d52[__obf_99f0e7230007b81b] = errors.New(__obf_823e6570785aac52)
	}

	return __obf_382efc6537d96d52
}

func __obf_513ecd3219b0baa9(errors []string, __obf_6585aef4313e6005 error) []string {
	switch __obf_823e6570785aac52 := __obf_6585aef4313e6005.(type) {
	case *Error:
		return append(errors, __obf_823e6570785aac52.Errors...)
	default:
		return append(errors, __obf_823e6570785aac52.Error())
	}
}
