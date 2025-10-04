package go_nowpay

import (
	"net/url"
	"testing"
	"time"
)

func TestDeposit(t *testing.T) {
	vLog := VLog{}
	//构造client
	cli := NewClient(vLog, &NowPayInitParams{MerchantInfo{MERCHANT_ID, ACCESS_KEY, BACK_KEY}, DEPOSIT_URL, WITHDRAW_URL})

	//发请求
	resp, err := cli.Deposit(GenDepositRequestDemo())
	if err != nil {
		cli.logger.Errorf("err:%s\n", err.Error())
		return
	}
	cli.logger.Infof("resp:%+v\n", resp)
}

func GenDepositRequestDemo() NowPayDepositReq {
	return NowPayDepositReq{
		OrderId:     url.QueryEscape("2025100443562675411"),
		OrderAmount: url.QueryEscape("10"),
		UserId:      url.QueryEscape("1"),
		OrderIp:     url.QueryEscape("127.0.0.1"),
		OrderTime:   url.QueryEscape(time.Now().Format(time.DateTime)),
		PayUserName: url.QueryEscape("jane"),
	}
}
