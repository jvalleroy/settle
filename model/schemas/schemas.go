package model

import "github.com/spolu/settle/lib/errors"

var schemas = map[string]map[string]string{
	"mint": map[string]string{},
}

func registerSchema(
	db string,
	table string,
	schema string,
) {
	schemas[db][table] = schema
}

const (
	toolsSQL = `
CREATE OR REPLACE FUNCTION utc_now() RETURNS TIMESTAMP AS $$
  SELECT CLOCK_TIMESTAMP() AT TIME ZONE 'utc'
$$ language sql;
`
)

func init() {
	registerSchema(
		"mint",
		"_tools",
		toolsSQL,
	)
}

// CreateMintDBTables creates the Mint DB tables if they don't exist.
func CreateMintDBTables() error {
	for t, sch := range schemas["mint"] {
		_, err := mintDB.Exec(sch)
		if err != nil {
			return errors.Trace(err)
		}
	}
}
