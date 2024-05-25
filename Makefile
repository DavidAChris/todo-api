build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/bootstrap main.go; zip -j bootstrap.zip ./bin/bootstrap

run: build
	@./bin/todo-api

clean:
	@rm -rf ./bin
	@rm bootstrap.zip
