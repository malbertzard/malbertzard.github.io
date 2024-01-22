dev:
	cd cmd && go run .

watch:
	cd cmd && ~/go/bin/air --tmp_dir ../tmp

build:
	cd cmd && go build -o ../tmp/ .

build-css:
	npx tailwindcss -i ./styles.css -o ./static/styles.css

watch-css:
	npx tailwindcss -i ./styles.css -o ./static/styles.css --watch
