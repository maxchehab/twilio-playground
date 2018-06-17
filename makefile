all:
	docker-compose up

build: 
	docker-compose build

install:
	mkdir -p secrets && touch secrets/AccountSid

prod:
	docker build ./db
	docker build ./api --build-arg app_env=production
	docker build ./frontend --build-arg app_env=production