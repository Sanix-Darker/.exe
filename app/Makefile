docker-build:
	docker build -t sanixdarker/d-exe:latest -f Dockerfile .

docker-run:
	docker run -it --rm -p 4321:4321 sanixdarker/d-exe:latest

run:
	go run main.go

build:
	go build -o ./.exe

exec: build
	./.exe

