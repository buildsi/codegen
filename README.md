# Codegen

This is a simple code generation library in Go, and one that we will use to generate
a bunch of (permutation based) binaries.

🚧️ **under development** 🚧️

## Usage

### Quick Start

To build the library:

```bash
$ make
```

Equivalently you can run it with go:

```bash
$ go run main.go -h
```
```
$ go run main.go -h
USAGE: codegen CMD [OPTIONS]

DESCRIPTION: Generate code for other langauges using Go

COMMANDS:

    NAME     ALIAS  DESCRIPTION
    gen      g      generate code from a codegen.yaml
    help     ?      Get help with a specific subcommand
    version  v      Print the version to the terminal.
```

To parse an example (this example renders one randomization):

```bash
$ go run main.go gen examples/cpp/simple/codegen.yaml 
```
```bash
// Printing [0:foo.h]
#pragma once

#include <cstdint>

void Function(unsigned int fpIntOnnjnhtahwkfzet, double * fpFloatEsybudtnmugtrh, signed __int128 fpIntZaziszlugirfp, _Complex long double * fpFloatIrwlzglwed, _Complex long double * fpFloatHhyjeeohgnerv, unsigned long long fpIntIvfzglggniy);
// Printing [1:foo.c]
#include <cstdio>
#include <ostream>
#include <iostream>
#include "foo.h"

void Function(unsigned int fpIntOnnjnhtahwkfzet, double * fpFloatEsybudtnmugtrh, signed __int128 fpIntZaziszlugirfp, _Complex long double * fpFloatIrwlzglwed, _Complex long double * fpFloatHhyjeeohgnerv, unsigned long long fpIntIvfzglggniy) {

     std::cout <<  fpIntOnnjnhtahwkfzet << std::endl;
     std::cout << &fpFloatEsybudtnmugtrh << std::endl;
     std::cout <<  fpIntZaziszlugirfp << std::endl;
     std::cout << &fpFloatIrwlzglwed << std::endl;
     std::cout << &fpFloatHhyjeeohgnerv << std::endl;
     std::cout <<  fpIntIvfzglggniy << std::endl;

}
// Printing [2:main.c]
#include "foo.h"
#include <iostream>

int main() {

     // Initialize each formal param
     unsigned int fpIntOnnjnhtahwkfzet = 886107401;
     double * fpFloatEsybudtnmugtrh = 88392455.99064827235;
     signed __int128 fpIntZaziszlugirfp;
     fpIntZaziszlugirfp = 0x4311818582100267;
     fpIntZaziszlugirfp << 64;
     fpIntZaziszlugirfp = 0x5577737429679036;

     _Complex long double * fpFloatIrwlzglwed = 3878151789.988004017;
     _Complex long double * fpFloatHhyjeeohgnerv = 10.745834759950897994;
     unsigned long long fpIntIvfzglggniy = 17840965526490475249;

     // bigcall(1, 2, 3, 4, 5, bigthing);
     Function(fpIntOnnjnhtahwkfzet, fpFloatEsybudtnmugtrh, fpIntZaziszlugirfp, fpFloatIrwlzglwed, fpFloatHhyjeeohgnerv, fpIntIvfzglggniy);
}
```

If you don't provide an output directory (to be shown next) only one example will be printed to the screen. To provide an output
directory, you can do:

```bash
$ go run main.go gen examples/cpp/simple/codegen.yaml --outdir examples/cpp/simple
```

That will create subfolders there, numbered from 1..N where N is the number that you've asked for in your codegen.yaml.

```bash
$ tree examples/cpp/simple/1/
examples/cpp/simple/1/
├── codegen.yaml
├── foo.c
├── foo.h
├── main.c
└── Makefile

0 directories, 5 files
```

You can also override the render type. For example, the codegen.yaml use in the example above has type "random:1", which says
to generate one random subfolder. Here is how to change that:

```bash
$ go run main.go gen examples/cpp/simple/codegen.yaml --outdir test/ --render-type random:5
```

### Writing a Template

A template is a folder with a codegen.yaml file and one or more files that are to be filled in (templates).
As an example, let's look at [examples/cpp/simple](examples/cpp/simple). For any template, you should include:


```cpp
#include <ostream>
#include <iostream>
```

if you use any of the printing functions.

#### codegen.yaml
The codegen.yaml file is going to tell us the following:

```yaml
generate:

  # Currently only cpp is supported!
  language: "cpp"

  # Files to parse during rendering (in the same directory as the codegen.yaml)
  files:
    - foo.h
    - foo.c
    - main.c

  # Generation type can be random with a number, or (something more controlled without replacement?)
  # the default of random is to generate one (e.g., "random:1" and you can imagine increasing this (e.g., random:100)
  type: "random"
      
  # Functions and other types (not yet implemented) to generate for the templates
  render:

    # "Function" is the identifier for this specific function to use in the template, e.g., {{ .Function }}
    Function:
      type: "function"
      parameters:
        min: 1
        max: 10
```

In the above, wherever we find reference to `{{ .Function }}` we will know it's a function type, and will have 1 to 10 parameters,
randomly generated from all the float and integral types that we understand. For each of the files in the listing there,
they will be parsed and populated with this information. For some examples, let's look through at an example.

#### main.c

The template looks like this:

```cpp
#include "struct.h"

int main() {

     // Initialize each formal param
     // This says "look up the function named "Function" and make declarations for all its required arguments (formal params)
{{ .Function | DeclareArgs }}

     // This says "look up the function named "Function" and print it's call with those same params
     // bigcall(1, 2, 3, 4, 5, bigthing);
     {{ .Function | CallFunction }}
}
```

You'll notice that things to be rendered appear in double brackets `{{ }}` and when we want to pass a name
to a function, we use a pipe. Thus, you need to know the functions that are supported for each type!

| Type | Function Name | Description |
|------|---------------|-------------|
|function| AsFormalParams | The comma separated list to write into a function declaration (e.g., without values defined) |
|function| CallFunction| Print the call to the function with all named parameters |
|function| DeclareArgs | Make multi-line declarations of parameters to pass into the function |
|function| GetFunctionName | Just print the name of the function |
|function| PrintArgs | just do a println of each named param, usually for debugging |

You'll also notice there is a Makefile in the folder - likely when we generate many of these
and save the renderings somewhere, we will copy over all the content here (including the Makefile, which
doesn't need parsing) to somewhere else to use.

## Notes

These are notes from [@thaines](https://github.com/thaines) about what we might want to
originally produce for testing cases:

```bash
# signedness = {unsigned, signed}
# integrals = {char, short, int, long, long long, size_t, __int128}
# cardinality = 14

# floats = {float, double, long double}
# complex = {_Complex}
# cardinality = 6

# total cardinality = 20

# Let's have up to 10 parameters, so nCr(20, 10) = 184756
```

## TODO

 - I'm not sure how _Complex works (need help initializing them)
 - need to be able to print an int128 without an error
 - double pointers seem to give me trouble (need help here too)

### License

This project is part of Spack. Spack is distributed under the terms of both the MIT license and the Apache License (Version 2.0). Users may choose either license, at their option.

All new contributions must be made under both the MIT and Apache-2.0 licenses.

See LICENSE-MIT, LICENSE-APACHE, COPYRIGHT, and NOTICE for details.

SPDX-License-Identifier: (Apache-2.0 OR MIT)

LLNL-CODE-811652
(base) vanessa@van
