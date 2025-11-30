package __obf_c953b7a5114a5dbe

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

func (__obf_d0a59ae5b9841a8a *Error) Error() string {
	__obf_81638d6b05b24a0e := make([]string, len(__obf_d0a59ae5b9841a8a.Errors))
	for __obf_507d21ee6c71f742, __obf_fb7fbd58154fa1dd := range __obf_d0a59ae5b9841a8a.Errors {
		__obf_81638d6b05b24a0e[__obf_507d21ee6c71f742] = fmt.Sprintf("* %s", __obf_fb7fbd58154fa1dd)
	}

	sort.Strings(__obf_81638d6b05b24a0e)
	return fmt.Sprintf(
		"%d error(s) decoding:\n\n%s",
		len(__obf_d0a59ae5b9841a8a.Errors), strings.Join(__obf_81638d6b05b24a0e, "\n"))
}

// WrappedErrors implements the errwrap.Wrapper interface to make this
// return value more useful with the errwrap and go-multierror libraries.
func (__obf_d0a59ae5b9841a8a *Error) WrappedErrors() []error {
	if __obf_d0a59ae5b9841a8a == nil {
		return nil
	}
	__obf_47a3e13e334b0913 := make([]error, len(__obf_d0a59ae5b9841a8a.Errors))
	for __obf_507d21ee6c71f742, __obf_d0a59ae5b9841a8a := range __obf_d0a59ae5b9841a8a.Errors {
		__obf_47a3e13e334b0913[__obf_507d21ee6c71f742] = errors.New(__obf_d0a59ae5b9841a8a)
	}

	return __obf_47a3e13e334b0913
}

func __obf_b59ff522c1d88f2e(errors []string, __obf_fb7fbd58154fa1dd error) []string {
	switch __obf_d0a59ae5b9841a8a := __obf_fb7fbd58154fa1dd.(type) {
	case *Error:
		return append(errors, __obf_d0a59ae5b9841a8a.Errors...)
	default:
		return append(errors, __obf_d0a59ae5b9841a8a.Error())
	}
}
