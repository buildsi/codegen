package cpp

import (
	"fmt"
	"log"

	"github.com/buildsi/codegen/utils"
)

// GetIntegralValue returns an integral value depending on the type
func GetIntegralValue(integralType string, isSigned bool, name string) string {

	switch integralType {
	case "int":
		return getIntValue(isSigned)
	case "char":
		return getCharValue()
	case "short":
		return getShortValue(isSigned)
	case "long":
		return getIntValue(isSigned)
	case "long long":
		return getLongLong(isSigned)
	case "size_t":
		return getSizeTValue()

	// This is multiple lines
	case "__int128":
		return getInt128Value(name)
	}

	log.Fatalf("Unrecognized integral type %s\n", integralType)
	return ""
}

// GetFloatValue returns a float value depending on type of float and if is complex
// TODO not written yet!
func GetFloatValue(floatType string, isComplex bool) string {
	return "123.33"
}

// https://docs.microsoft.com/en-us/cpp/c-language/cpp-integer-limits?view=msvc-160

// getCharValue returns a random char value
func getCharValue() string {
	return "'" + utils.RandomChar() + "'"
}

// getIntValue returns a random integer (signed or unsigned)
func getIntValue(isSigned bool) string {

	// Can we support negative integers?
	if isSigned {
		if utils.RandomBool() {
			return "-" + string(utils.RandomInt(2147483648))
		}
		return string(utils.RandomInt(2147483647))

	}
	return string(utils.RandomInt(4294967295))
}

// getShortValue returns a random short value
func getShortValue(isSigned bool) string {

	// short int and int: -32,767 to 32,767
	if isSigned {
		if utils.RandomBool() {
			return string(utils.RandomInt(32767))
		}
		return "-" + string(utils.RandomInt(32768))
	}
	return string(utils.RandomInt(65535))
}

// getLongLong returns a long long value
func getLongLong(isSigned bool) string {

	if isSigned {
		if utils.RandomBool() {
			return "-" + string(utils.RandomInt(9223372036854775807)) // should go to ..8
		}
		return string(utils.RandomInt(9223372036854775807))
	}
	// Max unsigned is 18446744073709551615
	return string(utils.RandomUint64())
}

// size_t on 64 bit will be 64 bit unsigned integer
func getSizeTValue() string {
	return string(utils.RandomInt(65535))
}

// get an int128 value
func getInt128Value(name string) string {

	//   __int128_t c;
	//   c = 0x0000000000000006;
	//   c = c << 64;
	//   c += 0x0000000000000007;

	// This needs custom parsing
	firstPart := "0x" + fmt.Sprintf("%12d", utils.RandomInt(9999999999999999))
	secondPart := "0x" + fmt.Sprintf("%12d", utils.RandomInt(9999999999999999))

	result := "\n     " + name + " = " + firstPart + ";\n"
	result += "     " + name + " << 64;\n"
	result += "     " + name + " = " + secondPart + ";\n"
	return result
}
