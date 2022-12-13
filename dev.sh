#!/bin/bash
pnpm nodemon --watch './**/*.go' --signal SIGTERM --exec 'go' run *.go
