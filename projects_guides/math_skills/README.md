# Math-Skills Project Guide

## 📋 Project Overview
Build a program that reads numerical data from a file and calculates four fundamental statistical measures: Average, Median, Variance, and Standard Deviation. This project combines file I/O, data parsing, and mathematical computation.

---

## 🎯 Learning Objectives

By completing this project, you will learn:
1. **Statistical Concepts**: Understanding and implementing core statistical measures
2. **File Processing**: Reading and parsing data from text files
3. **Data Conversion**: Converting strings to numbers
4. **Mathematical Computation**: Implementing statistical formulas
5. **Data Sorting**: Sorting numbers for median calculation
6. **Number Rounding**: Proper rounding of floating-point numbers
7. **Error Handling**: Handling invalid data and file errors
8. **Algorithm Implementation**: Translating mathematical formulas to code

---

## 📚 Prerequisites - Topics You Must Know

### 1. **Go File I/O**
- `os` package:
  - `os.ReadFile()` - Read entire file
  - `os.Open()` - Open file for reading
- `bufio` package:
  - `bufio.NewScanner()` - Read file line by line
  - `scanner.Scan()` - Read next line
  - `scanner.Text()` - Get current line

### 2. **String Manipulation**
- `strings` package:
  - `strings.TrimSpace()` - Remove whitespace
  - `strings.Split()` - Split strings
- Converting strings to numbers:
  - `strconv.Atoi()` - String to integer
  - `strconv.ParseFloat()` - String to float

### 3. **Math Operations**
- `math` package:
  - `math.Sqrt()` - Square root
  - `math.Pow()` - Power (x²)
  - `math.Round()` - Round to nearest integer
- Basic arithmetic: +, -, *, /

### 4. **Slices**
- Creating slices: `[]int{}`, `[]float64{}`
- Appending: `append(slice, value)`
- Accessing elements: `slice[i]`
- Length: `len(slice)`
- Iterating: `for i, v := range slice`

### 5. **Sorting**
- `sort` package:
  - `sort.Ints()` - Sort integer slice
  - `sort.Float64s()` - Sort float slice

### 6. **Command-Line Arguments**
- `os.Args` - Access command-line arguments
- `os.Args[0]` - Program name
- `os.Args[1]` - First argument (filename)

---

## 📐 Understanding the Statistics

Before coding, understand what you're calculating:

### **1. Average (Mean)**
**Definition**: Sum of all values divided by count

**Formula**: 
```
Average = (x₁ + x₂ + x₃ + ... + xₙ) / n
```

**Example**:
```
Numbers: 10, 20, 30, 40, 50
Average = (10 + 20 + 30 + 40 + 50) / 5 = 150 / 5 = 30
```

**Key Points**:
- Sum all numbers
- Divide by how many numbers there are
- Result can be decimal (needs rounding)

---

### **2. Median**
**Definition**: The middle value when numbers are sorted

**Formula**:
- If odd count: median = middle value
- If even count: median = average of two middle values

**Example 1 (Odd count)**:
```
Numbers: 10, 25, 30, 35, 50
Sorted: 10, 25, 30, 35, 50
Median = 30 (middle value at position 2)
```

**Example 2 (Even count)**:
```
Numbers: 10, 20, 30, 40
Sorted: 10, 20, 30, 40
Median = (20 + 30) / 2 = 25
```

**Key Points**:
- MUST sort numbers first
- Find middle position(s)
- If even count, average the two middle values

---

### **3. Variance**
**Definition**: Measures how spread out numbers are from the mean

**Formula**:
```
Variance = Σ(xᵢ - mean)² / n

Where:
- xᵢ = each value
- mean = average
- n = count of values
- Σ = sum of all
```

**Step-by-step calculation**:
```
Numbers: 10, 20, 30, 40, 50
Mean = 30

1. Calculate differences from mean:
   10 - 30 = -20
   20 - 30 = -10
   30 - 30 = 0
   40 - 30 = 10
   50 - 30 = 20

2. Square each difference:
   (-20)² = 400
   (-10)² = 100
   (0)² = 0
   (10)² = 100
   (20)² = 400

3. Sum of squared differences:
   400 + 100 + 0 + 100 + 400 = 1000

4. Divide by count:
   Variance = 1000 / 5 = 200
```

**Key Points**:
- Calculate average first
- Find difference of each number from average
- Square each difference
- Sum all squared differences
- Divide by count

---

### **4. Standard Deviation**
**Definition**: Square root of variance (shows spread in original units)

**Formula**:
```
Standard Deviation = √Variance
```

**Example**:
```
If Variance = 200
Standard Deviation = √200 ≈ 14.14 → rounds to 14
```

**Key Points**:
- Calculate variance first
- Take square root
- Result is in same units as original data

---

## 🛠️ Step-by-Step Implementation Guide

### **Phase 1: Project Setup** ✅

#### Step 1: Create Project Structure
```
math-skills/
├── main.go
├── go.mod
├── data.txt (test file)
└── README.md
```

#### Step 2: Initialize Go Module
```bash
go mod init math-skills
```

#### Step 3: Basic Main Function
```go
package main

import (
    "fmt"
    "os"
)

func main() {
    // Check command-line arguments
    if len(os.Args) != 2 {
        fmt.Println("Usage: go run . <filename>")
        return
    }
    
    filename := os.Args[1]
    fmt.Println("Reading file:", filename)
}
```

**Test**: Run `go run . data.txt` to verify argument handling.

---

### **Phase 2: File Reading** 📖

#### Step 4: Read File Content
Create a function to read the file:

```go
func readFile(filename string) ([]string, error) {
    // Option 1: Read entire file
    // Use os.ReadFile() and split by newlines
    
    // Option 2 (Recommended): Read line by line
    // Open file with os.Open()
    // Use bufio.NewScanner()
    // Scan each line
    // Return slice of lines
}
```

**Implementation Hints**:
```go
// Open file
file, err := os.Open(filename)
if err != nil {
    return nil, err
}
defer file.Close()

// Create scanner
scanner := bufio.NewScanner(file)

// Read line by line
lines := []string{}
for scanner.Scan() {
    line := scanner.Text()
    lines = append(lines, line)
}

// Check for scan errors
if err := scanner.Err(); err != nil {
    return nil, err
}

return lines, nil
```

**Test**: Print all lines to verify reading works.

---

#### Step 5: Parse Numbers from Lines
Create a function to convert strings to numbers:

```go
func parseNumbers(lines []string) ([]float64, error) {
    // Create empty slice for numbers
    // Iterate through each line
    // Trim whitespace from line
    // Skip empty lines
    // Convert string to float64
    // Handle conversion errors
    // Append to numbers slice
    // Return numbers
}
```

**Key Points**:
- Use `strings.TrimSpace()` to clean lines
- Use `strconv.ParseFloat(str, 64)` to convert
- Skip empty lines
- Handle errors (non-numeric data)

**Test Cases**:
```go
// Test with sample data
lines := []string{"10", "20", "30"}
numbers, err := parseNumbers(lines)
// Should return [10, 20, 30]

// Test with whitespace
lines := []string{"  10  ", "20", "  30"}
// Should handle correctly

// Test with invalid data
lines := []string{"10", "abc", "30"}
// Should return error
```

---

### **Phase 3: Statistical Calculations** 📊

#### Step 6: Calculate Average
Create a function to calculate the mean:

```go
func calculateAverage(numbers []float64) float64 {
    // Check if slice is empty (avoid division by zero)
    // Initialize sum variable
    // Loop through numbers and add to sum
    // Divide sum by count of numbers
    // Return result
}
```

**Algorithm**:
```
sum = 0
for each number in numbers:
    sum = sum + number
average = sum / length of numbers
return average
```

**Manual Test**:
```
Input: [10, 20, 30, 40, 50]
Sum: 10 + 20 + 30 + 40 + 50 = 150
Count: 5
Average: 150 / 5 = 30.0
```

---

#### Step 7: Calculate Median
Create a function to find the middle value:

```go
func calculateMedian(numbers []float64) float64 {
    // IMPORTANT: Make a copy of the slice (don't modify original)
    // Sort the copied slice
    // Find the length
    // If length is odd:
    //     return middle element
    // If length is even:
    //     return average of two middle elements
}
```

**Algorithm**:
```
1. Copy slice: sorted = make copy of numbers
2. Sort: sort.Float64s(sorted)
3. Get length: n = len(sorted)
4. If n is odd:
     index = n / 2
     return sorted[index]
5. If n is even:
     index1 = n/2 - 1
     index2 = n/2
     return (sorted[index1] + sorted[index2]) / 2
```

**Manual Test**:
```
Input (odd): [30, 10, 50, 20, 40]
Sorted: [10, 20, 30, 40, 50]
Length: 5 (odd)
Middle index: 5/2 = 2
Median: 30

Input (even): [30, 10, 40, 20]
Sorted: [10, 20, 30, 40]
Length: 4 (even)
Middle indices: 1 and 2
Median: (20 + 30) / 2 = 25
```

**Key Points**:
- Always sort FIRST
- Use copy to avoid modifying original data
- Handle both odd and even cases
- Middle index for odd: `n/2`
- Middle indices for even: `n/2-1` and `n/2`

---

#### Step 8: Calculate Variance
Create a function to measure spread:

```go
func calculateVariance(numbers []float64, mean float64) float64 {
    // If only one number, variance is 0
    // Initialize sum of squared differences
    // For each number:
    //     Calculate difference from mean
    //     Square the difference
    //     Add to sum
    // Divide sum by count of numbers
    // Return result
}
```

**Algorithm**:
```
sum = 0
for each number in numbers:
    difference = number - mean
    squared = difference * difference
    sum = sum + squared
variance = sum / length of numbers
return variance
```

**Manual Test**:
```
Input: [10, 20, 30, 40, 50]
Mean: 30

Differences: -20, -10, 0, 10, 20
Squared: 400, 100, 0, 100, 400
Sum: 400 + 100 + 0 + 100 + 400 = 1000
Variance: 1000 / 5 = 200
```

**Key Points**:
- Calculate mean first (pass as parameter)
- Difference can be negative
- Squaring makes all values positive
- Don't forget to divide by count

---

#### Step 9: Calculate Standard Deviation
Create a function for standard deviation:

```go
func calculateStdDev(variance float64) float64 {
    // Simply return the square root of variance
    // Use math.Sqrt()
}
```

**Algorithm**:
```
stddev = √variance
```

**Manual Test**:
```
Input variance: 200
Standard Deviation: √200 ≈ 14.142
```

---

### **Phase 4: Formatting and Output** 🖨️

#### Step 10: Round Results
Create a function to round to nearest integer:

```go
func roundToInt(value float64) int {
    // Use math.Round() to round to nearest integer
    // Convert to int
    // Return
}
```

**Rounding Rules**:
- 14.4 → 14
- 14.5 → 15
- 14.6 → 15
- -14.5 → -15

**Test Cases**:
```
35.2 → 35
35.5 → 36
35.7 → 36
4.4 → 4
4.5 → 5
```

---

#### Step 11: Print Results
Create a function to display results:

```go
func printResults(avg, median, variance, stddev float64) {
    // Round each value to integer
    // Print in the required format:
    // Average: <value>
    // Median: <value>
    // Variance: <value>
    // Standard Deviation: <value>
}
```

**Exact Format Required**:
```
Average: 35
Median: 4
Variance: 5
Standard Deviation: 65
```

**Key Points**:
- Use exact format (case-sensitive)
- Include colon and space after label
- Values must be rounded integers
- Each on separate line

---

### **Phase 5: Integration** 🔗

#### Step 12: Complete Main Function
Bring everything together:

```go
func main() {
    // 1. Check command-line arguments
    if len(os.Args) != 2 {
        fmt.Println("Usage: go run . <filename>")
        return
    }
    
    filename := os.Args[1]
    
    // 2. Read file
    lines, err := readFile(filename)
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }
    
    // 3. Parse numbers
    numbers, err := parseNumbers(lines)
    if err != nil {
        fmt.Println("Error parsing numbers:", err)
        return
    }
    
    // 4. Check if we have data
    if len(numbers) == 0 {
        fmt.Println("No data found")
        return
    }
    
    // 5. Calculate statistics
    avg := calculateAverage(numbers)
    median := calculateMedian(numbers)
    variance := calculateVariance(numbers, avg)
    stddev := calculateStdDev(variance)
    
    // 6. Print results
    printResults(avg, median, variance, stddev)
}
```

---

### **Phase 6: Testing** 🧪

#### Step 13: Create Test Files

**Test File 1: Simple Data (data1.txt)**
```
10
20
30
40
50
```
**Expected Results**:
```
Average: 30
Median: 30
Variance: 200
Standard Deviation: 14
```

**Manual Calculation**:
- Average: (10+20+30+40+50)/5 = 30
- Median: sorted [10,20,30,40,50], middle = 30
- Variance: ((−20)²+(−10)²+0²+10²+20²)/5 = 200
- Std Dev: √200 ≈ 14.14 → 14

---

**Test File 2: Even Count (data2.txt)**
```
10
20
30
40
```
**Expected Results**:
```
Average: 25
Median: 25
Variance: 125
Standard Deviation: 11
```

---

**Test File 3: Large Numbers (data3.txt)**
```
189
113
121
114
145
110
```
**Calculate yourself to verify**

---

**Test File 4: Decimal Results (data4.txt)**
```
5
15
25
35
45
```
**Expected**:
- Average: 25
- Median: 25
- Variance: 200
- Std Dev: 14

---

**Test File 5: Single Value (data5.txt)**
```
42
```
**Expected**:
- Average: 42
- Median: 42
- Variance: 0
- Std Dev: 0

---

#### Step 14: Edge Cases to Test

**Empty File**:
```
(empty file)
```
**Expected**: Error message or "No data found"

**Invalid Data**:
```
10
abc
30
```
**Expected**: Error parsing numbers

**Whitespace**:
```
  10  
  20  
  30  
```
**Expected**: Should handle correctly

**Negative Numbers**:
```
-10
-5
0
5
10
```
**Expected**: Should calculate correctly

---

### **Phase 7: Verification** ✅

#### Step 15: Manual Verification Steps

For each test file, manually calculate:

**Step-by-step for data1.txt [10, 20, 30, 40, 50]**:

1. **Average**:
   - Sum: 10 + 20 + 30 + 40 + 50 = 150
   - Count: 5
   - Average: 150 / 5 = 30 ✓

2. **Median**:
   - Sort: [10, 20, 30, 40, 50]
   - Count: 5 (odd)
   - Middle index: 5/2 = 2
   - Median: 30 ✓

3. **Variance**:
   - Differences from mean (30): [-20, -10, 0, 10, 20]
   - Squared: [400, 100, 0, 100, 400]
   - Sum: 1000
   - Variance: 1000 / 5 = 200 ✓

4. **Standard Deviation**:
   - √200 = 14.142...
   - Rounded: 14 ✓

Use online calculators to verify:
- Search "statistics calculator online"
- Enter your data
- Compare results

---

## 🐛 Common Issues and Solutions

### Issue 1: Wrong Median
**Problem**: Median is incorrect
**Solution**: 
- Ensure you're sorting the numbers first
- Make a copy before sorting (don't modify original)
- Check odd/even length handling
- Verify middle index calculation

### Issue 2: Variance Too High/Low
**Problem**: Variance doesn't match expected
**Solution**:
- Calculate average first, use it correctly
- Square the differences (not just multiply by 2)
- Don't forget to divide by count
- Check if you're using the mean correctly

### Issue 3: File Not Found
**Problem**: "no such file or directory"
**Solution**:
- Check file path is correct
- Run from correct directory
- Use relative or absolute paths correctly

### Issue 4: Parse Error
**Problem**: "invalid syntax" when parsing numbers
**Solution**:
- Trim whitespace with `strings.TrimSpace()`
- Skip empty lines
- Check for non-numeric characters
- Handle errors properly

### Issue 5: Division by Zero
**Problem**: Program crashes
**Solution**:
- Check if slice is empty before calculations
- Handle zero-length input gracefully

---

## 📋 Debugging Checklist

When results are wrong, check:

- [ ] Are you reading the file correctly?
- [ ] Are all lines parsed as numbers?
- [ ] Is the average calculation correct?
- [ ] Did you sort before finding median?
- [ ] Are you using the correct middle index for median?
- [ ] Did you calculate differences from mean correctly?
- [ ] Did you square the differences?
- [ ] Did you divide variance by count?
- [ ] Did you take square root for standard deviation?
- [ ] Are you rounding correctly?
- [ ] Is the output format exact?

---

## 🎓 Mathematical Formulas Summary

```
Given data: x₁, x₂, x₃, ..., xₙ (n values)

1. Average (μ or x̄):
   μ = (Σxᵢ) / n
   
2. Median (M):
   - Sort data first
   - If n is odd: M = x₍ₙ₊₁₎/₂
   - If n is even: M = (xₙ/₂ + x₍ₙ/₂₎₊₁) / 2

3. Variance (σ²):
   σ² = Σ(xᵢ - μ)² / n

4. Standard Deviation (σ):
   σ = √σ²
```

---

## ✅ Submission Checklist

**Functionality**:
- [ ] Reads filename from command-line argument
- [ ] Reads data from file correctly
- [ ] Parses numbers correctly
- [ ] Calculates average correctly
- [ ] Calculates median correctly (sorted)
- [ ] Calculates variance correctly
- [ ] Calculates standard deviation correctly
- [ ] Rounds to integers properly
- [ ] Prints in exact required format
- [ ] Handles errors gracefully
- [ ] Works with provided test data

**Code Quality**:
- [ ] Code is well-organized
- [ ] Functions are small and focused
- [ ] Variable names are clear
- [ ] Comments explain logic
- [ ] No hardcoded values
- [ ] Proper error handling everywhere
- [ ] No unused code

**Testing**:
- [ ] Tested with multiple datasets
- [ ] Tested edge cases (empty, single value)
- [ ] Tested with invalid data
- [ ] Manually verified calculations
- [ ] Results match online calculators

---

## 📖 Key Go Concepts Used

| Concept | Package/Function | Purpose |
|---------|-----------------|---------|
| File Reading | `os.Open()`, `bufio.Scanner` | Read file line by line |
| String Parsing | `strconv.ParseFloat()` | Convert strings to numbers |
| Math Operations | `math.Sqrt()`, `math.Round()` | Square root and rounding |
| Sorting | `sort.Float64s()` | Sort numbers for median |
| Slices | `append()`, `len()` | Store and manipulate data |
| Command-line Args | `os.Args` | Get filename from user |
| Error Handling | `if err != nil` | Handle file and parse errors |

---

## 🚀 Pro Tips

1. **Test with Small Data First**: Use 3-5 numbers you can calculate by hand
2. **Print Intermediate Values**: Print sum, mean, etc. to verify
3. **Use Online Calculators**: Verify your calculations
4. **Handle Edge Cases**: Empty file, single value, negative numbers
5. **Copy Before Sorting**: Don't modify the original slice
6. **Check Types**: Use `float64` for calculations, not `int`
7. **Understand the Math**: Don't just implement formulas blindly
8. **Test Rounding**: Ensure 0.5 rounds up consistently
9. **Validate Input**: Check for non-numeric data
10. **Keep Functions Pure**: Each function should do one thing well

---

## 📚 Resources for Understanding

**Statistics Concepts**:
- [Khan Academy - Statistics](https://www.khanacademy.org/math/statistics-probability)
- [Variance Explained Visually](https://www.youtube.com/watch?v=E4HAYd0QnRc)
- [Understanding Standard Deviation](https://www.mathsisfun.com/data/standard-deviation.html)

**Go Resources**:
- [Go by Example - File Reading](https://gobyexample.com/reading-files)
- [Go Documentation - math](https://pkg.go.dev/math)
- [Go Documentation - sort](https://pkg.go.dev/sort)

**Online Calculators** (for verification):
- [Calculator.net - Statistics](https://www.calculator.net/statistics-calculator.html)
- [Standard Deviation Calculator](https://www.calculator.net/standard-deviation-calculator.html)

---

## 💡 Extension Ideas

After completing basic requirements:

1. **Additional Statistics**:
   - Mode (most frequent value)
   - Range (max - min)
   - Quartiles (Q1, Q2, Q3)
   - Interquartile Range (IQR)

2. **Data Visualization**:
   - Create histogram in terminal
   - Draw distribution curve
   - Show box plot

3. **Multiple Files**:
   - Compare statistics across files
   - Aggregate multiple datasets

4. **Export Results**:
   - Save results to CSV
   - Generate report file

5. **Input Validation**:
   - Check for outliers
   - Detect data quality issues

6. **Performance**:
   - Calculate statistics in parallel
   - Optimize for large datasets

---

## 🧮 Example Walkthrough

**Complete Example: data.txt**
```
10
15
20
25
30
```

**Step-by-Step Calculation**:

1. **Average**:
   ```
   Sum = 10 + 15 + 20 + 25 + 30 = 100
   Count = 5
   Average = 100 / 5 = 20
   ```

2. **Median**:
   ```
   Already sorted: [10, 15, 20, 25, 30]
   Count = 5 (odd)
   Middle index = 5/2 = 2
   Median = 20
   ```

3. **Variance**:
   ```
   Mean = 20
   Differences: [10-20, 15-20, 20-20, 25-20, 30-20]
              = [-10, -5, 0, 5, 10]
   Squared: [100, 25, 0, 25, 100]
   Sum = 250
   Variance = 250 / 5 = 50
   ```

4. **Standard Deviation**:
   ```
   Std Dev = √50 = 7.071...
   Rounded = 7
   ```

**Final Output**:
```
Average: 20
Median: 20
Variance: 50
Standard Deviation: 7
```

---

**Remember**: This project is about understanding statistics, not just implementing formulas. Take time to understand what each measure means! 📊