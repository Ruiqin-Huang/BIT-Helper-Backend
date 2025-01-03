/*
 * @Author: flwfdd
 * @Date: 2024-06-06 16:42:56
 * @LastEditTime: 2024-11-18 00:00:05
 * @Description:
 * _(:з」∠)_
 */
package database

import (
	"BIT-Helper/util/config"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// 枚举用户类型(需要与数据库中定义一致)
const (
	Identity_Normal = iota
	Identity_Admin
)

// 基本模型
type Base struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"create_time"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"update_time"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"delete_time"`
}

// 用户
type User struct {
	Base
	Sid      string `gorm:"not null;uniqueIndex;size:42" json:"sid"`
	Nickname string `gorm:"not null;unique" json:"nickname"`
	Avatar   string `json:"avatar"`
	Intro    string `json:"intro"`
	Score    int    `json:"score"`
	Identity int    `json:"identity"`
	Phone    string `json:"phone"`
}

// 图片
type Image struct {
	Base
	Mid  string `gorm:"not null;uniqueIndex;size:233" json:"mid"`
	Size uint   `gorm:"not null" json:"size"`
	Uid  uint   `gorm:"not null" json:"uid"`
}

// 商品
type Goods struct {
	Base
	Type   int     `gorm:"not null" json:"type"`
	Uid    uint    `gorm:"not null" json:"uid"`
	Title  string  `gorm:"not null" json:"title"`
	Intro  string  `gorm:"not null" json:"intro"`
	Num    int     `gorm:"not null" json:"num"`
	Price  float32 `gorm:"not null" json:"price"`
	Images string  `json:"images"` //图片mids，以" "拼接
}

// 订单
type Order struct {
	Base
	State    int  `gorm:"not null" json:"state"`
	Goods    uint `gorm:"not null" json:"goods_id"`
	Receiver uint `gorm:"not null" json:"receiver_id"`
}

// 聊天
type Chats struct {
	Base
	Time     time.Time `gorm:"autoCreateTime" json:"time"`
	Type     string    `gorm:"not null" json:"type"`
	Uid_from uint      `gorm:"not null" json:"uid_from"`
	Content  string    `gorm:"not null" json:"content"`
	Uid_to   uint      `gorm:"not null" json:"uid_to"`
}

// 评论
type Comment struct {
	Base
	Obj        string `gorm:"not null;index" json:"obj"`      //评论对象
	Uid        uint   `gorm:"not null;index" json:"uid"`      //用户id
	Text       string `gorm:"not null" json:"text"`           //评论内容
	Anonymous  bool   `gorm:"default:false" json:"anonymous"` //是否匿名
	LikeNum    uint   `gorm:"default:0" json:"like_num"`      //点赞数
	CommentNum uint   `gorm:"default:0" json:"comment_num"`   //评论数
	ReplyObj   string `json:"reply_obj"`                      //回复对象
	ReplyUid   int    `gorm:"default:0" json:"reply_uid"`     //回复用户id
	Rate       uint   `gorm:"default:0" json:"rate"`          //评分
	Images     string `json:"images"`                         //图片mids，以" "拼接
}

// 点赞
type Like struct {
	Base
	Obj string `gorm:"not null;index" json:"obj"` //点赞对象
	Uid uint   `gorm:"not null;index" json:"uid"` //用户id
}

// TODO: ADD TOPIC :
// ADDITIONAL: 1. VOTE TOPIC 2. TOPIC RATE/VOTE COUNT 3. USER-DEFINED TOPIC TAGS
// -------------------------------------------------------

type Topic struct {
	Base
	Type       int     `gorm:"not null" json:"type"`         // 主题类型(“校园生活”、“电影”、“音乐”、“读书”之一)
	Uid        uint    `gorm:"not null" json:"uid"`          // 发布者ID
	Title      string  `gorm:"not null" json:"title"`        // 主题标题
	Content    string  `json:"content"`                      // 主题内容（允许为空）
	Images     string  `json:"images"`                       // 主题图片列表
	IsVote     bool    `gorm:"default:false" json:"is_vote"` // 是否为投票话题
	CommentNum uint    `gorm:"default:0" json:"comment_num"` // TODO: 评论数，注意每次有用户评论都需要+1
	LikeNum    uint    `gorm:"default:0" json:"like_num"`    // TODO: 点赞数, 注意每次有用户点赞都需要+1
	AvgRate    float32 `gorm:"default:0" json:"avg_rate"`    // 平均评分，每次有用户评论都需要更新一次
	AI         string  `json:"ai"`                           // AI总结
}

type VoteOption struct {
	Base
	TopicID uint   `gorm:"not null" json:"topic_id"`  // 所属主题ID
	VoteNum uint   `gorm:"default:0" json:"vote_num"` // 投票数统计
	Option  string `gorm:"not null" json:"option"`    // 投票选项内容
}

type VoteRecord struct {
	Base
	TopicID      uint `gorm:"not null" json:"topic_id"`       // 话题ID
	VoteOptionID uint `gorm:"not null" json:"vote_option_id"` // 投票选项ID
	UserID       uint `gorm:"not null" json:"user_id"`        // 用户ID
}

type TopicTag struct {
	Base
	TopicID uint   `gorm:"not null" json:"topic_id"` // 所属主题ID
	Tag     string `gorm:"not null" json:"tag"`      // 话题标签
}

// -------------------------------------------------------

func Init() {
	dsn := config.Config.Dsn
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = db

	// TODO: ADD TOPIC :
	// ADDITIONAL: 1. VOTE TOPIC 2. TOPIC RATE/LIKE COUNT 3. USER-DEFINED TOPIC TAGS
	// -------------------------------------------------------

	// err = db.AutoMigrate(
	// 	&User{}, &Image{}, &Goods{}, &Order{}, &Chats{}, &Comment{}, &Like{},
	// )
	err = db.AutoMigrate(
		&User{}, &Image{}, &Goods{}, &Order{}, &Chats{}, &Comment{}, &Like{}, &Topic{}, &VoteOption{}, &VoteRecord{}, &TopicTag{},
	)

	// -------------------------------------------------------

	if err != nil {
		panic(err)
	}
}
