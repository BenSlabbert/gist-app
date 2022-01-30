.PHONY: build fmt clean

default: all

all: clean fmt build

build:
	$(info ******************** build ********************)
	yarn build

fmt:
	$(info ******************** fmt ********************)
	./node_modules/.bin/prettier --write --plugin-search-dir=. ./**/*.svelte
	./node_modules/.bin/prettier --write --plugin-search-dir=. ./**/*.html
	./node_modules/.bin/prettier --write --plugin-search-dir=. ./**/*.js

clean:
	$(info ******************** clean ********************)
	rm -rf public/build
