package cpp

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/buildsi/codegen/config"
	"github.com/buildsi/codegen/utils"
	"text/template"
)

func Generate(conf config.Conf) {

	// Ensure that required files exist, and update to absolute path
	for i, file := range conf.Files {
		file = filepath.Join(conf.Root, file)
		if !utils.Exists(file) {
			log.Fatalf("%s does not exist.", file)
		}
		conf.Files[i] = file
	}

	// Currently only supported is random
	parts := strings.SplitN(conf.Type, ":", 2)

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
		GenerateCppRandom(conf, int(num))
	default:
		fmt.Printf("Type %s is not supported.", parts[0])
	}
}

// GenerateCppRandom generates randomized types for all templates
func GenerateCppRandom(conf config.Conf, num int) {

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

		// For each file, read it in...
		for i, file := range conf.Files {
			content := utils.ReadFile(file)
			templateName := filepath.Base(file)

			// Create a new template from the file content
			t := template.Must(template.New(templateName).Funcs(templateHelpers).Parse(content))

			if i != 0 {
				fmt.Printf("\n\n")
			}
			// And render the functions into it
			fmt.Println("//", templateName)
			t.Execute(os.Stdout, funcs)
			// To prevent from printint to the screen, t.Execute(ioutil.Discard, funcs)
		}
	}

}

// Generate a function from an entry
func GenerateFunction(name string, entry config.Render) Function {

	function := Function{Name: name}
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
		num = utils.RandomRange(entry.Parameters.Min, entry.Parameters.Max)
	}
	for i := 0; i < num; i++ {
		params = append(params, NewFormalParam())
	}

	function.FormalParams = params
	return function
}

// Functions to create new Formal Parameters (all random)
func NewFormalParam() FormalParam {
	switch utils.RandomChoice([]string{"integral", "float"}) {
	case "integral":
		return NewIntegral()
	case "float":
		return NewFloat()
	}
	return NewIntegral()
}

// NewIntegral returns a new integral type
func NewIntegral() FormalParam {

	// Get the type beforehand to derive a random value for it
	integralType := utils.RandomChoice(GetIntegralTypes())
	isSigned := utils.RandomBool()
	value := GetIntegralValue(integralType, isSigned)

	return IntegralFormalParam{Name: "fpInt" + strings.Title(utils.RandomName()),
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
		IsPointer: utils.RandomBool(),
	}
}
