BIN_NAME = sun
SRC_DIR = src
MAN_DIR = /usr/local/share/man/man1/
BIN_DIR = /usr/local/bin/
ICON_DIR = /usr/share/icons/
DESKTOP_DIR = /usr/share/applications/

.PHONY: all build install uninstall clean run doc install-doc uninstall-doc install-desktop uninstall-desktop

all: build

build:
	cd $(SRC_DIR) && go build -o $(BIN_NAME) . && cd ..

install: install-doc
	sudo cp -f ./${SRC_DIR}/$(BIN_NAME) $(BIN_DIR)

uninstall: uninstall-doc
	sudo rm -f $(BIN_DIR)/$(BIN_NAME)

clean:
	rm -f ./${SRC_DIR}/$(BIN_NAME)

run: build
	./${SRC_DIR}/$(BIN_NAME)

doc: build
	chmod +x doc/$(BIN_NAME).1

install-doc: doc
	sudo cp ./doc/$(BIN_NAME).1 $(MAN_DIR)

uninstall-doc:
	sudo rm -f $(MAN_DIR)/$(BIN_NAME).1

install-desktop:
	sudo cp assets/$(BIN_NAME).desktop $(DESKTOP_DIR)
	sudo cp assets/$(BIN_NAME).png $(ICON_DIR)

uninstall-desktop:
	sudo rm -f $(DESKTOP_DIR)/$(BIN_NAME).desktop
	sudo rm -f $(ICON_DIR)/$(BIN_NAME).
