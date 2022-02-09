// 基金信息表

package models

import "time"

// 字段定义的文档
// https://gorm.io/zh_CN/docs/models.html

// FUND_INFO  股票型基金信息数据表
type FUND_INFO struct {
	FUND_ID        int64      `gorm:"primaryKey;autoIncrement;not null"`
	FUND_CODE      string     `gorm:"type:varchar(50);not null;unique;index:idex_fund_code"` //	基金代码
	FUND_NAME      string     `gorm:"type:varchar(250);not null;index:idex_fund_name"`       // 基金简称
	FUND_DWJZ      float64    `gorm:"not null"`                                              //	单位净值
	FUND_LJJZ      float64    `gorm:"not null"`                                              //	累记净值
	RATE_1WEEK     float64    `gorm:"not null"`                                              //	近一周收益率
	RATE_1MONTH    float64    `gorm:"not null"`                                              //	近一月收益率
	RATE_3MONTH    float64    `gorm:"not null"`                                              //	近三月收益率
	RATE_6MONTH    float64    `gorm:"not null"`                                              //	近六月收益率
	RATE_1YEAR     float64    `gorm:"not null"`                                              //	近一年收益率
	RATE_2YEAR     float64    `gorm:"not null"`                                              //	近二年收益率
	RATE_3YEAR     float64    `gorm:"not null"`                                              //	近三年收益率
	RATE_THIS_YEAR float64    `gorm:"not null"`                                              //	今年以来收益率
	RATE_ALL_TIME  float64    `gorm:"not null"`                                              //	成立以来
	CREATED_ON     *time.Time `gorm:"type:datetime; NULL;"`
}

func (FUND_INFO) TableName() string {
	return "fund_info"
}
