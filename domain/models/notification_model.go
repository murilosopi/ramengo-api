package models

type NotificationModel struct {
	Message string
	User    *UserModel
}
