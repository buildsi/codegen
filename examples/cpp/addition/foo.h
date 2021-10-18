#pragma once

#include <cstdint>
#include <string>

{{ .FunctionA | DeclareStructs }} 
{{ .FunctionA | GetReturnType }} {{ .FunctionA | GetFunctionName }}({{ .FunctionA | AsFormalParams }});

{{ .FunctionB | DeclareStructs }} 
{{ .FunctionB | GetReturnType }} {{ .FunctionB | GetFunctionName }}({{ .FunctionB | AsFormalParams }});
