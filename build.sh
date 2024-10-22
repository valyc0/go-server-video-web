#!/bin/bash

# Nome del file eseguibile
BINARY_NAME="video_server"

# Directory dei file di destinazione
OUTPUT_DIR="./bin"

# Crea la directory di output se non esiste
mkdir -p $OUTPUT_DIR

# Compila per Linux
GOOS=linux GOARCH=amd64 go build -o $OUTPUT_DIR/$BINARY_NAME-linux-amd64 video_server.go
echo "Compilato per Linux in $OUTPUT_DIR/$BINARY_NAME-linux-amd64"

# Compila per Windows
GOOS=windows GOARCH=amd64 go build -o $OUTPUT_DIR/$BINARY_NAME-windows-amd64.exe video_server.go
echo "Compilato per Windows in $OUTPUT_DIR/$BINARY_NAME-windows-amd64.exe"

# Compila per macOS
GOOS=darwin GOARCH=amd64 go build -o $OUTPUT_DIR/$BINARY_NAME-darwin-amd64 video_server.go
echo "Compilato per macOS in $OUTPUT_DIR/$BINARY_NAME-darwin-amd64"

echo "Compilazione completata!"

