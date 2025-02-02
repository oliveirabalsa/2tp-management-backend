package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string  `gorm:"uniqueIndex;not null"`
	Password string  `gorm:"not null" json:"-"`
	Role     string  `gorm:"default:user"`
	Boards   []Board `gorm:"many2many:user_boards;" json:"boards,omitempty"`
}

type Board struct {
	gorm.Model
	Title   string   `gorm:"not null" json:"title"`
	AdminID uint     `gorm:"not null" json:"-"`
	Admin   User     `gorm:"foreignKey:AdminID" json:"-"`
	Users   []User   `gorm:"many2many:user_boards;" json:"users,omitempty"`
	Columns []Column `json:"columns,omitempty"`
}

type Column struct {
	gorm.Model
	Title   string `gorm:"not null"`
	BoardID uint   `gorm:"not null" json:"board_id"`
	Tasks   []Task `json:"tasks,omitempty"`
}

type Task struct {
	gorm.Model
	Title       string `gorm:"not null" json:"title"`
	Description string `json:"description"`
	AssignerID  uint   `json:"-"`
	Assigner    User   `gorm:"foreignKey:AssignerID" json:"assigner,omitempty"`
	ColumnID    uint   `json:"column_id"`
}
