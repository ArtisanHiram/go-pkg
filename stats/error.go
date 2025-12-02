package __obf_e7d13a84deed4934

type __obf_84f33fa02d002fc7 struct {
	__obf_de3ff84835a3e8a9 string
}

func (__obf_1e3975c9e116ae8c __obf_84f33fa02d002fc7,) Error() string {
	return __obf_1e3975c9e116ae8c.__obf_de3ff84835a3e8a9
}

func (__obf_1e3975c9e116ae8c __obf_84f33fa02d002fc7,) String() string {
	return __obf_1e3975c9e116ae8c.

	// These are the package-wide error values.
	// All error identification should use these values.
	// https://github.com/golang/go/wiki/Errors#naming
	__obf_de3ff84835a3e8a9
}

var (
	// ErrEmptyInput Input must not be empty
	ErrEmptyInput = __obf_84f33fa02d002fc7{"Input must not be empty."}
	// ErrNaN Not a number
	ErrNaN = __obf_84f33fa02d002fc7{"Not a number."}
	// ErrNegative Must not contain negative values
	ErrNegative = __obf_84f33fa02d002fc7{"Must not contain negative values."}
	// ErrZero Must not contain zero values
	ErrZero = __obf_84f33fa02d002fc7{"Must not contain zero values."}
	// ErrBounds Input is outside of range
	ErrBounds = __obf_84f33fa02d002fc7{"Input is outside of range."}
	// ErrSize Must be the same length
	ErrSize = __obf_84f33fa02d002fc7{"Must be the same length."}
	// ErrInfValue Value is infinite
	ErrInfValue = __obf_84f33fa02d002fc7{"Value is infinite."}
	// ErrYCoord Y Value must be greater than zero
	ErrYCoord = __obf_84f33fa02d002fc7{"Y Value must be greater than zero."}

	ErrIllegalNum = __obf_84f33fa02d002fc7{"Illegal numbers."}
)
