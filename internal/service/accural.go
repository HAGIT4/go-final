package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type GetOrderInfoResponse struct {
	Order   int     `json:"order"`
	Status  string  `json:"status"`
	Accural float32 `json:"accural,omitempty"`
}

type accuralClientInterface interface {
	GetOrderInfo(number int) (resp *GetOrderInfoResponse, err error)
}

type accuralClient struct {
	accuralAddress string
}

func (cl *accuralClient) GetOrderInfo(number int) (resp *GetOrderInfoResponse, err error) {
	url := fmt.Sprintf("%s/api/orders/%d", cl.accuralAddress, number)
	getResp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	switch getResp.StatusCode {
	case 200:
		body, err := io.ReadAll(getResp.Body)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(body, resp)
		if err != nil {
			return nil, err
		}
		return resp, nil
	case 429:
		err = errors.New("too many requests")
		return nil, err
	case 500:
		err = errors.New("internal error")
		return nil, err
	default:
		err = errors.New("unknown code")
		return nil, err
	}
}
