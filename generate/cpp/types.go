package cpp

// FormalParam holds a type of formal parameter
type FormalParam interface {
	GetName() string
	GetFieldName() string
	GetType() string
	GetRawType() string
	Prefix() string
	Print() string
	Reference() string
	DeclareFormalParam() string
	DeclareValue() string
	Declaration() string
	GetValue() string
}

// A Function holds a name and one or more formal params
type Function struct {
	Name         string
	FormalParams []FormalParam
}
