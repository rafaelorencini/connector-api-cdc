package domain

import (
	"gorm.io/gorm"
)

//type Users struct {
//	gorm.Model
//	ID          int
//	Name        string
//	Email       string
//	CreditCards []Tables `gorm:"foreignKey:UserRefer"`
//}
//
//type CreditCard struct {
//	gorm.Model
//	Number    string
//	UserRefer uint
//}

type Connector struct {
	gorm.Model
	ID            int      `gorm:"primaryKey;autoIncrement"`
	ConnectorName string   `valid:"required"`
	DatabaseName  string   `valid:"required"`
	Tables        []Tables `gorm:"ForeignKey:FkConnectorId"`
}

type Tables struct {
	gorm.Model
	ID            int    `gorm:"primaryKey;autoIncrement"`
	FkConnectorId int64  `valid:"required"`
	Table         string `valid:"required"`
}
