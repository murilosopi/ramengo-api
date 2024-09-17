package services

import "ramengo/domain/models"

type NotificationService interface {
	Send(user *models.UserModel, message string) *models.NotificationModel
}