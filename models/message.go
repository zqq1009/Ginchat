package models

import (
	"gorm.io/gorm"
)

// 消息
type Message struct {
	gorm.Model
	FormId   int64  //发送者
	TargetId int64  //接受者
	Type     int    //发送类型 1 私聊 2 群聊	3 广播
	Media    int    //消息类型 1 文字 2 表情包	3 图片 4 音频
	Content  string //消息内容
	Pic      string
	Url      string
	Desc     string
	Amount   int //其他数字统计
}

func (table *Message) TableName() string {
	return "message"
}
