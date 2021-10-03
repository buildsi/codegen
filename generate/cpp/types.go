package cpp

// A Function holds a name and one or more formal params
type Function struct {
	Name         string
	FormalParams []FormalParam
}

// TODO

// FormalParam holds a type of formal parameter
type FormalParam interface {
	GetName() string
	GetType() string
	Prefix() string
	Print() string
	Reference() string
	Declaration(withValue bool) string
	GetValue() string
}

// INTEGRAL Possible integral types
// TODO need to add size_t
func GetIntegralTypes() []string {
	return []string{"char", "short", "int", "long", "long long", "__int128"}
}

// Integral Types
type IntegralFormalParam struct {
	Name      string
	Type      string
	IsSigned  bool
	IsPointer bool
	Value     string // since we are printing to a template, this can be a string
}

// Declaration of an integral formal param
func (p IntegralFormalParam) Declaration(withValue bool) string {

	result := p.Prefix() + " " + p.Type + " " + p.Name

	if p.GetType() == "__int128" {
		if withValue {
			result += ";"
			result += p.Value
			return result
		}
	}

	// A declaration outside of function params
	if withValue {
		result += " = " + p.Value + ";"
	}
	return result
}

// GetValue returns the string representation of the value
func (p IntegralFormalParam) GetValue() string {
	if p.GetType() == "__int128" {
		return p.GetName()
	}
	return p.Value
}

// GetName returns the string representation of the value
func (p IntegralFormalParam) GetName() string {
	return p.Name
}

// Reference of an integral formal param
func (p IntegralFormalParam) Reference() string {
	if p.IsPointer {
		return "&" + p.Name
	}
	return p.Name
}

// Prefix of an integral formal param
func (p IntegralFormalParam) Prefix() string {
	if p.IsSigned {
		return "signed"
	}
	return "unsigned"
}

// GetType of an integral formal param
func (p IntegralFormalParam) GetType() string {
	if p.IsPointer {
		return p.Type + " *"
	}
	return p.Type
}

// Print prints an integral formal param
func (p IntegralFormalParam) Print() string {
	name := p.Name
	if p.IsPointer {
		name = "&" + p.Name
	}
	// TODO we will want custom printing based on the type here
	return "std::cout <<  " + name + " << std::endl;"
}

func GetFloatTypes() []string {
	return []string{"float", "double", "long double"}
}

// Integral Types
type FloatFormalParam struct {
	Name      string
	Type      string
	IsComplex bool
	IsPointer bool
	Value     string
}

// Declaration of a float formal param
func (p FloatFormalParam) Declaration(withValue bool) string {
	result := p.Prefix() + p.GetType() + " " + p.Name

	// A declaration outside of function params
	if withValue {
		result += " = " + p.Value + ";"
	}
	return result
}

// Reference of an integral formal param
func (p FloatFormalParam) Reference() string {
	if p.IsPointer {
		return "&" + p.Name
	}
	return p.Name
}

// Prefix of an integral formal param
func (p FloatFormalParam) Prefix() string {
	if p.IsComplex {
		return "_Complex "
	}
	return ""
}

// Type of an integral formal param
func (p FloatFormalParam) GetType() string {
	if p.IsPointer {
		return p.Type + " *"
	}
	return p.Type
}

// Value returns the string representation of the value
func (p FloatFormalParam) GetValue() string {
	return p.Value
}

// GetName returns the name
func (p FloatFormalParam) GetName() string {
	return p.Name
}

// Print prints an float formal param
func (p FloatFormalParam) Print() string {
	name := p.Name
	if p.IsPointer {
		name = "&" + p.Name
	}
	switch p.Type {
	case "long":
		return "printf(\"%ld\n\"," + name + ");"
	}
	// TODO we will want more custom formatting based on the type here
	return "std::cout << " + name + " << std::endl;"
}
