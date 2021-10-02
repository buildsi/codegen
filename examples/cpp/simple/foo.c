#include <cstdio>
#include "struct.h"

void {{ .Function | GetFunctionName }}({{ .Function | AsFormalParams }}) {

{{ .Function | PrintArgs }}
}
