BIN_NAME = sun
SRC_DIR = src
MAN_DIR = /usr/local/share/man/man1/
BIN_DIR = /usr/local/bin/
ICON_DIR = /usr/share/icons/
DESKTOP_DIR = /usr/share/applications/
DETECT_CMD = stats
INIT_Y =  init --y
BUILD_VERSION = 0.1

.PHONY: all build install uninstall clean run doc install-doc uninstall-doc install-desktop uninstall-desktop install-assets uninstall-assets

all: build 

build:
	@echo "Build…"
	@echo "Mode: local"
	@go build -ldflags "-X 'main.LocalInstall=true' -X 'main.SysInstall=false' -X 'main.PortableInstall=false' -X 'main.Version=$(BUILD_VERSION)'" .
	@echo "Done."
install: install-doc install-desktop install-assets
	@echo "Copying…"
	@sudo cp ./$(BIN_NAME) $(BIN_DIR)
	@echo "Done."

uninstall: uninstall-doc
	@sudo rm -f $(BIN_DIR)/$(BIN_NAME)

clean:
	@rm -f ./${SRC_DIR}/$(BIN_NAME)

test: build
	@./$(BIN_NAME) $(INIT_Y)
	@./$(BIN_NAME) $(DETECT_CMD)

test-help: build
	@./$(BIN_NAME) help init
	@./$(BIN_NAME) help read
	@./$(BIN_NAME) help stats
	@./$(BIN_NAME) help completion

run: build
	@./$(BIN_NAME)

doc:
	@chmod +x doc/$(BIN_NAME).1

install-doc: doc
	@sudo cp ./doc/$(BIN_NAME).1 $(MAN_DIR)

uninstall-doc:
	@sudo rm -f $(MAN_DIR)/$(BIN_NAME).1

install-desktop:
	@sudo cp ./assets/$(BIN_NAME).desktop $(DESKTOP_DIR)
	@sudo cp ./assets/$(BIN_NAME).png $(ICON_DIR)

uninstall-desktop:
	@sudo rm -f $(DESKTOP_DIR)/$(BIN_NAME).desktop
	@sudo rm -f $(ICON_DIR)/$(BIN_NAME).png

install-assets:
	@mkdir -p ~/.config/sun/
	@cp -rf ./assets/  ~/.config/sun/

uninstall-assets:
	@rm -rf ~/.config/sun/assets
