
mongo:
	docker run --name local-mongo -p 27017:27017 -e MONGO_INITDB_ROOT_USERNAME=root -e MONGO_INITDB_ROOT_PASSWORD=root -d mongo:4.4

run:
	go run main.go

.PHONY:mongo run