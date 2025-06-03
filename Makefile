all: build

build:
	go build -o sun ./src 

install:
	sudo cp -f ./sun /usr/local/bin

clean:
	rm -f ./sun

run: build
	./sun

	
