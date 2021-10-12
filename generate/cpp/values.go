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
	case "bool":
		return getBoolValue()
	case "std::string":
		return getStringValue()
	case "char":
		return getCharValue()
	case "short":
		return getShortValue(isSigned)
	case "long":
		return getIntValue(isSigned)
	case "long long":
		return getLongLong(isSigned)
	case "std::size_t":
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

	switch floatType {
	case "float":
		return getFloatValue()
	case "double":
		return getDoubleValue()
	case "long double":
		return getLongDoubleValue()
	}

	log.Fatalf("Unrecognized float type %s\n", floatType)
	return ""
}

// getFloatValue
func getFloatValue() string {
	// 1.17549e-038 to 3.40282e+038
	return insertDecimal(fmt.Sprintf("%d", utils.RandomFloat()))
}

func getDoubleValue() string {
	// TODO this range is probably too small, should be 2.22507e-308, 1.79769e308
	return insertDecimal(fmt.Sprintf("%d", utils.RandomFloat()))
}

func getLongDoubleValue() string {
	// TODO this range is probably too small, should be 2.22507e-308, 1.79769e+308
	return insertDecimal(fmt.Sprintf("%d", utils.RandomFloat()))
}

// insertDecimal adds a random decimal to somewhere in the string
func insertDecimal(value string) string {

	idx := utils.RandomIntRange(1, len(value)-1)
	result := ""
	for i, c := range value {
		result += string(c)
		if i == idx {
			result += "."
		}
	}
	return result
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
			return "-" + fmt.Sprintf("%d", utils.RandomInt(2147483648))
		}
		return fmt.Sprintf("%d", (utils.RandomInt(2147483647)))

	}
	return fmt.Sprintf("%d", (utils.RandomInt(4294967295)))
}

// getShortValue returns a random short value
func getShortValue(isSigned bool) string {

	// short int and int: -32,767 to 32,767
	if isSigned {
		if utils.RandomBool() {
			return fmt.Sprintf("%d", utils.RandomInt(32767))
		}
		return "-" + fmt.Sprintf("%d", utils.RandomInt(32768))
	}
	return fmt.Sprintf("%d", utils.RandomInt(65535))
}

// getLongLong returns a long long value
func getLongLong(isSigned bool) string {

	if isSigned {
		if utils.RandomBool() {
			return "-" + fmt.Sprintf("%d", utils.RandomInt(9223372036854775807)) // should go to ..8
		}
		return fmt.Sprintf("%d", utils.RandomInt(9223372036854775807))
	}
	// Max unsigned is 18446744073709551615
	return fmt.Sprintf("%d", utils.RandomUint64())
}

// size_t on 64 bit will be 64 bit unsigned integer
func getSizeTValue() string {
	return fmt.Sprintf("%d", utils.RandomInt(65535))
}

// Get a boolean value
func getBoolValue() string {
	return fmt.Sprintf("%v", utils.RandomBool())
}

// Get a string value
func getStringValue() string {
	return "\"" + utils.RandomName() + "\""
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
	result += "     " + name + " = " + secondPart + ";"
	return result
}
