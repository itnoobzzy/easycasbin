package global

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type BASE_MODEL struct {
	ID        uint           `gorm:"primarykey"` // 主键ID
	CreatedAt time.Time      // 创建时间
	UpdatedAt time.Time      // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
}

type JSON []byte

func (c JSON) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *JSON) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}
