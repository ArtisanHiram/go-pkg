package __obf_7bd99df3562420c2

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

func (__obf_8e9d6bf4baf0c0ba *Error) Error() string {
	__obf_ea73d57ec466ab80 := make([]string, len(__obf_8e9d6bf4baf0c0ba.Errors))
	for __obf_8d0ae84c7a3f3f55, __obf_bf6c24b903ce000b := range __obf_8e9d6bf4baf0c0ba.Errors {
		__obf_ea73d57ec466ab80[__obf_8d0ae84c7a3f3f55] = fmt.Sprintf("* %s", __obf_bf6c24b903ce000b)
	}

	sort.Strings(__obf_ea73d57ec466ab80)
	return fmt.Sprintf(
		"%d error(s) decoding:\n\n%s",
		len(__obf_8e9d6bf4baf0c0ba.Errors), strings.Join(__obf_ea73d57ec466ab80, "\n"))
}

// WrappedErrors implements the errwrap.Wrapper interface to make this
// return value more useful with the errwrap and go-multierror libraries.
func (__obf_8e9d6bf4baf0c0ba *Error) WrappedErrors() []error {
	if __obf_8e9d6bf4baf0c0ba == nil {
		return nil
	}
	__obf_6ad3d371cac0efc2 := make([]error, len(__obf_8e9d6bf4baf0c0ba.Errors))
	for __obf_8d0ae84c7a3f3f55, __obf_8e9d6bf4baf0c0ba := range __obf_8e9d6bf4baf0c0ba.Errors {
		__obf_6ad3d371cac0efc2[__obf_8d0ae84c7a3f3f55] = errors.New(__obf_8e9d6bf4baf0c0ba)
	}

	return __obf_6ad3d371cac0efc2
}

func __obf_8c0a23210109d8dd(errors []string, __obf_bf6c24b903ce000b error) []string {
	switch __obf_8e9d6bf4baf0c0ba := __obf_bf6c24b903ce000b.(type) {
	case *Error:
		return append(errors, __obf_8e9d6bf4baf0c0ba.Errors...)
	default:
		return append(errors, __obf_8e9d6bf4baf0c0ba.Error())
	}
}
