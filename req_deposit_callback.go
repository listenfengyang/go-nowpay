package go_nowpay

import (
	"encoding/json"
	"errors"
	"github.com/listenfengyang/go-nowpay/utils"
	"github.com/mitchellh/mapstructure"
)

// 充值-成功回调
func (cli *Client) DepositCallback(req NowPayDepositCallbackReq, processor func(NowPayDepositCallbackReq) error) error {
	//验证签名
	var params map[string]string
	//mapstructure.Decode(req, &params)
	params = map[string]string{"bill_no": req.BillNo, "sign": req.Sign}

	flag := utils.VerifyDepositCallback(params, cli.Params.BackKey)
	if !flag {
		//签名校验失败
		reqJson, _ := json.Marshal(req)
		cli.logger.Errorf("nowPay deposit back verify fail, req: %s", string(reqJson))
		return errors.New("sign verify error")
	}

	//开始处理
	return processor(req)
}

// 充值-取消回调
func (cli *Client) DepositCanceledCallback(req NowPayDepositCallbackReq, processor func(NowPayDepositCallbackReq) error) error {
	//验证签名
	var params map[string]string
	mapstructure.Decode(req, &params)
	delete(params, "amount")
	delete(params, "amount_usdt")

	flag := utils.VerifyDepositCallback(params, cli.Params.BackKey)
	if !flag {
		//签名校验失败
		reqJson, _ := json.Marshal(req)
		cli.logger.Errorf("nowPay deposit canceled back verify fail, req: %s", string(reqJson))
		return errors.New("sign verify error")
	}

	//开始处理
	return processor(req)
}
