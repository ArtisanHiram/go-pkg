package __obf_c357e2bf526d00f9

type __obf_03cb0607851e3afc struct {
	__obf_4a00730613aa7357 string
}

func (__obf_d4e2b66ebe9a2dbf __obf_03cb0607851e3afc) Error() string {
	return __obf_d4e2b66ebe9a2dbf.__obf_4a00730613aa7357
}

func (__obf_d4e2b66ebe9a2dbf __obf_03cb0607851e3afc) String() string {
	return __obf_d4e2b66ebe9a2dbf.__obf_4a00730613aa7357
}

// These are the package-wide error values.
// All error identification should use these values.
// https://github.com/golang/go/wiki/Errors#naming
var (
	// ErrEmptyInput Input must not be empty
	ErrEmptyInput = __obf_03cb0607851e3afc{"Input must not be empty."}
	// ErrNaN Not a number
	ErrNaN = __obf_03cb0607851e3afc{"Not a number."}
	// ErrNegative Must not contain negative values
	ErrNegative = __obf_03cb0607851e3afc{"Must not contain negative values."}
	// ErrZero Must not contain zero values
	ErrZero = __obf_03cb0607851e3afc{"Must not contain zero values."}
	// ErrBounds Input is outside of range
	ErrBounds = __obf_03cb0607851e3afc{"Input is outside of range."}
	// ErrSize Must be the same length
	ErrSize = __obf_03cb0607851e3afc{"Must be the same length."}
	// ErrInfValue Value is infinite
	ErrInfValue = __obf_03cb0607851e3afc{"Value is infinite."}
	// ErrYCoord Y Value must be greater than zero
	ErrYCoord = __obf_03cb0607851e3afc{"Y Value must be greater than zero."}

	ErrIllegalNum = __obf_03cb0607851e3afc{"Illegal numbers."}
)
