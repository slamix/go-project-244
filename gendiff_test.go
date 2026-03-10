package code

import (
	"os"
	"testing"

	"code/internal/parser"

	"github.com/stretchr/testify/assert"
)

func TestGenDiffStylish(t *testing.T) {
	data1, err := parser.Parse("testdata/fixture/nested_file1.json")
	assert.NoError(t, err)

	data2, err := parser.Parse("testdata/fixture/nested_file2.json")
	assert.NoError(t, err)

	expected, err := os.ReadFile("testdata/fixture/expected_nested_diff.txt")
	assert.NoError(t, err)

	result, err := GenDiff(data1, data2, "stylish")
	assert.NoError(t, err)
	assert.Equal(t, string(expected), result)
}

func TestGenDiffStylishYAML(t *testing.T) {
	data1, err := parser.Parse("testdata/fixture/nested_file1.yml")
	assert.NoError(t, err)

	data2, err := parser.Parse("testdata/fixture/nested_file2.yml")
	assert.NoError(t, err)

	expected, err := os.ReadFile("testdata/fixture/expected_nested_diff.txt")
	assert.NoError(t, err)

	result, err := GenDiff(data1, data2, "stylish")
	assert.NoError(t, err)
	assert.Equal(t, string(expected), result)
}

func TestGenDiffPlain(t *testing.T) {
	data1, err := parser.Parse("testdata/fixture/nested_file1.json")
	assert.NoError(t, err)

	data2, err := parser.Parse("testdata/fixture/nested_file2.json")
	assert.NoError(t, err)

	expected, err := os.ReadFile("testdata/fixture/expected_plain_diff.txt")
	assert.NoError(t, err)

	result, err := GenDiff(data1, data2, "plain")
	assert.NoError(t, err)
	assert.Equal(t, string(expected), result)
}

func TestGenDiffJSON(t *testing.T) {
	data1, err := parser.Parse("testdata/fixture/nested_file1.json")
	assert.NoError(t, err)

	data2, err := parser.Parse("testdata/fixture/nested_file2.json")
	assert.NoError(t, err)

	expected, err := os.ReadFile("testdata/fixture/expected_json_diff.txt")
	assert.NoError(t, err)

	result, err := GenDiff(data1, data2, "json")
	assert.NoError(t, err)
	assert.Equal(t, string(expected), result)
}
