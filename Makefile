APP_NAME = gopulse

build:
	go build -o bin/$(APP_NAME) ./cmd/gopulse
	
install:
	sudo cp bin/$(APP_NAME) /usr/local/bin/

