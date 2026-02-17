# 🎬 Grokking Go - Chapter 2
## Selection Sort, Arrays & Linked Lists

**Style:** Movie Seats, Memory Drawers, Real-World Examples!

---

## 📖 CHAPTER OVERVIEW

**What you'll learn:**
- How computer memory works (drawer analogy 🗄️)
- Arrays vs Linked Lists (when to use what!)
- Selection Sort algorithm 📊
- Big O for different data structures

**Key Concepts:**
- Random access vs Sequential access
- Insertion/Deletion trade-offs
- Your first sorting algorithm!

---

## 🎯 PROBLEM 1: The Movie Theater Analogy

### Scenario:
You and 4 friends want to see a movie. You need 5 seats TOGETHER.

**Question 1.1:** The theater uses **array-style seating** (everyone sits together). What happens if a 6th friend shows up?

**A)** Easy! Just add one more seat  
**B)** Must find 6 empty seats together elsewhere  
**C)** Split the group  
**D)** Cancel the movie  

<details><summary>💡 Answer</summary>

**B) Must find 6 empty seats together elsewhere**

**Explanation:**
- Arrays need **contiguous memory** (side-by-side)
- If current spot only fits 5, must move EVERYONE
- Like arrays: all elements must be adjacent

**Visual:**
```
Theater Row (Array):
[X][X][YOU][YOU][YOU][YOU][YOU][X][X][X]
         └──────5 friends─────┘

6th friend arrives... FULL on both sides!

Must move to new location:
[  ][YOU][YOU][YOU][YOU][YOU][YOU][  ][  ]
    └──────────6 friends────────────┘
```

**In Go code:**
```go
// Array - fixed size, contiguous memory
friends := [5]string{"Alice", "Bob", "Carol", "Dave", "Eve"}

// Want to add Frank?
// Can't! Array size is fixed
// Must create new array and copy everyone
newFriends := [6]string{}
copy(newFriends[:], friends[:])
newFriends[5] = "Frank"
```

**Cost:** O(n) - must copy all elements!

</details>

---

**Question 1.2:** Now imagine **linked list seating** (friends scattered, each knows next friend's seat). 6th friend arrives. What happens?

**A)** Still must move everyone  
**B)** Just update one "pointer" to next seat  
**C)** Impossible  
**D)** Same as arrays  

<details><summary>💡 Answer</summary>

**B) Just update one "pointer" to next seat!**

**Explanation:**
- Linked lists don't need contiguous space
- Each person just knows where the NEXT person sits
- Adding someone = update one link

**Visual:**
```
Before:
Alice(B3) → Bob(F7) → Carol(A12) → Dave(C2) → Eve(H5) → END
  ↑            ↑          ↑            ↑          ↑
Seat B3     Seat F7    Seat A12     Seat C2   Seat H5

After adding Frank at seat J9:
Alice(B3) → Bob(F7) → Carol(A12) → Dave(C2) → Eve(H5) → Frank(J9) → END
                                                           ↑
                                                        NEW!
```

**In Go code:**
```go
type Node struct {
    name string
    next *Node
}

// Add Frank to end
frank := &Node{name: "Frank", next: nil}
eve.next = frank  // Just update one pointer!
```

**Cost:** O(1) - constant time! ⚡

</details>

---

## 🎯 PROBLEM 2: Memory as Drawers

### Scenario:
Computer memory is like a giant chest of drawers. Each drawer has an address.

**Question 2.1:** You want to store your to-do list: [Buy milk, Walk dog, Code]. Using an **array**, where are these stored?

**A)** Anywhere in memory  
**B)** In consecutive drawers (e.g., 100, 101, 102)  
**C)** In random drawers  
**D)** On hard drive  

<details><summary>💡 Answer</summary>

**B) In consecutive drawers!**

**Explanation:**
- Arrays use **contiguous memory**
- Items are stored side-by-side
- Makes reading fast (simple math to find location)

**Visual:**
```
Memory (Drawers):
┌────┬────┬────┬────┬────┬────┬────┐
│ 99 │100 │101 │102 │103 │104 │105 │
├────┼────┼────┼────┼────┼────┼────┤
│    │Buy │Walk│Code│    │    │    │
│    │milk│ dog│    │    │    │    │
└────┴────┴────┴────┴────┴────┴────┘
      ↑    ↑    ↑
    array elements (consecutive!)
```

**Why this matters:**
```go
// If array starts at address 100
// And each string takes 1 "drawer"

arr[0] // address 100
arr[1] // address 101 (100 + 1)
arr[2] // address 102 (100 + 2)

// Simple math! O(1) access
```

</details>

---

**Question 2.2:** Using a **linked list**, where are items stored?

**A)** Consecutively  
**B)** Anywhere, each item points to next  
**C)** Only at beginning of memory  
**D)** Must be sorted  

<details><summary>💡 Answer</summary>

**B) Anywhere, each item points to next!**

**Explanation:**
- Linked lists don't need consecutive space
- Each node stores VALUE + ADDRESS of next node
- Like a treasure hunt!

**Visual:**
```
Memory (Drawers):
┌────┬────┬────┬────┬────┬────┬────┐
│ 50 │ 73 │100 │101 │150 │200 │201 │
├────┼────┼────┼────┼────┼────┼────┤
│    │Buy │    │    │Walk│Code│    │
│    │→150│    │    │→200│→NIL│    │
└────┴────┴────┴────┴────┴────┴────┘
      ↑              ↑    ↑
   node 1         node 2 node 3
   (@ 73)         (@ 150)(@ 200)
```

**In Go:**
```go
type TodoNode struct {
    task string
    next *TodoNode  // Points to next node
}

// Can be anywhere in memory!
node1 := &TodoNode{task: "Buy milk", next: nil}
node2 := &TodoNode{task: "Walk dog", next: nil}
node3 := &TodoNode{task: "Code", next: nil}

node1.next = node2  // Link them
node2.next = node3
```

</details>

---

## 🎯 PROBLEM 3: Arrays vs Linked Lists - Big O

### Comparison Table Question

**Question 3.1:** Complete this table:

| Operation | Array | Linked List |
|-----------|-------|-------------|
| Reading   | ?     | ?           |
| Insertion | ?     | ?           |
| Deletion  | ?     | ?           |

<details><summary>💡 Answer</summary>

| Operation | Array  | Linked List |
|-----------|--------|-------------|
| Reading   | O(1)   | O(n)        |
| Insertion | O(n)   | O(1)        |
| Deletion  | O(n)   | O(1)        |

**Detailed Explanation:**

**Reading (Random Access):**
```
Array: O(1)
arr[5] = address_start + (5 * element_size)
Simple math → instant!

Linked List: O(n)
Must traverse from head:
head → node1 → node2 → ... → node5
Could be anywhere → must follow chain
```

**Insertion (at beginning or middle):**
```
Array: O(n)
Insert at index 2? Must shift everything right:
[A][B][C][D] → [A][B][NEW][C][D]
               Shift C, D → slow!

Linked List: O(1)
Just update pointers:
A → B → C
A → B → NEW → C
     ↑      ↑
   2 pointer changes, that's it!
```

**Deletion:**
```
Array: O(n)
Delete index 2? Must shift left:
[A][B][C][D] → [A][B][D]
             Shift D left

Linked List: O(1)
Skip over node:
A → B → C → D
A → B ─────→ D
    (C removed, no shifting!)
```

</details>

---

**Question 3.2:** You're building Facebook. You need to insert new posts frequently and read them in order. Which data structure?

**A)** Array (insertion is O(n))  
**B)** Linked List (insertion is O(1))  
**C)** Doesn't matter  
**D)** Use a tree  

<details><summary>💡 Answer</summary>

**B) Linked List!**

**Reasoning:**
- **Lots of insertions** (people post constantly)
- Reading is sequential (scroll through feed)
- Don't need random access (not jumping to post #58,392)

**Why NOT array:**
- Inserting new post = O(n) (might need to shift millions)
- Linked list insertion = O(1) ⚡

**Real-world pattern:**
```go
type Post struct {
    content  string
    timestamp time.Time
    next     *Post
}

// Adding new post to feed
func AddPost(head *Post, content string) *Post {
    newPost := &Post{
        content: content,
        timestamp: time.Now(),
        next: head,  // Point to old head
    }
    return newPost  // New post is now head
}
// O(1) - super fast! No matter how many posts exist
```

**Facebook's actual approach:** Even more complex (hybrid structures), but linked lists for sequential data!

</details>

---

## 🎯 PROBLEM 4: Selection Sort Introduction

### Scenario:
You have a playlist of songs. You want to sort them from most to least played.

**Songs:** 
- "Go Gopher" (50 plays)
- "Pointers Forever" (120 plays)
- "Recursion Song" (80 plays)
- "Slice n' Dice" (200 plays)

**Question 4.1:** Selection sort finds the SMALLEST (or largest) element repeatedly. First step: which song has MOST plays?

**A)** Go Gopher  
**B)** Pointers Forever  
**C)** Recursion Song  
**D)** Slice n' Dice  

<details><summary>💡 Answer</summary>

**D) Slice n' Dice (200 plays)**

**How Selection Sort works:**
```
Step 1: Find maximum in list
Songs: [50, 120, 80, 200]
                      ↑
                    MAX!

Step 2: Remove it, add to sorted list
Sorted: [200]
Remaining: [50, 120, 80]

Step 3: Find max in remaining
Remaining: [50, 120, 80]
                  ↑
                 MAX!

Step 4: Remove, add to sorted
Sorted: [200, 120]
Remaining: [50, 80]

... continue until done
Final: [200, 120, 80, 50]
```

</details>

---

**Question 4.2:** How many times must you scan through the list to sort 4 songs?

**A)** 1  
**B)** 2  
**C)** 4  
**D)** 16  

<details><summary>💡 Answer</summary>

**C) 4 times!**

**Detailed Steps:**
```
n = 4 songs

Round 1: Check all 4 songs → find max
Round 2: Check remaining 3 → find max
Round 3: Check remaining 2 → find max
Round 4: Check remaining 1 → that's it!

Total checks: 4 + 3 + 2 + 1 = 10
```

**General formula:**
- For n items: n + (n-1) + (n-2) + ... + 1
- This is: n × (n+1) / 2 operations
- In Big O: **O(n²)**

**Why n²:**
```
Outer loop: n iterations (find each max)
Inner loop: average n/2 comparisons
Total: n × n/2 = n²/2 ≈ O(n²)
```

</details>

---

## 🎯 PROBLEM 5: Implementing Selection Sort

**Question 5.1:** Complete this Selection Sort implementation:

```go
func findSmallest(arr []int) int {
    smallest := arr[0]
    smallestIndex := 0
    
    for i := 1; i < len(arr); i++ {
        if arr[i] ___ smallest {  // A: fill in
            smallest = arr[i]
            smallestIndex = ___  // B: fill in
        }
    }
    return ___  // C: return what?
}
```

<details><summary>💡 Answer</summary>

```go
func findSmallest(arr []int) int {
    smallest := arr[0]
    smallestIndex := 0
    
    for i := 1; i < len(arr); i++ {
        if arr[i] < smallest {  // A: less than
            smallest = arr[i]
            smallestIndex = i   // B: current index
        }
    }
    return smallestIndex  // C: return INDEX, not value
}
```

**Why return index?**
- Need to remove element from original array
- Can't remove by value (might have duplicates)
- Need position to delete!

**Example:**
```go
arr := []int{5, 3, 6, 2, 10}
idx := findSmallest(arr)  // Returns 3 (index of 2)
smallest := arr[idx]      // Now get value: 2
```

</details>

---

**Question 5.2:** Complete the full Selection Sort:

```go
func selectionSort(arr []int) []int {
    newArr := []int{}
    
    for len(arr) > 0 {
        smallest := findSmallest(arr)
        newArr = ___  // A: add element
        arr = ___     // B: remove element
    }
    
    return newArr
}
```

<details><summary>💡 Answer</summary>

```go
func selectionSort(arr []int) []int {
    newArr := []int{}
    
    for len(arr) > 0 {
        smallest := findSmallest(arr)
        newArr = append(newArr, arr[smallest])  // A
        arr = append(arr[:smallest], arr[smallest+1:]...)  // B
    }
    
    return newArr
}
```

**Explanation:**

**A:** Append smallest to result
```go
newArr = append(newArr, arr[smallest])
// Adds element at index 'smallest' to new array
```

**B:** Remove from original
```go
arr = append(arr[:smallest], arr[smallest+1:]...)
// arr[:smallest] = everything before
// arr[smallest+1:] = everything after
// ... = spread operator
// Combines them, skipping arr[smallest]
```

**Example run:**
```go
arr := []int{5, 3, 6, 2, 10}

Round 1: smallest=3 (value 2)
  newArr: [2]
  arr: [5, 3, 6, 10]

Round 2: smallest=1 (value 3)
  newArr: [2, 3]
  arr: [5, 6, 10]

Round 3: smallest=0 (value 5)
  newArr: [2, 3, 5]
  arr: [6, 10]

Round 4: smallest=0 (value 6)
  newArr: [2, 3, 5, 6]
  arr: [10]

Round 5: smallest=0 (value 10)
  newArr: [2, 3, 5, 6, 10]
  arr: []

Done!
```

</details>

---

## 🎯 PROBLEM 6: Big O of Selection Sort

**Question 6.1:** For 5 elements, how many operations does selection sort perform?

**Hint:** Count comparisons in each round.

<details><summary>💡 Answer</summary>

**15 operations!**

**Breakdown:**
```
Round 1: Compare 5 elements → 4 comparisons
Round 2: Compare 4 elements → 3 comparisons
Round 3: Compare 3 elements → 2 comparisons
Round 4: Compare 2 elements → 1 comparison
Round 5: 1 element left → 0 comparisons

Total: 4 + 3 + 2 + 1 = 10 comparisons
Plus 5 removals = 15 operations
```

**General formula:**
```
For n elements:
(n-1) + (n-2) + ... + 1 = n(n-1)/2

For n=5: 5×4/2 = 10
```

</details>

---

**Question 6.2:** What's the Big O of selection sort for n elements?

**A)** O(n)  
**B)** O(n log n)  
**C)** O(n²)  
**D)** O(1)  

<details><summary>💡 Answer</summary>

**C) O(n²)**

**Why:**
```
Operations: n × (n-1) / 2 = (n² - n) / 2

Big O simplification:
- Drop constants: /2 doesn't matter
- Drop non-dominant terms: -n is tiny compared to n²
- Result: O(n²)
```

**Visual comparison:**
```
List size:     10      100      1,000    10,000
O(n):          10      100      1,000    10,000
O(n²):        100    10,000  1,000,000  100,000,000
               ↑        ↑          ↑           ↑
         Selection sort gets SLOW quickly!
```

**This is why selection sort is rarely used in practice!**

</details>

---

## 🎯 PROBLEM 7: When to Use Arrays vs Lists

### Scenario-based decisions

**Question 7.1:** Building a music player. Operations:
- Add songs to playlist (occasionally)
- Skip to track #25 (frequently)
- Shuffle (reads all tracks randomly)

Which structure?

**A)** Array  
**B)** Linked List  
**C)** Both work equally  
**D)** Neither  

<details><summary>💡 Answer</summary>

**A) Array!**

**Reasoning:**
- **Skip to track #25** → Need O(1) random access
- **Shuffle** → Need to access tracks randomly
- Adding songs is occasional → O(n) insertion OK

**Why NOT linked list:**
```go
// Array: Jump to any track instantly
tracks[24]  // Track #25 (0-indexed) → O(1)

// Linked List: Must traverse
head → track1 → track2 → ... → track25
     24 hops! → O(n)
```

**Real code:**
```go
type Playlist struct {
    tracks []Song  // Array/slice for random access
}

func (p *Playlist) PlayTrack(index int) {
    // O(1) - instant!
    player.Play(p.tracks[index])
}

func (p *Playlist) Shuffle() {
    // Can access all tracks randomly
    rand.Shuffle(len(p.tracks), func(i, j int) {
        p.tracks[i], p.tracks[j] = p.tracks[j], p.tracks[i]
    })
}
```

</details>

---

**Question 7.2:** Building Instagram feed. Operations:
- Insert new posts at top (constantly)
- Users scroll sequentially (never jump)
- Delete old posts (occasionally)

Which structure?

**A)** Array  
**B)** Linked List  
**C)** Both work  
**D)** Hash table  

<details><summary>💡 Answer</summary>

**B) Linked List!**

**Reasoning:**
- **Insert at top constantly** → Need O(1) insertion
- **Sequential scrolling** → Don't need random access
- **Delete posts** → O(1) with linked list

**Why NOT array:**
```go
// Array: Insert at beginning = disaster
// Must shift MILLIONS of posts
posts := make([]Post, 1000000)
// Insert new post at index 0?
// → Copy all 1,000,000 posts → O(n) 😱

// Linked List: Just update head
newPost.next = head
head = newPost  // Done! O(1) ⚡
```

**Real pattern (simplified Instagram):**
```go
type Post struct {
    image   string
    likes   int
    next    *Post
}

type Feed struct {
    head *Post
}

func (f *Feed) AddPost(image string) {
    newPost := &Post{
        image: image,
        next:  f.head,  // Point to old head
    }
    f.head = newPost  // New post becomes head
    // O(1) - instant! Even with millions of posts
}

func (f *Feed) ScrollFeed() {
    current := f.head
    for current != nil {
        displayPost(current)
        current = current.next  // Sequential access
    }
}
```

</details>

---

## 🎯 PROBLEM 8: Hybrid Approach

**Question 8.1:** You're building a contact list app:
- Store 1000 contacts
- Search by name frequently (binary search!)
- Add new contacts rarely

Which structure and why?

<details><summary>💡 Answer</summary>

**Array + Keep Sorted!**

**Strategy:**
1. Use **array** for storage (O(1) random access)
2. **Keep sorted** for binary search
3. Adding contact is rare, so O(n) insertion OK

**Implementation:**
```go
type ContactList struct {
    contacts []Contact  // Sorted array
}

// Search: O(log n) with binary search
func (c *ContactList) Find(name string) *Contact {
    return binarySearch(c.contacts, name)
}

// Add: O(n) but rare
func (c *ContactList) Add(contact Contact) {
    c.contacts = append(c.contacts, contact)
    sort.Slice(c.contacts, func(i, j int) bool {
        return c.contacts[i].name < c.contacts[j].name
    })
}
```

**Trade-off analysis:**
```
Operations frequency:
- Search: 1000 times/day → Must be fast → O(log n) ✓
- Add: 10 times/month → Can be slow → O(n) OK

Binary search only works on sorted data!
So maintain sorted array despite slower insertions.
```

</details>

---

## 🎯 PROBLEM 9: Real-World Facebook Problem

**Scenario:** Facebook has 2 billion users.

**Question 9.1:** User names stored in:
- **Option A:** Sorted array
- **Option B:** Linked list

Finding your friend "Alice":
- Option A takes how many steps?
- Option B takes how many steps?

<details><summary>💡 Answer</summary>

**Option A: ~31 steps (binary search)**  
**Option B: Up to 2 billion steps (sequential)**

**Calculations:**
```
Option A (Sorted Array):
Binary search: log₂(2,000,000,000) ≈ 31
Can jump directly to middle, eliminate half

Option B (Linked List):
Must traverse from beginning:
head → user1 → user2 → ... → Alice
Could be at position 1,999,999,999!
Average: 1 billion steps
Worst: 2 billion steps
```

**Why Facebook uses arrays (simplified):**
```go
type UserDatabase struct {
    users []User  // Sorted by name
}

// Find user in 31 steps max!
func (db *UserDatabase) FindUser(name string) *User {
    return binarySearch(db.users, name)
    // O(log n) → 31 steps for 2B users!
}
```

**This is why data structure choice MATTERS!**
- Array + Binary Search: Instant
- Linked List: Wait minutes! 🐌

</details>

---

## 📝 RECAP & KEY TAKEAWAYS

### ✅ Arrays:

**Pros:**
- ✅ O(1) random access (super fast reads!)
- ✅ Simple, cache-friendly
- ✅ Can use binary search (if sorted)

**Cons:**
- ❌ O(n) insertions (must shift elements)
- ❌ O(n) deletions (must shift)
- ❌ Need contiguous memory (might be unavailable)

**Use when:**
- Random access needed
- Few insertions/deletions
- Size known in advance

---

### ✅ Linked Lists:

**Pros:**
- ✅ O(1) insertions (just update pointers!)
- ✅ O(1) deletions
- ✅ No need for contiguous memory

**Cons:**
- ❌ O(n) access (must traverse)
- ❌ Can't use binary search
- ❌ Extra memory for pointers

**Use when:**
- Lots of insertions/deletions
- Sequential access only
- Size unknown/dynamic

---

### ✅ Selection Sort:

**How it works:**
1. Find smallest element
2. Remove it, add to sorted list
3. Repeat until done

**Big O:** O(n²)
- Slow for large lists
- Good learning algorithm
- Stepping stone to Quicksort!

**Code pattern:**
```go
func selectionSort(arr []int) []int {
    sorted := []int{}
    for len(arr) > 0 {
        smallest := findSmallest(arr)
        sorted = append(sorted, arr[smallest])
        arr = remove(arr, smallest)
    }
    return sorted
}
```

---

## 🎓 EXERCISES (Try Yourself!)

### Exercise 1:
Implement selection sort that sorts in DESCENDING order (largest first).

### Exercise 2:
You have 100 students in a class. You need to:
- Add students (frequently)
- Find student by ID (rarely)
Which structure? Why?

### Exercise 3:
Modify selection sort to return how many comparisons it made.

### Exercise 4:
Build a simple music playlist with:
- Add song
- Remove song
- Play next
Choose appropriate data structure and implement.

---

## 🔥 CHALLENGE PROBLEM

**The Hybrid Playlist:**

Build a smart playlist that:
1. Stores 10,000 songs in an array (for quick access by index)
2. Maintains a "recently played" linked list (max 10 songs)
3. Implements "shuffle" efficiently

**Requirements:**
- `PlayTrack(index int)` → O(1)
- `AddToRecent(song Song)` → O(1)
- `GetRecent()` → O(1) for last 10

<details><summary>💡 Solution Approach</summary>

```go
type SmartPlaylist struct {
    allSongs      []Song      // Array for random access
    recentHead    *RecentNode // Linked list for recent
    recentCount   int
}

type RecentNode struct {
    song *Song
    next *RecentNode
}

func (p *SmartPlaylist) PlayTrack(index int) {
    song := &p.allSongs[index]  // O(1) access
    p.AddToRecent(song)
    player.Play(song)
}

func (p *SmartPlaylist) AddToRecent(song *Song) {
    // Add to front of linked list
    newNode := &RecentNode{
        song: song,
        next: p.recentHead,
    }
    p.recentHead = newNode
    p.recentCount++
    
    // Keep only 10
    if p.recentCount > 10 {
        // Remove last (traverse to 9th, set next to nil)
        current := p.recentHead
        for i := 0; i < 8; i++ {
            current = current.next
        }
        current.next = nil
        p.recentCount = 10
    }
}

func (p *SmartPlaylist) Shuffle() {
    rand.Shuffle(len(p.allSongs), func(i, j int) {
        p.allSongs[i], p.allSongs[j] = p.allSongs[j], p.allSongs[i]
    })
}
```

**Why this works:**
- Array for main storage → fast random access for shuffle/play
- Linked list for recent → fast inserts at head
- Best of both worlds! 🎉

</details>

---

**🎉 Congratulations! You now understand Arrays, Linked Lists, and Selection Sort!**

**Next:** Chapter 3 - Recursion & The Call Stack 🔁