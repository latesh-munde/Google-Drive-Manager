package services

import (
	"context"
	"encoding/json"
	"os"

	"gin-drive-manager/config"

	"golang.org/x/oauth2"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

func GetDriveService() (*drive.Service, error) {
	conf := config.GetOAuthConfig()

	file, err := os.Open("tokens/token.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	token := &oauth2.Token{}
	json.NewDecoder(file).Decode(token)

	client := conf.Client(context.Background(), token)

	return drive.NewService(context.Background(), option.WithHTTPClient(client))
}
