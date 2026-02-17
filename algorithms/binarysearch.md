# 🔍 Grokking Go - Chapter 1
## Binary Search & Big O Notation

**Style:** Practical, Visual, Real-World Examples (Like Grokking Algorithms!)

---

## 📖 CHAPTER OVERVIEW

**What you'll learn:**
- How to search efficiently (Binary Search)
- How to measure algorithm speed (Big O)
- Why some algorithms are MUCH faster than others
- Real-world applications

**Key Concepts:**
- Simple Search vs Binary Search
- O(log n) vs O(n)
- When to use which algorithm

---

## 🎯 PROBLEM 1: The Phone Book Challenge

### Scenario:
You have a phone book with **1,000,000 names** (sorted alphabetically).

**Question 1.1:** Using **simple search** (checking each name one by one), what's the MAXIMUM number of steps to find a name?

**A)** 500,000  
**B)** 1,000,000  
**C)** 10,000  
**D)** 20  

<details><summary>💡 Answer</summary>

**B) 1,000,000**

**Explanation:**
- Simple search checks EVERY item one by one
- Worst case: name is at the END or not in book
- Maximum steps = total number of names = 1,000,000

**Visual:**
```
Names: [Aaron ... Zoe]
        ↑           ↑
     Start        End (worst case)
```

**Code (Simple Search):**
```go
func SimpleSearch(list []string, target string) int {
    for i := 0; i < len(list); i++ {
        if list[i] == target {
            return i  // Found!
        }
    }
    return -1  // Not found
}
// Worst case: checks ALL 1,000,000 names
```

</details>

---

**Question 1.2:** Using **binary search**, what's the MAXIMUM number of steps to find a name in 1,000,000 entries?

**A)** 1,000,000  
**B)** 500,000  
**C)** 100  
**D)** 20  

<details><summary>💡 Answer</summary>

**D) 20**

**Explanation:**
- Binary search cuts the list in HALF each time
- 1,000,000 → 500,000 → 250,000 → ... → 1
- log₂(1,000,000) ≈ 20 steps!

**How it works:**
```
Step 1: Check middle (500,000th name)
        Too low? Search right half
        Too high? Search left half

Step 2: Check middle of that half (250,000 items)

Step 3: Check middle of THAT half (125,000 items)

... keep halving until found!
```

**That's 50,000 times FASTER!** 🚀

**Code (Binary Search):**
```go
func BinarySearch(list []string, target string) int {
    low := 0
    high := len(list) - 1
    
    for low <= high {
        mid := (low + high) / 2
        guess := list[mid]
        
        if guess == target {
            return mid  // Found!
        }
        if guess > target {
            high = mid - 1  // Search left half
        } else {
            low = mid + 1   // Search right half
        }
    }
    return -1  // Not found
}
```

</details>

---

**Question 1.3:** If the phone book **doubles** to 2,000,000 names, how many MORE steps does binary search need?

**A)** 1,000,000 more steps  
**B)** 10 more steps  
**C)** Only 1 more step!  
**D)** Same number of steps  

<details><summary>💡 Answer</summary>

**C) Only 1 more step!**

**Explanation:**
- 1,000,000 names = 20 steps
- 2,000,000 names = 21 steps
- **Doubling the data adds just ONE step!**

**Why?**
```
log₂(1,000,000) = 20
log₂(2,000,000) = 21

One extra "cut" handles the doubled size!
```

**This is the POWER of logarithmic growth!** 📈

</details>

---

## 🎯 PROBLEM 2: Big O Notation

### Scenario:
You're comparing different search algorithms.

**Question 2.1:** Simple search has **O(n)** complexity. What does this mean?

**A)** It takes n seconds  
**B)** It might check up to n items  
**C)** It's faster than binary search  
**D)** It uses n memory  

<details><summary>💡 Answer</summary>

**B) It might check up to n items**

**Explanation:**
- **O(n)** means "linear time"
- For n items, worst case checks **all n**
- If n = 100, checks up to 100
- If n = 1,000,000, checks up to 1,000,000

**Visual:**
```
Items:        10      100     1,000   1,000,000
Max Steps:    10      100     1,000   1,000,000
              ↑        ↑        ↑         ↑
            Grows linearly with size!
```

**Key:** Big O describes HOW time grows with input size

</details>

---

**Question 2.2:** Binary search has **O(log n)** complexity. What does this mean?

**A)** Slower than O(n)  
**B)** Time grows logarithmically  
**C)** Checks every item  
**D)** Always takes log steps  

<details><summary>💡 Answer</summary>

**B) Time grows logarithmically**

**Explanation:**
- Each step cuts problem in HALF
- Adding data barely increases time!

**Comparison:**
```
Items:        128     256     1,024   1,000,000
O(n):         128     256     1,024   1,000,000
O(log n):       7       8        10          20
              
Binary is MUCH faster! 🚀
```

**Real Impact:**
- 128 items: 128 vs 7 steps (18x faster)
- 1M items: 1M vs 20 steps (50,000x faster!)

</details>

---

**Question 2.3:** Arrange these Big O complexities from FASTEST to SLOWEST:

**Options:** O(n²), O(1), O(n), O(log n), O(n log n)

<details><summary>💡 Answer</summary>

**From FASTEST to SLOWEST:**

1. **O(1)** - Constant (always same time)
2. **O(log n)** - Logarithmic (binary search)
3. **O(n)** - Linear (simple search)
4. **O(n log n)** - Log-linear (good sorting)
5. **O(n²)** - Quadratic (slow sorting)

**Visual (for n=8):**
```
O(1):      █ 1 step
O(log n):  ███ 3 steps  
O(n):      ████████ 8 steps
O(n log n):████████████████████████ 24 steps
O(n²):     ████████████████████████████████████████████████████████████████ 64 steps
```

**Remember:** Lower on list = slower!

</details>

---

## 🎯 PROBLEM 3: Real-World Applications

### Scenario:
You're building different features for an app.

**Question 3.1:** You need to find a user in a **sorted** list of 1 million usernames. Which algorithm?

**A)** Simple search  
**B)** Binary search  
**C)** Check every username  
**D)** Give up  

<details><summary>💡 Answer</summary>

**B) Binary search**

**Why:**
- ✅ List is **sorted** (required for binary search!)
- ✅ Need to find specific item
- ✅ 1M items → simple search too slow

**Code:**
```go
func FindUser(users []string, username string) int {
    // users is SORTED alphabetically
    low := 0
    high := len(users) - 1
    
    for low <= high {
        mid := (low + high) / 2
        
        if users[mid] == username {
            return mid
        }
        if users[mid] > username {
            high = mid - 1
        } else {
            low = mid + 1
        }
    }
    return -1
}

// Takes only ~20 steps for 1M users!
```

</details>

---

**Question 3.2:** You need to check if a number is EVEN or ODD. What's the Big O?

**A)** O(n)  
**B)** O(log n)  
**C)** O(1)  
**D)** O(n²)  

<details><summary>💡 Answer</summary>

**C) O(1)** - Constant time!

**Explanation:**
- Checking even/odd is ONE operation: `n % 2`
- Doesn't matter if n = 5 or n = 1,000,000
- **Always takes same time** → O(1)

**Code:**
```go
func IsEven(n int) bool {
    return n % 2 == 0  // One operation, always!
}

// O(1) - Lightning fast! ⚡
```

**Other O(1) examples:**
- Accessing array element: `arr[5]`
- Basic math: `a + b`
- Comparison: `x > 5`

</details>

---

**Question 3.3:** You need to print ALL items in a list. What's the Big O?

**A)** O(1)  
**B)** O(log n)  
**C)** O(n)  
**D)** O(n²)  

<details><summary>💡 Answer</summary>

**C) O(n)** - Linear time

**Explanation:**
- Must visit EVERY item to print it
- 10 items = 10 prints
- 1M items = 1M prints
- Time grows linearly with size

**Code:**
```go
func PrintAll(list []int) {
    for i := 0; i < len(list); i++ {
        fmt.Println(list[i])  // Must print each one
    }
}
// O(n) - can't avoid checking each item
```

</details>

---

## 🎯 PROBLEM 4: Algorithm Comparison

### Scenario:
Compare different approaches to same problem.

**Question 4.1:** You have a **sorted** list of numbers [1, 3, 5, 7, 9, 11, 13, 15]. Find the number 7.

**Simple Search Steps:** Count how many comparisons needed.

<details><summary>💡 Answer</summary>

**4 comparisons**

**Step-by-step:**
```
List: [1, 3, 5, 7, 9, 11, 13, 15]
       ↑
Step 1: Check 1 → Not 7, continue

List: [1, 3, 5, 7, 9, 11, 13, 15]
          ↑
Step 2: Check 3 → Not 7, continue

List: [1, 3, 5, 7, 9, 11, 13, 15]
             ↑
Step 3: Check 5 → Not 7, continue

List: [1, 3, 5, 7, 9, 11, 13, 15]
                ↑
Step 4: Check 7 → FOUND! ✓
```

**Total: 4 steps**

</details>

---

**Question 4.2:** Same list, find 7 using **binary search**. How many comparisons?

<details><summary>💡 Answer</summary>

**2 comparisons**

**Step-by-step:**
```
List: [1, 3, 5, 7, 9, 11, 13, 15]
                   ↑
Step 1: Check middle (index 4) → 9
        7 < 9, so search LEFT half

Left: [1, 3, 5, 7]
             ↑
Step 2: Check middle (index 2 of left) → 5
        Wait, recalculate: middle of [1,3,5,7] is index 1 or 2
        Let's say we check 5 first
        7 > 5, search RIGHT
        
Actually, let me recalculate properly:

Initial: low=0, high=7
Mid = 3 (value = 7) FOUND!

Wait, let me be more careful:
[1, 3, 5, 7, 9, 11, 13, 15]
 0  1  2  3  4   5   6   7

Step 1: low=0, high=7, mid=(0+7)/2=3
        list[3] = 7 ✓ FOUND!
```

**Total: 1 comparison!** Even faster! 🚀

</details>

---

**Question 4.3:** For a list of 16 items, what's the MAXIMUM steps for binary search?

**A)** 2  
**B)** 4  
**C)** 8  
**D)** 16  

<details><summary>💡 Answer</summary>

**B) 4 steps**

**Explanation:**
- log₂(16) = 4
- Each step halves: 16 → 8 → 4 → 2 → 1

**Visual:**
```
Step 1: 16 items → check 1, left with 8
Step 2:  8 items → check 1, left with 4
Step 3:  4 items → check 1, left with 2
Step 4:  2 items → check 1, left with 1
Step 5:  1 item  → found or not found!

Maximum = 5 steps actually... let me recalculate

Actually: log₂(16) = 4
- After 4 cuts, we're down to 1 item
- That's the answer: 4 steps
```

</details>

---

## 🎯 PROBLEM 5: When Binary Search FAILS

### Important Concept!

**Question 5.1:** Can you use binary search on this list: [5, 2, 9, 1, 7, 3]?

**A)** Yes  
**B)** No  
**C)** Only for number 5  
**D)** Only if we add more items  

<details><summary>💡 Answer</summary>

**B) No!**

**Why:** List is **NOT SORTED**!

**Explanation:**
- Binary search **requires sorted data**
- Works by assuming: left < middle < right
- Unsorted data breaks this assumption

**What happens if we try:**
```go
list := []int{5, 2, 9, 1, 7, 3}
// Looking for 1

Mid = 9 (index 2)
1 < 9, so search left: [5, 2]
  Mid = 2
  1 < 2, search left: [5]
  Check 5... not 1!
  
MISSED IT! (1 was actually at index 3, on the right)
```

**Solution:** Sort first!
```go
list := []int{1, 2, 3, 5, 7, 9}  // Sorted
// NOW binary search works!
```

</details>

---

**Question 5.2:** You need to find the LARGEST number in an unsorted list. What's the best approach?

**A)** Binary search  
**B)** Sort, then binary search  
**C)** Check each number  
**D)** Impossible  

<details><summary>💡 Answer</summary>

**C) Check each number** - O(n)

**Explanation:**
- Can't use binary search (unsorted)
- Could sort first, but that takes O(n log n)
- Finding max only needs ONE pass!

**Code:**
```go
func FindMax(list []int) int {
    if len(list) == 0 {
        return 0  // or error
    }
    
    max := list[0]
    for i := 1; i < len(list); i++ {
        if list[i] > max {
            max = list[i]
        }
    }
    return max
}
// O(n) - must check each item
```

**Why not sort first?**
- Sorting = O(n log n)
- One pass = O(n)
- O(n) < O(n log n) → faster!

</details>

---

## 🎯 PROBLEM 6: Implementation Challenge

**Question 6.1:** Complete this binary search implementation:

```go
func BinarySearch(list []int, target int) int {
    low := 0
    high := len(list) - 1
    
    for low <= high {
        mid := (low + high) / 2
        guess := list[mid]
        
        if guess == target {
            return ___  // A
        }
        if guess > target {
            ___ = mid - 1  // B
        } else {
            ___ = mid + 1  // C
        }
    }
    return ___  // D
}
```

**Fill in A, B, C, D:**

<details><summary>💡 Answer</summary>

```go
func BinarySearch(list []int, target int) int {
    low := 0
    high := len(list) - 1
    
    for low <= high {
        mid := (low + high) / 2
        guess := list[mid]
        
        if guess == target {
            return mid  // A - return index where found
        }
        if guess > target {
            high = mid - 1  // B - search left half
        } else {
            low = mid + 1  // C - search right half
        }
    }
    return -1  // D - not found
}
```

**Explanation:**
- **A:** Return index when found
- **B:** If guess too high, search lower (left)
- **C:** If guess too low, search higher (right)
- **D:** Loop ended without finding → return -1

</details>

---

## 🎯 PROBLEM 7: Big O in Real Life

**Question 7.1:** Facebook has 2 billion users. If you search for a friend by name (sorted alphabetically), how many steps maximum?

**A)** 2 billion  
**B)** 1 billion  
**C)** 1 million  
**D)** About 31  

<details><summary>💡 Answer</summary>

**D) About 31 steps!**

**Calculation:**
```
log₂(2,000,000,000) ≈ 31

2 billion → 1 billion → 500M → 250M → ...
... → 1

Only 31 cuts needed!
```

**Mind-blowing:**
- Simple search: 2 BILLION steps
- Binary search: 31 steps
- **65 million times faster!** 🤯

</details>

---

**Question 7.2:** You're checking if your username is available. The website has 100 million usernames (sorted). Binary search finds it in how many steps?

**A)** 100 million  
**B)** 50 million  
**C)** 1 million  
**D)** About 27  

<details><summary>💡 Answer</summary>

**D) About 27 steps**

**Calculation:**
```
log₂(100,000,000) ≈ 26.6 ≈ 27

100M → 50M → 25M → 12.5M → ...
```

**Instant feedback!** ⚡

**Code in action:**
```go
usernames := []string{...} // 100M sorted usernames

func IsAvailable(username string) bool {
    index := BinarySearch(usernames, username)
    return index == -1  // -1 means not found = available!
}
// Takes ~27 comparisons = milliseconds!
```

</details>

---

## 🎯 PROBLEM 8: Comparison Table

**Question 8.1:** Complete this table:

| List Size | Simple Search | Binary Search |
|-----------|---------------|---------------|
| 8         | 8             | ?             |
| 1,024     | ?             | 10            |
| ?         | 1,000,000     | 20            |

<details><summary>💡 Answer</summary>

| List Size | Simple Search | Binary Search |
|-----------|---------------|---------------|
| 8         | 8             | 3             |
| 1,024     | 1,024         | 10            |
| 1,000,000 | 1,000,000     | 20            |

**Calculations:**
- log₂(8) = 3
- Simple search always = n
- 2²⁰ = 1,048,576 ≈ 1,000,000

</details>

---

## 📝 RECAP & KEY TAKEAWAYS

### ✅ What You Should Know:

1. **Simple Search (O(n)):**
   - Checks each item one by one
   - Works on any list
   - Slow for large lists

2. **Binary Search (O(log n)):**
   - Cuts list in half each time
   - **Requires sorted list!**
   - MUCH faster for large lists

3. **Big O Notation:**
   - Describes how time grows with input
   - O(1) < O(log n) < O(n) < O(n²)
   - Focuses on worst-case scenario

### 💡 When to Use What:

**Use Simple Search when:**
- List is unsorted
- List is very small (< 100 items)
- Finding max/min (one pass needed)

**Use Binary Search when:**
- List is sorted
- List is large (> 100 items)
- Searching for specific value

### 🚀 Real-World Impact:

```
Algorithm:      Simple    Binary
───────────────────────────────────
100 users:        100        7
1K users:       1,000       10
1M users:   1,000,000       20
1B users: 1,000,000,000     30

Binary search scales AMAZINGLY!
```

---

## 🎓 EXERCISES (Try Yourself!)

### Exercise 1:
Implement binary search for strings (alphabetically sorted list).

### Exercise 2:
You have 4 billion phone numbers (sorted). Calculate max steps for binary search.

### Exercise 3:
Write a function that finds the first and last occurrence of a number in a sorted list.

### Exercise 4:
Given an unsorted list, when is it worth sorting first before searching multiple times?

---

## 🔥 CHALLENGE PROBLEM

**The Dictionary Problem:**

You have a dictionary app with 170,000 words (sorted alphabetically).

**Part A:** User searches for "algorithm". Maximum steps?

**Part B:** User searches for words starting with "z". Can binary search help find the START of the "z" section?

**Part C:** Auto-complete suggests words while typing. After typing "alg", you need to show all words starting with "alg". Best approach?

<details><summary>💡 Solutions</summary>

**Part A:** log₂(170,000) ≈ 18 steps

**Part B:** Yes! Binary search can find first word starting with "z", then iterate from there.

```go
// Find first "z" word
func FindFirstZ(dict []string) int {
    // Use modified binary search
    low, high := 0, len(dict)-1
    result := -1
    
    for low <= high {
        mid := (low + high) / 2
        if dict[mid][0] == 'z' {
            result = mid
            high = mid - 1  // Keep searching left for FIRST 'z'
        } else if dict[mid][0] < 'z' {
            low = mid + 1
        } else {
            high = mid - 1
        }
    }
    return result
}
```

**Part C:** Binary search to find first word with "alg" prefix, then collect all matching words linearly.

</details>

---

**🎉 Congratulations! You've mastered Binary Search and Big O Notation!**

**Next:** Chapter 2 - Selection Sort & Arrays vs Linked Lists