package cpp

// INTEGRAL Possible integral types
func GetIntegralTypes(includeSizeT bool) []string {
	if !includeSizeT {
		return []string{"char", "short", "int", "long", "long long", "__int128"}
	}
	return []string{"char", "short", "int", "std::size_t", "long", "long long", "__int128"}
}

// Integral Types
type IntegralFormalParam struct {
	Name      string
	Type      string
	IsSigned  bool
	IsPointer bool
	Value     string // since we are printing to a template, this can be a string
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
		result += p.Value
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
	if p.Type == "__int128" {
		return p.GetName()
	}
	return p.Value
}

// GetName returns the string representation of the value
func (p IntegralFormalParam) GetName() string {
	if p.IsPointer {
		return "&" + p.Name
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

	// size T always unsigned
	if p.Type == "std::size_t" {
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
	name := p.GetName()

	// TODO not sure how to do this one
	if p.Type == "__int128" {
		return ""
	}
	// TODO we will want custom printing based on the type here
	return "std::cout <<  " + name + " << std::endl;"
}
