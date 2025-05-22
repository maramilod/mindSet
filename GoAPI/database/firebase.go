package database

import (
	"context"
	"fmt"
	c "mind-set/config"
	"mind-set/internal/utils"
	"path/filepath"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var FirebaseApp *firebase.App

func initFirebase() {
	runDirectory := utils.GetRunPath()
	filePath := filepath.Join(runDirectory, c.Config.Firebase.ServiceAccountKey)
	fmt.Println(filePath)
	opt := option.WithCredentialsFile(filePath)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		fmt.Println("error initializing app: " + err.Error())
		panic("FirebaseAuth connection failed: " + err.Error())
	}
	FirebaseApp = app
}
