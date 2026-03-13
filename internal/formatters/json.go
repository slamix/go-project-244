package formatters

import (
	"encoding/json"
	"fmt"

	"code/internal/diff"
)

func FormatJSON(nodes []diff.DiffNode) (string, error) {
	data, err := json.MarshalIndent(toJSONMap(nodes), "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal diff to JSON: %w", err)
	}
	return string(data), nil
}

func toJSONMap(nodes []diff.DiffNode) map[string]interface{} {
	result := make(map[string]interface{})
	for _, n := range nodes {
		node := map[string]interface{}{
			"type": n.Type,
		}
		switch n.Type {
		case diff.NodeUnchanged:
			node["value"] = n.Value
		case diff.NodeAdded:
			node["new_value"] = n.NewValue
		case diff.NodeRemoved:
			node["old_value"] = n.OldValue
		case diff.NodeChanged:
			node["old_value"] = n.OldValue
			node["new_value"] = n.NewValue
		case diff.NodeNested:
			node["children"] = toJSONMap(n.Children)
		}
		result[n.Key] = node
	}
	return result
}
