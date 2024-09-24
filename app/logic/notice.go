package logic

import "github.com/injoyai/goutil/notice"

/*
通知模块


*/

type _notice struct{}

func (this *_notice) Popup(msg string) error {
	return notice.DefaultWindows.Publish(&notice.Message{
		Title:   "抢票通知",
		Content: msg,
	})
}

func (this *_notice) Notice(msg string) error {
	return notice.DefaultWindows.Publish(&notice.Message{
		Title:   "抢票通知",
		Content: msg,
	})
}

func (this *_notice) Voice(msg string) error {
	return notice.DefaultVoice.Speak(msg)
}
