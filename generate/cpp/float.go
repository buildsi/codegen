package cpp

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

// Declaration of a formal param
func (p FloatFormalParam) DeclareFormalParam() string {
	if p.IsPointer {
		return p.Type + " * " + p.Name
	}
	return p.Type + " " + p.GetName()
}

// DeclareValues includes the value
func (p FloatFormalParam) DeclareValue() string {
	return p.Prefix() + p.Type + " " + p.Name + " = " + p.Value
}

// Declaration of a float
func (p FloatFormalParam) Declaration() string {
	// This is a declaration for formal params (we need the * for pointer)
	return p.Type + " " + p.GetName()
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
		return "&" + p.Name
	}
	return p.Name
}

// GetFieldName returns the name without decoration
func (p FloatFormalParam) GetFieldName() string {
	return p.Name
}

// Print prints an float formal param
func (p FloatFormalParam) Print() string {
	name := p.GetName()
	// TODO we will want more custom formatting based on the type here
	return "std::cout << " + name + " << std::endl;"
}
