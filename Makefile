BIN_NAME = sun
SRC_DIR = src
MAN_DIR = /usr/local/share/man/man1/
BIN_DIR = /usr/local/bin/
ICON_DIR = /usr/share/icons/
DESKTOP_DIR = /usr/share/applications/
DETECT_CMD = detect
INIT_CMD = init --name="Sun" --language="Go" --author="Jellyfish & NotBitly" --author-email="sun.developpments@proton.me" --maintener="Jellyfish & NotBitly" --maintener-email="sun.developpments@proton.me" --platform="Github" --repo="https://github.com/SunDeveloppments/sun"
INIT_Y = init --y

.PHONY: all build install uninstall clean run doc install-doc uninstall-doc install-desktop uninstall-desktop

all: script build cobra_install

cobra_install:
	cd ./src && go get github.com/spf13/cobra && cd ..

script:
	@if command -v go >/dev/null 2>&1; then \
		cd src && go get -u github.com/spf13/cobra@latest; \
	else \
		echo "Cannot find any go installation."; \
	fi

build:
	@cd $(SRC_DIR) && go build -o ../$(BIN_NAME) . && cd ..

install: install-doc install-desktop install-assets
	@sudo cp -f ./$(BIN_NAME) $(BIN_DIR)

uninstall: uninstall-doc
	@sudo rm -f $(BIN_DIR)/$(BIN_NAME)

clean:
	@rm -f ./${SRC_DIR}/$(BIN_NAME)

test: build
	@./$(BIN_NAME) $(INIT_Y)
	@./$(BIN_NAME) $(DETECT_CMD)

run: build
	@./$(BIN_NAME)

doc: build
	@chmod +x doc/$(BIN_NAME).1

install-doc: doc
	@sudo cp ./doc/$(BIN_NAME).1 $(MAN_DIR)

uninstall-doc:
	@sudo rm -f $(MAN_DIR)/$(BIN_NAME).1

install-desktop:
	@sudo cp assets/$(BIN_NAME).desktop $(DESKTOP_DIR)
	@sudo cp assets/$(BIN_NAME).png $(ICON_DIR)

uninstall-desktop:
	@sudo rm -f $(DESKTOP_DIR)/$(BIN_NAME).desktop
	@sudo rm -f $(ICON_DIR)/$(BIN_NAME).png

install-assets:
	@mkdir -p ~/.config/sun/
	@echo "Copying files…"
	@cp -rf ./assets/  ~/.config/sun/

uninstall-assets:
	@rm -rf ~/.config/sun/assets
