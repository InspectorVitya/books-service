run:
	go run cmd/books/main.go

test:
	go test -v -race -timeout 30s ./...

generate:
	mkdir -p book
	protoc \
    		--proto_path=api \
    		--go_out=book \
    		--go-grpc_out=book \
    		api/*.proto

build-db:
	docker build -t mysql-books docker/mysql

run-db:
	docker run -d -p 3808:3306 --name mysql-books mysql-books

stop-db:
	docker stop mysql-books

remove-db:
	docker rm mysql-books

