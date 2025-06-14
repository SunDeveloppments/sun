all: build

build:
	cd src && go build -o sun .  && cd .. 

install:
	sudo cp -f ./src/sun /usr/local/bin

uninstall:
        sudo rm /usr/local/bin/sun

clean:
	rm -f ./src/sun

run: build
	./src/sun

	
