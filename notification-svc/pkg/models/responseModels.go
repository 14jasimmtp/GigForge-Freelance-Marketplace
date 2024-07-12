package models 

import(
	"time"
)

type NotificationModel struct {
	NotificaitonID     uint64
	UserID             uint64
	CreatedAt          time.Time
	CommentText        string
}

