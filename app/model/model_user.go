package model

type User struct {
	ID       int64  `json:"id"`                    //主键
	Username string `json:"username"`              //用户名
	Password string `json:"password"`              //密码
	InDate   int64  `json:"inDate" xorm:"created"` //创建时间
}
