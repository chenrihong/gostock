// 股票名称——A股

package models

// 字段定义的文档
// https://gorm.io/zh_CN/docs/models.html

// STOCK_INFO A股上市公司代码
type STOCK_INFO struct {
	STOCK_ID     int64  `gorm:"primaryKey;autoIncrement;not null"`
	COMPANY_CODE string `gorm:"type:varchar(50);not null;unique"` //	公司代码
	COMPANY_ABBR string `gorm:"type:varchar(50);not null"`        // 	公司简称
	LISTING_DATE string `gorm:"type:varchar(50);null"`            //	上市日期
	COMPANY_KIND string `gorm:"type:varchar(50);null"`            //	上市地址（sh，sz)
	COMPANY_SITE string `gorm:"type:varchar(250);null"`           //	公司网站
	COMPANY_ATTR string `gorm:"type:varchar(50);null`             //    所属行业
}

func (STOCK_INFO) TableName() string {
	return "stock_info"
}
