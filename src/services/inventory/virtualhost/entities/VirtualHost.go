package entities

import "gorm.io/gorm"

type VirtualHost struct {
	gorm.Model
	ClusterId   uint   `json:"ClusterId" gorm:"index:idx_unique_username_by_host,unique"`
	Name        string `json:"Username" gorm:"index:idx_unique_vhost_by_host"`
	Description string `json:"Description"`
}

func (VirtualHost) TableName() string {
	return "VirtualHost"
}
