package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DBMaster struct {
	Id         uuid.UUID `gorm:"primaryKey,column:id"`
	Name       string    `gorm:"column:name;type:varchar(100)"`
	Algorithm  string    `gorm:"column:algorithm;type:varchar(100);not null"`
	Email      string    `gorm:"column:email;type:varchar(100);not null"`
	SpecialKey string    `gorm:"column:special_key;not null"`
	Count      int       `gorm:"column:count"`
}

func (DBMaster) TableName() string {
	return "master_tbl"
}

func (*DBMaster) BeforeCreate(tx *gorm.DB) error {
	uuid := uuid.New().String()
	tx.Statement.SetColumn("Id", uuid)
	return nil
}

type DBUser struct {
	Id         uuid.UUID `gorm:"primaryKey,column:id"`
	Name       string    `gorm:"column:name;type:varchar(100)"`
	Email      string    `gorm:"column:email;type:varchar(100)"`
	Password   string    `gorm:"column:password;type:varchar(100);not null"`
	PublicKey  string    `gorm:"column:public_key"`
	PrivateKey string    `gorm:"column:private_key"`
	SpecialKey string    `gorm:"column:special_key;not null"`
	MasterId   uuid.UUID `gorm:"column:master_id;type:uuid;not null;constraint:OnDelete:CASCADE"`
}

func (DBUser) TableName() string {
	return "user_tbl"
}

func (*DBUser) BeforeCreate(tx *gorm.DB) error {
	uuid := uuid.New().String()
	tx.Statement.SetColumn("Id", uuid)
	return nil
}

type DbPassword struct {
	Id          uuid.UUID `gorm:"primaryKey,column:id"`
	WebisteName string    `gorm:"column:website_name;type:varchar(100)"`
	Password    string    `gorm:"column:password;type:varchar(100)"`
	UserId      uuid.UUID `gorm:"column:user_id;type:uuid;not null;constraint:OnDelete:CASCADE"`
}

func (DbPassword) TableName() string {
	return "psswrd_tbl"
}

func (*DbPassword) BeforeCreate(tx *gorm.DB) error {
	uuid := uuid.New().String()
	tx.Statement.SetColumn("Id", uuid)
	return nil
}

type DBLogin struct {
	IsLogin  bool      `gorm:"column:is_login;not null"`
	UserId   uuid.UUID `gorm:"column:user_id;type:uuid;not null;constraint:OnDelete:CASCADE"`
	MasterId uuid.UUID `gorm:"column:master_id;type:uuid;not null;constraint:OnDelete:CASCADE"`
}

func (DBLogin) TableName() string {
	return "login_tbl"
}
