# ASCII-Art Project Guide

> **Before you start:** Open each of the three banner files (`standard.txt`, `shadow.txt`, `thinkertoy.txt`) and study them carefully. The entire project depends on understanding the file format before writing any code.

---

## Objectives

By completing this project you will learn:

1. **File System API** — Reading files with Go's `os` package and processing their content
2. **Data Mapping** — Using a mathematical formula to locate data inside a structured file
3. **String Building** — Constructing multi-line output by combining pieces row by row
4. **Rune Handling** — Iterating over a string as Unicode code points, not bytes
5. **Input Parsing** — Splitting on a custom delimiter (`\n` as a two-character sequence)
6. **Edge Case Thinking** — Handling empty input, special characters, and unknown characters gracefully

---

## Prerequisites — Topics You Must Know Before Starting

### 1. Go Basics
- Functions, return values, error handling
- Slices — indexing and ranging
- `for` loops and `range`

### 2. Strings and Runes
- The difference between a `byte` and a `rune` in Go
- How to iterate over a string and get individual characters as runes
- `strings.Split` — splitting on a custom separator
- `strings.Builder` — efficient string construction in a loop

### 3. File Operations
- `os.ReadFile` — read an entire file at once
- Converting `[]byte` to `string`

### 4. ASCII
- What the ASCII table is
- What "printable ASCII characters" means (codes 32–126)
- The decimal value of the space character

**If any of these are unfamiliar, read about them before writing any code.**

- Search: **"ASCII table printable characters"**
- Search: **"golang rune vs byte"**
- https://pkg.go.dev/os
- https://pkg.go.dev/strings

---

## Project Structure

```
ascii-art/
├── main.go
├── standard.txt
├── shadow.txt
├── thinkertoy.txt
└── go.mod
```

---

## Milestone 1 — Understand the Banner File Format

**This milestone has no code.** Do not skip it.

Open `standard.txt` and answer every question below before writing anything.

**Questions to answer:**
- How many lines tall is each character's art?
- What separates one character from the next in the file?
- Which character appears first in the file? What is its ASCII value?
- Which character appears second? What is its ASCII value?
- How many total lines does one character occupy in the file (including the separator)?
- If the space character `' '` starts at line index 1, where does `'!'` start? Where does `'"'` start?
- Write the formula: given a character `c`, at what line index does its art begin?

**Verify:** You should be able to state the formula clearly before moving on:
```
startLine = (ASCII value of c  -  32)  *  9  +  1
```
If you reached a different formula, re-examine the file. Count carefully.

---

## Milestone 2 — Load the Banner File

**Goal:** Read a banner file and return its lines as a slice of strings.

**Questions to answer:**
- What does `os.ReadFile` return, and what do you need to do before splitting it into lines?
- Should you split on `"\n"` or `"\r\n"`? What happens on different operating systems?
- How many lines should `standard.txt` produce when split? Calculate it from the formula before running the code.

**Code Placeholder:**
```go
func loadBanner(filename string) ([]string, error) {
    // Read the file into memory

    // Convert to string

    // Split into lines on "\n"
    // Note: if lines contain "\r", trim it from each line

    // Return the lines slice
}
```

**Resources:**
- https://pkg.go.dev/os#ReadFile
- Search: **"golang strings TrimRight carriage return"**

**Verify:**
- Call `loadBanner("standard.txt")` and print the length of the result
- Calculate the expected length manually first — does it match?

---

## Milestone 3 — Extract One Character's Art

**Goal:** Given the full lines slice and a character, return the 8 lines of art for that character.

**Questions to answer:**
- Using your formula from Milestone 1, what is the start line index for `' '`? For `'A'`? For `'~'`?
- How many lines do you take starting from that index?

**Code Placeholder:**
```go
func getCharLines(lines []string, c rune) []string {
    // Calculate the start line index using the formula

    // Collect the next 8 lines starting from startLine

    // Return them as a slice
}
```

**Verify (add temporary code in main to test):**
- Extract `'H'` and print its 8 lines — should look like a large H
- Extract `' '` and print its 8 lines — should be 8 lines of spaces
- Extract `'!'` and print its 8 lines — should look like a large exclamation mark

---

## Milestone 4 — Render a Single Line of Text

**Goal:**
```
go run . "Hi"
```
Prints H and i side by side across 8 rows.

**Questions to answer:**
- Why can you not print one character at a time from top to bottom?
- What does the output look like if you loop over rows first, then characters?
- On each row, what exactly do you concatenate together?

**Code Placeholder:**
```go
func renderLine(banner []string, text string) {
    // Loop over rows 0 to 7:
    //   For each row:
    //     Loop over each character (rune) in text:
    //       Get that character's art lines
    //       Append the current row's line to a result string
    //     Print the result string for this row
}
```

**Resources:**
- Search: **"golang range string rune"**
- Search: **"golang strings Builder"** — consider using it instead of `+=`

**Verify:**
- `go run . "Hi"` renders H and i side by side
- `go run . "Hello"` matches the spec output exactly
- `go run . ""` prints nothing and does not crash

---

## Milestone 5 — Handle `\n` in the Input

**Goal:**
```
go run . "Hello\nThere"    → Hello rendered, then There below it
go run . "Hello\n\nThere"  → Hello, a blank line, then There
go run . "\n"              → one blank line
```

**Questions to answer:**
- When the user types `"Hello\nThere"` in the shell, what does your program actually receive — a real newline or two characters `\` and `n`?
- What does splitting on `"\\n"` give you for the input `"Hello\nThere"`?
- What does splitting on `"\\n"` give you for `"Hello\n\nThere"`? How many parts?
- What should an empty part in the result render as?

**Code Placeholder:**
```go
func render(banner []string, input string) {
    // Split input on the two-character sequence "\n" (backslash + n)
    // This is NOT the same as splitting on a real newline

    // For each part:
    //   If the part is empty: print one blank line
    //   Otherwise: call renderLine for that part
}
```

**Verify using `cat -e`** (the `$` marks line endings — every line must end exactly there):
```bash
go run . "" | cat -e
go run . "\n" | cat -e
go run . "Hello\nThere" | cat -e
go run . "Hello\n\nThere" | cat -e
```
Compare your output character by character against the spec examples.

---

## Milestone 6 — Final main.go

**Goal:** Connect everything. Handle the argument, load the banner, render and print.

**Questions to answer:**
- How many arguments does your program expect?
- What should happen if the banner file cannot be found?

**Code Placeholder:**
```go
func main() {
    // 1. Check that exactly 1 argument was provided
    //    If not, print usage and return

    // 2. Load the banner file (use "standard.txt" for now)
    //    Handle the error

    // 3. Call render with the banner and the input argument
}
```

**Verify:** Run every example from the spec and confirm each matches exactly:
```bash
go run . "" | cat -e
go run . "\n" | cat -e
go run . "Hello\n" | cat -e
go run . "hello" | cat -e
go run . "HeLlO" | cat -e
go run . "Hello There" | cat -e
go run . "1Hello 2There" | cat -e
go run . "{Hello There}" | cat -e
go run . "Hello\nThere" | cat -e
go run . "Hello\n\nThere" | cat -e
```

---

## Milestone 7 — Edge Cases

Test each of these and make sure your program handles them without crashing:

- `go run . "!@#$%"` — special characters are valid printable ASCII
- `go run . "123"` — numbers work automatically if your formula is correct
- `go run . "\n\n\n"` — three blank lines
- `go run . "Hello World"` — space character must render as 8 lines of spaces

**Question for each one:** Before running it, predict the output. If reality differs from your prediction, find out why.

---

## Milestone 8 — Unit Tests

Write at least these:

**Code Placeholder:**
```go
// main_test.go

func TestLoadBanner(t *testing.T) {
    // Test that the file loads without error
    // Test that len(lines) matches your expected value
}

func TestGetCharLines(t *testing.T) {
    // Test that ' ' returns exactly 8 lines
    // Test that 'A' returns exactly 8 lines
    // Test that the start line index for '!' is correct (calculate it manually first)
}

func TestRenderEmpty(t *testing.T) {
    // Test that rendering an empty string produces no output (or 8 empty lines?)
    // Check the spec to determine the expected behavior
}
```

**Resource:** https://go.dev/doc/tutorial/add-a-test

---

## Debugging Checklist

Before asking for help, go through this:

- Does the output look visually correct but `cat -e` shows extra spaces? The banner file lines may have trailing `\r`. Trim them when loading.
- Are characters appearing at the wrong position? Print the start line index for `' '`, `'!'`, and `'"'` — they should be 1, 10, and 19. If they are not, your formula is wrong.
- Does the program panic with "index out of range"? A character in your input may be outside the range 32–126. Check that every rune is a valid printable ASCII character before calling `getCharLines`.
- Is `\n` in the input not creating separate rendered lines? Make sure you split on the two-character string `"\\n"`, not on `"\n"`.

---

## Key Packages

| Package | What You Use It For | Docs |
|---|---|---|
| `os` | Read the banner file, read args | https://pkg.go.dev/os |
| `strings` | Split on `\n`, TrimRight, Builder | https://pkg.go.dev/strings |
| `fmt` | Print rendered rows | https://pkg.go.dev/fmt |

---

## Submission Checklist

- [ ] `go run . ""` outputs nothing without crashing
- [ ] `go run . "\n"` outputs one blank line
- [ ] `go run . "Hello"` matches spec output exactly (verified with `cat -e`)
- [ ] `go run . "Hello\nThere"` renders two blocks of ASCII art
- [ ] `go run . "Hello\n\nThere"` has a blank line between the two blocks
- [ ] Numbers, spaces, and special characters render correctly
- [ ] `{`, `}` and other bracket characters render correctly
- [ ] No trailing spaces on any output line (verified with `cat -e`)
- [ ] No crash on any valid printable ASCII input
- [ ] Missing banner file produces a meaningful error message
- [ ] Unit tests written and passing
- [ ] Only standard Go packages used
