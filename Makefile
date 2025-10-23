APP_NAME = gopulse

build:
	go build -o bin/$(APP_NAME) ./cmd/gopulse
	
install:
	sudo install -m755 bin/$(APP_NAME) /usr/local/bin/$(APP_NAME)

uninstall:
	sudo rm -f /usr/local/bin/${APP_NAME}

