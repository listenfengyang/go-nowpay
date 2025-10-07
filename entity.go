package go_nowpay

type NowPayInitParams struct {
	MerchantInfo `yaml:",inline" mapstructure:",squash"`

	DepositUrl  string `json:"depositUrl" mapstructure:"depositUrl" config:"depositUrl"  yaml:"depositUrl"`
	WithdrawUrl string `json:"withdrawUrl" mapstructure:"withdrawUrl" config:"withdrawUrl"  yaml:"withdrawUrl"`
}

type MerchantInfo struct {
	MerchantId int    `json:"merchantId" mapstructure:"merchantId" config:"merchantId"  yaml:"merchantId"` // merchantId
	AccessKey  string `json:"accessKey" mapstructure:"accessKey" config:"accessKey"  yaml:"accessKey"`     // accessKey
	BackKey    string `json:"backKey" mapstructure:"backKey" config:"backKey"  yaml:"backKey"`             //backKey
}

//============================================================

// nowpay入金
type NowPayDepositReq struct {
	OrderId     string `json:"orderId" mapstructure:"order_id"`          //商户orderNo
	OrderAmount string `json:"orderAmount" mapstructure:"order_amount"`  //订单金额
	UserId      string `json:"userId" mapstructure:"user_id"`            //客户唯一标识
	OrderIp     string `json:"orderIp" mapstructure:"order_ip"`          //客户下单的IP地址
	OrderTime   string `json:"orderTime" mapstructure:"order_time"`      //请求下单时间（2022-10-30 17:56:45）
	PayUserName string `json:"payUserName" mapstructure:"pay_user_name"` //付款人姓名

	//这个不需要业务侧使用,而是sdk帮计算和补充
	//SysNo       string `json:"sysNo" mapstructure:"sysNo"`             //商户编号
	//Sign        string `json:"sign" mapstructure:"sign"`               //签名
}

type NowPayDepositRsp struct {
	Code int         `json:"code" mapstructure:"code"`
	Data DepositData `json:"data" mapstructure:"data"`
	Msg  string      `json:"msg" mapstructure:"msg"`
}

type DepositData struct {
	OrderNo string `json:"order_no" mapstructure:"order_no"` //商户orderNo
	OrderId string `json:"order_id" mapstructure:"order_id"` //订单金额
	SendUrl string `json:"send_url" mapstructure:"send_url"` //客户唯一标识
	UserId  string `json:"user_id" mapstructure:"user_id"`   //客户下单的IP地址
}

// 入金回调
type NowPayDepositCallbackReq struct {
	BillNo     string `json:"bill_no" mapstructure:"bill_no"`         //唯一订单号，商户下单时传过来的order_id
	BillStatus int    `json:"bill_status" mapstructure:"bill_status"` //订单状态：1=订单已取消 2=订单已激活
	Sign       string `json:"sign" mapstructure:"sign"`
	SysNo      string `json:"sys_no" mapstructure:"sys_no"` //商户编号
	//成功参数
	Amount     string `json:"amount" mapstructure:"amount"`           //订单金额
	AmountUsdt string `json:"amount_usdt" mapstructure:"amount_usdt"` //订单USDT数量
}

// nowpay出金
type WithdrawData struct {
	UserName    string `json:"user_name" mapstructure:"user_name"`       //真实姓名
	BankcardNo  string `json:"bankcard_no" mapstructure:"bankcard_no"`   //卡号
	SerialNo    string `json:"serial_no" mapstructure:"serial_no"`       //订单号
	BankAddress string `json:"bank_address" mapstructure:"bank_address"` //支行地址
	Amount      string `json:"amount" mapstructure:"amount"`             //金额
}
type NowPayWithdrawReq struct {
	Data []WithdrawData `json:"data" mapstructure:"data"`
	//这个不需要业务侧使用,而是sdk帮计算和补充
	//SysNo       string `json:"sys_no" mapstructure:"sys_no"`             //商户编号
	//Sign        string `json:"sign" mapstructure:"sign"`               //签名
}
type NowPayWithdrawRsp struct {
	Code int    `json:"code" mapstructure:"code"` //200
	Msg  string `json:"msg" mapstructure:"msg"`   //成功
}

// 出金回调
type NowPayWithdrawCallbackReq struct {
	BillNo     string `json:"bill_no" mapstructure:"bill_no"`         //唯一订单号，商户下单时传过来的order_id
	BillStatus int    `json:"bill_status" mapstructure:"bill_status"` //订单状态：1=订单已取消 2=订单已激活
	Sign       string `json:"sign" mapstructure:"sign"`
	SysNo      string `json:"sys_no" mapstructure:"sys_no"` //商户编号
	//成功参数
	Amount string `json:"amount" mapstructure:"amount"` //订单金额
}

type NowPayWithdrawCallbackRsp struct {
	Code int    `json:"code" mapstructure:"code"` //200
	Msg  string `json:"msg" mapstructure:"msg"`   //成功
}
