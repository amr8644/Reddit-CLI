run:
	go run main.go requests.go auth.go users.go subreddits.go

build:
	go build main.go requests.go auth.go users.go subreddits.go


clean:
	rm -rf output
