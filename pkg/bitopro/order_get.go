package bitopro

import (
	"encoding/json"
	"fmt"

	"github.com/bitoex/bitopro-api-go/internal"
)

// OrderInfo struct
type OrderInfo struct {
	ID                string `json:"id"`
	Pair              string `json:"pair"`
	Price             string `json:"price"`
	AvgExecutionPrice string `json:"avgExecutionPrice"`
	Action            string `json:"action"`
	Type              string `json:"type"`
	Timestamp         string `json:"timestamp"`
	CreatedTimestamp  int64  `json:"createdTimestamp"`
	UpdatedTimestamp  int64  `json:"updatedTimestamp"`
	Status            int    `json:"status"`
	OriginalAmount    string `json:"originalAmount"`
	RemainingAmount   string `json:"remainingAmount"`
	ExecutedAmount    string `json:"executedAmount"`
	Fee               string `json:"fee"`
	FeeSymbol         string `json:"feeSymbol"`
	BitoFee           string `json:"bitoFee"`
	Total             string `json:"total"`
	Seq               string `json:"seq"`
	TimeInForce       string `json:"timeInForce"`
	StatusCode
}

// GetOrder Ref. https://developer.bitopro.com/docs#operation/getOrderStatus
func (api *AuthAPI) GetOrder(pair string, orderID int) *OrderInfo {
	var data OrderInfo

	code, res := internal.ReqWithoutBody(api.identity, api.key, api.secret, "GET", fmt.Sprintf("%s/%s/%d", "v3/orders", pair, orderID))

	if err := json.Unmarshal([]byte(res), &data); err != nil {
		data.Error = res
	}

	data.Code = code

	return &data
}
