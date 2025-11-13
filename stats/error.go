package __obf_9d1e97a8e05fedb6

type __obf_7bf50a38d611eef1 struct {
	__obf_43faab21d338d2b9 string
}

func (__obf_56bdc23e16a7cc2a __obf_7bf50a38d611eef1) Error() string {
	return __obf_56bdc23e16a7cc2a.__obf_43faab21d338d2b9
}

func (__obf_56bdc23e16a7cc2a __obf_7bf50a38d611eef1) String() string {
	return __obf_56bdc23e16a7cc2a.__obf_43faab21d338d2b9
}

// These are the package-wide error values.
// All error identification should use these values.
// https://github.com/golang/go/wiki/Errors#naming
var (
	// ErrEmptyInput Input must not be empty
	ErrEmptyInput = __obf_7bf50a38d611eef1{"Input must not be empty."}
	// ErrNaN Not a number
	ErrNaN = __obf_7bf50a38d611eef1{"Not a number."}
	// ErrNegative Must not contain negative values
	ErrNegative = __obf_7bf50a38d611eef1{"Must not contain negative values."}
	// ErrZero Must not contain zero values
	ErrZero = __obf_7bf50a38d611eef1{"Must not contain zero values."}
	// ErrBounds Input is outside of range
	ErrBounds = __obf_7bf50a38d611eef1{"Input is outside of range."}
	// ErrSize Must be the same length
	ErrSize = __obf_7bf50a38d611eef1{"Must be the same length."}
	// ErrInfValue Value is infinite
	ErrInfValue = __obf_7bf50a38d611eef1{"Value is infinite."}
	// ErrYCoord Y Value must be greater than zero
	ErrYCoord = __obf_7bf50a38d611eef1{"Y Value must be greater than zero."}

	ErrIllegalNum = __obf_7bf50a38d611eef1{"Illegal numbers."}
)
