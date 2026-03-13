package formatters

import (
	"fmt"

	"code/internal/diff"
)

func Format(nodes []diff.DiffNode, format string) (string, error) {
	switch format {
	case "stylish", "":
		return FormatStylish(nodes), nil
	case "plain":
		return FormatPlain(nodes), nil
	case "json":
		return FormatJSON(nodes)
	default:
		return "", fmt.Errorf("unknown format: %s", format)
	}
}
