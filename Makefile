build:
	@go build -o ./bin/todo-api

run: build
	@./bin/todo-api

clean:
	@rm -rf ./bin