all: build

build:
	cd src && go build -o sun .  && cd .. 

install:
	sudo cp -f ./src/sun /usr/local/bin

clean:
	rm -f ./src/sun

run: build
	./src/sun

	
