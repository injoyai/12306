package main

import (
	"fmt"
	"github.com/injoyai/12306/app/logic"
	"github.com/injoyai/goutil/g"
	"github.com/injoyai/goutil/oss"
	"github.com/injoyai/goutil/times"
	"github.com/injoyai/logs"
	"time"
)

func main() {

	logs.SetFormatterWithTime()

	defer func() {
		if e := recover(); e != nil {
			logs.Err(e)
			g.Input("按回车键退出...")
		}
	}()

	/*
		1. 输入日期
		2. 输入车次




	*/

	for {

		lastDate := times.IntegerDay(time.Now())
		for {
			date := g.Input("请输入日期,", fmt.Sprintf("默认为(%s)", lastDate.Format(g.TimeDate)))
			if len(date) == 0 {
				break
			}

			t, err := time.Parse(g.TimeDate, date)
			if err == nil {
				if t.Sub(times.IntegerDay(time.Now())) > 0 {
					lastDate = t
					break
				}
				err = fmt.Errorf("日期必须大于等于今天")
			}
			logs.Err(err)
		}

		lastFrom := "杭州东"
		for {
			from := g.Input("请输入出发地,", fmt.Sprintf("默认为(%s)", lastFrom))
			if len(from) == 0 {
				from = lastFrom
			}
			if _, err := logic.CheckStationName(from); err == nil {
				lastFrom = from
				break
			}
			logs.Err("无效站台名称: ", from)
		}

		lastTo := "温州南"
		for {
			to := g.Input("请输入目的地,", fmt.Sprintf("默认为(%s)", lastTo))
			if len(to) == 0 {
				to = lastTo
			}
			if _, err := logic.CheckStationName(to); err == nil {
				lastTo = to
				break
			}
			logs.Err("无效站台名称: ", to)
		}

		list, err := logic.GetTicketList(lastDate, lastFrom, lastTo)
		if err != nil {
			logs.Err(err)
			continue
		}

		if len(list) == 0 {
			logs.Warn("未查询到车次信息")
			continue
		}

		m := list.TrainNoMap()
		fmt.Printf("\n获取到车次信息:\n%s\n\n", list.String())

		for {
			i := g.InputVar("选择需要抢购的车次序号...").Int()
			if _, ok := m[i]; !ok {
				logs.Warn("无效车次序号: ", i)
				continue
			}

			logs.Infof("\n您要抢购的车次信息为:\n%s\n\n", list[i-1].String())

			//开始抢票
			fmt.Printf("\n等待%s 08:00:00开始抢票...\n", lastDate.Format(g.TimeDate))
			for {
				sub := lastDate.Add(time.Hour*8).Sub(time.Now()) / 1e9 * 1e9
				if sub <= 0 {
					do()
					break
				}
				fmt.Printf("\r倒计时: %s    ", lastDate.Add(time.Hour*8).Sub(time.Now())/1e9*1e9)
				<-time.After(time.Second)
			}

		}

	}

}

func do() {
	logs.Debug("正在抢票中...")
	oss.Wait()
}
