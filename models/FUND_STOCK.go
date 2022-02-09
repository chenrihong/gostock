// 基金信息表

package models

import "time"

// 字段定义的文档
// https://gorm.io/zh_CN/docs/models.html

// FUND_STOCK  股票型基金信息数据表
type FUND_STOCK struct {
	ID         int64      `gorm:"primaryKey;autoIncrement;not null"`
	FUND_CODE  string     `gorm:"type:varchar(50);not null;index:idex_fund_code"` //	基金代码
	STOCK_CODE string     `gorm:"not null"`                                       //	股票代码
	STOCK_NAME string     `gorm:"not null"`                                       //	股票名称
	STOCK_RATE float64    `gorm:"not null"`                                       //	股票仓位
	CREATED_ON *time.Time `gorm:"type:datetime; NULL;"`
}

func (FUND_STOCK) TableName() string {
	return "fund_stock"
}
