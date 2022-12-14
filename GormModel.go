package main

import (
	"database/sql"
	"time"
)

// PROJECT_NAME:Test
// DATE:2022/10/19 11:16
// USER:21005
type User1 struct {
	ID           uint
	Name         string
	Email        *string
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
type User2 struct {
	//Name string `gorm:"<-:create"`          // 允许读和创建
	//Name string `gorm:"<-:update"`          // 允许读和更新
	//Name string `gorm:"<-"`                 // 允许读和写（创建和更新）
	//Name string `gorm:"<-:false"`           // 允许读，禁止写
	//Name string `gorm:"->"`                 // 只读（除非有自定义配置，否则禁止写）
	//Name string `gorm:"->;<-:create"`       // 允许读和写
	//Name string `gorm:"->:false;<-:create"` // 仅创建（禁止从 db 读）
	//Name string `gorm:"-"`                  // 通过 struct 读写会忽略该字段
	//Name string `gorm:"-:all"`              // 通过 struct 读写、迁移会忽略该字段
	//Name string `gorm:"-:migration"`        // 通过 struct 迁移会忽略该字段
}
type User3 struct {
	//CreatedAt time.Time // 在创建时，如果该字段值为零值，则使用当前时间填充
	//UpdatedAt int       // 在创建时该字段值为零值或者在更新时，使用当前时间戳秒数填充
	//Updated   int64 `gorm:"autoUpdateTime:nano"` // 使用时间戳填纳秒数充更新时间
	//Updated   int64 `gorm:"autoUpdateTime:milli"` // 使用时间戳毫秒数填充更新时间
	//Created   int64 `gorm:"autoCreateTime"`      // 使用时间戳秒数填充创建时间
}

func main() {

}
