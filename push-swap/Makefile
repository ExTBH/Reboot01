.PHONY: all clean

all: push-swap checker

push-swap:
	go build -o push-swap cmd/push-swap/main.go

checker:
	go build -o checker cmd/checker/main.go

clean:
	rm push-swap checker
