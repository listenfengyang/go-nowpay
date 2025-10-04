package go_nowpay

import (
	"testing"
)

func TestWithdraw(t *testing.T) {
	vLog := VLog{}
	//构造client
	cli := NewClient(vLog, &NowPayInitParams{MerchantInfo{MERCHANT_ID, ACCESS_KEY, BACK_KEY}, DEPOSIT_URL, WITHDRAW_URL})

	//发请求
	resp, err := cli.WithdrawReq(GenWithdrawRequestDemo())
	if err != nil {
		cli.logger.Errorf("err:%s\n", err.Error())
		return
	}
	cli.logger.Infof("resp:%+v\n", resp)
}

func GenWithdrawRequestDemo() NowPayWithdrawReq {
	data := make([]WithdrawData, 1, 2)
	data[0] = WithdrawData{
		UserName:    "张三",
		BankcardNo:  "353236326",
		SerialNo:    "2025642422446",
		BankAddress: "fgww",
		Amount:      "11",
	}
	return NowPayWithdrawReq{data}
}
