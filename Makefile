.PHONY: build-local build templ notify-templ-proxy run

-include .env

build-local:
	@go build -o ./bin/tudo-tools ./cmd

build:
	@npm run build

templ:
	@templ generate --watch --proxy=http://localhost:$(APP_PORT) --proxyport=$(TEMPL_PROXY_PORT) --open-browser=false --proxybind='0.0.0.0'

notify-templ-proxy:
	@templ generate --notify-proxy --proxyport=$(TEMPL_PROXY_PORT)

run:
	@make templ & sleep 1
	@air