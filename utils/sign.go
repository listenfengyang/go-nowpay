package utils

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/samber/lo"
	"github.com/spf13/cast"
	"net/url"
	"sort"
	"strings"
)

func Sign(params map[string]string, key string) (string, error) {
	// 1. Validate key
	if key == "" {
		return "", errors.New("APP_KEY 参数为空，请填写")
	}

	// 2. Get and sort keys
	keys := lo.Keys(params)
	sort.Strings(keys) // ASCII ascending order

	// 3. Build sign string
	var sb strings.Builder
	for _, k := range keys {
		value := cast.ToString(params[k])
		if k != "sign" && value != "" {
			//只有非空才可以参与签名
			sb.WriteString(fmt.Sprintf("%s=%s&", k, url.QueryEscape(value)))
		}
	}
	signStr := sb.String()
	signStr = strings.Trim(signStr, "&")
	signStr += fmt.Sprintf("%s", key)

	fmt.Printf("[rawString]%s\n", signStr)

	// 4. Generate MD5
	hash := md5.Sum([]byte(signStr))
	signResult := hex.EncodeToString(hash[:])

	// Debug print (optional)
	//fmt.Printf("验签str: %s\n结果: %s\n", signStr, signResult)

	return signResult, nil
}

func Verify(params map[string]string, signKey string) (bool, error) {
	// Check if signature exists in params
	signature, exists := params["sign"]
	if !exists {
		return false, nil
	}

	// Remove signature from params for verification
	delete(params, "sign")

	// Generate current signature
	currentSignature, err := Sign(params, signKey)
	if err != nil {
		return false, fmt.Errorf("signature generation failed: %w", err)
	}

	// Compare signatures
	return signature == currentSignature, nil
}

// 入金&出金回调-成功-验签
func VerifyCallback(params map[string]string, signKey string) bool {
	signature := params["sign"]

	keys := lo.Keys(params)

	var sb strings.Builder
	var value string
	for _, k := range keys {
		value = cast.ToString(params[k])
		if k != "sign" && value != "" {
			sb.WriteString(fmt.Sprintf("%s", url.QueryEscape(value)))
		}
	}

	sb.WriteString(fmt.Sprintf("%s", signKey))
	signStr := sb.String()
	fmt.Printf("[rawString]%s\n", signStr)

	hash := md5.Sum([]byte(signStr))
	//fmt.Printf("sign: %s\n", signature)
	//fmt.Printf("md5 sign: %s\n", hex.EncodeToString(hash[:]))
	return signature == hex.EncodeToString(hash[:])
}

// 入金&出金回调-取消-验签
func VerifyCanceledCallback(params map[string]interface{}, signKey string) bool {
	signature := params["sign"]

	keys := []string{"bill_no", "bill_status", "sys_no"}

	var sb strings.Builder
	var value string
	for _, k := range keys {
		value = cast.ToString(params[k])
		if k != "sign" && value != "" {
			sb.WriteString(fmt.Sprintf("%s=%s&", k, url.QueryEscape(value)))
		}
	}

	signStr := sb.String()
	signStr = strings.Trim(signStr, "&")
	signStr += fmt.Sprintf("%s", signKey)

	fmt.Printf("[rawString]%s\n", signStr)

	hash := md5.Sum([]byte(signStr))
	//fmt.Printf("sign: %s\n", signature)
	//fmt.Printf("md5 sign: %s\n", hex.EncodeToString(hash[:]))
	return signature == hex.EncodeToString(hash[:])
}

// 出金
func SignWithdraw(params map[string]string, key string) (string, error) {
	// 1. Validate key
	if key == "" {
		return "", errors.New("APP_KEY 参数为空，请填写")
	}

	// 2. Get and sort keys
	keys := lo.Keys(params)
	sort.Strings(keys) // ASCII ascending order

	// 3. Build sign string
	var sb strings.Builder
	var value string
	for _, k := range keys {
		value = cast.ToString(params[k])
		if k != "sign" && value != "" {
			//只有非空才可以参与签名
			sb.WriteString(fmt.Sprintf("%s", value))
		}
	}
	sb.WriteString(fmt.Sprintf("%s", key))
	signStr := sb.String()
	fmt.Printf("[rawString]%s\n", signStr)

	// 4. Generate MD5
	hash := md5.Sum([]byte(signStr))
	signResult := hex.EncodeToString(hash[:])

	// Debug print (optional)
	//fmt.Printf("验签str: %s\n结果: %s\n", signStr, signResult)

	return signResult, nil
}

func VerifySignWithdraw(params map[string]string, signKey string) (bool, error) {
	// Check if signature exists in params
	signature, exists := params["sign"]
	if !exists {
		return false, nil
	}

	// Remove signature from params for verification
	delete(params, "sign")

	// Generate current signature
	currentSignature, err := SignWithdraw(params, signKey)
	if err != nil {
		return false, fmt.Errorf("signature generation failed: %w", err)
	}

	// Compare signatures
	return signature == currentSignature, nil
}
