update:
	go get -u
	go mod tidy
	go mod vendor
	go mod verify

test:
	go test -v ./...
	# go test `go list ./... | grep -v e2e_test`
run:
	go run main.go
build:
	go build -o main main.go

docker_test: docker-build_test 
	sudo docker compose up -d 
	sudo docker compsoe exec -T http go test ./...
	sudo docker compose down 
docker_build:
	sudo docker build -t shahriarraka/null_pointer_goapi:latest .
	# docker build -t shahriarraka/null_pointer_goapi:latest .
docker_push:
	sudo docker push shahriarraka/null_pointer_goapi:latest 
docker_build_test:
	docker build . -t service_test --target=test

docker_run:
	sudo docker stop null_pointer_goapi || true && sudo docker rm null_pointer_goapi || true
	sudo docker run --publish 8080:8080 --name null_pointer_goapi1 shahriarraka/null_pointer_goapi:latest  
