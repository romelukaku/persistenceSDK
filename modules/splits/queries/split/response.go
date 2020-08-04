package split

import (
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappers"
)

type queryResponse struct {
	Splits mappers.Splits `json:"splits" valid:"required~required field splits missing"`
}

var _ helpers.QueryResponse = (*queryResponse)(nil)

func queryResponsePrototype() helpers.QueryResponse {
	return queryResponse{}
}

func newQueryResponse(splits mappers.Splits) helpers.QueryResponse {
	return queryResponse{Splits: splits}
}