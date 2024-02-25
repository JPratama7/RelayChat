# RelayChat
Simple project to forward whatsapp messages to destination whatsapp number
with additional functionality to send message to a specific number from destination whatsapp number.


## Environment Variable
```bash
export POSTGRESQL="postgres uri"
export PHONENUM="phone number to listing message"
export PHONENUMDEST"phone number to send and receive message"
```

## Run
```bash
git clone git@github.com:JPratama7/RelayChat.git
cd RelayChat
go run cmd/responder/main.go
```