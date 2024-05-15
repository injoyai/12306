package logic

import (
	"encoding/base64"
	"errors"
	"github.com/injoyai/base/g"
	"github.com/injoyai/goutil/net/http"
)

/*
用户模块


*/

var User = &user{}

type user struct {
	Cookies
}

func (this *user) url(url string) *http.Request {
	return http.Url(url).AddCookie(this.Cookies...)
}

// Check 判断是否登录
func (this *user) Check() (bool, error) {
	url := "https://kyfw.12306.cn/otn/login/conf"
	resp := this.url(url).Get()
	if resp.Err() != nil {
		return false, resp.Err()
	}
	return resp.GetBodyDMap().GetString("data.is_login") == "Y", nil
}

// Login 登录操作
func (this *user) Login() ([]*http.Cookie, error) { return nil, nil }

// LoginByQR 二维码登录
func (this *user) LoginByQR() ([]*http.Cookie, error) { return nil, nil }

// QR 获取登录二维码
func (this *user) QR() ([]byte, error) {
	res := &struct {
		Image   string `json:"image"`          //base64的二维码图片
		Message string `json:"result_message"` //接口结果消息
		Code    string `json:"result_code"`    //接口状态码
		UUID    string `json:"uuid"`           //uuid
	}{}
	url := "https://kyfw.12306.cn/passport/web/create-qr64?appid=otn"
	resp := this.url(url).Bind(res).Get()
	if resp.Err() != nil {
		return nil, resp.Err()
	}
	if res.Code != "0" {
		return nil, errors.New(res.Message)
	}
	return base64.StdEncoding.DecodeString(res.Image)
}

// CheckLoginByQR 校验二维码登录是否成功,返回是否登录成功,是否过期,和错误信息
func (this *user) CheckLoginByQR(uuid string) (succ, overdue bool, err error) {
	url := "https://kyfw.12306.cn/passport/web/checkqr"
	resp := this.url(url).SetQuerys(g.M{
		"appid": "otn",
		"uuid":  uuid,
	}).Get()
	if resp.Err() != nil {
		return false, false, nil
	}
	m := resp.GetBodyDMap()
	code := m.GetString("result_code")
	msg := m.GetString("result_message")
	switch code {
	case "0":
		return false, false, nil
	case "2":
		return true, false, nil
	case "3":
		return false, true, nil
	default:
		return false, false, errors.New(msg)
	}
}

// Info 用户信息
func (this *user) Info() {

}

// Logout 退出登录
func (this *user) Logout() error {
	url := "https://kyfw.12306.cn/otn/login/loginOut"
	resp := http.Url(url).SetHeader("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8").Get()
	if resp.Err() != nil {
		return resp.Err()
	}
	this.Cookies = nil
	return nil
}
