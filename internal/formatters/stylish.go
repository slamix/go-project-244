package formatters

import (
	"fmt"
	"sort"
	"strings"
)

func FormatStylish(nodes []DiffNode) string {
	return "{\n" + renderNodes(nodes, 1) + "\n}"
}

func renderNodes(nodes []DiffNode, depth int) string {
	pad := strings.Repeat(" ", depth*4-2)
	var lines []string
	for _, node := range nodes {
		switch node.Type {
		case "nested":
			lines = append(lines, fmt.Sprintf("%s  %s: {", pad, node.Key))
			lines = append(lines, renderNodes(node.Children, depth+1))
			lines = append(lines, strings.Repeat(" ", depth*4)+"}")
		case "unchanged":
			lines = append(lines, fmt.Sprintf("%s  %s: %s", pad, node.Key, renderValue(node.Value, depth)))
		case "added":
			lines = append(lines, fmt.Sprintf("%s+ %s: %s", pad, node.Key, renderValue(node.NewValue, depth)))
		case "removed":
			lines = append(lines, fmt.Sprintf("%s- %s: %s", pad, node.Key, renderValue(node.OldValue, depth)))
		case "changed":
			lines = append(lines, fmt.Sprintf("%s- %s: %s", pad, node.Key, renderValue(node.OldValue, depth)))
			lines = append(lines, fmt.Sprintf("%s+ %s: %s", pad, node.Key, renderValue(node.NewValue, depth)))
		}
	}
	return strings.Join(lines, "\n")
}

func renderValue(val interface{}, depth int) string {
	if val == nil {
		return "null"
	}
	m, isMap := val.(map[string]interface{})
	if !isMap {
		return fmt.Sprint(val)
	}
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var lines []string
	lines = append(lines, "{")
	for _, k := range keys {
		innerPad := strings.Repeat(" ", (depth+1)*4)
		lines = append(lines, fmt.Sprintf("%s%s: %s", innerPad, k, renderValue(m[k], depth+1)))
	}
	lines = append(lines, strings.Repeat(" ", depth*4)+"}")
	return strings.Join(lines, "\n")
}
