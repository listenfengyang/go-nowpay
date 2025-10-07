package go_nowpay

import (
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
		OrderId:     "2025100443562675418",
		OrderAmount: "10",
		UserId:      "1",
		OrderIp:     "127.0.0.1",
		OrderTime:   time.Now().Format(time.DateTime),
		PayUserName: "张三",
	}
}
