.PHONY: build fmt clean install_deps

default: all

all: clean install_deps fmt build

build:
	$(info ******************** build ********************)
	yarn build

install_deps:
	$(info ******************** install_deps ********************)
	yarn install

fmt:
	$(info ******************** fmt ********************)
	./node_modules/.bin/prettier --write --plugin-search-dir=. ./**/*.svelte
	./node_modules/.bin/prettier --write --plugin-search-dir=. ./**/*.html
	./node_modules/.bin/prettier --write --plugin-search-dir=. ./**/*.js

clean:
	$(info ******************** clean ********************)
	rm -rf public/build
