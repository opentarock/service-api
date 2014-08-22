#!/usr/bin/enc bash

go run go/cli/oauth2cli/main.go access_token "client_id" "client_secret" "{\"grant_type\":\"refresh_token\", \"refresh_token\":\"$1\"}"
