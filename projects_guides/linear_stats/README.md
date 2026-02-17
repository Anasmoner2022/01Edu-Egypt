# Linear-Stats Project Guide

## 📋 Project Overview
Build a program that reads numerical data from a file and calculates two key statistical measures: the Linear Regression Line (best-fit line) and the Pearson Correlation Coefficient (strength of linear relationship). This project extends your statistical knowledge from math-skills into predictive analytics and relationship measurement.

---

## 🎯 Learning Objectives

By completing this project, you will learn:
1. **Linear Regression**: Finding the best-fit line through data points
2. **Correlation Analysis**: Measuring strength of linear relationships
3. **Slope and Intercept Calculation**: Understanding line equations
4. **Predictive Modeling**: Using regression for prediction
5. **Statistical Significance**: Interpreting correlation values
6. **Precision Formatting**: Outputting numbers with exact decimal places
7. **Data Indexing**: Treating line numbers as x-coordinates
8. **Mathematical Formulas**: Implementing complex statistical calculations

---

## 📚 Prerequisites - Topics You Must Know

### 1. **Math-Skills Foundation**
You should understand from the previous project:
- Average (Mean)
- Variance
- Standard Deviation
- Summation (Σ notation)

### 2. **File I/O** (Same as math-skills)
- Reading files line by line
- Parsing strings to numbers
- Error handling

### 3. **Mathematical Concepts**
- **Coordinate System**: (x, y) points
- **Line Equation**: y = mx + b
  - m = slope
  - b = y-intercept
- **Covariance**: Measure of joint variability
- **Correlation**: Standardized covariance

### 4. **Number Formatting**
- `fmt.Printf()` with format specifiers
- `%.6f` - 6 decimal places
- `%.10f` - 10 decimal places

### 5. **Slices and Iteration**
- Creating parallel slices for x and y values
- Iterating with indices
- Accessing elements by index

---

## 📐 Understanding the Mathematics

Before coding, you must deeply understand what you're calculating:

### **What is Linear Regression?**

**Definition**: Finding the straight line that best fits a set of points.

**Visual Example**:
```
Y axis
  ↑
200|     •
   |   •   •
150|  •  •   •
   | •   •  •
100|•   •  •
   |  •  •
 50|   •
   |_______________→ X axis
    0  2  4  6  8

Goal: Find line y = mx + b that minimizes distance to all points
```

**Why It Matters**:
- Predicts future values
- Shows trends in data
- Quantifies relationships

---

### **Understanding the Data Structure**

**Input File**:
```
189  ← line 0, value 189 → point (0, 189)
113  ← line 1, value 113 → point (1, 113)
121  ← line 2, value 121 → point (2, 121)
114  ← line 3, value 114 → point (3, 114)
145  ← line 4, value 145 → point (4, 145)
110  ← line 5, value 110 → point (5, 110)
```

**Key Insight**: 
- X values are the line numbers (0, 1, 2, 3, ...)
- Y values are the numbers from the file
- You're finding the relationship between position and value

---

### **Linear Regression Formula**

The regression line is: **y = mx + b**

Where:
- **m** = slope (how much y changes per unit of x)
- **b** = y-intercept (value of y when x = 0)

**Calculating Slope (m)**:
```
         Σ((xᵢ - x̄)(yᵢ - ȳ))
m = ────────────────────────────
         Σ((xᵢ - x̄)²)

Where:
- xᵢ = each x value
- yᵢ = each y value
- x̄ = mean of x values
- ȳ = mean of y values
- Σ = sum of all
```

**Alternative Formula** (easier to compute):
```
     n·Σ(xᵢ·yᵢ) - Σ(xᵢ)·Σ(yᵢ)
m = ──────────────────────────
     n·Σ(xᵢ²) - (Σ(xᵢ))²

Where n = number of data points
```

**Calculating Y-Intercept (b)**:
```
b = ȳ - m·x̄
```

---

### **Step-by-Step Calculation Example**

**Data**: 
```
Line (x) | Value (y)
---------|----------
   0     |   189
   1     |   113
   2     |   121
   3     |   114
   4     |   145
   5     |   110
```

**Step 1: Calculate Means**
```
x̄ = (0 + 1 + 2 + 3 + 4 + 5) / 6 = 15/6 = 2.5

ȳ = (189 + 113 + 121 + 114 + 145 + 110) / 6 = 792/6 = 132
```

**Step 2: Calculate Required Sums**

For slope formula, we need:
- Σ(xᵢ·yᵢ) = sum of products
- Σ(xᵢ) = sum of x values
- Σ(yᵢ) = sum of y values
- Σ(xᵢ²) = sum of x squared

```
| x  | y   | x·y | x²  |
|----|-----|-----|-----|
| 0  | 189 |   0 |  0  |
| 1  | 113 | 113 |  1  |
| 2  | 121 | 242 |  4  |
| 3  | 114 | 342 |  9  |
| 4  | 145 | 580 | 16  |
| 5  | 110 | 550 | 25  |
|----|-----|-----|-----|
Sum: 15   792  1827  55
```

**Step 3: Calculate Slope (m)**
```
     n·Σ(xᵢ·yᵢ) - Σ(xᵢ)·Σ(yᵢ)
m = ──────────────────────────
     n·Σ(xᵢ²) - (Σ(xᵢ))²

     6·1827 - 15·792
m = ────────────────
     6·55 - 15²

     10962 - 11880
m = ───────────────
     330 - 225

     -918
m = ──────
     105

m = -8.742857... ≈ -8.742857
```

**Step 4: Calculate Y-Intercept (b)**
```
b = ȳ - m·x̄
b = 132 - (-8.742857)·2.5
b = 132 + 21.857143
b = 153.857143
```

**Step 5: Final Linear Regression Line**
```
y = -8.742857x + 153.857143
```

**Interpretation**:
- Slope is negative (-8.74), so values are decreasing
- For each increase in x (line number), y decreases by ~8.74
- When x = 0, y starts at ~153.86

---

### **Pearson Correlation Coefficient**

**Definition**: Measures the strength and direction of the linear relationship between x and y.

**Range**: -1 ≤ r ≤ 1

**Interpretation**:
- **r = 1**: Perfect positive linear relationship
- **r = 0.7 to 0.9**: Strong positive correlation
- **r = 0.4 to 0.7**: Moderate positive correlation
- **r = 0**: No linear correlation
- **r = -0.4 to -0.7**: Moderate negative correlation
- **r = -0.7 to -0.9**: Strong negative correlation
- **r = -1**: Perfect negative linear relationship

**Formula**:
```
           Σ((xᵢ - x̄)(yᵢ - ȳ))
r = ─────────────────────────────────
    √(Σ(xᵢ - x̄)²) · √(Σ(yᵢ - ȳ)²)
```

**Alternative Formula** (easier to compute):
```
         n·Σ(xᵢ·yᵢ) - Σ(xᵢ)·Σ(yᵢ)
r = ────────────────────────────────────────
    √(n·Σ(xᵢ²) - (Σ(xᵢ))²) · √(n·Σ(yᵢ²) - (Σ(yᵢ))²)
```

**Relationship to Regression**:
- r² = coefficient of determination
- Shows what proportion of variance is explained by the line
- If r² = 0.8, the line explains 80% of the variance

---

### **Pearson Coefficient Calculation Example**

Using same data as above, we need additional sum:
- Σ(yᵢ²) = sum of y squared

```
| y   | y²    |
|-----|-------|
| 189 | 35721 |
| 113 | 12769 |
| 121 | 14641 |
| 114 | 12996 |
| 145 | 21025 |
| 110 | 12100 |
|-----|-------|
Sum:    109252
```

**Calculate r**:
```
         n·Σ(xᵢ·yᵢ) - Σ(xᵢ)·Σ(yᵢ)
r = ────────────────────────────────────────
    √(n·Σ(xᵢ²) - (Σ(xᵢ))²) · √(n·Σ(yᵢ²) - (Σ(yᵢ))²)

         6·1827 - 15·792
r = ────────────────────────────────────────
    √(6·55 - 15²) · √(6·109252 - 792²)

         10962 - 11880
r = ──────────────────────────────────────
    √(330 - 225) · √(655512 - 627264)

         -918
r = ───────────────────
    √105 · √28248

         -918
r = ──────────────
    10.247 · 168.08

         -918
r = ────────────
    1722.04

r = -0.5330784094...
```

**Interpretation**: 
- r ≈ -0.533
- Moderate negative correlation
- As x increases, y tends to decrease
- About 28% of variance explained (r² = 0.284)

---

## 🛠️ Step-by-Step Implementation Guide

### **Phase 1: Project Setup** ✅

#### Step 1: Create Project Structure
```
linear-stats/
├── main.go
├── go.mod
├── data.txt (test file)
└── README.md
```

#### Step 2: Initialize Go Module
```bash
go mod init linear-stats
```

#### Step 3: Basic Main Function
```go
package main

import (
    "fmt"
    "os"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Println("Usage: go run . <filename>")
        return
    }
    
    filename := os.Args[1]
    fmt.Println("Reading file:", filename)
}
```

---

### **Phase 2: Data Reading** 📖

#### Step 4: Read and Parse File
Create functions to read data (similar to math-skills):

```go
func readFile(filename string) ([]float64, error) {
    // Open file
    // Read line by line
    // Parse each line to float64
    // Return slice of y values
}
```

**Key Difference from math-skills**: 
You need to create x values automatically:
```go
func createDataPoints(yValues []float64) ([]float64, []float64) {
    n := len(yValues)
    xValues := make([]float64, n)
    
    // Generate x values: 0, 1, 2, 3, ...
    for i := 0; i < n; i++ {
        xValues[i] = float64(i)
    }
    
    return xValues, yValues
}
```

**Test**:
```go
yValues := []float64{189, 113, 121, 114, 145, 110}
xValues, _ := createDataPoints(yValues)
// xValues should be [0, 1, 2, 3, 4, 5]
```

---

### **Phase 3: Statistical Calculations** 📊

#### Step 5: Calculate Basic Statistics
Create helper functions:

```go
func calculateMean(values []float64) float64 {
    // Sum all values
    // Divide by count
    // Return mean
}

func calculateSum(values []float64) float64 {
    // Sum all values
    // Return sum
}
```

**Test**:
```go
x := []float64{0, 1, 2, 3, 4, 5}
meanX := calculateMean(x)  // Should be 2.5

y := []float64{189, 113, 121, 114, 145, 110}
meanY := calculateMean(y)  // Should be 132
```

---

#### Step 6: Calculate Products and Squares
Create functions for the sums needed:

```go
func calculateSumOfProducts(x, y []float64) float64 {
    // Σ(xᵢ·yᵢ)
    // For each pair (xᵢ, yᵢ):
    //   multiply them
    //   add to sum
    // Return sum
}

func calculateSumOfSquares(values []float64) float64 {
    // Σ(xᵢ²)
    // For each value:
    //   square it
    //   add to sum
    // Return sum
}
```

**Implementation Hints**:
```go
func calculateSumOfProducts(x, y []float64) float64 {
    sum := 0.0
    for i := 0; i < len(x); i++ {
        sum += x[i] * y[i]
    }
    return sum
}
```

**Test**:
```go
x := []float64{0, 1, 2, 3, 4, 5}
y := []float64{189, 113, 121, 114, 145, 110}
sumXY := calculateSumOfProducts(x, y)  // Should be 1827
sumX2 := calculateSumOfSquares(x)      // Should be 55
sumY2 := calculateSumOfSquares(y)      // Should be 109252
```

---

### **Phase 4: Linear Regression** 📈

#### Step 7: Calculate Slope
Implement the slope formula:

```go
func calculateSlope(x, y []float64) float64 {
    n := float64(len(x))
    
    // Calculate required sums
    sumX := calculateSum(x)
    sumY := calculateSum(y)
    sumXY := calculateSumOfProducts(x, y)
    sumX2 := calculateSumOfSquares(x)
    
    // Apply formula:
    //      n·Σ(xᵢ·yᵢ) - Σ(xᵢ)·Σ(yᵢ)
    // m = ──────────────────────────
    //      n·Σ(xᵢ²) - (Σ(xᵢ))²
    
    numerator := n*sumXY - sumX*sumY
    denominator := n*sumX2 - sumX*sumX
    
    if denominator == 0 {
        return 0  // Handle edge case
    }
    
    return numerator / denominator
}
```

**Manual Verification**:
```
Data: x=[0,1,2,3,4,5], y=[189,113,121,114,145,110]
n = 6
sumX = 15
sumY = 792
sumXY = 1827
sumX2 = 55

numerator = 6*1827 - 15*792 = 10962 - 11880 = -918
denominator = 6*55 - 15*15 = 330 - 225 = 105
slope = -918 / 105 = -8.742857142857143
```

---

#### Step 8: Calculate Y-Intercept
Implement the intercept formula:

```go
func calculateIntercept(x, y []float64, slope float64) float64 {
    // b = ȳ - m·x̄
    meanX := calculateMean(x)
    meanY := calculateMean(y)
    
    return meanY - slope*meanX
}
```

**Manual Verification**:
```
meanX = 2.5
meanY = 132
slope = -8.742857142857143

intercept = 132 - (-8.742857142857143)*2.5
          = 132 + 21.857142857142858
          = 153.857142857142858
```

---

#### Step 9: Format Regression Line
Create function to format output:

```go
func formatRegressionLine(slope, intercept float64) string {
    // Format: "Linear Regression Line: y = <slope>x + <intercept>"
    // slope and intercept need 6 decimal places
    // Use fmt.Sprintf with %.6f
    
    return fmt.Sprintf("Linear Regression Line: y = %.6fx + %.6f", slope, intercept)
}
```

**Expected Output**:
```
Linear Regression Line: y = -8.742857x + 153.857143
```

**Important**: Handle negative intercept correctly:
```go
// If intercept is negative:
// y = 2.5x + -3.2  ← WRONG
// y = 2.5x - 3.2   ← CORRECT

func formatRegressionLine(slope, intercept float64) string {
    if intercept >= 0 {
        return fmt.Sprintf("Linear Regression Line: y = %.6fx + %.6f", slope, intercept)
    } else {
        return fmt.Sprintf("Linear Regression Line: y = %.6fx - %.6f", slope, -intercept)
    }
}
```

---

### **Phase 5: Pearson Correlation** 📉

#### Step 10: Calculate Pearson Coefficient
Implement the correlation formula:

```go
func calculatePearsonCorrelation(x, y []float64) float64 {
    n := float64(len(x))
    
    // Calculate required sums
    sumX := calculateSum(x)
    sumY := calculateSum(y)
    sumXY := calculateSumOfProducts(x, y)
    sumX2 := calculateSumOfSquares(x)
    sumY2 := calculateSumOfSquares(y)
    
    // Apply formula:
    //          n·Σ(xᵢ·yᵢ) - Σ(xᵢ)·Σ(yᵢ)
    // r = ────────────────────────────────────────
    //     √(n·Σ(xᵢ²) - (Σ(xᵢ))²) · √(n·Σ(yᵢ²) - (Σ(yᵢ))²)
    
    numerator := n*sumXY - sumX*sumY
    
    denomX := n*sumX2 - sumX*sumX
    denomY := n*sumY2 - sumY*sumY
    
    if denomX <= 0 || denomY <= 0 {
        return 0  // Handle edge case
    }
    
    denominator := math.Sqrt(denomX) * math.Sqrt(denomY)
    
    if denominator == 0 {
        return 0
    }
    
    return numerator / denominator
}
```

**Manual Verification**:
```
Data: x=[0,1,2,3,4,5], y=[189,113,121,114,145,110]
n = 6
sumX = 15
sumY = 792
sumXY = 1827
sumX2 = 55
sumY2 = 109252

numerator = 6*1827 - 15*792 = -918

denomX = 6*55 - 15*15 = 105
denomY = 6*109252 - 792*792 = 655512 - 627264 = 28248

denominator = √105 * √28248 = 10.247 * 168.08 = 1722.04

r = -918 / 1722.04 = -0.5330784094
```

---

#### Step 11: Format Pearson Output
Create function to format correlation:

```go
func formatPearsonCorrelation(r float64) string {
    // Format: "Pearson Correlation Coefficient: <r>"
    // r needs 10 decimal places
    // Use fmt.Sprintf with %.10f
    
    return fmt.Sprintf("Pearson Correlation Coefficient: %.10f", r)
}
```

**Expected Output**:
```
Pearson Correlation Coefficient: -0.5330784094
```

---

### **Phase 6: Integration** 🔗

#### Step 12: Complete Main Function
Bring everything together:

```go
package main

import (
    "fmt"
    "math"
    "os"
)

func main() {
    // 1. Check command-line arguments
    if len(os.Args) != 2 {
        fmt.Println("Usage: go run . <filename>")
        return
    }
    
    filename := os.Args[1]
    
    // 2. Read y values from file
    yValues, err := readFile(filename)
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }
    
    // 3. Check if we have data
    if len(yValues) < 2 {
        fmt.Println("Need at least 2 data points")
        return
    }
    
    // 4. Create x values (0, 1, 2, 3, ...)
    xValues, yValues := createDataPoints(yValues)
    
    // 5. Calculate linear regression
    slope := calculateSlope(xValues, yValues)
    intercept := calculateIntercept(xValues, yValues, slope)
    
    // 6. Calculate Pearson correlation
    correlation := calculatePearsonCorrelation(xValues, yValues)
    
    // 7. Print results
    fmt.Println(formatRegressionLine(slope, intercept))
    fmt.Println(formatPearsonCorrelation(correlation))
}
```

---

### **Phase 7: Testing** 🧪

#### Step 13: Create Test Cases

**Test File 1: Simple Increasing (data1.txt)**
```
0
10
20
30
40
50
```

**Expected Calculation**:
```
x: [0, 1, 2, 3, 4, 5]
y: [0, 10, 20, 30, 40, 50]

This is a perfect line: y = 10x + 0
Slope: 10.000000
Intercept: 0.000000
Correlation: 1.000000 (perfect positive)
```

---

**Test File 2: Simple Decreasing (data2.txt)**
```
50
40
30
20
10
0
```

**Expected Calculation**:
```
x: [0, 1, 2, 3, 4, 5]
y: [50, 40, 30, 20, 10, 0]

Perfect negative line: y = -10x + 50
Slope: -10.000000
Intercept: 50.000000
Correlation: -1.000000 (perfect negative)
```

---

**Test File 3: Horizontal Line (data3.txt)**
```
25
25
25
25
25
```

**Expected Calculation**:
```
x: [0, 1, 2, 3, 4]
y: [25, 25, 25, 25, 25]

Flat line: y = 0x + 25
Slope: 0.000000
Intercept: 25.000000
Correlation: undefined (or 0)
```

---

**Test File 4: Given Example (data4.txt)**
```
189
113
121
114
145
110
```

**Expected Output**:
```
Linear Regression Line: y = -8.742857x + 153.857143
Pearson Correlation Coefficient: -0.5330784094
```

---

**Test File 5: Large Numbers (data5.txt)**
```
1000
2000
3000
4000
5000
```

**Expected**: Perfect positive correlation (r = 1.0)

---

#### Step 14: Verify with Online Calculators

Use these tools to verify your calculations:
1. Search "linear regression calculator"
2. Enter your x and y values
3. Compare slope, intercept, and r value

**Recommended Calculators**:
- [Calculator.net Linear Regression](https://www.calculator.net/linear-regression-calculator.html)
- [Desmos Graphing Calculator](https://www.desmos.com/calculator)
- [Stat Trek Linear Regression](https://stattrek.com/online-calculator/linear-regression.aspx)

---

### **Phase 8: Edge Cases** ⚠️

#### Step 15: Handle Special Cases

**Case 1: Two Points**
```
Minimum data for regression
Should work correctly
```

**Case 2: Vertical Spread (same x, different y)**
```
Not possible with our data structure
(x is always 0, 1, 2, 3, ...)
```

**Case 3: No Spread in Y (all same)**
```
y: [5, 5, 5, 5]
Slope should be 0
Correlation undefined or NaN
Handle division by zero
```

**Case 4: Large Dataset**
```
Test with 1000+ points
Check performance
Ensure no overflow
```

**Edge Case Handling**:
```go
func calculatePearsonCorrelation(x, y []float64) float64 {
    // ... calculations ...
    
    if denominator == 0 {
        // No variation in data
        return 0  // or math.NaN()
    }
    
    result := numerator / denominator
    
    // Handle floating point errors
    if result > 1.0 {
        result = 1.0
    }
    if result < -1.0 {
        result = -1.0
    }
    
    return result
}
```

---

## 🐛 Common Issues and Solutions

### Issue 1: Wrong Sign in Intercept
**Problem**: Output shows "y = 2x + -3" instead of "y = 2x - 3"
**Solution**: Handle negative intercept in formatting:
```go
if intercept >= 0 {
    return fmt.Sprintf("... + %.6f", intercept)
} else {
    return fmt.Sprintf("... - %.6f", -intercept)
}
```

### Issue 2: Correlation Outside [-1, 1]
**Problem**: r = 1.0000000001 due to floating point errors
**Solution**: Clamp the value:
```go
if r > 1.0 { r = 1.0 }
if r < -1.0 { r = -1.0 }
```

### Issue 3: Wrong Decimal Places
**Problem**: Output has wrong number of decimals
**Solution**: 
- Regression: Use `%.6f` (6 decimals)
- Pearson: Use `%.10f` (10 decimals)

### Issue 4: Division by Zero
**Problem**: Crash when all y values are the same
**Solution**: Check denominators before division:
```go
if denominator == 0 {
    return 0  // or handle appropriately
}
```

### Issue 5: Wrong X Values
**Problem**: Using line numbers starting from 1 instead of 0
**Solution**: Ensure x starts at 0:
```go
for i := 0; i < n; i++ {
    xValues[i] = float64(i)  // NOT float64(i+1)
}
```

---

## 📋 Testing Checklist

**Basic Functionality**:
- [ ] Reads file correctly
- [ ] Creates x values as 0, 1, 2, 3, ...
- [ ] Calculates slope correctly
- [ ] Calculates intercept correctly
- [ ] Calculates Pearson coefficient correctly
- [ ] Formats output with correct decimal places
- [ ] Handles negative intercept correctly

**Edge Cases**:
- [ ] Works with 2 points (minimum)
- [ ] Handles horizontal line (slope = 0)
- [ ] Handles perfect correlation (r = 1 or -1)
- [ ] No crashes on division by zero
- [ ] Large datasets (100+ points)

**Output Format**:
- [ ] Exact format: "Linear Regression Line: y = <slope>x + <intercept>"
- [ ] Exact format: "Pearson Correlation Coefficient: <value>"
- [ ] 6 decimal places for slope and intercept
- [ ] 10 decimal places for correlation
- [ ] Handles + and - signs correctly

**Verification**:
- [ ] Results match online calculators
- [ ] Manual calculations verified
- [ ] All test files produce correct output

---

## ✅ Submission Checklist

**Code Quality**:
- [ ] Clean, readable code
- [ ] Meaningful variable names
- [ ] Comments explain formulas
- [ ] No hardcoded values
- [ ] Error handling
- [ ] Efficient algorithms

**Functionality**:
- [ ] Reads filename from command-line
- [ ] Parses data correctly
- [ ] Calculates regression correctly
- [ ] Calculates correlation correctly
- [ ] Outputs exact format
- [ ] Handles edge cases gracefully

**Testing**:
- [ ] Tested with multiple datasets
- [ ] Verified with calculators
- [ ] Checked decimal precision
- [ ] Tested edge cases
- [ ] No crashes or errors

---

## 📖 Key Formulas Summary

### **Linear Regression**
```
Slope (m):
     n·Σ(xᵢ·yᵢ) - Σ(xᵢ)·Σ(yᵢ)
m = ──────────────────────────
     n·Σ(xᵢ²) - (Σ(xᵢ))²

Intercept (b):
b = ȳ - m·x̄

Line:
y = mx + b
```

### **Pearson Correlation**
```
         n·Σ(xᵢ·yᵢ) - Σ(xᵢ)·Σ(yᵢ)
r = ────────────────────────────────────────
    √(n·Σ(xᵢ²) - (Σ(xᵢ))²) · √(n·Σ(yᵢ²) - (Σ(yᵢ))²)

Range: -1 ≤ r ≤ 1
```

### **Required Sums**
```
Σ(xᵢ)    = sum of x values
Σ(yᵢ)    = sum of y values
Σ(xᵢ·yᵢ) = sum of products
Σ(xᵢ²)   = sum of x squared
Σ(yᵢ²)   = sum of y squared
```

---

## 📊 Interpreting Results

### **Slope (m)**
- **Positive**: y increases as x increases
- **Negative**: y decreases as x increases
- **Zero**: No change (horizontal line)
- **Magnitude**: How steep the line is

### **Intercept (b)**
- Value of y when x = 0
- Where line crosses y-axis
- Starting point of regression

### **Pearson Coefficient (r)**
| Value | Interpretation |
|-------|----------------|
| 1.0 | Perfect positive correlation |
| 0.7 to 1.0 | Strong positive correlation |
| 0.3 to 0.7 | Moderate positive correlation |
| -0.3 to 0.3 | Weak or no correlation |
| -0.7 to -0.3 | Moderate negative correlation |
| -1.0 to -0.7 | Strong negative correlation |
| -1.0 | Perfect negative correlation |

### **R-Squared (r²)**
- Square of correlation coefficient
- Proportion of variance explained
- Example: r = 0.8 → r² = 0.64 → 64% variance explained

---

## 🚀 Pro Tips

1. **Verify Formulas**: Manually calculate with small dataset first
2. **Use High Precision**: Use float64, not float32
3. **Check Signs**: Pay attention to + and - in formulas
4. **Test Extremes**: Perfect correlation, no correlation, negative
5. **Compare Results**: Use multiple online calculators
6. **Understand Math**: Don't just implement blindly
7. **Handle Edge Cases**: Division by zero, identical values
8. **Format Carefully**: Exact decimal places matter
9. **Keep It Simple**: Don't overcomplicate the implementation
10. **Document**: Comment the formulas you're using

---

## 💡 Extension Ideas

After completing basic requirements:

1. **Visualization**:
   - Generate scatter plot with regression line
   - Show residuals (errors)

2. **Additional Statistics**:
   - R-squared value
   - Standard error of estimate
   - Confidence intervals
   - P-value for significance

3. **Predictions**:
   - Predict y for given x values
   - Forecast future values

4. **Multiple Regression**:
   - Handle multiple independent variables
   - y = b₀ + b₁x₁ + b₂x₂ + ...

5. **Quality Metrics**:
   - Mean Absolute Error (MAE)
   - Root Mean Square Error (RMSE)
   - Residual analysis

---

## 📚 Additional Resources

**Linear Regression**:
- [Khan Academy - Linear Regression](https://www.khanacademy.org/math/statistics-probability/describing-relationships-quantitative-data)
- [StatQuest - Linear Regression](https://www.youtube.com/watch?v=nk2CQITm_eo)
- [Understanding Least Squares](https://setosa.io/ev/ordinary-least-squares-regression/)

**Pearson Correlation**:
- [Understanding Correlation](https://www.mathsisfun.com/data/correlation.html)
- [Correlation vs Causation](https://tylervigen.com/spurious-correlations)
- [Interpreting Correlation](https://statistics.laerd.com/statistical-guides/pearson-correlation-coefficient-statistical-guide.php)

**Go Resources**:
- [Go Math Package](https://pkg.go.dev/math)
- [Formatting in Go](https://gobyexample.com/string-formatting)

---

## 🎓 Complete Example Walkthrough

**File: example.txt**
```
10
20
30
40
50
```

**Step-by-Step Calculation**:

**1. Create Data Points**:
```
x: [0, 1, 2, 3, 4]
y: [10, 20, 30, 40, 50]
n = 5
```

**2. Calculate Sums**:
```
Σx = 0+1+2+3+4 = 10
Σy = 10+20+30+40+50 = 150
Σ(xy) = 0*10 + 1*20 + 2*30 + 3*40 + 4*50 = 0+20+60+120+200 = 400
Σ(x²) = 0+1+4+9+16 = 30
Σ(y²) = 100+400+900+1600+2500 = 5500
```

**3. Calculate Slope**:
```
numerator = n·Σ(xy) - Σx·Σy
          = 5*400 - 10*150
          = 2000 - 1500
          = 500

denominator = n·Σ(x²) - (Σx)²
            = 5*30 - 10²
            = 150 - 100
            = 50

slope = 500 / 50 = 10.0
```

**4. Calculate Intercept**:
```
x̄ = 10/5 = 2
ȳ = 150/5 = 30

intercept = ȳ - slope·x̄
          = 30 - 10*2
          = 30 - 20
          = 10.0
```

**5. Regression Line**:
```
y = 10.000000x + 10.000000
```

**6. Calculate Pearson**:
```
numerator = 500 (same as slope numerator)

denomX = 50 (same as slope denominator)
denomY = n·Σ(y²) - (Σy)²
       = 5*5500 - 150²
       = 27500 - 22500
       = 5000

denominator = √50 * √5000
            = 7.071 * 70.711
            = 500.0

r = 500 / 500 = 1.0000000000
```

**Final Output**:
```
Linear Regression Line: y = 10.000000x + 10.000000
Pearson Correlation Coefficient: 1.0000000000
```

**Interpretation**: Perfect positive linear relationship!

---

**Remember**: Linear regression is about finding patterns in data. Take time to understand the mathematics behind it, not just implement formulas! 📈📊