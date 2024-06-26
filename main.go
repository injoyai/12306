package main

import (
	"github.com/injoyai/conv/cfg"
	"github.com/injoyai/conv/codec"
	"github.com/injoyai/logs"
)

func main() {

	/*

		1. 用户设置想抢的票(车次,日期,预期位置,截止日期)
		2. 服务校验是否登录,没登录需要用户进行登录操作
		3. 临近开票时间(提前15天,0点刷新?)开始刷新接口
		4. 发现有票则提交订单
		5. 发送通知提醒用户进行付款(有票),或提醒第一波失败(无票),可自行候补(或等下一波放票,会持续放票)



		据了解，铁路12306放票的时间通常是早上8点到下午16点之间。一般会选择在半点或者整点放出新票。
		在正常情况下，会一次性放出所有的车票。
		但在节假日或者高峰期出行时，可能会采用分批放票的方式。
		如果有乘客办理改签或退票手续，这些车票会重新放回票库中，并作为新票再发放出去。

		12306线上售票平台的放票时间是每天的八点整。
		也就是说，想要抢到新一天的车票，你需要在早上八点准时登录12306网站或手机APP进行查询和购买。

	*/

	cfg.Default = cfg.New("./config/config.yaml", codec.Yaml)
	logs.Debug(cfg.GetString("test.cookieStr"))

}
