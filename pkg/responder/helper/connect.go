package helper

import "go.mau.fi/whatsmeow"

func ConnectClient(client *whatsmeow.Client) error {
	if !client.IsConnected() {
		return client.Connect()
	}
	return nil
}
