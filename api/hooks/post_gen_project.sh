#!/usr/bin/env bash

set -eu -o pipefail

cp .env.example .env

go mod init github.com/{{ cookiecutter.github.org }}/{{ cookiecutter.github.repository }}
go get -u ./...

git init
git add .
git commit -m "Init"

pre-commit install
pre-commit

echo "[DONE] Your API is ready. Run 'cd {{ cookiecutter.app.name }} && make' to see all available commands."