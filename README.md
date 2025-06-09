# Hello Go

A simple Go application that greets the world in multiple languages.

## Features

- Support for 6 different languages
- Command-line interface with language selection
- Clean and simple Go code structure

## Supported Languages

- `el` - Greek (Χαίρετε Κόσμε)
- `en` - English (Hello world) - Default
- `fr` - French (Bonjour le monde)
- `he` - Hebrew (שלום עולם)
- `ur` - Urdu (ہیلو دنیا)
- `vi` - Vietnamese (Xin chào Thế Giới)

## Usage

### Running with Go

```bash
# Default greeting in English
go run main.go

# Greeting in Vietnamese
go run main.go -lang=vi

# Greeting in French
go run main.go -lang=fr

# List available options
go run main.go -help
```

### Building and Running

```bash
# Build the application
go build -o hello-go

# Run the built binary
./hello-go -lang=ur
```

## Requirements

- Go 1.24.3 or later

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/phihdn/hello-go.git
   cd hello-go
   ```

2. Run the application:

   ```bash
   go run main.go -lang=vi
   ```

## Examples

```bash
$ go run main.go
Hello world

$ go run main.go -lang=vi
Xin chào Thế Giới

$ go run main.go -lang=el
Χαίρετε Κόσμε

$ go run main.go -lang=xyz
unsupported language: "xyz"
```
