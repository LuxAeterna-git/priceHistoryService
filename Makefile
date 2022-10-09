run:
	go run cmd/main.go
up:
	docker run --name mngPriceStorage -e MONGO_DATABASE=files -e MONGO_INITDB_ROOT_USERNAME=root -e MONGO_INITDB_ROOT_PASSWORD=qwerty -p 27017:27017 --rm -d mongo
down:
	docker stop mngPriceStorage