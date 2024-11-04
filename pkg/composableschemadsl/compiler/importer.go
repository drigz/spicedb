package compiler

import (
	"fmt"
	"os"
	"path"

	"github.com/authzed/spicedb/pkg/composableschemadsl/input"
	"github.com/authzed/spicedb/pkg/genutil/mapz"
	"github.com/rs/zerolog/log"
)

const SchemaFileSuffix = ".zed"

func ImportFile(pathSegments []string, existingNames *mapz.Set[string]) (*CompiledSchema, error) {
	filepath := constructFilePath(pathSegments)

	var schemaBytes []byte
	schemaBytes, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to read schema file: %w", err)
	}
	// TODO: do we want this sort of logging here?
	log.Trace().Str("schema", string(schemaBytes)).Str("file", filepath).Msg("read schema from file")

	compiled, err := Compile(InputSchema{
		// TODO: should this point to the schema file? What is this for?
		Source:       input.Source("schema"),
		SchemaString: string(schemaBytes),
	}, AllowUnprefixedObjectType(), ExistingNames(existingNames.AsSlice()))
	if err != nil {
		return nil, err
	}
	return compiled, nil
}

func constructFilePath(segments []string) string {
	return path.Join(segments...) + SchemaFileSuffix
}
