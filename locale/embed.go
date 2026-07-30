package locale

import "embed"

//go:embed translations/*.json
var LocalFS embed.FS
