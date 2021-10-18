#include "foo.h"
#include <iostream>
#include <string>

int main() {
{{ .FunctionA | DeclareArgs }}
     {{ .FunctionA | CallFunction }}
{{ .FunctionB | DeclareArgs }}
     {{ .FunctionB | CallFunction }}
}
