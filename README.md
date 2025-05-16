# Thermogo CLI

<img align="right" width="90px" src="assets/logo.svg">

Thermogo CLI is a command-line tool to convert temperatures between **Celsius**, **Fahrenheit**, and **Kelvin**.
It uses interactive prompts for selecting units and inputting temperature values.
Built with [Cobra](https://github.com/spf13/cobra) and [PromptUI](https://github.com/manifoldco/promptui), written in Go.

## 📦 Installation

```sh
go install github.com/luissena/thermogo-cli@latest
```

Or clone and build manually:

```sh
git clone https://github.com/luissena/thermogo-cli.git
cd thermogo-cli
go build -o thermogo-cli
```

## 🚀 Usage

```sh
thermogo-cli convert
```

You will be prompted to:

- Select the input temperature unit
- Select the target unit
- Enter the temperature value

### Example

```sh
thermogo-cli convert
```

```
? Select input unit: Fahrenheit
? Select output unit: Celsius
? Enter temperature value: 98.6

Result: 98.60 Fahrenheit → 37.00 Celsius
```

## ⚠️ Error Handling

- Only numeric input values are accepted.
- If a non-numeric value is entered, an error message will be displayed and the program will exit.

### Error Example

```sh
? Enter temperature value: abc
# Error: failed to read temperature value: strconv.ParseFloat: parsing "abc": invalid syntax
```

## Shell Autocompletion

Thermogo CLI supports autocompletion scripts for `sh`, `zsh`, `fish`, and `powershell`.

You can generate the completion script using:

```sh
thermogo-cli completion [sh|zsh|fish|powershell]
```

### Example: Enable completion for Zsh

```sh
thermogo-cli completion zsh > /usr/local/share/zsh/site-functions/_thermogo-cli
```

Refer to Cobra or your shell's documentation for complete setup.

## 📁 Project Structure

```plaintext
.
├── cmd
│   ├── root.go
│   └── convert.go
├── assets
│   └── logo.svg
├── main.go
├── go.mod
└── go.sum
```

## 📚 Requirements

- Go 1.20 or newer
- [Cobra](https://github.com/spf13/cobra)
- [PromptUI](https://github.com/manifoldco/promptui)

## 📄 License

Licensed under the MIT License.
