package __obf_ec2f65f16fa88470

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

func (__obf_a640902971632406 *Error) Error() string {
	__obf_7bfbe95d7012106a := make([]string, len(__obf_a640902971632406.Errors))
	for __obf_6cbf6487a01daf31, __obf_6d8ef670b0943b3c := range __obf_a640902971632406.Errors {
		__obf_7bfbe95d7012106a[__obf_6cbf6487a01daf31] = fmt.Sprintf("* %s", __obf_6d8ef670b0943b3c)
	}

	sort.Strings(__obf_7bfbe95d7012106a)
	return fmt.Sprintf(
		"%d error(s) decoding:\n\n%s",
		len(__obf_a640902971632406.Errors), strings.Join(__obf_7bfbe95d7012106a, "\n"))
}

// WrappedErrors implements the errwrap.Wrapper interface to make this
// return value more useful with the errwrap and go-multierror libraries.
func (__obf_a640902971632406 *Error) WrappedErrors() []error {
	if __obf_a640902971632406 == nil {
		return nil
	}
	__obf_bb9d1909c7b1c125 := make([]error, len(__obf_a640902971632406.Errors))
	for __obf_6cbf6487a01daf31, __obf_a640902971632406 := range __obf_a640902971632406.Errors {
		__obf_bb9d1909c7b1c125[__obf_6cbf6487a01daf31] = errors.New(__obf_a640902971632406)
	}

	return __obf_bb9d1909c7b1c125
}

func __obf_9ddf59feabf88c65(errors []string, __obf_6d8ef670b0943b3c error) []string {
	switch __obf_a640902971632406 := __obf_6d8ef670b0943b3c.(type) {
	case *Error:
		return append(errors, __obf_a640902971632406.Errors...)
	default:
		return append(errors, __obf_a640902971632406.Error())
	}
}
