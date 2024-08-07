build:
	env GOOS=linux env go build -ldflags="-s -w" -o bin/bootstrap main.go; zip -jr bootstrap.zip ./bin/bootstrap

run: build
	@./bin/todo-api

clean:
	@rm -rf ./bin
	@rm bootstrap.zip
