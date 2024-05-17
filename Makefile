create-all: stop-all
	@echo Build all images
	docker-compose up --build -d

create-server: stop-server
	@echo Building and starting server
	docker-compose up --build -d

start-server:
	@echo Starting backend server
	docker-compose up -d 

stop-server:
	@echo Stopping backend
	docker-compose down backend

stop-all:
	@echo Stopping all
	docker-compose down	

debug-server:
	@echo Starting debug local server
	cd server && go run cmd/api/main.go