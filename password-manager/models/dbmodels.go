package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DBMaster struct {
	Id         uuid.UUID `gorm:"primaryKey,column:id"`
	CreatedAt  time.Time `gorm:"column:created_at;not_null"`
	Name       string    `gorm:"column:name;type:varchar(100)"`
	Algorithm  string    `gorm:"column:algorithm;type:varchar(100);not null"`
	Email      string    `gorm:"column:email;type:varchar(100);not null"`
	SpecialKey string    `gorm:"column:special_key;not null"`
	Plan       string    `gorm:"column:plan;not null"`
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

type DBRSAUser struct {
	Id         uuid.UUID  `gorm:"primaryKey,column:id"`
	CreatedAt  time.Time  `gorm:"column:created_at;not_null"`
	CreatedBy  uuid.UUID  `gorm:"column:created_by;not null"`
	DeletedBy  *uuid.UUID `gorm:"column:deleted_by"`
	Name       string     `gorm:"column:name;type:varchar(100)"`
	Email      string     `gorm:"column:email;type:varchar(100)"`
	Password   string     `gorm:"column:password;type:varchar(100);not null"`
	PublicKey  string     `gorm:"column:public_key"`
	PrivateKey string     `gorm:"column:private_key"`
	SpecialKey string     `gorm:"column:special_key;not null"`
	IsMaster   bool       `gorm:"column:is_master"`
	MasterId   uuid.UUID  `gorm:"column:master_id;type:uuid;not null;constraint:OnDelete:CASCADE"`
}

func (DBRSAUser) TableName() string {
	return "user_tbl"
}

func (*DBRSAUser) BeforeCreate(tx *gorm.DB) error {
	uuid := uuid.New().String()
	tx.Statement.SetColumn("Id", uuid)
	return nil
}

type DBASAUser struct {
	Id          uuid.UUID  `gorm:"primaryKey,column:id"`
	CreatedAt   time.Time  `gorm:"column:created_at;not_null"`
	CreatedBy   uuid.UUID  `gorm:"column:created_by;not null"`
	DeletedBy   *uuid.UUID `gorm:"column:deleted_by"`
	Name        string     `gorm:"column:name;type:varchar(100)"`
	Email       string     `gorm:"column:email;type:varchar(100)"`
	Password    string     `gorm:"column:password;type:varchar(100);not null"`
	PasswordKey string     `gorm:"column:password_key;type:varchar(100);not null"`
	SpecialKey  string     `gorm:"column:special_key;not null"`
	IsMaster    bool       `gorm:"column:is_master"`
	MasterId    uuid.UUID  `gorm:"column:master_id;type:uuid;not null;constraint:OnDelete:CASCADE"`
}

func (DBASAUser) TableName() string {
	return "user_tbl"
}

func (*DBASAUser) BeforeCreate(tx *gorm.DB) error {
	uuid := uuid.New().String()
	tx.Statement.SetColumn("Id", uuid)
	return nil
}

type DbPassword struct {
	Id          uuid.UUID `gorm:"primaryKey,column:id"`
	WebisteName string    `gorm:"column:website_name;type:varchar(100)"`
	Password    string    `gorm:"column:password;type:varchar(500)"`
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
	Id       uuid.UUID `gorm:"primaryKey,column:id"`
	UserId   uuid.UUID `gorm:"column:user_id;type:uuid;not null"`
	IsLogin  bool      `gorm:"not null"`
	IsMaster bool      `gorm:"not null"`
}

func (DBLogin) TableName() string {
	return "login_tbl"
}

func (*DBLogin) BeforeCreate(tx *gorm.DB) error {
	uuid := uuid.New().String()
	tx.Statement.SetColumn("Id", uuid)
	return nil
}
