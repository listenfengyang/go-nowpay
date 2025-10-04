package go_nowpay

import (
	"github.com/go-resty/resty/v2"
	"github.com/listenfengyang/go-nowpay/utils"
)

type Client struct {
	Params *NowPayInitParams

	ryClient  *resty.Client
	debugMode bool //是否调试模式
	logger    utils.Logger
}

func NewClient(logger utils.Logger, params *NowPayInitParams) *Client {
	return &Client{
		Params: params,

		ryClient:  resty.New(), //client实例
		debugMode: false,
		logger:    logger,
	}
}

func (cli *Client) SetDebugModel(debugModel bool) {
	cli.debugMode = debugModel
}

func (cli *Client) SetMerchantInfo(merchant MerchantInfo) {
	cli.Params.MerchantInfo = merchant
}
