package go_nowpay

import (
	"encoding/json"
	"fmt"
	"testing"
)

type VLog struct {
}

func (l VLog) Debugf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Infof(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Warnf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Errorf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}

func TestCallback(t *testing.T) {
	vLog := VLog{}
	//构造client
	cli := NewClient(vLog, &NowPayInitParams{MerchantInfo{MERCHANT_ID, ACCESS_KEY, BACK_KEY}, DEPOSIT_URL, WITHDRAW_URL})

	//1. 获取请求
	req := GenCallbackRequestDemo() //提现的返回
	var backReq NowPayDepositCallbackReq
	err := json.Unmarshal([]byte(req), &backReq)
	if err != nil {
		cli.logger.Errorf("Error:%s", err.Error())
		return
	}

	//2. 处理请求
	//err = cli.DepositCallback(backReq, func(NowPayDepositCallbackReq) error { return nil })
	//if err != nil {
	//	cli.logger.Errorf("Error:%s", err.Error())
	//	return
	//}
	//cli.logger.Infof("resp:%+v\n", backReq)

	// 处理取消回调
	err = cli.DepositCanceledCallback(backReq, func(NowPayDepositCallbackReq) error { return nil })
	if err != nil {
		cli.logger.Errorf("Error:%s", err.Error())
		return
	}
	cli.logger.Infof("resp:%+v\n", backReq)
}

func GenCallbackRequestDemo() string {
	// bill_status 1=订单已取消 2=订单已激活
	return `{"sign":"2c89857a90e2773f27583d954c91f40c","bill_no":"2025100443562675418","bill_status":1,"sys_no":"505299"}`
}
