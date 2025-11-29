#!/bin/bash

# Build script for Todo List CLI

echo "🔨 Building Todo List CLI..."

# Build for current platform
go build -o todo-list-cli main.go

if [ $? -eq 0 ]; then
    echo "✅ Build successful!"
    echo "Run with: ./todo-list-cli help"
else
    echo "❌ Build failed!"
    exit 1
fi
