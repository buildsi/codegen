package cpp

import (
	"github.com/buildsi/codegen/utils"
)

// INTEGRAL Possible integral types NOTE __int128 removed for now, need help to define!
func GetIntegralTypes(withinStruct bool) []string {
	if withinStruct {
		return []string{"char", "short", "int", "long", "long long", "std::string", "bool"}
	}
	return []string{"char", "short", "int", "std::size_t", "long", "long long", "std::string", "bool"}
}

// Integral numeric types
func GetIntegralNumericTypes() []string {
	return []string{"short", "int", "long", "long long"}
}

// Integral Types
type IntegralFormalParam struct {
	Name      string `json:"name"`
	Type      string `json:"type"`
	IsSigned  bool   `json:"is_signed"`
	IsPointer bool   `json:"is_pointer"`
	Value     string `json:"value"`
}

// Declaration of a formal param
func (p IntegralFormalParam) DeclareFormalParam() string {
	if p.IsPointer {
		return p.Prefix() + " " + p.Type + " * " + p.Name
	}
	return p.Prefix() + " " + p.GetType() + " " + p.GetName()
}

// DeclareValues includes the value
func (p IntegralFormalParam) DeclareValue() string {
	var result string
	if p.GetType() == "__int128" {
		result = p.Prefix() + p.GetType() + " " + p.Name + ";"
		result += p.GetValue()
		return result
	}
	return p.Prefix() + p.Type + " " + p.Name + " = " + p.Value
}

// Declaration of an integral formal param
func (p IntegralFormalParam) Declaration() string {
	result := p.Prefix() + p.Type + " " + p.Name
	if p.GetRawType() == "__int128" {
		return p.Prefix() + p.GetType() + " " + p.Name
	}
	return result
}

// GetValue returns the string representation of the value
func (p IntegralFormalParam) GetValue() string {
	return p.Value
}

// GetName returns the string representation of the value
func (p IntegralFormalParam) GetName() string {
	if p.IsPointer {
		return "* " + p.Name
	}
	return p.Name
}
func (p IntegralFormalParam) GetFieldName() string {
	return p.Name
}

// Reference of an integral formal param
func (p IntegralFormalParam) Reference() string {

	// TODO need help creating these types
	if p.Type == "double" {
		return p.Name
	}
	if p.IsPointer {
		return "&" + p.Name
	}
	return p.Name
}

// Prefix of an integral formal param
func (p IntegralFormalParam) Prefix() string {

	// These are integral types that are not signed/unsigned
	skips := []string{"std::string", "std::size_t", "bool"}

	// size T, string, and bool always unsigned
	if utils.IncludesString(p.Type, skips) {
		return ""
	}
	if p.IsSigned {
		return "signed "
	}
	return "unsigned "
}

// GetType of an integral formal param
func (p IntegralFormalParam) GetType() string {
	if p.IsPointer {
		return p.Type + " *"
	}
	return p.Type
}

func (p IntegralFormalParam) GetRawType() string {
	return p.Type
}

// Print prints an integral formal param
func (p IntegralFormalParam) Print() string {
	// TODO not sure how to do this one
	if p.Type == "__int128" {
		return ""
	}
	// TODO we will want custom printing based on the type here
	return "std::cout << \"" + p.Name + "\" << " + p.Reference() + " << std::endl;"
}
