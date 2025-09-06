# MAYI-CLI

A powerful CLI tool for managing codeowners files with intelligent conflict detection, rule analysis, and automatic optimization. Powered by [mayi](https://github.com/MoonMoon1919/mayi)

## Features

- 🖊️ Add, update, and remove rules
- ✅ Validate rules and accuracy
- 🔎 Check who can review files & directories
- 📄 Generate files


## Quick Start

### Installation

```sh
go install github.com/MoonMoon1919/mayi-cli@latest
```

### Basic commands

```sh
mayi-cli create

# Add a rule
mayi-cli add rule --pattern 'docs/*' --owner '@MoonMoon1919' --owner '@toastsandwich123'

# Search for owners
mayi-cli get owners --pattern docs/

# Add another owner
mayi-cli add owner --pattern 'docs/*' --owner '@example'

# Analyze the file
mayi-cli analyze
```

## Usage

### Creating files

```sh
# Defaults to .github/CODEOWNERS
mayi-cli create

# Or override to your desired path
mayi-cli create --path CODEOWNERS
```

### Adding rules

```sh
# Add a basic rule to have one owner for docs
mayi-cli add rule --pattern 'docs/*' --owner '@MoonMoon1919'

# Except for samples - require no owner for those
mayi-cli add rule --pattern 'docs/internal/samples/*' --owner '' --action exclude
```

### Adding rule owners

```sh
mayi-cli add owner --pattern 'docs/*' --owner @example
```

### Removing rules

```sh
mayi-cli delete rule --pattern 'docs/internal/samples/*'
```

### Removing rule owners

```sh
mayi-cli delete owner --pattern 'docs/*' --owner @example
```

### Searching for rules

```sh
mayi-cli get owners --pattern docs/
```

### Moving rules

```sh
# Add a random rule
mayi-cli add rule --pattern '*.md' --owner @MoonMoon1919

# Then move it
mayi-cli move --source-pattern '*.md' --destination-pattern 'docs/*' --direction before
```

### Analyzing

```sh
# Analyze without fixing
mayi-cli analyze

# Or fix any errors the analyzer encounters
mayi-cli analyze --fix
	
```

## Contributing

See [CONTRIBUTING](./CONTRIBUTING.md) for details.

## License

MIT License - see [LICENSE](./LICENSE) for details.

## Disclaimers

This work does not represent the interests or technologies of any employer, past or present. It is a personal project only.
