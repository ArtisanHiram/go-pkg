package __obf_0c4a280c8ca42808

type __obf_6fc215de815dc8fd struct {
	__obf_9e05075c97d73684 string
}

func (__obf_232c61af255d55a3 __obf_6fc215de815dc8fd) Error() string {
	return __obf_232c61af255d55a3.__obf_9e05075c97d73684
}

func (__obf_232c61af255d55a3 __obf_6fc215de815dc8fd) String() string {
	return __obf_232c61af255d55a3.__obf_9e05075c97d73684
}

// These are the package-wide error values.
// All error identification should use these values.
// https://github.com/golang/go/wiki/Errors#naming
var (
	// ErrEmptyInput Input must not be empty
	ErrEmptyInput = __obf_6fc215de815dc8fd{"Input must not be empty."}
	// ErrNaN Not a number
	ErrNaN = __obf_6fc215de815dc8fd{"Not a number."}
	// ErrNegative Must not contain negative values
	ErrNegative = __obf_6fc215de815dc8fd{"Must not contain negative values."}
	// ErrZero Must not contain zero values
	ErrZero = __obf_6fc215de815dc8fd{"Must not contain zero values."}
	// ErrBounds Input is outside of range
	ErrBounds = __obf_6fc215de815dc8fd{"Input is outside of range."}
	// ErrSize Must be the same length
	ErrSize = __obf_6fc215de815dc8fd{"Must be the same length."}
	// ErrInfValue Value is infinite
	ErrInfValue = __obf_6fc215de815dc8fd{"Value is infinite."}
	// ErrYCoord Y Value must be greater than zero
	ErrYCoord = __obf_6fc215de815dc8fd{"Y Value must be greater than zero."}

	ErrIllegalNum = __obf_6fc215de815dc8fd{"Illegal numbers."}
)
