package __obf_ce59006f265ba976

type __obf_bc04d09b6c6c4792 struct {
	__obf_2a4fcc8899864600 string
}

func (__obf_65d938f48d53f189 __obf_bc04d09b6c6c4792) Error() string {
	return __obf_65d938f48d53f189.__obf_2a4fcc8899864600
}

func (__obf_65d938f48d53f189 __obf_bc04d09b6c6c4792) String() string {
	return __obf_65d938f48d53f189.__obf_2a4fcc8899864600
}

// These are the package-wide error values.
// All error identification should use these values.
// https://github.com/golang/go/wiki/Errors#naming
var (
	// ErrEmptyInput Input must not be empty
	ErrEmptyInput = __obf_bc04d09b6c6c4792{"Input must not be empty."}
	// ErrNaN Not a number
	ErrNaN = __obf_bc04d09b6c6c4792{"Not a number."}
	// ErrNegative Must not contain negative values
	ErrNegative = __obf_bc04d09b6c6c4792{"Must not contain negative values."}
	// ErrZero Must not contain zero values
	ErrZero = __obf_bc04d09b6c6c4792{"Must not contain zero values."}
	// ErrBounds Input is outside of range
	ErrBounds = __obf_bc04d09b6c6c4792{"Input is outside of range."}
	// ErrSize Must be the same length
	ErrSize = __obf_bc04d09b6c6c4792{"Must be the same length."}
	// ErrInfValue Value is infinite
	ErrInfValue = __obf_bc04d09b6c6c4792{"Value is infinite."}
	// ErrYCoord Y Value must be greater than zero
	ErrYCoord = __obf_bc04d09b6c6c4792{"Y Value must be greater than zero."}

	ErrIllegalNum = __obf_bc04d09b6c6c4792{"Illegal numbers."}
)
