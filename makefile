all:
	docker build ./db
	docker build ./api --build-arg app_env=development 

install:
	go get -u github.com/gorilla/mux
	go get -u github.com/jinzhu/gorm
	go get -u github.com/lib/pq

	mkdir -p secrets && touch secrets/AccountSid
