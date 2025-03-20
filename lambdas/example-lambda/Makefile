test: 
		

local:
	npm install -D @types/aws-lambda esbuild

build:
	rm -rf dist
	npx esbuild handler.ts --bundle --minify --sourcemap --platform=node --target=es2020 --outfile=dist/handler.js

	cd dist && zip -r handler.zip handler.js*



.PHONY: test build local