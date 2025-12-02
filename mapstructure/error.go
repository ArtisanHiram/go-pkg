package __obf_8873749254e9e83f

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

func (__obf_6a4bc82dbdc21f1b *Error) Error() string {
	__obf_c90264b783f49b8f := make([]string, len(__obf_6a4bc82dbdc21f1b.Errors))
	for __obf_4e5b08e68a148355, __obf_ce3ada6ea4dc41de := range __obf_6a4bc82dbdc21f1b.Errors {
		__obf_c90264b783f49b8f[__obf_4e5b08e68a148355] = fmt.Sprintf("* %s", __obf_ce3ada6ea4dc41de)
	}

	sort.Strings(__obf_c90264b783f49b8f)
	return fmt.Sprintf(
		"%d error(s) decoding:\n\n%s",
		len(__obf_6a4bc82dbdc21f1b.Errors), strings.Join(__obf_c90264b783f49b8f, "\n"))
}

// WrappedErrors implements the errwrap.Wrapper interface to make this
// return value more useful with the errwrap and go-multierror libraries.
func (__obf_6a4bc82dbdc21f1b *Error) WrappedErrors() []error {
	if __obf_6a4bc82dbdc21f1b == nil {
		return nil
	}
	__obf_a9459758946d2eaa := make([]error, len(__obf_6a4bc82dbdc21f1b.Errors))
	for __obf_4e5b08e68a148355, __obf_6a4bc82dbdc21f1b := range __obf_6a4bc82dbdc21f1b.Errors {
		__obf_a9459758946d2eaa[__obf_4e5b08e68a148355] = errors.New(__obf_6a4bc82dbdc21f1b)
	}

	return __obf_a9459758946d2eaa
}

func __obf_8cd4c4c9f4ceb7f3(errors []string, __obf_ce3ada6ea4dc41de error) []string {
	switch __obf_6a4bc82dbdc21f1b := __obf_ce3ada6ea4dc41de.(type) {
	case *Error:
		return append(errors, __obf_6a4bc82dbdc21f1b.Errors...)
	default:
		return append(errors, __obf_6a4bc82dbdc21f1b.Error())
	}
}
