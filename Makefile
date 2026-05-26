NAME1 = push-swap
NAME2 = checker

build:
	go build -o $(NAME1) ./cmd/push-swap
	go build -o $(NAME2) ./cmd/checker

run:
	go run ./cmd/push-swap

test:
	go test ./...

random:
	./tests/random_test.sh

benchmark:
	./tests/benchmark.sh

clean:
	rm -f $(NAME1) $(NAME2)

re: clean build