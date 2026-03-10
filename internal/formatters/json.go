package formatters

import (
	"encoding/json"
	"fmt"
)

func FormatJSON(nodes []DiffNode) (string, error) {
	data, err := json.MarshalIndent(toJSONNodes(nodes), "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal diff to JSON: %w", err)
	}
	return string(data), nil
}

func toJSONNodes(nodes []DiffNode) []map[string]interface{} {
	result := make([]map[string]interface{}, 0, len(nodes))
	for _, n := range nodes {
		node := map[string]interface{}{
			"key":  n.Key,
			"type": n.Type,
		}
		switch n.Type {
		case "unchanged":
			node["value"] = n.Value
		case "added":
			node["new_value"] = n.NewValue
		case "removed":
			node["old_value"] = n.OldValue
		case "changed":
			node["old_value"] = n.OldValue
			node["new_value"] = n.NewValue
		case "nested":
			node["children"] = toJSONNodes(n.Children)
		}
		result = append(result, node)
	}
	return result
}
