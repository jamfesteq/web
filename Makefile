VERSION ?= 1.0.0
NAME := jamfesteq.web.app

.PHONY: server
server:
	@hugo server -b http://localhost:1313/
.PHONY: build
build:
	@#rm -rf public/*
	@hugo -b https://jamfesteq.web.app/
relogin:
	firebase logout
	firebase login
	firebase use xackery
deploy: build
	@firebase deploy
set-version:
	@echo "VERSION=${VERSION}" >> $$GITHUB_ENV

.PHONY: scss
scss:
	@(echo "/* This file is autogenerated via 'make scss'. DO NOT EDIT */" > static/css/tpe-scss.css)
	docker run --rm -v $(PWD):$(PWD) -w $(PWD) jbergknoff/sass scss/app.scss >> static/css/tpe-scss.css

.PHONY: spells
spells:
	go run scripts/spells/main.go

.PHONY: rule
rule:
	go run scripts/rule/main.go

.PHONY: rule-inject
rule-inject:
	cd .. && ./yakuku inject web/data/rule.sql	