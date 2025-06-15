all: build

build:
	cd src && go build -o sun .  && cd .. 

install: install-doc
	sudo cp -f ./src/sun /usr/local/bin

uninstall:
	sudo rm /usr/local/bin/sun

clean:
	rm -f ./src/sun

run: build
	./src/sun

doc: build

	chmod +x doc/sun.1

install-doc: doc

	sudo cp ./doc/sun.1 /usr/local/share/man/man1/ 
