# ogd77cwgen

A command-line tool for converting text to Morse code (CW) and playing it back as an audio tone, using the OpenGD77 tone table and ITU timing standards.

## Directory Structure

```
.
├── README.md
└── src
    ├── go.mod
    ├── go.sum
    ├── morse
    │   ├── encoder.go
    │   └── table.go
    └── ogd77cwgen.go
```

## Features

- Converts input text to Morse code using ITU standard timing.
- Plays Morse code as an audio tone using a sine wave (FM synth) through your sound card.
- Lets you select code speed (WPM) and tone (note) via optional command-line flags.
- Defaults to 25 WPM and note 38 (A# 932.3 Hz) if no flags are provided.

## Requirements

- Go 1.18 or newer
- [faiface/beep](https://github.com/faiface/beep) audio library

## Installation

1. Clone this repository.
2. Change to the `src` directory:
   ```sh
   cd src
   ```
3. Download Go dependencies:
   ```sh
   go mod tidy
   ```

## Usage

### Run the program

```sh
go run ogd77cwgen.go
```

### Optional flags

- `--wpm` &nbsp;&nbsp;&nbsp;&nbsp;Set Morse code speed in words per minute (default: 25)
- `--tone` &nbsp;&nbsp;Set the note value for the Morse tone (see table below, default: 38 for A# 932.3 Hz)

#### Examples

```sh
go run ogd77cwgen.go
go run ogd77cwgen.go --wpm 20 --tone 40
```

### Interactive Usage

1. Enter the text you want to convert to Morse code when prompted.
2. The program will display the Morse code as a sequence of tone/duration pairs.
3. You will be asked if you want to hear the Morse code. Enter `y` to play it.

## Tone Table

| Value | Note | Frequency (Hz) |
|-------|------|---------------|
| 1     | A    | 110.0         |
| 2     | A#   | 116.5         |
| 3     | B    | 123.5         |
| ...   | ...  | ...           |
| 38    | A#   | 932.3         |
| 40    | C    | 1046.5        |
| ...   | ...  | ...           |
| 45    | F    | 1397.0        |

(See `ogd77cwgen.go` for the full table.)

## Morse Code Mapping

- ITU standard for character and word spacing is used.
- The Morse code mapping is in `src/morse/table.go`.

## Building Binaries

You can build standalone binaries for your platform using Go's cross-compilation support.

### Linux (x86_64, .deb and .rpm based)

From the `src` directory:

```sh
# Build for Linux x86_64
GOOS=linux GOARCH=amd64 go build -o ogd77cwgen ogd77cwgen.go
```

You can then copy `ogd77cwgen` to `/usr/local/bin` or any directory in your `$PATH`.

#### For .deb-based systems (Debian, Ubuntu, Mint, etc.)

```sh
sudo cp ogd77cwgen /usr/local/bin/
```

#### For .rpm-based systems (Fedora, CentOS, RHEL, openSUSE, etc.)

```sh
sudo cp ogd77cwgen /usr/local/bin/
```

### Windows

```sh
# Build for Windows x86_64
GOOS=windows GOARCH=amd64 go build -o ogd77cwgen.exe ogd77cwgen.go
```

You can then run `ogd77cwgen.exe` on any Windows 64-bit system.

### MacOS

```sh
# Build for MacOS x86_64
GOOS=darwin GOARCH=amd64 go build -o ogd77cwgen ogd77cwgen.go
```

You can then run `ogd77cwgen` on any Intel-based Mac.

---

**Note:**  
- For Apple Silicon (M1/M2), use `GOARCH=arm64`.
- You need Go installed on your system for building.  
- The resulting binary is standalone and does not require Go to be installed on the target system.

---

## License

MIT

---

**73!**