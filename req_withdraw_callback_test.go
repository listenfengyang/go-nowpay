package go_nowpay

import (
	"encoding/json"
	"testing"
)

func TestWithdrawCallback(t *testing.T) {
	vLog := VLog{}
	//构造client
	cli := NewClient(vLog, &NowPayInitParams{MerchantInfo{MERCHANT_ID, ACCESS_KEY, BACK_KEY}, DEPOSIT_URL, WITHDRAW_URL})

	//1. 获取请求
	req := GenWdCallbackRequestDemo() //提现的返回
	var backReq NowPayWithdrawCallbackReq
	err := json.Unmarshal([]byte(req), &backReq)
	if err != nil {
		cli.logger.Errorf("Error:%s", err.Error())
		return
	}

	//2. 处理请求
	//err = cli.WithdrawCallback(backReq, func(NowPayWithdrawCallbackReq) error { return nil })
	//if err != nil {
	//	cli.logger.Errorf("Error:%s", err.Error())
	//	return
	//}
	//cli.logger.Infof("resp:%+v\n", backReq)

	//2. 处理请求
	err = cli.WithdrawCanceledCallback(backReq, func(NowPayWithdrawCallbackReq) error { return nil })
	if err != nil {
		cli.logger.Errorf("Error:%s", err.Error())
		return
	}
	cli.logger.Infof("resp:%+v\n", backReq)
}

func GenWdCallbackRequestDemo() string {
	//return `{"amount":"1","sign":"45fc45fb8b1115f5ce339e15328bed56","bill_no":"2025100417244604954","sys_no":"505299"}`
	return `{"bill_no":"202510090656090207","bill_status":1,"sign":"f6f7f6d0540167f7241c56df4bb318ee","sys_no":"505299","amount":""}`
}
