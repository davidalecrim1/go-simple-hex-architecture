build:
	docker build -t go_simple_hex_architecture:latest .

local_run:
	docker run --name go_simple_hex_architecture -d -p 8080:8080 go_simple_hex_architecture:latest

local_inside:
	docker exec -it go_simple_hex_architecture /bin/sh

local_stop:
	docker stop go_simple_hex_architecture
	docker rm go_simple_hex_architecture