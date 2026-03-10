package formatters

import "fmt"

type DiffNode struct {
	Key      string
	Type     string
	Value    interface{}
	OldValue interface{}
	NewValue interface{}
	Children []DiffNode
}

func Format(nodes []DiffNode, format string) (string, error) {
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
