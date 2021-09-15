package generator

import (
	"testing"

	"github.com/stretchr/testify/require"

	v0 "github.com/authzed/spicedb/internal/proto/authzed/api/v0"
	"github.com/authzed/spicedb/pkg/namespace"
	"github.com/authzed/spicedb/pkg/schemadsl/compiler"
	"github.com/authzed/spicedb/pkg/schemadsl/input"
)

func TestGenerator(t *testing.T) {
	type generatorTest struct {
		name     string
		input    *v0.NamespaceDefinition
		expected string
		okay     bool
	}

	tests := []generatorTest{
		{
			"empty",
			namespace.Namespace("foos/test"),
			"definition foos/test {}",
			true,
		},
		{
			"simple relation",
			namespace.Namespace("foos/test",
				namespace.Relation("somerel", nil, &v0.RelationReference{
					Namespace: "foos/bars",
					Relation:  "hiya",
				}),
			),
			`definition foos/test {
	relation somerel: foos/bars#hiya
}`,
			true,
		},
		{
			"simple permission",
			namespace.Namespace("foos/test",
				namespace.Relation("someperm", namespace.Union(
					namespace.ComputedUserset("anotherrel"),
				)),
			),
			`definition foos/test {
	permission someperm = anotherrel
}`,
			true,
		},
		{
			"complex permission",
			namespace.Namespace("foos/test",
				namespace.Relation("someperm", namespace.Union(
					namespace.Rewrite(
						namespace.Exclusion(
							namespace.ComputedUserset("rela"),
							namespace.ComputedUserset("relb"),
							namespace.TupleToUserset("rely", "relz"),
						),
					),
					namespace.ComputedUserset("relc"),
				)),
			),
			`definition foos/test {
	permission someperm = (rela - relb - rely->relz) + relc
}`,
			true,
		},
		{
			"legacy relation",
			namespace.Namespace("foos/test",
				namespace.Relation("somerel", namespace.Union(
					namespace.This(),
					namespace.ComputedUserset("anotherrel"),
				), &v0.RelationReference{
					Namespace: "foos/bars",
					Relation:  "hiya",
				}),
			),
			`definition foos/test {
	relation somerel: foos/bars#hiya = /* _this unsupported here. Please rewrite into a relation and permission */ + anotherrel
}`,
			false,
		},
		{
			"missing type information",
			namespace.Namespace("foos/test",
				namespace.Relation("somerel", nil),
			),
			`definition foos/test {
	relation somerel: /* missing allowed types */
}`,
			false,
		},

		{
			"full example",
			namespace.NamespaceWithComment("foos/document", `/**
* Some comment goes here
*/`,
				namespace.Relation("owner", nil,
					&v0.RelationReference{
						Namespace: "foos/user",
						Relation:  "...",
					},
				),
				namespace.RelationWithComment("reader", "//foobar", nil,
					&v0.RelationReference{
						Namespace: "foos/user",
						Relation:  "...",
					},
					&v0.RelationReference{
						Namespace: "foos/group",
						Relation:  "member",
					},
				),
				namespace.Relation("read", namespace.Union(
					namespace.ComputedUserset("reader"),
					namespace.ComputedUserset("owner"),
				)),
			),
			`/** Some comment goes here */
definition foos/document {
	relation owner: foos/user

	// foobar
	relation reader: foos/user | foos/group#member
	permission read = reader + owner
}`,
			true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require := require.New(t)
			source, ok := GenerateSource(test.input)
			require.Equal(test.expected, source)
			require.Equal(test.okay, ok)
		})
	}
}

func TestFormatting(t *testing.T) {
	type formattingTest struct {
		name     string
		input    string
		expected string
	}

	tests := []formattingTest{
		{
			"empty",
			"definition foos/test {}",
			"definition foos/test {}",
		},
		{
			"with comment",
			`/** some def */definition foos/test {}`,
			`/** some def */
definition foos/test {}`,
		},
		{
			"with rel comment",
			`/** some def */definition foos/test {

				// some rel
				relation somerel: foos/bars;
			}`,
			`/** some def */
definition foos/test {
	// some rel
	relation somerel: foos/bars
}`,
		},
		{
			"with multiple rel comment",
			`/** some def */definition foos/test {

				// some rel
				/* another comment */
				relation somerel: foos/bars;
			}`,
			`/** some def */
definition foos/test {
	// some rel
	/* another comment */
	relation somerel: foos/bars
}`,
		},
		{
			"with multiple rels with comment",
			`/** some def */definition foos/test {

				// some rel
				relation somerel: foos/bars;
				// another perm
				permission someperm = somerel
			}`,
			`/** some def */
definition foos/test {
	// some rel
	relation somerel: foos/bars

	// another perm
	permission someperm = somerel
}`,
		},

		{
			"becomes single line comment",
			`definition foos/test {
				/**
				 * hi there
				 */
				relation somerel: foos/bars;
			}`,
			`definition foos/test {
	/** hi there */
	relation somerel: foos/bars
}`,
		},

		{
			"full example",
			`
/** the document */
definition foos/document {
	/** some super long comment goes here and therefore should be made into a multiline comment */
	relation reader: foos/user

	/** multiline
comment */
	relation  writer: foos/user

	// writers are also readers
	permission read = reader + writer + another
	permission write = writer
	permission minus = rela - relb - relc
}
`,
			`/** the document */
definition foos/document {
	/**
	 * some super long comment goes here and therefore should be made into a multiline comment
	 */
	relation reader: foos/user

	/**
	 * multiline
	 * comment
	 */
	relation writer: foos/user

	// writers are also readers
	permission read = reader + writer + another
	permission write = writer
	permission minus = (rela - relb) - relc
}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require := require.New(t)
			defs, err := compiler.Compile([]compiler.InputSchema{
				{input.InputSource(test.name), test.input},
			}, nil)
			require.NoError(err)

			source, _ := GenerateSource(defs[0])
			require.Equal(test.expected, source)
		})
	}
}
