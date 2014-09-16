#!/usr/bin/env bash

go run go/cli/oauth2cli/main.go access_token "client_id" "client_secret" '{"grant_type":"password", "username":"email@example.com", "password":"password"}'
