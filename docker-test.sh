#!/bin/bash
docker run -it -v "$PWD":/app -w /app --rm usb-test:dev go run main.go