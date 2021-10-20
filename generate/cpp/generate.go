package cpp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/buildsi/codegen/config"
	"github.com/buildsi/codegen/utils"
	"text/template"
)

func Generate(conf config.Conf, outdir string, renderType string) {

	// Ensure that required files exist, and update to absolute path
	for _, file := range conf.Files {
		if !utils.Exists(filepath.Join(conf.Root, file)) {
			log.Fatalf("%s does not exist.", file)
		}
	}

	// Currently only supported is random
	if renderType == "" {
		renderType = conf.Type
	}
	parts := strings.SplitN(renderType, ":", 2)

	// Number of generations to do
	num := uint64(1)
	if len(parts) > 1 {
		parsed, err := strconv.ParseUint(parts[1], 10, 64)
		if err != nil {
			log.Fatalf("Cannot parse number into int: %x\n", err)
		}
		num = parsed
	}

	switch parts[0] {
	case "random":
		GenerateCppRandom(conf, int(num), outdir)
	default:
		fmt.Printf("Type %s is not supported.", parts[0])
	}
}

// GenerateCppRandom generates randomized types for all templates
func GenerateCppRandom(conf config.Conf, num int, outdir string) {

	// Get other files in output directory
	paths := utils.FindDiff(utils.ListDir(conf.Root, false, true), conf.Files)

	// TODO if we generate multiple, can eventually use co-routine or channels?
	for i := 0; i < num; i++ {

		// Generate functions with params to render (eventually can add more types)
		funcs := map[string]Function{}

		for name, entry := range conf.Renders {
			switch entry.Type {
			case "function":
				funcs[name] = GenerateFunction(name, entry)
			}
		}

		// If we don't have an output directory,  make it (start at 1, not 0)
		outsubdir := filepath.Join(outdir, fmt.Sprintf("%x", i+1))
		if outdir != "" {
			os.MkdirAll(outsubdir, os.ModePerm)
		}

		// For each file, read it in...
		for i, templateName := range conf.Files {

			file := filepath.Join(conf.Root, templateName)

			// Create a new template from the file content
			t := template.Must(template.New(templateName).Funcs(templateHelpers).ParseFiles(file))

			// And render the functions into it
			if outdir != "" {
				fmt.Printf("// Writing [%x:%s]\n", i, templateName)
				WriteTemplate(filepath.Join(outsubdir, templateName), t, &funcs)
			} else {
				fmt.Printf("// Printing [%x:%s]\n", i, templateName)
				t.Execute(os.Stdout, funcs)
				// To prevent from printint to the screen, t.Execute(ioutil.Discard, funcs)
			}
		}

		// Copy the remaining files there
		if outdir != "" {
			for _, path := range paths {
				utils.CopyFile(filepath.Join(conf.Root, path), filepath.Join(outsubdir, path))
			}

			// Save json for functions
			output, _ := json.MarshalIndent(funcs, "", " ")
			_ = ioutil.WriteFile(filepath.Join(outsubdir, "codegen.json"), output, 0644)

			// If not writing to file, only allow printing one
		} else {
			break
		}
	}

}

// WriteTemplate to a filepath
func WriteTemplate(path string, t *template.Template, funcs *map[string]Function) {

	var buf bytes.Buffer
	if err := t.Execute(&buf, funcs); err != nil {
		log.Fatalf("Cannot write template to buffer: %x", err)
	}
	utils.WriteFile(path, buf.String())
}

// Generate a function from an entry
func GenerateFunction(name string, entry config.Render) Function {

	function := Function{Name: name}
	params := GenerateFormalParams(entry, 0, false, false)
	function.FormalParams = params
	return function
}

// Generate formal params, either for a function or structure
func GenerateFormalParams(entry config.Render, nestingCount int, withinStruct bool, withinUnion bool) []FormalParam {

	params := []FormalParam{}

	// How many to generate? First choice goes to exact
	num := entry.Parameters.Exact
	if num == 0 {
		if entry.Parameters.Min == 0 && entry.Parameters.Max == 0 {
			log.Fatalf("exact, or min AND max must be defined under function parameter settings.")
		}
		if entry.Parameters.Min >= entry.Parameters.Max {
			log.Fatalf("Function parameter min cannot be >= max.")
		}
		num = utils.RandomIntRange(entry.Parameters.Min, entry.Parameters.Max)
	}
	for i := 0; i < num; i++ {
		params = append(params, NewFormalParam(entry, nestingCount, withinStruct, withinUnion))
	}
	return params
}

// Functions to create new Formal Parameters (all random)
func NewFormalParam(entry config.Render, nestingCount int, withinStruct bool, withinUnion bool) FormalParam {

	// Only allow 1 level of structs (and make float, integral more likely)
	choices := []string{"integral", "float", "struct", "integral", "float"}
	if nestingCount > 0 {
		choices = []string{"integral", "float"}
	}

	// If we only want numeric
	if entry.Numeric {
		choices = []string{"numeric", "float"}
	}

	switch utils.RandomChoice(choices) {
	case "numeric":
		return NewIntegralNumeric()
	case "integral":
		return NewIntegral(withinStruct, withinUnion)
	case "float":
		return NewFloat()
	case "struct":
		return NewStruct(entry, nestingCount)
	}
	return NewIntegral(withinStruct, withinUnion)
}

// NewStruct returns a new struct type, which also includes its own set of fields
func NewStruct(entry config.Render, nestingCount int) FormalParam {

	isUnion := utils.RandomBool()
	typeName := "struct"
	if isUnion {
		typeName = "union"
	}

	// A struct has its own list of fields
	nestingCount += 1
	fields := GenerateFormalParams(entry, nestingCount, false, isUnion)

	// Get the type beforehand to derive a random value for it
	name := typeName + strings.Title(utils.RandomName())
	return StructureParam{Name: name,
		Type:      typeName,
		IsUnion:   isUnion,
		IsPointer: utils.RandomBool(),
		Fields:    fields}
}

// NewIntegral returns a new integral type
func NewIntegral(withinStruct bool, withinUnion bool) FormalParam {

	// Get the type beforehand to derive a random value for it
	name := "fpInt" + strings.Title(utils.RandomName())
	integralType := utils.RandomChoice(GetIntegralTypes(withinStruct, withinUnion))
	isSigned := utils.RandomBool()
	value := GetIntegralValue(integralType, isSigned, name)

	return IntegralFormalParam{Name: name,
		Type:      integralType,
		IsSigned:  isSigned,
		Value:     value,
		IsPointer: utils.RandomBool()}
}

// NewIntegralNumeric returns a new integral numeric type
func NewIntegralNumeric() FormalParam {

	// Get the type beforehand to derive a random value for it
	name := "fpInt" + strings.Title(utils.RandomName())
	integralType := utils.RandomChoice(GetIntegralNumericTypes())
	isSigned := utils.RandomBool()
	value := GetIntegralValue(integralType, isSigned, name)

	return IntegralFormalParam{Name: name,
		Type:      integralType,
		IsSigned:  isSigned,
		Value:     value,
		IsPointer: utils.RandomBool()}
}

func NewFloat() FormalParam {

	floatType := utils.RandomChoice(GetFloatTypes())
	isComplex := utils.RandomBool()
	value := GetFloatValue(floatType, isComplex)

	// TODO get float value here and add to structure
	return FloatFormalParam{
		Name:      "fpFloat" + strings.Title(utils.RandomName()),
		Type:      floatType,
		Value:     value,
		IsComplex: isComplex,
		IsPointer: false,
	}
}
