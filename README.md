# Emoji Game

A simple browser game where you have to try and guess an emoji. Try to aim for the highest streak possible

Made with Golang, HTMX, Alpine.js and Chi

## Self Hosting Guide

1. Clone this repo

```bash
$ git clone https://github.com/grqphical/emoji-game
```

2. Make sure you have Go 1.21 installed

3. Create a .env file and define HOST_ADDR as an enviroment variable with the address you want to host on. The example below will host on port 8000

```
HOST_ADDR=:8000
```

3. Run the makefile

```bash
$ make run
```

## License

This project is licensed under the MIT license
