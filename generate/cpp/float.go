package cpp

func GetFloatTypes() []string {
	return []string{"float", "double", "long double"}
}

// Integral Types
type FloatFormalParam struct {
	Name      string `json:"name"`
	Type      string `json:"type"`
	IsComplex bool   `json:"is_complex"`
	IsPointer bool   `json:"is_pointer"`
	Value     string `json:"value"`
}

// Declaration of a formal param
func (p FloatFormalParam) DeclareFormalParam() string {
	return p.Type + " " + p.Reference()
}

// DeclareValues includes the value
func (p FloatFormalParam) DeclareValue() string {
	return p.Prefix() + p.Type + " " + p.Name + " = " + p.Value
}

// Declaration of a float
func (p FloatFormalParam) Declaration() string {
	// This is a declaration for formal params (we need the * for pointer)
	return p.Type + " " + p.Reference()
}

// TODO the above returns a & but for struct we don't want that?
// also look up ow to define struct with pointer...

// Reference of an integral formal param
func (p FloatFormalParam) Reference() string {
	if p.IsPointer {
		return "&" + p.Name
	}
	return p.Name
}

// Prefix of an integral formal param
func (p FloatFormalParam) Prefix() string {

	// TODO need to add complex back
	//	if p.IsComplex {
	//		return "_Complex "
	//	}
	return ""
}

// Type of an integral formal param
func (p FloatFormalParam) GetType() string {
	if p.IsPointer {
		return p.Type + " *"
	}
	return p.Type
}

// Raw type
func (p FloatFormalParam) GetRawType() string {
	return p.Type
}

// Value returns the string representation of the value
func (p FloatFormalParam) GetValue() string {
	return p.Value
}

// GetName returns the name
func (p FloatFormalParam) GetName() string {
	if p.IsPointer {
		return "*" + p.Name
	}
	return p.Name
}

// GetFieldName returns the name without decoration
func (p FloatFormalParam) GetFieldName() string {
	return p.Name
}

// Print prints an float formal param
func (p FloatFormalParam) Print() string {
	// TODO we will want more custom formatting based on the type here
	return "std::cout << \"" + p.Name + " \" << " + p.Reference() + " << std::endl;"
}
