APP_NAME = "gin_bbs"

default:
	go build -o ${APP_NAME}
	# env GOOS=linux GOARCH=amd64 go build -o ${APP_NAME}

install:
	go mod download

dev:
  # go get github.com/pilu/fresh
	fresh -c ./fresh.conf

api-doc:
  # go get -u github.com/swaggo/swag/cmd/swag
	swag init

mock:
	go run ./main.go -m

clean:
	if [ -f ${APP_NAME} ]; then rm ${APP_NAME}; fi

help:
	@echo "make - compile the source code"
	@echo "make install - install dep"
	@echo "make dev - run go fresh"
	@echo "make mock - mock data"
	@echo "make clean - remove binary file"
