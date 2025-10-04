package go_nowpay

import (
	"encoding/json"
	"errors"
	"github.com/listenfengyang/go-nowpay/utils"
	"github.com/mitchellh/mapstructure"
	"log"
)

// 出金-成功回调
func (cli *Client) WithdrawCallback(req NowPayWithdrawCallbackReq, processor func(req NowPayWithdrawCallbackReq) error) error {
	//验证签名
	var params map[string]string
	params = map[string]string{"bill_no": req.BillNo}

	// Verify signature
	flag, err := utils.Verify(params, cli.Params.BackKey)
	if err != nil {
		log.Printf("Signature verification error: %v", err)
		return err
	}
	if !flag {
		//签名校验失败
		reqJson, _ := json.Marshal(req)
		cli.logger.Errorf("nowpay successfull back verify fail, req: %s", string(reqJson))
		return errors.New("sign verify error")
	}

	//开始处理
	return processor(req)
}

// 出金-取消回调
func (cli *Client) WithdrawCanceledCallback(req NowPayWithdrawCallbackReq, processor func(req NowPayWithdrawCallbackReq) error) error {
	//验证签名
	var params map[string]string
	mapstructure.Decode(req, &params)
	delete(params, "amount")
	delete(params, "amount_usdt")

	// Verify signature
	flag, err := utils.Verify(params, cli.Params.BackKey)
	if err != nil {
		log.Printf("Signature canceled verification error: %v", err)
		return err
	}
	if !flag {
		//签名校验失败
		reqJson, _ := json.Marshal(req)
		cli.logger.Errorf("nowpay canceled back verify fail, req: %s", string(reqJson))
		return errors.New("sign canceled verify error")
	}

	//开始处理
	return processor(req)
}
