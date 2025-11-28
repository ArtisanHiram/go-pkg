package __obf_0d0d05e597ad9adf

type __obf_5433bb37143e7b75 struct {
	__obf_7e8938093f134bb5 string
}

func (__obf_53454715f79fbdf4 __obf_5433bb37143e7b75) Error() string {
	return __obf_53454715f79fbdf4.__obf_7e8938093f134bb5
}

func (__obf_53454715f79fbdf4 __obf_5433bb37143e7b75) String() string {
	return __obf_53454715f79fbdf4.__obf_7e8938093f134bb5
}

// These are the package-wide error values.
// All error identification should use these values.
// https://github.com/golang/go/wiki/Errors#naming
var (
	// ErrEmptyInput Input must not be empty
	ErrEmptyInput = __obf_5433bb37143e7b75{"Input must not be empty."}
	// ErrNaN Not a number
	ErrNaN = __obf_5433bb37143e7b75{"Not a number."}
	// ErrNegative Must not contain negative values
	ErrNegative = __obf_5433bb37143e7b75{"Must not contain negative values."}
	// ErrZero Must not contain zero values
	ErrZero = __obf_5433bb37143e7b75{"Must not contain zero values."}
	// ErrBounds Input is outside of range
	ErrBounds = __obf_5433bb37143e7b75{"Input is outside of range."}
	// ErrSize Must be the same length
	ErrSize = __obf_5433bb37143e7b75{"Must be the same length."}
	// ErrInfValue Value is infinite
	ErrInfValue = __obf_5433bb37143e7b75{"Value is infinite."}
	// ErrYCoord Y Value must be greater than zero
	ErrYCoord = __obf_5433bb37143e7b75{"Y Value must be greater than zero."}

	ErrIllegalNum = __obf_5433bb37143e7b75{"Illegal numbers."}
)
