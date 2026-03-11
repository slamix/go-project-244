package formatters

import (
	"encoding/json"
	"fmt"
)

func FormatJSON(nodes []DiffNode) (string, error) {
	data, err := json.MarshalIndent(toJSONMap(nodes), "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal diff to JSON: %w", err)
	}
	return string(data), nil
}

func toJSONMap(nodes []DiffNode) map[string]interface{} {
	result := make(map[string]interface{})
	for _, n := range nodes {
		node := map[string]interface{}{
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
			node["children"] = toJSONMap(n.Children)
		}
		result[n.Key] = node
	}
	return result
}
