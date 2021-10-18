#include <cstdio>
#include <ostream>
#include <iostream>
#include "foo.h"

{{ .FunctionA | GetReturnType }} {{ .FunctionA | GetFunctionName }}({{ .FunctionA | AsFormalParams }}) {
    return {{ .FunctionA | AddArgs }}
}

{{ .FunctionB | GetReturnType }} {{ .FunctionB | GetFunctionName }}({{ .FunctionB | AsFormalParams }}) {
    return {{ .FunctionB | AddArgs }}
}
