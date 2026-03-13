package formatters

import (
	"fmt"
	"strings"

	"code/internal/diff"
)

func FormatPlain(nodes []diff.DiffNode) string {
	return formatPlainNodes(nodes, "")
}

func formatPlainNodes(nodes []diff.DiffNode, path string) string {
	var lines []string
	for _, node := range nodes {
		fullPath := node.Key
		if path != "" {
			fullPath = path + "." + node.Key
		}
		switch node.Type {
		case diff.NodeNested:
			inner := formatPlainNodes(node.Children, fullPath)
			if inner != "" {
				lines = append(lines, inner)
			}
		case diff.NodeAdded:
			lines = append(lines, fmt.Sprintf("Property '%s' was added with value: %s", fullPath, plainValue(node.NewValue)))
		case diff.NodeRemoved:
			lines = append(lines, fmt.Sprintf("Property '%s' was removed", fullPath))
		case diff.NodeChanged:
			lines = append(lines, fmt.Sprintf("Property '%s' was updated. From %s to %s", fullPath, plainValue(node.OldValue), plainValue(node.NewValue)))
		}
	}
	return strings.Join(lines, "\n")
}

func plainValue(val interface{}) string {
	if val == nil {
		return "null"
	}
	if _, isMap := val.(map[string]interface{}); isMap {
		return "[complex value]"
	}
	if s, isString := val.(string); isString {
		return "'" + s + "'"
	}
	return fmt.Sprint(val)
}
