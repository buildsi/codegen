#pragma once

#include <cstdint>

void {{ .Function | GetFunctionName }}({{ .Function | AsFormalParams }});
