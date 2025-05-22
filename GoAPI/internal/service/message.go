package service

import (
	"context"
	"errors"
	"fmt"
	"mind-set/database"
	"mind-set/internal/model"

	"firebase.google.com/go/messaging"
)

type MessageService struct {
	model.BaseModel
}

func NewMessageService() *MessageService {
	return &MessageService{}
}

func (u *MessageService) SendMessage(token string, title string, body string) error {
	fcm_token := ""
	salePointModel := model.NewStudent()
	resultSalePoint := u.DB().Model(salePointModel).Where("token = ?", token).First(salePointModel)
	if resultSalePoint.Error == nil {
		fcm_token = salePointModel.Fcmtoken
	} /*else {
		agentModel := model.NewAgent()
		resultAgent := u.DB().Model(agentModel).Where("token = ?", token).First(agentModel)
		if resultAgent.Error == nil {
			fcm_token = agentModel.Fcmtoken
		} else {
			adminModel := model.NewAdminUser()
			resultAdmin := u.DB().Model(adminModel).Where("token = ?", token).First(adminModel)
			if resultAdmin.Error == nil {
				fcm_token = adminModel.Fcmtoken
			}
		}
	}*/

	ctx := context.Background()
	app := database.FirebaseApp
	client, err := app.Messaging(ctx)
	if err != nil {
		return err
	}
	//	registrationToken := "ePYqhEcCQl6yPXNRE-X0nG:APA91bEb2GdQIaGkATEQ-D991wr23zByng5QEu2na8Mrr2RfvqjwlTbBv6vUM6yKK8CnYdDaFBMNVwKZXiA_vBY0-l4kKl1QyXM5WECLsMKGbD_T-QHdX1Oxj8fqxT_QDe6rSkUgyYdN"
	if fcm_token == "" || fcm_token == "null" {
		return errors.New("registrationToken is empty")
	}

	messages := &messaging.Message{
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
		Token: fcm_token,
	}

	response, err := client.Send(ctx, messages)
	if err != nil {
		return err
	}
	fmt.Println("Successfully sent message:" + response)
	//logger.Logger.Info("Successfully sent message:" + response)

	return nil
}
