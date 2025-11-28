package __obf_13f6e310b0abe7e3

type __obf_ad913de13b75d2bc struct {
	__obf_9ee970d69bb6dfc4 string
}

func (__obf_516af7fe712d1b0f __obf_ad913de13b75d2bc) Error() string {
	return __obf_516af7fe712d1b0f.__obf_9ee970d69bb6dfc4
}

func (__obf_516af7fe712d1b0f __obf_ad913de13b75d2bc) String() string {
	return __obf_516af7fe712d1b0f.__obf_9ee970d69bb6dfc4
}

// These are the package-wide error values.
// All error identification should use these values.
// https://github.com/golang/go/wiki/Errors#naming
var (
	// ErrEmptyInput Input must not be empty
	ErrEmptyInput = __obf_ad913de13b75d2bc{"Input must not be empty."}
	// ErrNaN Not a number
	ErrNaN = __obf_ad913de13b75d2bc{"Not a number."}
	// ErrNegative Must not contain negative values
	ErrNegative = __obf_ad913de13b75d2bc{"Must not contain negative values."}
	// ErrZero Must not contain zero values
	ErrZero = __obf_ad913de13b75d2bc{"Must not contain zero values."}
	// ErrBounds Input is outside of range
	ErrBounds = __obf_ad913de13b75d2bc{"Input is outside of range."}
	// ErrSize Must be the same length
	ErrSize = __obf_ad913de13b75d2bc{"Must be the same length."}
	// ErrInfValue Value is infinite
	ErrInfValue = __obf_ad913de13b75d2bc{"Value is infinite."}
	// ErrYCoord Y Value must be greater than zero
	ErrYCoord = __obf_ad913de13b75d2bc{"Y Value must be greater than zero."}

	ErrIllegalNum = __obf_ad913de13b75d2bc{"Illegal numbers."}
)
