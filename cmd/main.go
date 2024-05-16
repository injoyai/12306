package main

import (
	"fmt"
	"github.com/injoyai/12306/app/logic"
	"github.com/injoyai/goutil/g"
	"github.com/injoyai/goutil/times"
	"github.com/injoyai/logs"
	"strings"
	"time"
)

func main() {

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

		lastDate := time.Now()
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

		fmt.Printf("\n获取到车次信息:\n%s\n",
			strings.Join(func() []string {
				ls := []string(nil)
				for _, v := range list {
					ls = append(ls, v.String())
				}
				return ls
			}(), "\n"))

	}

}
