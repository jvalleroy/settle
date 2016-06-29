package mint

import (
	"fmt"
	"math/big"

	"github.com/spolu/settle/model"

	"golang.org/x/net/context"
)

// AssetResource is the representation of an asset in the mint API.
type AssetResource struct {
	ID       string `json:"id"`
	Created  int64  `json:"created"`
	Livemode bool   `json:"livemode"`

	Name   string `json:"name"`
	Issuer string `json:"issuer"`
	Code   string `json:"code"`
	Scale  int8   `json:"scale"`
}

// NewAssetResource generates a new resource given the required models and
// information.
func NewAssetResource(
	ctx context.Context,
	asset *model.Asset,
	issuer *model.User,
	host string,
) AssetResource {
	return AssetResource{
		ID:       asset.Token,
		Created:  asset.Created.UnixNano() / (1000 * 1000),
		Livemode: asset.Livemode,
		Name: fmt.Sprintf(
			"%s@%s:%s.%d",
			issuer.Username, host, asset.Code, asset.Scale,
		),
		Issuer: fmt.Sprintf(
			"%s@%s", issuer.Username, host,
		),
		Code:  asset.Code,
		Scale: asset.Scale,
	}
}

// OperationResource is the representation of an operation in the mint API.
type OperationResource struct {
	ID       string `json:"id"`
	Created  int64  `json:"created"`
	Livemode bool   `json:"livemode"`

	Asset       AssetResource `json:"asset"`
	Source      *string       `json:"source"`
	Destination string        `json:"destination"`
	Amount      *big.Int      `json:"amount"`
}

// NewOperationResource generates a new resource given the required models and
// information.
func NewOperationResource(
	ctx context.Context,
	operation *model.Operation,
	assetResource AssetResource,
) OperationResource {
	return OperationResource{
		ID:          operation.Token,
		Created:     operation.Created.UnixNano() / (1000 * 1000),
		Livemode:    operation.Livemode,
		Asset:       assetResource,
		Source:      operation.Source,
		Destination: operation.Destination,
		Amount:      (*big.Int)(&operation.Amount),
	}
}