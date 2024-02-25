package responder

import (
	"fmt"
	"github.com/rs/zerolog"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types"
	waLog "go.mau.fi/whatsmeow/util/log"
	"os"
	"os/signal"
	"regexp"
	"responder/pkg/responder/formatter"
	"responder/pkg/responder/forwarder"
	"responder/pkg/responder/helper"
	"responder/pkg/responder/message"
	"responder/pkg/responder/postgres"
	"responder/pkg/responder/store"
	"syscall"
)

var logger = zerolog.New(os.Stdout).With().Timestamp().Logger()

func Run() {
	zerolog.SetGlobalLevel(zerolog.WarnLevel)

	con, err := postgres.NewDatabase(os.Getenv("POSTGRESQL"))
	if err != nil {
		panic(err)
	}

	storer, err := store.NewStore(con, "postgres", waLog.Zerolog(logger))
	if err != nil {
		panic(err)
	}

	phoneNum := os.Getenv("PHONENUM")
	if phoneNum == "" {
		panic("PHONENUM not set")
	}

	// TODO: here should use NewJID rather than NewADJID because device id is not set
	//jid := types.NewJID(phoneNum, types.DefaultUserServer)
	jid := types.NewADJID(phoneNum, 0, 5)
	device, _ := storer.GetDevice(jid)
	logger.Warn().Msgf("Device: %v", jid)
	if device == nil {
		if device == nil {
			device = storer.NewDevice()
		}
	}

	logger.Warn().Msgf("Device: %v", device)
	client := whatsmeow.NewClient(device, waLog.Zerolog(logger))

	if client.Store.ID == nil {
		logger.Warn().Msgf("Pairing with %s", jid.User)
		if helper.ConnectClient(client) != nil {
			panic("cannot connected")
		}

		code, err := client.PairPhone(phoneNum, true, whatsmeow.PairClientUnknown, "Chrome (Windows)")
		if err != nil {
			panic(err)
		}
		fmt.Printf("Pairing code: %s\n", code)
	}
	format := formatter.NewFormat()

	senderPattern := `Sender:\s*(\d+)`
	convPattern := `----Message----\n(.+)`

	senderParser := regexp.MustCompile(senderPattern)
	convParser := regexp.MustCompile(convPattern)

	forward := forwarder.NewMessage(client, types.NewJID(os.Getenv("PHONENUMDEST"), types.DefaultUserServer), format.FormatString)

	err = helper.ConnectClient(client)
	if err != nil {
		panic(err)
	}

	client.AddEventHandler(message.EventListener(client, senderParser, convParser, forward))

	sigC := make(chan os.Signal, 1)
	signal.Notify(sigC, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)
	<-sigC

	if err = device.Save(); err != nil {
		panic(err)
	}

}
