package response

import (
	"eticketing/entity"
)

type User struct {
	UserID    uint   `json:"user_id"`
	Username  string `json:"username"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

func BuildUser(user entity.User) User {
	return User{
		UserID:    user.ID,
		Username:  user.Username,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
		DeletedAt: user.DeletedAt.Time.String(),
	}
}
