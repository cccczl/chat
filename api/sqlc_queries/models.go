// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package sqlc_queries

import (
	"encoding/json"
	"time"
)

type AuthUser struct {
	ID          int32
	Password    string
	LastLogin   time.Time
	IsSuperuser bool
	Username    string
	FirstName   string
	LastName    string
	Email       string
	IsStaff     bool
	IsActive    bool
	DateJoined  time.Time
}

type ChatMessage struct {
	ID              int32
	Uuid            string
	ChatSessionUuid string
	Role            string
	Content         string
	Score           float64
	UserID          int32
	CreatedAt       time.Time
	UpdatedAt       time.Time
	CreatedBy       int32
	UpdatedBy       int32
	Raw             json.RawMessage
}

type ChatPrompt struct {
	ID              int32
	Uuid            string
	ChatSessionUuid string
	Role            string
	Content         string
	Score           float64
	UserID          int32
	CreatedAt       time.Time
	UpdatedAt       time.Time
	CreatedBy       int32
	UpdatedBy       int32
}

type ChatSession struct {
	ID          int32
	UserID      int32
	Uuid        string
	Topic       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Active      bool
	MaxLength   int32
	Temperature float64
	MaxTokens   int32
}

type UserActiveChatSession struct {
	ID              int32
	UserID          int32
	ChatSessionUuid string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
