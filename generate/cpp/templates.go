package cpp

import (
	"text/template"
)

var templateHelpers template.FuncMap = map[string]interface{}{

	// AsFormalParams renders one or more formal params into declarations
	"AsFormalParams": func(f Function) string {
		render := ""
		for i, param := range f.FormalParams {
			render += param.DeclareFormalParam()
			if i != len(f.FormalParams)-1 {
				render += ", "
			}
		}
		return render
	},

	// PrintArgs prints each formal param as a print statement
	"PrintArgs": func(f Function) string {

		// TODO indent should come from template
		render := ""
		for _, param := range f.FormalParams {
			render += "     " + param.Print() + "\n"
		}
		return render
	},

	// DeclareArgs declares a variable for each argument
	"DeclareArgs": func(f Function) string {
		render := ""
		for i, param := range f.FormalParams {
			render += "     " + param.DeclareValue() + ";"
			if i != len(f.FormalParams)-1 {
				render += "\n"
			}
		}
		return render
	},

	// DeclareStructs declares each structure
	"DeclareStructs": func(f Function) string {
		render := ""
		for _, param := range f.FormalParams {
			if param.GetRawType() == "struct" {
				render += "     " + param.Declaration() + ";"
			}
		}
		return render
	},

	// GetFunctionName returns the name
	"GetFunctionName": func(f Function) string {
		return f.Name
	},

	// CallFunction calls the function with the args
	"CallFunction": func(f Function) string {
		render := f.Name + "("
		for i, param := range f.FormalParams {
			render += param.Reference()
			if i != len(f.FormalParams)-1 {
				render += ", "
			}
		}
		render += ");"
		return render
	},
}
