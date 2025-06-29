# 📦 Go Reloaded

Welcome to **Go Reloaded**, a project designed to sharpen your skills in the Go programming language by building a string transformation engine that reads from a text file, processes transformation tags, and outputs the final result.

---

## 📖 Project Overview

This program takes a `.txt` file as input, scans for specific formatting tags within the text, and replaces or transforms the words accordingly. It emphasizes core Go concepts like:

- String manipulation
- File I/O
- Error handling
- Modular code structure
- Unit testing

---

## 🧠 Transformation Rules

The program supports the following transformation tags:

| Tag        | Description                                  | Example Input              | Example Output   |
|------------|----------------------------------------------|----------------------------|------------------|
| `(up)`     | Converts the previous word to UPPERCASE      | `very (up)happy(up)`       | `very HAPPY`     |
| `(low)`    | Converts the previous word to lowercase      | `Was (low)GREAT(low)`      | `Was great`      |
| `(cap)`    | Capitalizes the first letter                 | `such a (cap)wonderful(cap)`| `such a Wonderful` |
| `(hex)`    | Converts the number to hexadecimal           | `value is (hex)255(hex)`   | `value is ff`    |
| `(bin)`    | Converts the number to binary                | `ID is (bin)13(bin)`       | `ID is 1101`     |

---

## 📂 Project Structure

go-reloaded/
├── internal/ # Contains core logic (e.g., parsers, transformers)
├── test/ # Unit tests for different components
├── go.mod # Go module file
├── main.go # Program entry point
├── sample.txt # Input file with tags
├── result.txt # Output file after transformations
└── README.md # This file

---

## 🚀 How to Run

### 1. Clone the repository

```bash
git clone https://github.com/Simon4-ik/go-reloaded.git
cd go-reloaded
2. Prepare your input file
Edit or create sample.txt with content like this:


I am (cap)excited(cap) about this (up)project(up).
The result is (hex)170(hex) and (bin)8(bin).
3. Run the program

go run main.go sample.txt result.txt
The output will be printed to the terminal and/or saved to output.txt.

✅ Example Output

I am Excited about this PROJECT.
The result is aa and 1000.
🧪 Testing
If you've added tests (recommended), run them with:

go test ./...

👨‍💻 Author
Made by 💻 Muhammadabdulloh using Go for educational purposes.

