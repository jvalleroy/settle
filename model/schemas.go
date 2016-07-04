// OWNER: stan

package model

import (
	"fmt"

	"github.com/spolu/peer-currencies/lib/errors"
)

var schemas = map[string]map[string]string{
	"mint": map[string]string{},
}

const (
	toolsSQL = `
CREATE OR REPLACE FUNCTION utc_now() RETURNS TIMESTAMP AS $$
  SELECT CLOCK_TIMESTAMP() AT TIME ZONE 'utc'
$$ language sql;
`
)

func init() {
	ensureMintDB()
	RegisterSchema(
		"mint",
		"_tools",
		toolsSQL,
	)
}

// RegisterSchema lets schemas register themselves.
func RegisterSchema(
	db string,
	table string,
	schema string,
) {
	schemas[db][table] = schema
}

// CreateMintDBTables creates the Mint DB tables if they don't exist.
func CreateMintDBTables() error {
	ensureMintDB()
	for name, sch := range schemas["mint"] {
		fmt.Printf("Executing schema: %s\n", name)
		_, err := mintDB.Exec(sch)
		if err != nil {
			return errors.Trace(err)
		}
	}
	return nil
}
