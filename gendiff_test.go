package code

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenDiffStylish(t *testing.T) {
	expected, err := os.ReadFile("testdata/fixture/expected_nested_diff.txt")
	assert.NoError(t, err)

	result, err := GenDiff("testdata/fixture/nested_file1.json", "testdata/fixture/nested_file2.json", "stylish")
	assert.NoError(t, err)
	assert.Equal(t, string(expected), result)
}

func TestGenDiffStylishYAML(t *testing.T) {
	expected, err := os.ReadFile("testdata/fixture/expected_nested_diff.txt")
	assert.NoError(t, err)

	result, err := GenDiff("testdata/fixture/nested_file1.yml", "testdata/fixture/nested_file2.yml", "stylish")
	assert.NoError(t, err)
	assert.Equal(t, string(expected), result)
}

func TestGenDiffPlain(t *testing.T) {
	expected, err := os.ReadFile("testdata/fixture/expected_plain_diff.txt")
	assert.NoError(t, err)

	result, err := GenDiff("testdata/fixture/nested_file1.json", "testdata/fixture/nested_file2.json", "plain")
	assert.NoError(t, err)
	assert.Equal(t, string(expected), result)
}

func TestGenDiffJSON(t *testing.T) {
	expected, err := os.ReadFile("testdata/fixture/expected_json_diff.txt")
	assert.NoError(t, err)

	result, err := GenDiff("testdata/fixture/nested_file1.json", "testdata/fixture/nested_file2.json", "json")
	assert.NoError(t, err)
	assert.Equal(t, string(expected), result)
}
