#include <cstdio>
#include <ostream>
#include <iostream>
#include <string>
#include "foo.h"
#include <assert.h>

void {{ .Function | GetFunctionName }}({{ .Function | AsFormalParams }}) {

{{ .Function | AssertArgs }}
}
