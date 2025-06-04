# Go Morse Code Application

This project is a simple Go application that converts text input into Morse code tones. It adheres to ITU standards for character spacing and tone representation.

## Project Structure

```
go-morse-code-app
├── src
│   ├── main.go          # Entry point of the application
│   ├── morse
│   │   ├── encoder.go   # Function to encode text to Morse code
│   │   └── table.go     # Constants and mappings for Morse code
├── go.mod               # Module definition
└── README.md            # Project documentation
```

## Installation

To run this application, ensure you have Go installed on your machine. Clone the repository and navigate to the project directory:

```bash
git clone <repository-url>
cd go-morse-code-app
```

## Running the Application

To run the application, use the following command:

```bash
go run src/main.go
```

## Usage

After running the application, you will be prompted to enter text. The application will output the corresponding Morse code tones.

### Example

**Input:**
```
SOS
```

**Output:**
```
Dot, Dot, Dot, SilenceDot, Dash, Dash, Dash, SilenceDot, Dot, Dot, Dot
```

## Contributing

Feel free to submit issues or pull requests if you have suggestions or improvements for the project.