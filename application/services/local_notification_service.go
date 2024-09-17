package services

import (
	"fmt"
	"ramengo/domain/models"
)

type LocalNotificationService struct {
	NotificationsSent []*models.NotificationModel
}

func NewLocalNotificationService() *LocalNotificationService {
	return &LocalNotificationService{}
}

func (service *LocalNotificationService) Send(user *models.UserModel, message string) (notification *models.NotificationModel) {
	notification = &models.NotificationModel{
		User: user,
		Message: message,
	}

	fmt.Printf("New message sent: %v\nTo: %v (ID: %v)\n", message, user.Name, user.Id)

	service.NotificationsSent = append(service.NotificationsSent, notification)

	return
}
