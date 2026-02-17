# Go-Reloaded Project Guide

## 📋 Project Overview
A text completion/editing/auto-correction tool that reads an input file, applies various text transformations, and writes the result to an output file.

---

## 🎯 Learning Objectives

By completing this project, you will learn:
1. **File I/O Operations**: Reading from and writing to files in Go
2. **String Manipulation**: Advanced string processing and transformation
3. **Number System Conversion**: Binary and hexadecimal to decimal conversion
4. **Pattern Matching**: Identifying and replacing patterns in text
5. **Text Parsing**: Breaking down text into words and processing them
6. **Data Structures**: Using slices and string builders efficiently
7. **Error Handling**: Proper error management in Go applications

---

## 📚 Prerequisites - Topics You Must Know

### 1. **Go Basics**
- Variables, constants, and data types
- Functions and parameters
- Control structures (if, for, switch)
- Arrays and slices

### 2. **String Operations**
- `strings` package functions:
  - `strings.Split()`, `strings.Join()`
  - `strings.ToUpper()`, `strings.ToLower()`, `strings.Title()`
  - `strings.HasPrefix()`, `strings.HasSuffix()`
  - `strings.TrimSpace()`, `strings.Contains()`

### 3. **File I/O**
- `os` package:
  - `os.ReadFile()` - Read entire file content
  - `os.WriteFile()` - Write content to file
  - `os.Args` - Command-line arguments
- Error handling with files

### 4. **Number Conversions**
- `strconv` package:
  - `strconv.ParseInt()` - Convert string to integer with different bases
  - `strconv.Itoa()` - Convert integer to string
- Understanding number bases (binary, hexadecimal, decimal)

### 5. **Regular Expressions (Optional but helpful)**
- `regexp` package basics for pattern matching

---

## 🛠️ Step-by-Step Implementation Guide

### **Phase 1: Project Setup** ✅

#### Step 1: Create Project Structure
```
go-reloaded/
├── main.go
├── go.mod
├── sample.txt (for testing)
└── result.txt (output)
```

#### Step 2: Initialize Go Module
```bash
go mod init go-reloaded
```

#### Step 3: Basic Main Function Setup
```go
package main

import (
    "fmt"
    "os"
)

func main() {
    // Check command-line arguments
    if len(os.Args) != 3 {
        fmt.Println("Usage: go run . <input_file> <output_file>")
        return
    }
    
    inputFile := os.Args[1]
    outputFile := os.Args[2]
    
    fmt.Printf("Input: %s, Output: %s\n", inputFile, outputFile)
}
```

**Test**: Run `go run . sample.txt result.txt` to verify argument handling.

---

### **Phase 2: File Reading** 📖

#### Step 4: Read Input File
Create a function to read the file content:

```go
func readFile(filename string) (string, error) {
    // Use os.ReadFile to read entire file
    // Return content as string and error
}
```

**Key Points**:
- Use `os.ReadFile()` which returns `[]byte`
- Convert `[]byte` to `string`
- Handle errors properly

**Test**: Print the file content to verify reading works.

---

### **Phase 3: Text Processing Core** 🔧

#### Step 5: Split Text into Words
Create a function that splits text into words:

```go
func splitIntoWords(text string) []string {
    // Split by spaces
    // Return slice of words
}
```

**Key Points**:
- Use `strings.Fields()` or `strings.Split()`
- Handle multiple spaces

**Manual Test**: Create test strings and print split results.

---

#### Step 6: Number Conversion Functions

Create three helper functions:

**A. Hexadecimal to Decimal**
```go
func hexToDecimal(hex string) string {
    // Use strconv.ParseInt with base 16
    // Convert result back to string
    // Handle errors
}
```

**B. Binary to Decimal**
```go
func binToDecimal(bin string) string {
    // Use strconv.ParseInt with base 2
    // Convert result back to string
    // Handle errors
}
```

**Manual Test Examples**:
- "1E" (hex) → "30"
- "10" (bin) → "2"
- "FF" (hex) → "255"
- "1010" (bin) → "10"

---

#### Step 7: Case Conversion Functions

**A. Uppercase Function**
```go
func toUpperCase(word string) string {
    // Use strings.ToUpper()
}
```

**B. Lowercase Function**
```go
func toLowerCase(word string) string {
    // Use strings.ToLower()
}
```

**C. Capitalize Function**
```go
func capitalize(word string) string {
    // Capitalize first letter, lowercase the rest
    // Handle empty strings
}
```

**Manual Test Examples**:
- "hello" → "HELLO" (upper)
- "WORLD" → "world" (lower)
- "brooklyn" → "Brooklyn" (cap)

---

### **Phase 4: Main Processing Logic** 🎯

#### Step 8: Process Modifiers
Create the main processing function:

```go
func processText(words []string) []string {
    result := []string{}
    
    for i := 0; i < len(words); i++ {
        word := words[i]
        
        // Check for modifiers: (hex), (bin), (up), (low), (cap)
        // Check for modifiers with numbers: (up, 2), (low, 3), etc.
        // Apply transformations to previous words
        
        result = append(result, word)
    }
    
    return result
}
```

**Implementation Strategy**:
1. Iterate through words
2. When you find a modifier (contains parentheses):
   - Check what type: (hex), (bin), (up), (low), (cap)
   - Check if it has a number: (up, 2)
   - Apply to previous word(s)
   - Remove the modifier from output
3. Otherwise, add word to result

**Key Logic**:
```
Current word: "(hex)"
Action: 
  - Go back to previous word in result
  - Convert it using hexToDecimal()
  - Replace previous word with converted value
  - Don't add "(hex)" to result
```

**Test Cases**:
- "42 (hex)" → "66"
- "go (up)" → "GO"
- "exciting (up, 2)" with "so exciting" → "SO EXCITING"

---

### **Phase 5: Punctuation Handling** ✨

#### Step 9: Fix Punctuation Spacing
Create a function to handle punctuation:

```go
func fixPunctuation(words []string) []string {
    // Handle: . , ! ? : ;
    // Remove space before punctuation
    // Add space after punctuation
}
```

**Algorithm**:
1. Iterate through words
2. If word is punctuation or starts with punctuation:
   - Attach to previous word (no space before)
   - Ensure space after (next word not punctuation)
3. Handle groups: "..." , "!?" , "!!"

**Test Cases**:
- "there ,and" → "there, and"
- "BAMM !!" → "BAMM!!"
- "thinking ..." → "thinking..."

---

#### Step 10: Handle Single Quotes
Create a function for quote handling:

```go
func fixQuotes(words []string) []string {
    // Find pairs of single quotes
    // Attach them to words inside
    // Remove spaces
}
```

**Algorithm**:
1. Find first `'`
2. Find matching closing `'`
3. Attach opening `'` to right of next word (no space)
4. Attach closing `'` to left of previous word (no space)
5. Remove standalone `'` marks

**Test Cases**:
- "' awesome '" → "'awesome'"
- "' I am happy '" → "'I am happy'"

---

### **Phase 6: A/An Correction** 📝

#### Step 11: Fix Article Usage
Create a function to correct "a" to "an":

```go
func fixArticles(words []string) []string {
    vowels := "aeiouAEIOUhH"
    
    // Iterate through words
    // If word is "a" or "A"
    // Check if next word starts with vowel or 'h'
    // Change to "an" or "An"
}
```

**Test Cases**:
- "a amazing" → "an amazing"
- "A apple" → "An apple"
- "a hour" → "an hour"
- "a banana" → "a banana" (no change)

---

### **Phase 7: Integration** 🔗

#### Step 12: Combine All Functions
Create the main processing pipeline:

```go
func processFullText(text string) string {
    // 1. Split into words
    words := splitIntoWords(text)
    
    // 2. Process modifiers (hex, bin, up, low, cap)
    words = processText(words)
    
    // 3. Fix punctuation
    words = fixPunctuation(words)
    
    // 4. Fix quotes
    words = fixQuotes(words)
    
    // 5. Fix articles (a/an)
    words = fixArticles(words)
    
    // 6. Join back into string
    return strings.Join(words, " ")
}
```

---

#### Step 13: Write Output File
Complete the main function:

```go
func writeFile(filename string, content string) error {
    // Use os.WriteFile
    // Return error if any
}

func main() {
    // Check arguments
    // Read input file
    content, err := readFile(inputFile)
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }
    
    // Process text
    result := processFullText(content)
    
    // Write output file
    err = writeFile(outputFile, result)
    if err != nil {
        fmt.Println("Error writing file:", err)
        return
    }
}
```

---

### **Phase 8: Testing** 🧪

#### Step 14: Create Test Cases
Create test files with different scenarios:

**Test 1: Basic Modifiers**
```
sample.txt: "hello (up) world (low) test (cap)"
Expected: "HELLO world Test"
```

**Test 2: Number Conversions**
```
sample.txt: "1E (hex) and 10 (bin)"
Expected: "30 and 2"
```

**Test 3: Multiple Word Modifiers**
```
sample.txt: "this is great (up, 2)"
Expected: "this IS GREAT"
```

**Test 4: Punctuation**
```
sample.txt: "Hello ,world !How are you ?"
Expected: "Hello, world! How are you?"
```

**Test 5: Quotes**
```
sample.txt: "He said : ' hello world '"
Expected: "He said: 'hello world'"
```

**Test 6: Articles**
```
sample.txt: "a apple a day keeps a doctor away"
Expected: "an apple a day keeps a doctor away"
```

**Test 7: Complex (Provided Example)**
```
it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) ...
```

---

### **Phase 9: Debugging Tips** 🐛

#### Common Issues and Solutions:

1. **Index Out of Range**
   - Always check if index exists before accessing
   - Use `if i > 0` before accessing `words[i-1]`

2. **Modifiers Not Applied**
   - Make sure to remove the modifier from output
   - Check that you're modifying the correct previous word(s)

3. **Punctuation Still Has Spaces**
   - Ensure you're joining words back correctly
   - Check that punctuation is attached to previous word

4. **Quotes Not Working**
   - Track quote state (opened/closed)
   - Handle multiple words between quotes

5. **Number Conversion Errors**
   - Validate input before conversion
   - Handle strconv.ParseInt errors

#### Debugging Strategy:
```go
// Print intermediate results
fmt.Println("After split:", words)
fmt.Println("After modifiers:", words)
fmt.Println("After punctuation:", words)
fmt.Println("After quotes:", words)
fmt.Println("Final:", result)
```

---

## 🎓 Advanced Challenges (Optional)

Once basic functionality works:

1. **Optimization**: Process text in single pass instead of multiple passes
2. **Edge Cases**: Handle empty files, special characters, unicode
3. **Performance**: Use `strings.Builder` instead of concatenation
4. **Validation**: Validate hex/bin numbers before conversion
5. **Testing**: Write unit tests for each function

---

## ✅ Checklist Before Submission

- [ ] Program accepts two command-line arguments
- [ ] Reads input file correctly
- [ ] Handles (hex) conversion
- [ ] Handles (bin) conversion
- [ ] Handles (up), (low), (cap) modifiers
- [ ] Handles modifiers with numbers: (up, 2)
- [ ] Fixes punctuation spacing correctly
- [ ] Handles punctuation groups (..., !?)
- [ ] Fixes single quotes correctly
- [ ] Corrects a/an usage
- [ ] Writes output file correctly
- [ ] Handles errors properly
- [ ] Code follows Go best practices
- [ ] All provided test cases pass
- [ ] Created own test cases

---

## 📖 Key Go Concepts Used

| Concept | Package/Function | Purpose |
|---------|-----------------|---------|
| File Reading | `os.ReadFile()` | Read entire file into memory |
| File Writing | `os.WriteFile()` | Write string to file |
| String Splitting | `strings.Fields()` or `strings.Split()` | Break text into words |
| String Joining | `strings.Join()` | Combine words back into text |
| Case Conversion | `strings.ToUpper/ToLower()` | Change text case |
| Number Parsing | `strconv.ParseInt()` | Convert string to number with base |
| Number Formatting | `strconv.Itoa()` | Convert number to string |
| Command-line Args | `os.Args` | Get input/output filenames |
| Error Handling | `if err != nil` | Handle file and conversion errors |

---

## 🚀 Pro Tips

1. **Start Small**: Implement one feature at a time and test it
2. **Use Helper Functions**: Break down complex logic into smaller functions
3. **Test Frequently**: Run tests after each feature
4. **Print Debug Info**: Use fmt.Println to see what's happening
5. **Handle Edge Cases**: Empty strings, first/last words, etc.
6. **Read Go Documentation**: Understand how each function works
7. **Ask Questions**: If stuck, review the concept, don't jump to AI

---

## 📚 Resources

- [Go by Example - Files](https://gobyexample.com/reading-files)
- [Go Documentation - strings](https://pkg.go.dev/strings)
- [Go Documentation - strconv](https://pkg.go.dev/strconv)
- [Go Documentation - os](https://pkg.go.dev/os)

---

**Good luck! Remember: The goal is to learn, not just to finish. Take your time and understand each step.** 🎯