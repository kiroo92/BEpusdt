package model

type Config struct {
	K string `gorm:"column:k;type:varchar(32);primaryKey"`
	V string `gorm:"column:v;type:varchar(255)"`
}

func (c Config) TableName() string {

	return "config"
}

func SetK(k, v string) {
	DB.Save(&Config{K: k, V: v})
}

func GetK(k string) string {
	var row Config

	var tx = DB.Where("k = ?", k).First(&row)
	if tx.Error == nil {

		return row.V
	}

	return ""
}
