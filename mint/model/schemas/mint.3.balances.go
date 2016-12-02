// OWNER: stan

package schemas

import "github.com/spolu/settle/mint/model"

const (
	balancesSQL = `
CREATE TABLE IF NOT EXISTS balances(
  owner VARCHAR(256) NOT NULL,  -- owner address
  token VARCHAR(256) NOT NULL,  -- token
  created TIMESTAMP NOT NULL,

  asset VARCHAR(256) NOT NULL,  -- asset name
  holder VARCHAR(256) NOT NULL, -- balance holder address
  value VARCHAR(64) NOT NULL,   -- balance value

  PRIMARY KEY(owner, token),
  CONSTRAINT balances_asset_holder_u UNIQUE (asset, holder) -- not propagated
);
`
)

func init() {
	model.RegisterSchema(
		"mint",
		"balances",
		balancesSQL,
	)
}
