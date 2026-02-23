# Base16 Builder

A Go implementation of a Base16 Builder that follows the guidelines at [Base16](https://github.com/chriskempson/base16).

## Prerequisites
* **Go** (1.16 or later) - [Install Guide](https://go.dev/doc/install)
* **Environment**: Ensure your Go binary directory is in your system's `PATH`.
    * **Linux/macOS**: Add `export PATH=$PATH:$(go env GOPATH)/bin` to your `.bashrc` or `.zshrc`.
    * **Windows**: Add `%USERPROFILE%\go\bin` to your Path via *System Environment Variables*.

## Installation

### Via Go CLI (Recommended)
```bash
go install github.com/mxilinas/base16-builder-go/cmd/base16-builder@latest
```

### Build Manually
```
git clone https://github.com/mxilinas/base16-builder-go
cd base16-builder-go
go build -o base16-builder
```
*Note: You would need to run the binary using ./base16-builder*

## Usage

```bash
cat scheme.yaml | base16-builder --template template.mustache > theme.file
```
