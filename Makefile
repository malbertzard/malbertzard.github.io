dev:
	cd server && go run .

watch:
	cd server && ~/go/bin/air --tmp_dir ../tmp

build:
	cd server && go build -o ../tmp/ .

build-css:
	npx tailwindcss -i ./styles.css -o ./public/styles.css

watch-css:
	npx tailwindcss -i ./styles.css -o ./public/styles.css --watch
