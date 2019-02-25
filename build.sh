#!/bin/bash

GOOS=darwin go build -o bin/terminal-mac main.go
GOOS=linux go build -o bin/terminal-linux main.go