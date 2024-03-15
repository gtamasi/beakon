build:
	@echo "Building beakon"
	go build -o ./bin/beakon ./cmd/beakon/main.go


run:
	@	echo "Running beakon dev"
	go run ./cmd/beakon/main.go

tailwind_build:
	@echo "Building tailwind"
	@npx tailwindcss build -o ./assets/styles/styles.css -i ./templates/styles.css

tailwind_watch:
	@echo "Watching tailwind"
	@npx tailwindcss build -o ./assets/styles/styles.css -i ./templates/styles/styles.css --watch

templ_build:
	@echo "Building templates"
	templ generate -path templates
templ_watch:
	@echo "Building templates"
	templ generate -path templates -watch
