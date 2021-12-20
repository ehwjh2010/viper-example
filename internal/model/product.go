package model

import (
	"github.com/ehwjh2010/viper/db/rdb"
	"github.com/ehwjh2010/viper/types"
)

type Product struct {
	rdb.BaseModel
	Name       string           `gorm:"column:name;type:varchar(128);not null;default:'';index:idx_name_total_count;comment:商品名称" json:"name"`
	Price      float64          `gorm:"column:price;type:decimal(10,4);not null;default:0;comment:商品价格" json:"price"`
	TotalCount int64            `gorm:"column:total_count;type:int unsigned;not null;default:0;index:idx_name_total_count;comment:商品库存" json:"totalCount"`
	Brand      types.NullString `gorm:"column:brand;type:varchar(128);comment:品牌" json:"brand" swaggertype:"primitive,string"`
}

func NewProduct() *Product {
	return &Product{}
}

//TableName 指定Product结构体对应的数据表为product
func (p Product) TableName() string {
	return "product"
}
