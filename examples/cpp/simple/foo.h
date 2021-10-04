#pragma once

#include <cstdint>

// Structs used in the function should be declared first
{{ .Function | DeclareStructs }} 

void {{ .Function | GetFunctionName }}({{ .Function | AsFormalParams }});
