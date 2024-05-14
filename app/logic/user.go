package logic

import (
	"encoding/base64"
	"errors"
	"github.com/injoyai/goutil/net/http"
	"strings"
)

/*
用户模块


*/

var User = &user{}

type user struct {
	Cookies []*http.Cookie
}

func (this *user) url(url string) *http.Request {
	return http.Url(url).AddCookie(this.Cookies...)
}

// decodeCookie 解析cookie,测试用
func (this *user) decodeCookie(s string) []*http.Cookie {
	cookies := []*http.Cookie(nil)
	for _, v := range strings.Split(s, ";") {
		if list := strings.SplitN(v, "=", 2); len(list) == 2 {
			cookies = append(cookies, &http.Cookie{Name: list[0], Value: list[1]})
		}
	}
	return cookies
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

// ByQR 二维码登录
func (this *user) ByQR() ([]*http.Cookie, error) { return nil, nil }

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

// Info 用户信息
func (this *user) Info() {

}

// Logout 退出登录
func (this *user) Logout() error {

	this.Cookies = nil
	return nil
}
