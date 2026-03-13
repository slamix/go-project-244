package code

import (
	"fmt"
	"reflect"
	"sort"

	"code/internal/diff"
	"code/internal/formatters"
	"code/internal/parser"
)

func buildDiff(data1, data2 map[string]interface{}) []diff.DiffNode {
	keys := collectKeys(data1, data2)
	sort.Strings(keys)

	var nodes []diff.DiffNode
	for _, key := range keys {
		val1, in1 := data1[key]
		val2, in2 := data2[key]

		if in1 && in2 {
			m1, isMap1 := val1.(map[string]interface{})
			m2, isMap2 := val2.(map[string]interface{})
			switch {
			case isMap1 && isMap2:
				nodes = append(nodes, diff.DiffNode{Key: key, Type: diff.NodeNested, Children: buildDiff(m1, m2)})
			case reflect.DeepEqual(val1, val2):
				nodes = append(nodes, diff.DiffNode{Key: key, Type: diff.NodeUnchanged, Value: val1})
			default:
				nodes = append(nodes, diff.DiffNode{Key: key, Type: diff.NodeChanged, OldValue: val1, NewValue: val2})
			}
		} else if in1 {
			nodes = append(nodes, diff.DiffNode{Key: key, Type: diff.NodeRemoved, OldValue: val1})
		} else {
			nodes = append(nodes, diff.DiffNode{Key: key, Type: diff.NodeAdded, NewValue: val2})
		}
	}
	return nodes
}

func GenDiff(file1Path, file2Path string, format ...string) (string, error) {
	f := "stylish"
	if len(format) > 0 && format[0] != "" {
		f = format[0]
	}
	data1, err := parser.Parse(file1Path)
	if err != nil {
		return "", fmt.Errorf("parsing %s: %w", file1Path, err)
	}
	data2, err := parser.Parse(file2Path)
	if err != nil {
		return "", fmt.Errorf("parsing %s: %w", file2Path, err)
	}
	diff := buildDiff(data1, data2)
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
