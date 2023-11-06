#!/usr/bin/env bash

go mod download
go install github.com/cosmtrek/air@latest


for file in ./addon/*; \
    do \
      if [ -f "$file/main.go" ]; \
        then \
          go build -buildmode=plugin -o $file/main.so $file/main.go; \
      else \
        echo $file; \
      fi \
    done

go build -o main .

go run .