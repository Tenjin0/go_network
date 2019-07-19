include .env

FILES = $(shell ls $(ENV_FOLDER) | grep .go | grep -iv main.go | sed 's/^/$(ENV_FOLDER)\//')

run:
	go run $(ENV_FOLDER)/main.go $(FILES) $(target) $(host)