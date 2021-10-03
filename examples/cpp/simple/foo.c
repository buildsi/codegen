#include <cstdio>
#include <ostream>
#include <iostream>
#include "foo.h"

void {{ .Function | GetFunctionName }}({{ .Function | AsFormalParams }}) {

{{ .Function | PrintArgs }}
}
