package cpp

// Structure Type
type StructureParam struct {
	Type      string        `json:"type"`
	Value     string        `json:"value"`
	Name      string        `json:"name"`
	IsPointer bool          `json:"is_pointer"`
	Fields    []FormalParam `json:"fields"`
}

// Declaration of a structure
func (p StructureParam) DeclareValue() string {
	return "STRUCT" + p.GetFieldName() + " " + p.GetFieldName() + ";"
}

// Declaration is for a separate declaration of just the type
func (p StructureParam) Declaration() string {

	result := "\nstruct STRUCT" + p.Name + " {\n"

	// TODO Declaration cannot have & - need different func
	for _, field := range p.Fields {
		result = result + "   " + field.Declaration() + " = " + field.GetValue() + ";\n"
	}
	result += "}"
	return result
}

// DeclareFormalParams just declares the type (and if a pointer)
func (p StructureParam) DeclareFormalParam() string {
	return "STRUCT" + p.GetFieldName() + " " + p.GetName()
}

// Prefix of an structure formal param
func (p StructureParam) Prefix() string {
	return ""
}

// Value returns the string representation of the value
func (p StructureParam) GetValue() string {
	return p.Value
}

// GetName returns the name
func (p StructureParam) GetName() string {
	if p.IsPointer {
		return "* " + p.Name
	}
	return p.Name
}

// GetFieldName returns the field name
func (p StructureParam) GetFieldName() string {
	return p.Name
}

// GetType for a general param
func (p StructureParam) GetType() string {
	if p.IsPointer {
		return p.Type + " *"
	}
	return p.Type
}

// Get the raw type
func (p StructureParam) GetRawType() string {
	return p.Type
}

// Reference of a structure param
func (p StructureParam) Reference() string {
	if p.IsPointer {
		return "&" + p.Name
	}
	return p.Name
}

// Print a general param
func (p StructureParam) Print() string {
	result := ""
	sep := "."
	if p.IsPointer {
		sep = "->"
	}
	for _, field := range p.Fields {
		result += "    std::cout << \"" + p.Name + "\" << " + p.Name + sep + field.GetFieldName() + " << std::endl;\n"
	}
	return result
}
