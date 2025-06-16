all: build

build:
	cd src && go build -o sun .  && cd .. 

install: install-doc
	sudo cp -f ./src/sun /usr/local/bin

uninstall: uninstall-doc
	sudo rm /usr/local/bin/sun

clean:
	rm -f ./src/sun

run: build
	./src/sun

doc: build

	chmod +x doc/sun.1

install-doc: doc

	sudo cp ./doc/sun.1 /usr/local/share/man/man1/ 

uninstall-doc:

	sudo rm /usr/local/share/man/man1/sun.1

install-desktop:

	sudo cp assets/sun.desktop /usr/share/applications/
	sudo cp assets/sun.png /usr/share/icons/

uninstall-desktop:

	sudo rm /usr/share/applications/sun.desktop
	sudo rm /usr/share/icons/sun.png
