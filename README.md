# Morse Code Translator

This is a simple program that translates text to Morse code and vice versa. It allows you to input a line of text and converts it to Morse code if it is regular text, or decodes it from Morse code to text if it is Morse code.

## Installation

To use this program, you need to have Go installed on your system. If you don't have Go installed, you can download it from the official website: [https://golang.org](https://golang.org)

1. Clone the repository:

```bash
git clone https://github.com/hussainalboori/morse-code-encoder-decoder.git
```

2. Navigate to the project directory:

```bash
cd morse-code-encoder-decoder
```

3. Build the executable:

```bash
go build
```

4. Run the program:

```bash
./morse-code-encoder-decoder
```

## Usage

1. Enter the line you want to translate when prompted:

```bash
Enter the line you want to translate:
```

2. The program will determine if the input is Morse code or regular text. If it is Morse code, it will decode it and display the corresponding text. If it is regular text, it will encode it into Morse code and display the result.

## Examples

### Example 1: Encoding Text to Morse Code

```
Enter the line you want to translate: Hello, World!
.... . .-.. .-.. --- --..-- / .-- --- .-. .-.. -.. -.-.-- 
```

### Example 2: Decoding Morse Code to Text

```
Enter the line you want to translate: .... . .-.. .-.. --- --..-- / .-- --- .-. .-.. -.. -.-.-- 
Hello, World!
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

Special thanks to [MSK17A](https://github.com/MSK17A/morse_code_decoder) for insprtion for the intial logic and and the idea.

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvement, please create a pull request or submit an issue in the [GitHub repository](https://github.com/hussainalboori/morse-code-encoder-decoder.git).