all:
	go build -o dist/App
	./dist/App

install:
	go get -u github.com/gorilla/mux
	mkdir -p secrets && touch secrets/AccountSid
