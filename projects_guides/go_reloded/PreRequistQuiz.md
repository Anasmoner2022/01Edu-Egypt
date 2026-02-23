# 🎯 Go-Reloaded Prerequisites Quiz
## strings · regexp · File I/O · strconv · Text Processing

**Time Limit:** 55 minutes  
**Total Questions:** 32  
**Passing Score:** 26/32 (81%)

> Questions are tagged: 🟢 Easy · 🟡 Medium · 🔴 Hard  
> All topics are general — no specific project knowledge required.

---

## 📋 SECTION 1: THE strings PACKAGE (10 Questions)

### Q1 🟢 — What does `strings.ToUpper("hello")` return?

**A)** `"Hello"`  
**B)** `"HELLO"`  
**C)** `"hello"` — strings are immutable, so nothing changes  
**D)** An error  

<details><summary>💡 Answer</summary>

**B) `"HELLO"`**

```go
fmt.Println(strings.ToUpper("hello"))   // HELLO
fmt.Println(strings.ToLower("WORLD"))   // world
fmt.Println(strings.Title("hello world")) // Hello World (deprecated — use golang.org/x/text/cases)
```

String functions in Go always return a NEW string — the original is unchanged. Go strings are immutable; you can never modify them in place.

</details>

---

### Q2 🟢 — What does `strings.TrimSpace("  hello  ")` return?

**A)** `"hello"` — removes all leading and trailing whitespace  
**B)** `"  hello"` — only trims the right  
**C)** `"hello  "` — only trims the left  
**D)** `"hello"` with a single space on each side  

<details><summary>💡 Answer</summary>

**A) `"hello"` — removes all leading and trailing whitespace**

```go
strings.TrimSpace("  hello  ")   // "hello"
strings.TrimSpace("\t\nhello\n") // "hello"
strings.Trim("***hello***", "*") // "hello" — trims specific chars
strings.TrimLeft("  hi", " ")   // "hi" — left only
strings.TrimRight("hi  ", " ")  // "hi" — right only
```

`TrimSpace` is the go-to for cleaning user input and file lines. It removes spaces, tabs, newlines, and carriage returns from both ends.

</details>

---

### Q3 🟢 — What does `strings.Contains("hello world", "world")` return?

**A)** The index where "world" starts  
**B)** `true` — because "world" is a substring of "hello world"  
**C)** The number of times "world" appears  
**D)** `"world"`  

<details><summary>💡 Answer</summary>

**B) `true`**

```go
strings.Contains("hello world", "world")  // true
strings.Contains("hello world", "earth")  // false
strings.HasPrefix("hello", "he")          // true
strings.HasSuffix("hello", "lo")          // true
strings.Index("hello", "ll")              // 2 (first occurrence index)
strings.Count("hello", "l")              // 2 (number of occurrences)
```

These are the core string search functions. Use `Contains` for "does it exist?", `Index` for "where is it?", `Count` for "how many?".

</details>

---

### Q4 🟢 — What does `strings.Split("a,b,c", ",")` return?

**A)** `"a b c"`  
**B)** `["a", "b", "c"]` — a `[]string` with 3 elements  
**C)** `"a,b,c"` — commas are left in  
**D)** An error  

<details><summary>💡 Answer</summary>

**B) `[]string{"a", "b", "c"}`**

```go
parts := strings.Split("a,b,c", ",")
// ["a", "b", "c"]

strings.Split("hello", "")
// ["h", "e", "l", "l", "o"] — split on empty string = individual chars

strings.SplitN("a,b,c,d", ",", 2)
// ["a", "b,c,d"] — limit to 2 parts

strings.Join([]string{"a", "b", "c"}, "-")
// "a-b-c" — inverse of Split
```

`Split` and `Join` are complementary — one breaks a string into a slice, the other reassembles it.

</details>

---

### Q5 🟡 — What does `strings.Replace("aabbcc", "b", "X", 1)` return?

**A)** `"aaXXcc"` — replaces all occurrences  
**B)** `"aaXbcc"` — replaces only the first occurrence (limit = 1)  
**C)** `"aabbcc"` — nothing changes  
**D)** An error  

<details><summary>💡 Answer</summary>

**B) `"aaXbcc"` — limit of 1 means only the first**

```go
strings.Replace("aabbcc", "b", "X", 1)   // "aaXbcc" — first only
strings.Replace("aabbcc", "b", "X", 2)   // "aaXXcc" — first two
strings.Replace("aabbcc", "b", "X", -1)  // "aaXXcc" — all (use -1 for all)
strings.ReplaceAll("aabbcc", "b", "X")   // "aaXXcc" — same as -1, cleaner
```

The last argument is the count: `n > 0` replaces first n occurrences, `-1` replaces all. `ReplaceAll` is `Replace` with `-1` — prefer it when you always want all replacements.

</details>

---

### Q6 🟡 — What does `strings.Fields("  hello   world  ")` return?

**A)** `["hello", "world"]` — splits on whitespace and removes empty entries  
**B)** `["", "", "hello", "", "", "world", "", ""]`  
**C)** `["hello   world"]` — only strips the outer whitespace  
**D)** Same as `strings.Split("  hello   world  ", " ")`  

<details><summary>💡 Answer</summary>

**A) `[]string{"hello", "world"}` — splits on any whitespace, skips empty**

```go
strings.Fields("  hello   world  ")
// ["hello", "world"]

// Compare with Split:
strings.Split("  hello   world  ", " ")
// ["", "", "hello", "", "", "world", "", ""]  ← lots of empty strings

strings.Fields("one\ttwo\nthree")
// ["one", "two", "three"]  ← handles tabs and newlines too
```

`Fields` is the right tool when you want to tokenize human-readable text with variable whitespace. `Split` with `" "` produces empty strings for consecutive delimiters.

</details>

---

### Q7 🟡 — What does `strings.Builder` do and why use it instead of `+` concatenation?

**A)** It validates string contents  
**B)** It's a mutable byte buffer optimized for building strings piece-by-piece — avoids the O(n²) allocation cost of repeated `+` concatenation  
**C)** It's the same as `fmt.Sprintf`  
**D)** It encrypts the string  

<details><summary>💡 Answer</summary>

**B) A mutable buffer for efficient string building**

```go
// BAD — each + creates a new string allocation:
result := ""
for i := 0; i < 1000; i++ {
    result += fmt.Sprintf("item%d,", i)  // 1000 allocations
}

// GOOD — one buffer, one final string:
var b strings.Builder
for i := 0; i < 1000; i++ {
    fmt.Fprintf(&b, "item%d,", i)
}
result := b.String()
```

For small concatenations (2–3 strings), `+` is fine. For loops or many pieces, use `strings.Builder` or `strings.Join`. The performance difference is significant for large data.

</details>

---

### Q8 🟡 — What does `strings.Repeat("ab", 3)` return?

**A)** `"ab3"`  
**B)** `"ababab"`  
**C)** `"ab ab ab"`  
**D)** An error  

<details><summary>💡 Answer</summary>

**B) `"ababab"`**

```go
strings.Repeat("ab", 3)   // "ababab"
strings.Repeat("-", 20)   // "--------------------"
strings.Repeat(".", 0)    // "" — zero repetitions
```

Useful for generating separators, padding, or test data. `n` must be non-negative — negative n panics.

</details>

---

### Q9 🔴 — What is the output?

```go
s := "hello"
s = strings.Replace(s, "l", "r", -1)
t := strings.ToUpper(s[1:3])
fmt.Println(s, t)
```

**A)** `"herro ER"`  
**B)** `"hello EL"`  
**C)** Compile error  
**D)** `"herro RR"`  

<details><summary>💡 Answer</summary>

**A) `"herro ER"`**

Step by step:
1. `strings.Replace("hello", "l", "r", -1)` → `"herro"` (both `l`s replaced)
2. `s[1:3]` = bytes at index 1 and 2 = `"er"`
3. `strings.ToUpper("er")` = `"ER"`
4. Output: `herro ER`

String slicing `s[i:j]` gives bytes from index `i` to `j-1`. This works correctly for ASCII. For multi-byte Unicode, use `[]rune(s)` to slice by character.

</details>

---

### Q10 🔴 — What is the difference between `strings.EqualFold("Go", "go")` and `strings.ToLower("Go") == strings.ToLower("go")`?

**A)** `EqualFold` is wrong — always use `ToLower` for comparison  
**B)** Both produce the same result (`true`), but `EqualFold` is more efficient — it compares directly without allocating new strings  
**C)** `EqualFold` is case-sensitive; `ToLower` comparison is not  
**D)** `EqualFold` only works for ASCII; `ToLower` handles Unicode  

<details><summary>💡 Answer</summary>

**B) Same result, but `EqualFold` avoids string allocation**

```go
strings.EqualFold("Go", "GO")    // true — case-insensitive, no allocation
strings.EqualFold("café", "CAFÉ") // true — Unicode-aware
strings.ToLower("Go") == strings.ToLower("GO") // also true, but creates 2 new strings
```

`EqualFold` is the idiomatic Go way to do case-insensitive string comparison. It's faster and uses less memory than the `ToLower`/`ToUpper` approach. It's also correctly Unicode-aware.

</details>

---

## 📋 SECTION 2: REGULAR EXPRESSIONS (8 Questions)

### Q11 🟢 — What does `regexp.MustCompile(pattern)` do differently from `regexp.Compile(pattern)`?

**A)** `MustCompile` is faster  
**B)** `MustCompile` panics if the pattern is invalid; `Compile` returns an error — use `MustCompile` for patterns known at compile time  
**C)** `MustCompile` matches more patterns  
**D)** They are identical  

<details><summary>💡 Answer</summary>

**B) `MustCompile` panics on invalid pattern; `Compile` returns `(Regexp, error)`**

```go
// For package-level (startup) compilation — panic is appropriate:
var re = regexp.MustCompile(`\d+`)

// For runtime patterns (user input) — always check the error:
re, err := regexp.Compile(userInput)
if err != nil {
    return fmt.Errorf("invalid pattern: %w", err)
}
```

Never use `MustCompile` with user-supplied patterns — a bad pattern panics your program. Use it only for literal patterns you control and that you've already verified are valid.

</details>

---

### Q12 🟢 — What does `re.MatchString(s)` return?

**A)** The matched substring  
**B)** `true` if the pattern matches anywhere in `s`, `false` otherwise — also returns an error  
**C)** All matches as a `[]string`  
**D)** The number of matches  

<details><summary>💡 Answer</summary>

**B) `(bool, error)` — true if pattern matches anywhere in the string**

```go
re := regexp.MustCompile(`\d+`)
matched, _ := re.MatchString("hello 42 world") // true
matched, _ = re.MatchString("no digits here")  // false

// Shorthand for simple checks:
ok, _ := regexp.MatchString(`^\d+$`, "12345") // true — entire string is digits
```

`MatchString` checks if the pattern appears *anywhere* in the string. To match the whole string, anchor with `^` (start) and `$` (end).

</details>

---

### Q13 🟡 — What is the difference between `FindString` and `FindAllString`?

**A)** `FindString` is case-sensitive; `FindAllString` is not  
**B)** `FindString` returns the first (leftmost) match; `FindAllString` returns all non-overlapping matches as a `[]string`  
**C)** `FindAllString` is deprecated  
**D)** `FindString` uses the whole string; `FindAllString` searches word-by-word  

<details><summary>💡 Answer</summary>

**B) `FindString` = first match; `FindAllString` = all matches**

```go
re := regexp.MustCompile(`\d+`)

re.FindString("cat 3 dog 17 bird 5")
// "3" — only the first match

re.FindAllString("cat 3 dog 17 bird 5", -1)
// ["3", "17", "5"] — all matches (-1 means no limit)

re.FindAllString("cat 3 dog 17", 2)
// ["3", "17"] — at most 2 matches
```

The second argument to `FindAll*` is the maximum count: `-1` means all, `n > 0` means at most n matches.

</details>

---

### Q14 🟡 — What do capturing groups `()` return in `FindAllStringSubmatch`?

**A)** They are ignored — same result as without groups  
**B)** Each match returns a `[]string` where `[0]` is the full match and `[1]`, `[2]`... are the capture groups  
**C)** Compile error  
**D)** Only the captured groups, not the full match  

<details><summary>💡 Answer</summary>

**B) `[0]` = full match, `[1]`, `[2]`... = capture groups**

```go
re := regexp.MustCompile(`(\w+)\s*=\s*(\w+)`)
matches := re.FindAllStringSubmatch("x = 10, y = 20", -1)

// matches[0] = ["x = 10", "x", "10"]
// matches[1] = ["y = 20", "y", "20"]

for _, m := range matches {
    key := m[1]   // "x", then "y"
    val := m[2]   // "10", then "20"
    fmt.Printf("%s → %s\n", key, val)
}
```

Capture groups are the main reason to use regex over simple string functions. They extract structured data from patterned text in one pass.

</details>

---

### Q15 🟡 — What does `re.ReplaceAllString(src, repl)` do?

**A)** Replaces only the first match  
**B)** Replaces every non-overlapping match of `re` in `src` with `repl`, returning a new string  
**C)** Modifies `src` in place  
**D)** Returns a `[]byte`  

<details><summary>💡 Answer</summary>

**B) Replaces all matches, returns a new string**

```go
re := regexp.MustCompile(`\d+`)
result := re.ReplaceAllString("cat3dog17", "NUM")
// "catNUMdogNUM"

// Use $1, $2 to reference capture groups in replacement:
re2 := regexp.MustCompile(`(\w+)\s(\w+)`)
result2 := re2.ReplaceAllString("John Smith", "$2, $1")
// "Smith, John"
```

`$1`, `$2` in the replacement string refer to capture group contents. This makes regex replacement extremely powerful for reformatting text.

</details>

---

### Q16 🟡 — What does `re.ReplaceAllStringFunc(src, func(string) string)` allow that `ReplaceAllString` cannot do?

**A)** Nothing extra  
**B)** The replacement is computed by calling a function on each match — enables dynamic replacements based on the matched text  
**C)** It works on `[]byte` instead of strings  
**D)** It handles overlapping matches  

<details><summary>💡 Answer</summary>

**B) Dynamic replacement via a function called on each match**

```go
re := regexp.MustCompile(`\d+`)

// Double every number found:
result := re.ReplaceAllStringFunc("cat 3 and 17", func(match string) string {
    n, _ := strconv.Atoi(match)
    return strconv.Itoa(n * 2)
})
// "cat 6 and 34"

// Convert hex to decimal:
hexRe := regexp.MustCompile(`0x[0-9a-fA-F]+`)
result2 := hexRe.ReplaceAllStringFunc("value is 0xFF", func(m string) string {
    n, _ := strconv.ParseInt(m[2:], 16, 64)
    return strconv.FormatInt(n, 10)
})
// "value is 255"
```

This is one of the most powerful regex features — transforms matches using arbitrary Go logic.

</details>

---

### Q17 🔴 — What is the output?

```go
re := regexp.MustCompile(`(ha)+`)
fmt.Println(re.FindString("hahaha"))
fmt.Println(re.FindAllString("ha haha hahaha", -1))
```

**A)** `"ha"` and `["ha", "ha", "ha"]`  
**B)** `"hahaha"` and `["ha", "haha", "hahaha"]`  
**C)** `"hahaha"` and `["ha", "haha", "hahaha"]`  
**D)** `"hahaha"` and `["ha", "haha", "hahaha"]` — but `FindAllString` finds only the longest  

<details><summary>💡 Answer</summary>

**B) `"hahaha"` and `["ha", "haha", "hahaha"]`**

`(ha)+` matches one or more consecutive `"ha"` sequences. It's greedy — always matches as much as possible.
- In `"hahaha"`: greedily matches the full `"hahaha"`
- In `"ha haha hahaha"`: three separate words, each matched in full: `"ha"`, `"haha"`, `"hahaha"`

Regex is greedy by default — quantifiers (`+`, `*`, `?`) consume as many characters as possible. Add `?` after the quantifier for lazy (minimal) matching: `(ha)+?`.

</details>

---

### Q18 🔴 — What do the anchors `^` and `$` mean in a regex?

**A)** `^` means "not", `$` means end of file  
**B)** `^` matches the start of the string (or line in multiline mode); `$` matches the end — together they assert the whole string matches the pattern  
**C)** `^` matches uppercase; `$` is a variable reference  
**D)** They are invalid in Go regex  

<details><summary>💡 Answer</summary>

**B) `^` = start of string, `$` = end of string**

```go
re := regexp.MustCompile(`^\d+$`)
re.MatchString("12345")    // true  — all digits, nothing else
re.MatchString("123abc")   // false — has non-digit after digits
re.MatchString("  123  ")  // false — has spaces

// Without anchors:
re2 := regexp.MustCompile(`\d+`)
re2.MatchString("abc123def")  // true — digits found anywhere
```

Without anchors, a regex matches if the pattern appears *anywhere* in the string. With `^` and `$`, the entire string must match. This is critical for input validation.

</details>

---

## 📋 SECTION 3: FILE I/O (7 Questions)

### Q19 🟢 — What is the simplest way to read an entire file into memory as a `[]byte`?

**A)** `io.Read("file.txt")`  
**B)** `os.ReadFile("file.txt")` — returns `([]byte, error)`  
**C)** `file.ReadAll()`  
**D)** `bufio.ReadFile("file.txt")`  

<details><summary>💡 Answer</summary>

**B) `os.ReadFile("file.txt")` — returns `([]byte, error)`**

```go
data, err := os.ReadFile("input.txt")
if err != nil {
    return fmt.Errorf("reading file: %w", err)
}
content := string(data)  // convert to string for text processing
```

`os.ReadFile` (added in Go 1.16) reads the whole file in one call. Use it when the file fits in memory. For very large files, use `bufio.Scanner` to read line by line instead.

</details>

---

### Q20 🟢 — What is the simplest way to write a string to a file, creating it if it doesn't exist?

**A)** `io.WriteFile("out.txt", data)`  
**B)** `os.WriteFile("out.txt", []byte(content), 0644)` — creates or overwrites  
**C)** `file.Write("out.txt", content)`  
**D)** `os.Create("out.txt"); file.WriteString(content)`  

<details><summary>💡 Answer</summary>

**B) `os.WriteFile("out.txt", []byte(content), 0644)`**

```go
content := "processed output\n"
err := os.WriteFile("output.txt", []byte(content), 0644)
if err != nil {
    return fmt.Errorf("writing file: %w", err)
}
```

`0644` is the Unix file permission: owner can read+write, group and others can only read. `os.WriteFile` creates the file if it doesn't exist and truncates (overwrites) it if it does. For appending, open with `os.OpenFile` and `os.O_APPEND`.

</details>

---

### Q21 🟡 — How do you read a text file line by line without loading the whole file into memory?

**A)** `strings.Split(os.ReadFile(name), "\n")`  
**B)** Use `bufio.Scanner`:
```go
f, _ := os.Open("file.txt")
scanner := bufio.NewScanner(f)
for scanner.Scan() {
    line := scanner.Text()
}
```
**C)** `io.ReadLines("file.txt")`  
**D)** `os.ReadLines("file.txt")`  

<details><summary>💡 Answer</summary>

**B) `bufio.Scanner` with `Scan()` + `Text()` loop**

```go
f, err := os.Open("file.txt")
if err != nil { return err }
defer f.Close()

scanner := bufio.NewScanner(f)
for scanner.Scan() {
    line := scanner.Text()  // line without the trailing newline
    // process line
}
if err := scanner.Err(); err != nil {
    return err  // always check scanner errors after the loop
}
```

`scanner.Text()` returns the current line without the `\n`. `scanner.Scan()` returns `false` at EOF. Always check `scanner.Err()` after the loop — a `false` from `Scan()` could mean EOF or an error.

</details>

---

### Q22 🟡 — What is the correct way to handle command-line arguments in Go?

**A)** `args := os.Argv`  
**B)** `args := os.Args` — a `[]string` where `[0]` is the program name and `[1:]` are the user arguments  
**C)** `args := flag.Args()`  
**D)** `func main(args []string)`  

<details><summary>💡 Answer</summary>

**B) `os.Args` — `[0]` is the binary name, `[1:]` are user arguments**

```go
// go run . input.txt output.txt
// os.Args = ["./program", "input.txt", "output.txt"]

func main() {
    if len(os.Args) != 3 {
        fmt.Fprintln(os.Stderr, "Usage: program <input> <output>")
        os.Exit(1)
    }
    inputFile := os.Args[1]   // "input.txt"
    outputFile := os.Args[2]  // "output.txt"
}
```

Always validate `len(os.Args)` before indexing — accessing `os.Args[1]` without checking panics if no argument was given.

</details>

---

### Q23 🟡 — What does `defer f.Close()` do and why place it immediately after opening a file?

**A)** Closes the file immediately  
**B)** Schedules the file to be closed when the surrounding function returns — placing it right after open guarantees cleanup even if the function returns early or panics  
**C)** Flushes the file buffer  
**D)** Makes the file read-only  

<details><summary>💡 Answer</summary>

**B) Deferred close — guaranteed cleanup when the function exits**

```go
f, err := os.Open("file.txt")
if err != nil { return err }
defer f.Close()  // guaranteed to run when function returns — even on early returns

// ... rest of the function — no need to remember to close
```

Without `defer`, every early return path needs a manual `f.Close()` — easy to miss. With `defer`, one line covers all exit paths. Not closing files leaks file descriptors — on Linux, the default limit is 1024 open files per process.

</details>

---

### Q24 🔴 — What is the difference between `os.Open` and `os.OpenFile`?

**A)** `os.Open` is for reading; `os.OpenFile` gives control over flags (read, write, append, create) and permissions  
**B)** They are identical  
**C)** `os.OpenFile` is deprecated  
**D)** `os.Open` creates the file if it doesn't exist; `os.OpenFile` doesn't  

<details><summary>💡 Answer</summary>

**A) `os.Open` = read-only shortcut; `os.OpenFile` = full control**

```go
// os.Open — read-only (O_RDONLY):
f, err := os.Open("file.txt")

// os.OpenFile — full control:
f, err := os.OpenFile("file.txt",
    os.O_RDWR|os.O_CREATE|os.O_APPEND,  // flags
    0644,                                 // permissions
)

// Common flag combinations:
// os.O_RDONLY              — read only
// os.O_WRONLY|os.O_CREATE|os.O_TRUNC  — write, create, overwrite (like os.Create)
// os.O_WRONLY|os.O_CREATE|os.O_APPEND — append to file
// os.O_RDWR                           — read and write
```

Use `os.Open` for reading, `os.Create` for creating/overwriting, and `os.OpenFile` when you need append or specific combinations.

</details>

---

### Q25 🔴 — What happens if `bufio.Scanner` is used to read a line longer than its buffer (default 64KB)?

**A)** It reads the line in chunks automatically  
**B)** `scanner.Scan()` returns `false` and `scanner.Err()` returns `bufio.ErrTooLong`  
**C)** It panics  
**D)** The line is silently truncated  

<details><summary>💡 Answer</summary>

**B) `scanner.Scan()` returns `false`, error is `bufio.ErrTooLong`**

```go
scanner := bufio.NewScanner(f)

// Fix: increase buffer for files with very long lines:
buf := make([]byte, 1024*1024)           // 1MB buffer
scanner.Buffer(buf, len(buf))            // max token size = 1MB

for scanner.Scan() {
    // process line
}
if err := scanner.Err(); err != nil {
    log.Fatal(err)  // catches ErrTooLong and other errors
}
```

This is a gotcha for text processing — if input files have very long lines (e.g. minified JSON or SVG), the default scanner will fail. Always check `scanner.Err()` after the loop.

</details>

---

## 📋 SECTION 4: strconv — TYPE CONVERSION (7 Questions)

### Q26 🟢 — How do you convert the string `"42"` to the integer `42`?

**A)** `int("42")`  
**B)** `strconv.Atoi("42")` — returns `(int, error)`  
**C)** `fmt.Sscanf("42", "%d")`  
**D)** `(int)("42")`  

<details><summary>💡 Answer</summary>

**B) `strconv.Atoi("42")` — returns `(int, error)`**

```go
n, err := strconv.Atoi("42")
if err != nil {
    // "42" was not a valid integer
}
// n == 42

// For int64:
n64, err := strconv.ParseInt("42", 10, 64)

// For float:
f, err := strconv.ParseFloat("3.14", 64)
```

Always check the error — if the string is not a valid integer, the error is non-nil. `Atoi` is equivalent to `ParseInt(s, 10, 0)` — base 10, platform int size.

</details>

---

### Q27 🟢 — How do you convert the integer `42` to the string `"42"`?

**A)** `string(42)` — convert with type cast  
**B)** `strconv.Itoa(42)`  
**C)** `fmt.Sprint(42)`  
**D)** Both B and C work; B is more efficient  

<details><summary>💡 Answer</summary>

**D) Both work — `strconv.Itoa` is more efficient**

```go
strconv.Itoa(42)      // "42" — fast, no allocation beyond the string
fmt.Sprintf("%d", 42) // "42" — works but slower (uses reflection internally)
fmt.Sprint(42)        // "42" — also works

// WRONG — do NOT use:
string(42)   // this gives the Unicode character with code point 42 ("*"), NOT "42"
```

`string(42)` is the classic trap — it creates a single-character string with Unicode code point 42, not the string `"42"`. Always use `strconv.Itoa` or `fmt.Sprintf`.

</details>

---

### Q28 🟡 — What does `strconv.ParseInt("FF", 16, 64)` return?

**A)** An error — "FF" is not a number  
**B)** `255` — parses "FF" as base-16 (hexadecimal)  
**C)** `0xFF`  
**D)** `"FF"` converted to ASCII codes  

<details><summary>💡 Answer</summary>

**B) `255` — hex FF = decimal 255**

```go
n, err := strconv.ParseInt("FF", 16, 64)  // 255, nil
n, err = strconv.ParseInt("10", 2, 64)   // 2 (binary 10)
n, err = strconv.ParseInt("10", 8, 64)   // 8 (octal 10)
n, err = strconv.ParseInt("10", 10, 64)  // 10 (decimal)

// Arguments: (string, base, bitSize)
// base: 2, 8, 10, 16 (or 0 to detect from prefix: "0x", "0", "0b")
// bitSize: 8, 16, 32, 64 — used for overflow detection
```

The third argument (bitSize) controls overflow checking, not the output type. The return type is always `int64` — cast to smaller types if needed.

</details>

---

### Q29 🟡 — `strconv.Atoi` returns `(int, error)`. What error does it return for `strconv.Atoi("abc")`?

**A)** `nil` — it silently returns 0  
**B)** A `*strconv.NumError` with `.Err == strconv.ErrSyntax`  
**C)** A generic `errors.New("invalid")` error  
**D)** A panic  

<details><summary>💡 Answer</summary>

**B) `*strconv.NumError` with `ErrSyntax`**

```go
n, err := strconv.Atoi("abc")
if err != nil {
    numErr := err.(*strconv.NumError)
    fmt.Println(numErr.Func)  // "Atoi"
    fmt.Println(numErr.Num)   // "abc"
    fmt.Println(numErr.Err)   // strconv.ErrSyntax or strconv.ErrRange
}
// n == 0 when err != nil

// ErrSyntax  — input is not a valid number
// ErrRange   — number is valid but too large for the type
```

In practice, just check `if err != nil` and handle it — you rarely need to inspect the specific error type unless you want to distinguish syntax vs range errors.

</details>

---

### Q30 🟡 — What does `strconv.Quote("hello\nworld")` return?

**A)** `hello\nworld` — removes the escape sequence  
**B)** `"hello\nworld"` — a Go string literal with the newline as `\n` and surrounding quotes  
**C)** An error  
**D)** `hello world`  

<details><summary>💡 Answer</summary>

**B) `"hello\nworld"` — produces a valid Go string literal**

```go
strconv.Quote("hello\nworld")  // `"hello\nworld"` (includes the quotes!)
strconv.Quote(`tab	here`)     // `"tab\there"`
strconv.Unquote(`"hello"`)     // "hello", nil

// Useful for:
// - Debugging: printing strings with invisible characters visible
// - Code generation: producing valid Go string literals
fmt.Println(strconv.Quote("line1\nline2"))
// Output: "line1\nline2"   ← the \n is printed as backslash-n, not a newline
```

`Quote` is invaluable for debugging text processing — it makes invisible characters (newlines, tabs, carriage returns) visible.

</details>

---

### Q31 🔴 — What is the output?

```go
s := "123abc"
n, err := strconv.Atoi(s)
fmt.Println(n, err != nil)

s2 := "99999999999999999999"
n2, err2 := strconv.Atoi(s2)
fmt.Println(n2, err2 != nil)
```

**A)** `0 true` then `0 true`  
**B)** `123 false` then `99999999999999999999 false`  
**C)** Panic on first call  
**D)** `0 true` then some very large number `false`  

<details><summary>💡 Answer</summary>

**A) `0 true` then `0 true`**

- `"123abc"` is not a valid integer (mixed content) → `ErrSyntax`, returns `0`
- `"99999999999999999999"` is too large for `int` (max ~9.2×10¹⁸ on 64-bit) → `ErrRange`, returns the max or min `int`

Actually for ErrRange, `Atoi` returns `math.MaxInt` (not 0), so: `9223372036854775807 true` for the second. The key point: **both return non-nil errors**. Always check the error before using the returned value.

</details>

---

### Q32 🔴 — You need to replace every number in a string with its doubled value. Which approach works correctly?

**A)**
```go
strings.ReplaceAll(s, `\d+`, func(m string) string {
    n, _ := strconv.Atoi(m); return strconv.Itoa(n * 2)
})
```
**B)**
```go
re := regexp.MustCompile(`\d+`)
re.ReplaceAllStringFunc(s, func(m string) string {
    n, _ := strconv.Atoi(m); return strconv.Itoa(n * 2)
})
```
**C)**
```go
for _, c := range s {
    if c >= '0' && c <= '9' { /* double it */ }
}
```
**D)** Both A and B work  

<details><summary>💡 Answer</summary>

**B) `regexp.ReplaceAllStringFunc` — the correct tool for dynamic replacements**

```go
re := regexp.MustCompile(`\d+`)
result := re.ReplaceAllStringFunc("cat 3 and 17 dogs", func(m string) string {
    n, _ := strconv.Atoi(m)
    return strconv.Itoa(n * 2)
})
// "cat 6 and 34 dogs"
```

Option A is wrong: `strings.ReplaceAll` takes a string replacement, not a function. Option C processes digit-by-digit, so `"17"` would be processed as `"1"` then `"7"` separately — incorrect for multi-digit numbers. Option B is the correct pattern: regex finds complete number tokens, function transforms each one.

</details>

---

## 📊 Score Interpretation

| Score | Result |
|---|---|
| 30–32 ✅ | **Exceptional** — string manipulation and regex mastered. |
| 26–29 ✅ | **Ready** — review any missed sections and start. |
| 20–25 ⚠️ | **Study first** — identify your weakest section and work through it. |
| Below 20 ❌ | **Not ready** — spend time with the `strings`, `regexp`, and `strconv` package docs and examples. |

---

## 🔍 Review Map

| Missed | Topic to Study |
|---|---|
| Q1–Q10 | `strings.ToUpper/Lower`, `TrimSpace`, `Contains`, `Split/Join`, `Replace/ReplaceAll`, `Fields`, `Builder`, `Repeat`, `EqualFold` |
| Q11–Q18 | `regexp.Compile` vs `MustCompile`, `MatchString`, `FindString`, `FindAllString`, capture groups, `ReplaceAllStringFunc`, anchors `^`/`$` |
| Q19–Q25 | `os.ReadFile`, `os.WriteFile`, `bufio.Scanner`, `os.Args`, `defer f.Close()`, `os.Open` vs `os.OpenFile`, scanner buffer |
| Q26–Q32 | `strconv.Atoi`, `strconv.Itoa`, `ParseInt` bases, `NumError`, `Quote`, combining regex + strconv |