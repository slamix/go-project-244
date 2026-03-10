package code

import (
	"fmt"
	"sort"

	"code/internal/formatters"
)

func BuildDiff(data1, data2 map[string]interface{}) []formatters.DiffNode {
	keys := collectKeys(data1, data2)
	sort.Strings(keys)

	var nodes []formatters.DiffNode
	for _, key := range keys {
		val1, in1 := data1[key]
		val2, in2 := data2[key]

		if in1 && in2 {
			m1, isMap1 := val1.(map[string]interface{})
			m2, isMap2 := val2.(map[string]interface{})
			switch {
			case isMap1 && isMap2:
				nodes = append(nodes, formatters.DiffNode{Key: key, Type: "nested", Children: BuildDiff(m1, m2)})
			case fmt.Sprint(val1) == fmt.Sprint(val2):
				nodes = append(nodes, formatters.DiffNode{Key: key, Type: "unchanged", Value: val1})
			default:
				nodes = append(nodes, formatters.DiffNode{Key: key, Type: "changed", OldValue: val1, NewValue: val2})
			}
		} else if in1 {
			nodes = append(nodes, formatters.DiffNode{Key: key, Type: "removed", OldValue: val1})
		} else {
			nodes = append(nodes, formatters.DiffNode{Key: key, Type: "added", NewValue: val2})
		}
	}
	return nodes
}

func GenDiff(data1, data2 map[string]interface{}, format ...string) (string, error) {
	f := "stylish"
	if len(format) > 0 && format[0] != "" {
		f = format[0]
	}
	diff := BuildDiff(data1, data2)
	return formatters.Format(diff, f)
}

func collectKeys(data1, data2 map[string]interface{}) []string {
	seen := make(map[string]bool)
	for k := range data1 {
		seen[k] = true
	}
	for k := range data2 {
		seen[k] = true
	}
	keys := make([]string, 0, len(seen))
	for k := range seen {
		keys = append(keys, k)
	}
	return keys
}
