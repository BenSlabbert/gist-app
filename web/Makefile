.PHONY: build
build:
	@yarn build

.PHONY: fmt
fmt:
	@./node_modules/.bin/prettier --write --plugin-search-dir=. ./**/*.svelte
	@./node_modules/.bin/prettier --write --plugin-search-dir=. ./**/*.html
	@./node_modules/.bin/prettier --write --plugin-search-dir=. ./**/*.js
