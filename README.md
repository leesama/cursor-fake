# Cursor Fake

[English](./README.md) | [中文](./README_zh.md)

A tool designed to reset device IDs for Cursor IDE. It helps you generate new machine IDs and device IDs when needed.

## Features

- Reset Cursor IDE device identifiers
- Generate new macMachineId and machineId
- Create new devDeviceId
- Cross-platform support (Windows, macOS, Linux)
- Simple one-click operation
- Automated releases for all platforms

## What It Does

This tool helps you:
- Reset the Cursor IDE device identification
- Generate new machine IDs for your system
- Create fresh device IDs for Cursor
- Modify the storage.json configuration

The tool will generate:
- New macMachineId
- New machineId
- New devDeviceId

## Quick Start

1. Download the appropriate version for your system from [Releases](../../releases):
   - Windows: `cursor-fake-windows-amd64.exe`
   - macOS (Apple Silicon): `cursor-fake-darwin-arm64`
   - macOS (Intel): `cursor-fake-darwin-amd64`
   - Linux: `cursor-fake-linux-amd64`

2. Run the program:
   - Windows: Double click the `.exe` file
   - macOS/Linux: Open terminal and run:
     ```bash
     chmod +x cursor-fake-*  # Make it executable (first time only)
     ./cursor-fake-*  # Run the program
     ```

## Usage

1. Close Cursor IDE if it's running
2. Run the program as described above
3. The tool will automatically:
   - Locate your Cursor storage file
   - Generate new IDs
   - Update the configuration
4. Press Enter to exit after completion

## For Developers

If you want to build from source, you'll need:
- Go 1.21 or higher
- Make (for build automation)

### Build from Source

1. Clone the repository
```bash
git clone [repository-url]
cd cursor-fake
```

2. Build for your platform
```bash
make build
```

Or build for all platforms:
```bash
make build-all
```

### Available Commands

- `make build` - Build for current platform
- `make build-all` - Build for all platforms
- `make run` - Run the program
- `make test` - Run tests
- `make clean` - Clean build artifacts
- `make deps` - Update dependencies

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License

MIT License 