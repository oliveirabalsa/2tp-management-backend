package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uuid.UUID      `gorm:"type:varchar(36);primary_key" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type User struct {
	BaseModel
	Username string  `gorm:"uniqueIndex;not null" json:"username" binding:"required,min=1"`
	Password string  `gorm:"not null" json:"-" binding:"required,min=1"`
	Role     string  `gorm:"default:user" json:"-"`
	Boards   []Board `gorm:"many2many:user_boards;" json:"boards,omitempty"`
}

type Board struct {
	BaseModel
	Title   string    `gorm:"not null" json:"title"`
	AdminID uuid.UUID `gorm:"type:uuid;not null" json:"-"`
	Admin   User      `gorm:"foreignKey:AdminID" json:"-"`
	Users   []User    `gorm:"many2many:user_boards;" json:"users,omitempty"`
	Columns []Column  `json:"columns,omitempty"`
}

type Column struct {
	BaseModel
	Title   string    `gorm:"not null" json:"title"`
	BoardID uuid.UUID `gorm:"type:uuid;not null" json:"-"`
	Board   Board     `gorm:"foreignKey:BoardID" json:"-"`
	Tasks   []Task    `gorm:"foreignKey:ColumnID;constraint:OnDelete:CASCADE" json:"tasks"`
}

type Task struct {
	BaseModel
	Title       string     `gorm:"not null" json:"title"`
	Description string     `json:"description"`
	AssignerID  *uuid.UUID `gorm:"type:uuid" json:"-"`
	Assigner    *User      `gorm:"foreignKey:AssignerID" json:"assigner,omitempty" serializer:"limited"`
	ColumnID    uuid.UUID  `gorm:"type:uuid;not null" json:"-"`
	Column      Column     `gorm:"foreignKey:ColumnID" json:"-"`
	Priority    string     `gorm:"default:0" json:"priority"`
}

type LimitedUser struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Password string    `json:"-"` // Hide password from JSON
}

// Add this method to User struct
func (u User) ToLimited() LimitedUser {
	return LimitedUser{
		ID:       u.ID,
		Username: u.Username,
	}
}

// Add this method to BaseModel
func (b BaseModel) GetID() uuid.UUID {
	return b.ID
}
