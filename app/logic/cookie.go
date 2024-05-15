package logic

import (
	"github.com/injoyai/goutil/net/http"
	"strings"
)

func ParseCookie(s string) Cookies {
	cookies := []*http.Cookie(nil)
	for _, v := range strings.Split(s, ";") {
		if list := strings.SplitN(v, "=", 2); len(list) == 2 {
			cookies = append(cookies, &http.Cookie{Name: list[0], Value: list[1]})
		}
	}
	return cookies
}

type Cookies []*http.Cookie

func (this *Cookies) String() string {
	var s string
	for _, v := range *this {
		s += v.Name + "=" + v.Value + "; "
	}
	return s
}

func (this *Cookies) Get(key string) string {
	for _, v := range *this {
		if v.Name == key {
			return v.Value
		}
	}
	return ""
}

func (this *Cookies) Set(cookie *http.Cookie) {
	has := false
	for i, v := range *this {
		if v.Name == cookie.Name {
			(*this)[i] = cookie
			has = true
		}
	}
	if !has {
		*this = append(*this, cookie)
	}
}

func (this *Cookies) Add(cookie *http.Cookie) {
	*this = append(*this, cookie)
}

func (this *Cookies) Del(key string) {
	for i, v := range *this {
		if v.Name == key {
			*this = append((*this)[:i], (*this)[i+1:]...)
			break
		}
	}
}
