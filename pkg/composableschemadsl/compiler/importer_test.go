package compiler_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/assert"

	"github.com/authzed/spicedb/pkg/composableschemadsl/input"
	"github.com/authzed/spicedb/pkg/composableschemadsl/compiler"
	"github.com/authzed/spicedb/pkg/composableschemadsl/generator"
)

type importerTest struct {
	name string
	folder string
}

func (it *importerTest) input() string {
	b, err := os.ReadFile(fmt.Sprintf("importer-test/%s/root.zed", it.folder))
	if err != nil {
		panic(err)
	}

	return string(b)
}

func (it *importerTest) expected() string {
	b, err := os.ReadFile(fmt.Sprintf("importer-test/%s/expected.zed", it.folder))
	if err != nil {
		panic(err)
	}

	return string(b)
}

func (it *importerTest) writeExpected(schema string) {
	err := os.WriteFile(fmt.Sprintf("importer-test/%s/expected.zed", it.folder), []byte(schema), 0o_600)
	if err != nil {
		panic(err)
	}
}

func TestImporter(t *testing.T) {
	importerTests := []importerTest{
		{"simple local import", "simple-local"},
		{"simple local import", "simple-local-with-hop"},
		{"nested local import", "nested-local"},
		{"nested local two layers deep import", "nested-two-layer-local"},
	}

	for _, test := range importerTests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			inputSchema := test.input()

			compiled, err := compiler.Compile(compiler.InputSchema{
				Source:       input.Source("schema"),
				SchemaString: inputSchema,
			}, compiler.AllowUnprefixedObjectType())
			require.NoError(t, err)

			generated, _, err := generator.GenerateSchema(compiled.OrderedDefinitions)
			require.NoError(t, err)

			if os.Getenv("REGEN") == "true" {
				test.writeExpected(generated)
			} else {
				expected := test.expected()
				if !assert.Equal(t, expected, generated, test.name) {
					t.Log(generated)
				}
			}
		})
	}

}
