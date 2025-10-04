package go_nowpay

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/listenfengyang/go-nowpay/utils"
	"github.com/spf13/cast"
)

func (cli *Client) WithdrawReq(req NowPayWithdrawReq) (*NowPayWithdrawRsp, error) {

	rawURL := cli.Params.WithdrawUrl
	// 2. Convert struct to map for signing
	var params map[string]string
	params = map[string]string{}
	b, _ := json.Marshal(req.Data)
	params["data"] = cast.ToString(b)
	params["sys_no"] = cast.ToString(cli.Params.MerchantId)

	// Generate signature
	signStr, _ := utils.SignWithdraw(params, cli.Params.AccessKey)
	params["sign"] = signStr
	var result NowPayWithdrawRsp
	fmt.Println(params)
	resp2, err := cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetCloseConnection(true).
		R().
		SetFormData(params).
		//SetBody(params).
		SetHeaders(getHeaders()).
		SetDebug(cli.debugMode).
		SetResult(&result).
		SetError(&result).
		Post(rawURL)

	restLog, _ := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(utils.GetRestyLog(resp2))
	cli.logger.Infof("PSPResty#nowpay#withdraw->%s", string(restLog))

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
