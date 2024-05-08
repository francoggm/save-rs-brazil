start-server:
	@echo Starting backend server
	docker-compose up --build -d 

stop-server:
	@echo Stopping backend
	docker-compose down

server-debug:
	@echo Start debug in server
	cd src && go run main.go