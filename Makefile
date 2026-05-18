NAME1 = push-swap
NAME2 = checker

build:
    go build -o $(NAME1) ./push-swap
	go build -o $(NAME2) ./checker

run:
    go run ./push-swap

test:
    go test ./...

clean:
    rm -f $(NAME1) $(NAME2)

re: clean build