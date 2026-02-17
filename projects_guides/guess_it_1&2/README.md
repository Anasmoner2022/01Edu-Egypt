# Guess-It-1 Project Guide

## 📋 Project Overview
Build a predictive program that reads numbers from standard input one at a time and predicts a range for the next number. The program combines statistical analysis with prediction algorithms to balance accuracy and range size. This is a streaming/online learning problem where you update predictions as new data arrives.

---

## 🎯 Learning Objectives

By completing this project, you will learn:
1. **Predictive Modeling**: Using statistics to predict future values
2. **Online/Streaming Algorithms**: Processing data one item at a time
3. **Statistical Prediction**: Moving averages, trends, and standard deviation
4. **Performance Optimization**: Efficient algorithms for real-time processing
5. **Balancing Trade-offs**: Accuracy vs. range size optimization
6. **Standard Input/Output**: Reading from stdin, writing to stdout
7. **Shell Scripting**: Creating executable scripts for automation
8. **Testing & Scoring**: Understanding evaluation metrics

---

## 📚 Prerequisites - Topics You Must Know

### 1. **Math-Skills Foundation**
You should have completed the math-skills project and understand:
- Average (Mean)
- Median
- Variance
- Standard Deviation

### 2. **Streaming Input**
- Reading from standard input (stdin)
- `bufio.Scanner` for line-by-line input
- Continuous reading until EOF
- Writing to standard output (stdout)

### 3. **Running Statistics**
- Updating average without re-summing all values
- Maintaining count and sum
- Computing variance incrementally
- Keeping history of values

### 4. **Data Structures**
- Slices for storing history
- Queues (FIFO) for moving windows
- Circular buffers (optional)

### 5. **Prediction Concepts**
- Trend detection (increasing, decreasing, stable)
- Moving averages (simple, weighted)
- Confidence intervals
- Linear regression basics

### 6. **Shell Scripting**
- Creating executable scripts
- Shebang (`#!/bin/sh`)
- Running programs from scripts
- File permissions (`chmod +x`)

---

## 🎲 Understanding the Problem

### **What You're Building**
```
Input Stream:  189 → 113 → 121 → 114 → 145 → 110 → ...
Your Program:  
  Reads: 189
  Predicts: "120 200" (range for next number)
  
  Reads: 113  (was 113 in range [120, 200]? NO - bad prediction)
  Predicts: "160 230" (range for next number)
  
  Reads: 121  (was 121 in range [160, 230]? NO - bad prediction)
  Predicts: "110 140" (range for next number)
  
  Reads: 114  (was 114 in range [110, 140]? YES - good prediction!)
  Predicts: "100 200" (range for next number)
  ...
```

### **Key Insights**
1. **You don't know the next number** - you're predicting it
2. **You only have past numbers** - use history to predict future
3. **Smaller range = higher score** - if you're correct
4. **Wrong prediction = zero points** - accuracy matters
5. **Find the balance** - not too small (miss), not too large (low score)

---

## 📊 Prediction Strategies

### **Strategy 1: Mean ± Standard Deviation**
**Concept**: Next number will be within X standard deviations of the mean

```
Algorithm:
1. Calculate mean of all previous numbers
2. Calculate standard deviation
3. Range = [mean - k*stddev, mean + k*stddev]
   where k is a multiplier (e.g., 2 or 3)
```

**Pros**: Simple, mathematically sound
**Cons**: Assumes normal distribution, no trend detection

**Example**:
```
Data: [189, 113, 121]
Mean: 141
StdDev: 38.7
Range (k=2): [141 - 2*38.7, 141 + 2*38.7] = [63, 219]
```

---

### **Strategy 2: Moving Average + Variance**
**Concept**: Use recent data more than old data

```
Algorithm:
1. Keep last N numbers (e.g., N=10)
2. Calculate average of last N
3. Calculate standard deviation of last N
4. Range = [avg - k*stddev, avg + k*stddev]
```

**Pros**: Adapts to trends, responsive to changes
**Cons**: Needs tuning of window size N

**Example**:
```
Last 5 numbers: [113, 121, 114, 145, 110]
Mean: 120.6
StdDev: 14.8
Range (k=2): [91, 150]
```

---

### **Strategy 3: Trend-Aware Prediction**
**Concept**: Detect if data is increasing, decreasing, or stable

```
Algorithm:
1. Calculate slope of recent data (linear regression)
2. Predict next value = current + slope
3. Add confidence interval based on variance
4. Range = [predicted - margin, predicted + margin]
```

**Pros**: Follows trends, more accurate
**Cons**: More complex, needs more data

**Example**:
```
Data: [100, 110, 120, 130]  (increasing by 10)
Predicted next: 140
Margin: ±20
Range: [120, 160]
```

---

### **Strategy 4: Adaptive Range**
**Concept**: Adjust range based on recent prediction accuracy

```
Algorithm:
1. Start with conservative range
2. If last predictions were correct: decrease range
3. If last predictions were wrong: increase range
4. Balance between risk and reward
```

**Pros**: Self-correcting, optimizes over time
**Cons**: Needs careful tuning

---

## 🛠️ Step-by-Step Implementation Guide

### **Phase 1: Project Setup** ✅

#### Step 1: Create Project Structure
```
guess-it-1/
├── student/
│   ├── main.go (or your language)
│   └── script.sh
├── go.mod
└── test_data/
    └── sample.txt
```

#### Step 2: Create Basic Input/Output
```go
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    
    // Read first number
    scanner.Scan()
    firstNum := parseFloat(scanner.Text())
    
    fmt.Println("lower upper")  // Print range for next number
    
    // Continue reading numbers
    for scanner.Scan() {
        num := parseFloat(scanner.Text())
        // Process number and predict next range
        fmt.Println("lower upper")
    }
}
```

**Test**: 
```bash
echo -e "189\n113\n121" | go run main.go
```

---

### **Phase 2: Data Storage** 📊

#### Step 3: Store Historical Data
Create a structure to maintain history:

```go
type Predictor struct {
    history []float64  // All previous numbers
    count   int        // How many numbers seen
}

func NewPredictor() *Predictor {
    return &Predictor{
        history: []float64{},
        count:   0,
    }
}

func (p *Predictor) AddNumber(num float64) {
    p.history = append(p.history, num)
    p.count++
}
```

**Key Points**:
- Store every number received
- Maintain count for efficiency
- Can limit history size if needed (e.g., last 100 numbers)

---

#### Step 4: Calculate Running Statistics
Implement efficient statistics:

```go
func (p *Predictor) CalculateMean() float64 {
    if p.count == 0 {
        return 0
    }
    sum := 0.0
    for _, v := range p.history {
        sum += v
    }
    return sum / float64(p.count)
}

func (p *Predictor) CalculateStdDev() float64 {
    if p.count < 2 {
        return 0
    }
    mean := p.CalculateMean()
    variance := 0.0
    for _, v := range p.history {
        diff := v - mean
        variance += diff * diff
    }
    variance /= float64(p.count)
    return math.Sqrt(variance)
}
```

**Optimization Tip**: 
For large datasets, maintain running sum and sum of squares:
```go
type Predictor struct {
    history    []float64
    count      int
    sum        float64  // Running sum
    sumSquares float64  // Running sum of squares
}
```

---

### **Phase 3: Basic Prediction** 🔮

#### Step 5: Implement Simple Range Prediction
Start with Strategy 1 (Mean ± StdDev):

```go
func (p *Predictor) PredictRange() (lower, upper float64) {
    // Need at least 2 numbers to calculate stddev
    if p.count < 2 {
        // First number: return wide range
        return 0, 200
    }
    
    mean := p.CalculateMean()
    stddev := p.CalculateStdDev()
    
    // Use 2 standard deviations (covers ~95% of normal distribution)
    margin := 2 * stddev
    
    lower = mean - margin
    upper = mean + margin
    
    // Ensure range is at least some minimum width
    if upper - lower < 10 {
        lower = mean - 5
        upper = mean + 5
    }
    
    return lower, upper
}
```

**Test with Sample Data**:
```
Input: 189
Predict: 0 200 (no history yet)

Input: 113
Mean: 189
Predict: Wide range (not enough data)

Input: 121
Mean: (189+113)/2 = 151
StdDev: 38
Predict: [151-76, 151+76] = [75, 227]
```

---

### **Phase 4: Improved Predictions** 📈

#### Step 6: Implement Moving Average
Use recent data for better predictions:

```go
func (p *Predictor) GetRecentNumbers(n int) []float64 {
    if n > p.count {
        n = p.count
    }
    start := p.count - n
    return p.history[start:]
}

func (p *Predictor) PredictRangeMoving(windowSize int) (lower, upper float64) {
    if p.count < 2 {
        return 0, 200
    }
    
    // Use last N numbers
    recent := p.GetRecentNumbers(windowSize)
    
    // Calculate mean of recent
    sum := 0.0
    for _, v := range recent {
        sum += v
    }
    mean := sum / float64(len(recent))
    
    // Calculate stddev of recent
    variance := 0.0
    for _, v := range recent {
        diff := v - mean
        variance += diff * diff
    }
    variance /= float64(len(recent))
    stddev := math.Sqrt(variance)
    
    // Predict range
    margin := 2 * stddev
    lower = mean - margin
    upper = mean + margin
    
    return lower, upper
}
```

**Window Size Selection**:
- Small (5-10): Very responsive, but can be unstable
- Medium (15-20): Good balance
- Large (30+): Smooth, but slow to adapt

**Test Different Window Sizes**:
```go
// Try window sizes: 5, 10, 15, 20
// Measure which gives best score on test data
```

---

#### Step 7: Detect and Follow Trends
Add trend detection:

```go
func (p *Predictor) CalculateTrend() float64 {
    if p.count < 3 {
        return 0
    }
    
    // Simple trend: average change between consecutive numbers
    recent := p.GetRecentNumbers(5)
    if len(recent) < 2 {
        return 0
    }
    
    changes := 0.0
    for i := 1; i < len(recent); i++ {
        changes += recent[i] - recent[i-1]
    }
    
    return changes / float64(len(recent)-1)
}

func (p *Predictor) PredictRangeTrend() (lower, upper float64) {
    if p.count < 2 {
        return 0, 200
    }
    
    // Get last value and trend
    lastValue := p.history[p.count-1]
    trend := p.CalculateTrend()
    
    // Predict next value
    predicted := lastValue + trend
    
    // Calculate uncertainty
    stddev := p.CalculateStdDev()
    margin := 2 * stddev
    
    lower = predicted - margin
    upper = predicted + margin
    
    return lower, upper
}
```

**Example**:
```
Recent: [100, 110, 120, 130, 140]
Trend: +10 per step
Last: 140
Predicted: 140 + 10 = 150
StdDev: 15.8
Range: [150 - 31.6, 150 + 31.6] = [118, 182]
```

---

### **Phase 5: Optimization** ⚡

#### Step 8: Tune Parameters
Create configurable parameters:

```go
type Config struct {
    WindowSize      int     // For moving average
    StdDevMultiplier float64 // How many stddevs (1, 2, 3)
    MinRangeSize    float64 // Minimum range width
    MaxRangeSize    float64 // Maximum range width
}

func (p *Predictor) PredictRangeOptimized(cfg Config) (lower, upper float64) {
    if p.count < 2 {
        return 0, 200
    }
    
    // Use moving window
    recent := p.GetRecentNumbers(cfg.WindowSize)
    
    // Calculate statistics
    mean := calculateMean(recent)
    stddev := calculateStdDev(recent, mean)
    
    // Apply multiplier
    margin := cfg.StdDevMultiplier * stddev
    
    // Predict range
    lower = mean - margin
    upper = mean + margin
    
    // Enforce constraints
    rangeSize := upper - lower
    if rangeSize < cfg.MinRangeSize {
        center := (lower + upper) / 2
        lower = center - cfg.MinRangeSize/2
        upper = center + cfg.MinRangeSize/2
    }
    if rangeSize > cfg.MaxRangeSize {
        center := (lower + upper) / 2
        lower = center - cfg.MaxRangeSize/2
        upper = center + cfg.MaxRangeSize/2
    }
    
    return lower, upper
}
```

**Parameter Tuning Process**:
1. Test with Data 1, 2, 3
2. Try different combinations:
   - WindowSize: 5, 10, 15, 20
   - StdDevMultiplier: 1.5, 2.0, 2.5, 3.0
3. Measure score for each combination
4. Choose best performing configuration

---

#### Step 9: Adaptive Strategy
Adjust based on recent performance:

```go
type Predictor struct {
    history       []float64
    count         int
    recentMisses  int     // How many recent predictions were wrong
    currentMargin float64 // Current multiplier
}

func (p *Predictor) CheckPrediction(actual, lower, upper float64) bool {
    inRange := actual >= lower && actual <= upper
    
    if inRange {
        // Good prediction - can tighten range
        p.recentMisses = 0
        p.currentMargin *= 0.95 // Decrease by 5%
        if p.currentMargin < 1.5 {
            p.currentMargin = 1.5
        }
    } else {
        // Bad prediction - widen range
        p.recentMisses++
        if p.recentMisses > 2 {
            p.currentMargin *= 1.1 // Increase by 10%
            if p.currentMargin > 3.0 {
                p.currentMargin = 3.0
            }
        }
    }
    
    return inRange
}

func (p *Predictor) PredictRangeAdaptive() (lower, upper float64) {
    // Use currentMargin instead of fixed 2.0
    mean := p.CalculateMean()
    stddev := p.CalculateStdDev()
    margin := p.currentMargin * stddev
    
    lower = mean - margin
    upper = mean + margin
    
    return lower, upper
}
```

---

### **Phase 6: Integration** 🔗

#### Step 10: Complete Main Program
```go
package main

import (
    "bufio"
    "fmt"
    "math"
    "os"
    "strconv"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    predictor := NewPredictor()
    
    // Configuration (tune these values)
    config := Config{
        WindowSize:       15,
        StdDevMultiplier: 2.0,
        MinRangeSize:     10,
        MaxRangeSize:     400,
    }
    
    // Read and process each number
    for scanner.Scan() {
        line := scanner.Text()
        num, err := strconv.ParseFloat(line, 64)
        if err != nil {
            continue
        }
        
        // Predict range for NEXT number
        lower, upper := predictor.PredictRangeOptimized(config)
        
        // Print prediction
        fmt.Printf("%d %d\n", int(math.Round(lower)), int(math.Round(upper)))
        
        // Add this number to history
        predictor.AddNumber(num)
    }
}
```

---

### **Phase 7: Shell Script** 📜

#### Step 11: Create Executable Script
In `student/script.sh`:

```bash
#!/bin/sh
# Script to run the guess-it-1 program
# Must be executable: chmod +x script.sh

# For Go:
go run ./student/main.go

# For Python:
# python3 ./student/solution.py

# For JavaScript:
# node ./student/solution.js

# For Rust:
# cd ./student && cargo run --release
```

**Make it executable**:
```bash
chmod +x student/script.sh
```

**Test the script**:
```bash
# From root directory
echo -e "189\n113\n121\n114" | ./student/script.sh
```

---

### **Phase 8: Testing & Scoring** 🎯

#### Step 12: Understanding the Scoring System

**Score Calculation**:
```
For each prediction:
  if actual is in [lower, upper]:
    score += 1 / (upper - lower)
  else:
    score += 0

Total Score = sum of all scores
```

**Example**:
```
Prediction: [100, 200], Actual: 150
  In range: YES
  Range size: 100
  Score: 1 / 100 = 0.01

Prediction: [140, 160], Actual: 150
  In range: YES
  Range size: 20
  Score: 1 / 20 = 0.05  (BETTER! Smaller range)

Prediction: [170, 180], Actual: 150
  In range: NO
  Score: 0  (MISS! No points)
```

**Key Insight**: 
- Small range + correct = high score
- Large range + correct = low score
- Any range + wrong = zero score

---

#### Step 13: Test with Provided Data
Download the tester and test:

```bash
# Download tester
wget https://assets.01-edu.org/guess-it/guess-it-dockerized.zip
unzip guess-it-dockerized.zip

# Place your student/ folder in tester root
cp -r student/ guess-it-dockerized/

# Run tests
cd guess-it-dockerized
./run_tests.sh
```

**Test Datasets**:
- Data 1: First test dataset
- Data 2: Second test dataset
- Data 3: Third test dataset

**Analyze Results**:
```
Data 1 Score: 23.45
Data 2 Score: 19.12
Data 3 Score: 31.78
Average: 24.78

Check which predictions were wrong
Adjust parameters and re-test
```

---

### **Phase 9: Optimization Loop** 🔄

#### Step 14: Iterative Improvement Process

**Step-by-step optimization**:

1. **Establish Baseline**:
   ```go
   // Start with simple mean ± 2*stddev
   // Record score on all 3 datasets
   ```

2. **Try Moving Average**:
   ```go
   // Test window sizes: 5, 10, 15, 20, 25
   // Record which gives best score
   ```

3. **Adjust Multiplier**:
   ```go
   // Test multipliers: 1.5, 2.0, 2.5, 3.0
   // Find balance between accuracy and range size
   ```

4. **Add Trend Detection**:
   ```go
   // Test trend-aware prediction
   // Compare with moving average
   ```

5. **Combine Strategies**:
   ```go
   // Use different strategies for different data patterns
   // E.g., trend for increasing data, mean for stable data
   ```

**Testing Script**:
```bash
#!/bin/bash
# test_all.sh

for window in 5 10 15 20; do
  for mult in 1.5 2.0 2.5 3.0; do
    echo "Testing: window=$window, mult=$mult"
    # Update config in code
    # Run tests
    # Record scores
  done
done
```

---

## 🐛 Common Issues and Solutions

### Issue 1: Range Always Too Small
**Problem**: Many predictions are wrong
**Solution**:
- Increase standard deviation multiplier
- Use larger window size
- Add minimum range size constraint

### Issue 2: Range Always Too Large
**Problem**: Low scores despite being correct
**Solution**:
- Decrease standard deviation multiplier
- Use trend detection
- Implement adaptive narrowing

### Issue 3: Program Crashes on First Input
**Problem**: Division by zero or empty history
**Solution**:
- Handle first 1-2 numbers specially
- Return safe default range (e.g., [0, 200])
- Check count before calculations

### Issue 4: Slow Performance
**Problem**: Timeout on large datasets
**Solution**:
- Use running statistics (don't recalculate from scratch)
- Limit history size (e.g., last 100 numbers)
- Optimize calculations

### Issue 5: Script Not Executable
**Problem**: Permission denied
**Solution**:
```bash
chmod +x student/script.sh
```

---

## 📋 Testing Checklist

**Basic Functionality**:
- [ ] Reads from standard input
- [ ] Outputs two numbers (lower upper) separated by space
- [ ] Processes numbers continuously until EOF
- [ ] Handles first few numbers gracefully
- [ ] No crashes or errors

**Prediction Quality**:
- [ ] Most predictions include actual value
- [ ] Range sizes are reasonable (not too wide)
- [ ] Adapts to data patterns
- [ ] Handles different data types (stable, trending, random)

**Performance**:
- [ ] Processes input quickly (<1ms per number)
- [ ] Uses efficient algorithms
- [ ] No memory leaks

**Submission Requirements**:
- [ ] `student/` folder exists
- [ ] All required files in `student/`
- [ ] `script.sh` is executable
- [ ] Script runs from tester root directory
- [ ] Tested with official tester

---

## 📊 Performance Benchmarks

**Target Scores** (approximate):
- Beginner: 15-20 points per dataset
- Intermediate: 20-30 points per dataset
- Advanced: 30-40 points per dataset
- Expert: 40+ points per dataset

**Factors Affecting Score**:
1. **Accuracy**: How often you're right
2. **Range Size**: Smaller is better (if correct)
3. **Data Pattern**: Some datasets are easier than others
4. **Strategy**: More sophisticated algorithms score higher

---

## ✅ Submission Checklist

**Code Quality**:
- [ ] Clean, readable code
- [ ] Comments explain logic
- [ ] No hardcoded values (use config)
- [ ] Error handling
- [ ] Efficient algorithms

**Functionality**:
- [ ] Reads stdin correctly
- [ ] Outputs correct format
- [ ] Handles all edge cases
- [ ] Tested on all 3 datasets
- [ ] Scores are competitive

**Submission**:
- [ ] `student/` folder created
- [ ] All source files in `student/`
- [ ] `script.sh` present and executable
- [ ] Script path is correct
- [ ] Tested with official tester
- [ ] README.md explaining approach (optional but recommended)

---

## 📖 Statistical Concepts Reference

### **Normal Distribution**
- 68% of data within 1 standard deviation of mean
- 95% of data within 2 standard deviations
- 99.7% of data within 3 standard deviations

### **Moving Average**
```
Simple Moving Average (SMA):
SMA = (x₁ + x₂ + ... + xₙ) / n

Weighted Moving Average (WMA):
WMA = (w₁*x₁ + w₂*x₂ + ... + wₙ*xₙ) / (w₁ + w₂ + ... + wₙ)
```

### **Linear Regression** (for trend)
```
Slope = Σ((xᵢ - x̄)(yᵢ - ȳ)) / Σ((xᵢ - x̄)²)
Where x = position, y = value
```

### **Confidence Interval**
```
CI = mean ± (z * stddev / √n)
Where z depends on confidence level:
  90% confidence: z = 1.645
  95% confidence: z = 1.96
  99% confidence: z = 2.576
```

---

## 🚀 Pro Tips

1. **Start Simple**: Get basic mean ± stddev working first
2. **Use Test Data**: Run tester frequently during development
3. **Log Predictions**: Save predictions to file to analyze mistakes
4. **Visualize Data**: Plot data to understand patterns
5. **Test Edge Cases**: First numbers, small datasets, large jumps
6. **Balance Risk**: Better to have slightly wider range than miss
7. **Profile Performance**: Ensure no bottlenecks
8. **Document Approach**: Explain your strategy in comments
9. **Iterate**: Test → Analyze → Improve → Repeat
10. **Compare Strategies**: Keep multiple versions to compare

---

## 💡 Advanced Strategies

### **1. Ensemble Prediction**
Combine multiple strategies:
```go
func EnsemblePredict() (lower, upper float64) {
    // Get predictions from different strategies
    l1, u1 := PredictMovingAvg()
    l2, u2 := PredictTrend()
    l3, u3 := PredictMedian()
    
    // Take the union (widest range covering all)
    lower = min(l1, l2, l3)
    upper = max(u1, u2, u3)
    
    return lower, upper
}
```

### **2. Pattern Detection**
Detect and handle different patterns:
```go
func DetectPattern() string {
    // Calculate variance of differences
    // Low variance: stable
    // High positive: increasing
    // High negative: decreasing
    // High both: volatile
}
```

### **3. Outlier Detection**
Remove outliers before prediction:
```go
func RemoveOutliers(data []float64) []float64 {
    // Calculate Q1, Q3
    // IQR = Q3 - Q1
    // Remove values outside [Q1 - 1.5*IQR, Q3 + 1.5*IQR]
}
```

### **4. Kalman Filter** (Advanced)
Use Kalman filter for optimal prediction with noise

---

## 📚 Additional Resources

**Statistics**:
- [Moving Average Explained](https://www.investopedia.com/terms/m/movingaverage.asp)
- [Standard Deviation and Prediction Intervals](https://en.wikipedia.org/wiki/Prediction_interval)
- [Linear Regression](https://www.youtube.com/watch?v=zPG4NjIkCjc)

**Go Resources**:
- [Reading from Stdin](https://gobyexample.com/reading-files)
- [Bufio Scanner](https://pkg.go.dev/bufio#Scanner)
- [Math Package](https://pkg.go.dev/math)

**Optimization**:
- [Algorithm Complexity](https://www.bigocheatsheet.com/)
- [Profiling Go Programs](https://go.dev/blog/pprof)

---

## 🎓 Example Walkthrough

**Sample Data**: [189, 113, 121, 114, 145, 110]

**Step-by-step Predictions**:

**Iteration 1**:
```
Input: 189
History: [189]
Count: 1
Prediction: [0, 200] (no history)
```

**Iteration 2**:
```
Input: 113
History: [189, 113]
Count: 2
Mean: 151
StdDev: 38
Prediction: [151 - 76, 151 + 76] = [75, 227]
Actual next (121): In range? YES
Score: 1/152 = 0.0066
```

**Iteration 3**:
```
Input: 121
History: [189, 113, 121]
Count: 3
Mean: 141
StdDev: 38.7
Prediction: [141 - 77.4, 141 + 77.4] = [64, 218]
Actual next (114): In range? YES
Score: 1/154 = 0.0065
```

**Iteration 4** (with moving window = 2):
```
Input: 114
Recent: [113, 121]
Mean: 117
StdDev: 4
Prediction: [117 - 8, 117 + 8] = [109, 125]
Actual next (145): In range? NO
Score: 0
```

**Key Observation**: Narrow range missed! Need to adjust multiplier.

---

**Remember**: This project is about finding the optimal balance between accuracy and precision. Experiment, test, and iterate! 🎯📊