package cpp

import (
	"github.com/buildsi/codegen/utils"
	"strings"
)

// Structure or Union Type
type StructureParam struct {
	Type      string        `json:"type"`
	Value     string        `json:"value"`
	Name      string        `json:"name"`
	IsPointer bool          `json:"is_pointer"`
	Fields    []FormalParam `json:"fields"`
	IsUnion   bool          `json:"is_union"`
}

// Declaration of a structure
func (p StructureParam) DeclareValue() string {
	prefix := "STRUCT"
	if p.IsUnion {
		prefix = "UNION"
	}
	return prefix + p.GetFieldName() + " " + p.GetFieldName() + ";"
}

// Declaration is for a separate declaration of just the type
func (p StructureParam) Declaration() string {
	if p.IsUnion {
		return p.declarationUnion()
	}
	return p.declarationStruct()
}

func (p StructureParam) declarationStruct() string {
	prefix := "STRUCT"
	result := "\n" + strings.ToLower(prefix) + " " + prefix + p.Name + " {\n"
	for _, field := range p.Fields {
		result = result + "   " + field.Declaration() + " = " + field.GetValue() + ";\n"
	}
	result += "}"
	return result
}

func (p StructureParam) declarationUnion() string {
	prefix := "UNION"
	result := "\n" + strings.ToLower(prefix) + " " + prefix + p.Name + " {\n"

	// Randomly choose just one to give a value to
	idx := utils.RandomIntRange(0, len(p.Fields))
	for i, field := range p.Fields {
		if i == idx {
			result = result + "   " + field.Declaration() + " = " + field.GetValue() + ";\n"
		} else {
			result = result + "   " + field.Declaration() + ";\n"
		}
	}
	result += "}"
	return result
}

// DeclareFormalParams just declares the type (and if a pointer)
func (p StructureParam) DeclareFormalParam() string {
	prefix := "STRUCT"
	if p.IsUnion {
		prefix = "UNION"
	}
	return prefix + p.GetFieldName() + " " + p.GetName()
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

// Assert a structure param
func (p StructureParam) Assert() string {
	result := ""
	sep := "."
	if p.IsPointer {
		sep = "->"
	}
	for _, field := range p.Fields {
		result += "    assert (" + p.GetName() + sep + field.GetFieldName() + " == " + field.GetValue() + ");\n"
	}
	return result
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
