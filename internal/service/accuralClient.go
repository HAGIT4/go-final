package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type GetOrderInfoResponse struct {
	Order   string  `json:"order"`
	Status  string  `json:"status"`
	Accural float32 `json:"accrual,omitempty"`
	Action  string
}

type accuralClientInterface interface {
	GetOrderInfo(number int) (resp *GetOrderInfoResponse)
}

type accuralClient struct {
	accuralAddress string
}

func NewAccuralClient(address string) (client *accuralClient, err error) {
	client = &accuralClient{
		accuralAddress: address,
	}
	return client, nil
}

func (cl *accuralClient) GetOrderInfo(number int) (resp *GetOrderInfoResponse) {
	resp = &GetOrderInfoResponse{}
	url := fmt.Sprintf("http://%s/api/orders/%d", cl.accuralAddress, number)
	getResp, err := http.Get(url)
	if err != nil {
		resp.Action = "retry"
		return resp
	}
	switch getResp.StatusCode {
	case 200:
		body, err := io.ReadAll(getResp.Body)
		if err != nil {
			resp.Action = "retry"
			return resp
		}
		err = json.Unmarshal(body, resp)
		if err != nil {
			resp.Action = "retry"
			return resp
		}
		resp.Action = "ok"
		return resp
	case 429:
		resp.Action = "retry"
		return resp
	case 500:
		resp.Action = "retry"
		return resp
	default:
		resp.Action = "retry"
		return resp
	}
}
