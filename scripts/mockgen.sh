#!/usr/bin/env bash

mockgen -source=./go/client/lobby.go -destination=./go/mock/clientmock/lobby.go -package=clientmock
