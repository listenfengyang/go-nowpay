package go_nowpay

import (
	"crypto/tls"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/listenfengyang/go-nowpay/utils"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/cast"
)

// 下单
func (cli *Client) Deposit(req NowPayDepositReq) (*NowPayDepositRsp, error) {

	rawURL := cli.Params.DepositUrl

	var params map[string]string
	mapstructure.Decode(req, &params)

	//补充字段
	params["sys_no"] = cast.ToString(cli.Params.MerchantId)

	// Generate signature
	signStr, _ := utils.Sign(params, cli.Params.AccessKey)
	params["sign"] = signStr
	fmt.Println(params)
	var result NowPayDepositRsp

	resp2, err := cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetCloseConnection(true).
		R().
		SetFormData(params).
		SetBody(params).
		SetHeaders(getHeaders()).
		SetDebug(cli.debugMode).
		SetResult(&result).
		SetError(&result).
		Post(rawURL)

	restLog, _ := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(utils.GetRestyLog(resp2))
	cli.logger.Infof("PSPResty#nowpay#deposit->%s", string(restLog))

	if err != nil {
		return nil, err
	}

	if resp2.StatusCode() != 200 {
		//反序列化错误会在此捕捉
		return nil, fmt.Errorf("status code: %d", resp2.StatusCode())
	}

	if resp2.Error() != nil {
		//反序列化错误会在此捕捉
		return nil, fmt.Errorf("%v, body:%s", resp2.Error(), resp2.Body())
	}

	return &result, nil
}
