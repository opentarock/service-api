#!/usr/bin/env bash

go run go/cli/usercli/main.go register_user '{"display_name":"Example Name", "email":"email@example.com", "password":"password"}' "https://redirect.example.com"
