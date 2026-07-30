package locale

import "embed"

// Files loaded for localization.
//
//go:embed translations/*.json
var LocalFS embed.FS
