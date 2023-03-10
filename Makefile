all: clean build run

build:
	go build -o orca

run:
	./orca

clean:
	rm -f orca
