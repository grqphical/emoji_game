run:
	templ generate
	go run .

build:
	templ generate
	go build -o build/emoji_game .