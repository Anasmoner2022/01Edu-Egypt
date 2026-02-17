# Groupie-Tracker Project Guide

## 📋 Project Overview
Build a web application that fetches data from an external API about music artists and displays it in an interactive, user-friendly website. The project combines API consumption, data manipulation, web server creation, and dynamic HTML rendering.

---

## 🎯 Learning Objectives

By completing this project, you will learn:
1. **API Consumption**: Fetching and parsing data from external REST APIs
2. **JSON Handling**: Unmarshaling JSON data into Go structs
3. **HTTP Server**: Creating a web server using Go's `net/http` package
4. **HTML Templates**: Dynamic HTML generation using Go templates
5. **Client-Server Architecture**: Understanding request-response cycles
6. **Data Structures**: Organizing and relating complex data
7. **Error Handling**: Graceful error handling for server and API requests
8. **Frontend Basics**: HTML, CSS for creating user interfaces
9. **Routing**: Handling different URL paths and HTTP methods

---

## 📚 Prerequisites - Topics You Must Know

### 1. **Go HTTP Package**
- `net/http` package:
  - `http.HandleFunc()` - Register handlers for routes
  - `http.ListenAndServe()` - Start HTTP server
  - `http.Get()` - Make HTTP requests
  - `http.ResponseWriter` - Write HTTP responses
  - `http.Request` - Handle incoming requests
  - `http.Error()` - Send error responses

### 2. **JSON in Go**
- `encoding/json` package:
  - `json.Unmarshal()` - Parse JSON into structs
  - `json.Marshal()` - Convert structs to JSON
- Struct tags: `json:"fieldname"`
- Handling nested JSON structures

### 3. **HTML Templates**
- `html/template` package:
  - `template.ParseFiles()` - Load HTML templates
  - `template.Execute()` - Render templates with data
  - Template syntax: `{{.}}`, `{{range}}`, `{{if}}`
- Passing data to templates

### 4. **HTTP Concepts**
- HTTP methods: GET, POST
- Status codes: 200, 404, 500
- Request/Response cycle
- URL routing
- Query parameters

### 5. **Data Structures**
- Structs and nested structs
- Slices and maps
- Struct methods

### 6. **HTML & CSS Basics**
- HTML structure: tags, attributes
- CSS styling: selectors, properties
- Responsive design basics
- Forms and input handling

### 7. **Error Handling**
- Error checking patterns
- Custom error messages
- Graceful degradation

---

## 🛠️ Step-by-Step Implementation Guide

### **Phase 1: Understanding the API** 🔍

#### Step 1: Explore the API
Before coding, understand what data is available:

**Visit these endpoints in your browser:**
- Main API: `https://groupietrackers.herokuapp.com/api`
- Artists: `https://groupietrackers.herokuapp.com/api/artists`
- Locations: `https://groupietrackers.herokuapp.com/api/locations`
- Dates: `https://groupietrackers.herokuapp.com/api/dates`
- Relations: `https://groupietrackers.herokuapp.com/api/relation`

**Task**: Write down the structure of each JSON response:
- What fields does each endpoint have?
- What data types are used?
- How are they related?

**Example Analysis**:
```
Artists endpoint returns:
- id (number)
- image (string - URL)
- name (string)
- members (array of strings)
- creationDate (number)
- firstAlbum (string)
- locations (string - URL to locations)
- concertDates (string - URL to dates)
- relations (string - URL to relations)
```

---

### **Phase 2: Project Setup** ✅

#### Step 2: Create Project Structure
```
groupie-tracker/
├── main.go
├── handlers/
│   ├── home.go
│   └── artist.go
├── api/
│   └── fetch.go
├── models/
│   └── structs.go
├── templates/
│   ├── index.html
│   ├── artist.html
│   └── error.html
├── static/
│   ├── css/
│   │   └── style.css
│   └── js/
│       └── script.js
└── go.mod
```

#### Step 3: Initialize Go Module
```bash
go mod init groupie-tracker
```

#### Step 4: Create Basic Main Function
```go
package main

import (
    "fmt"
    "net/http"
)

func main() {
    // Setup routes
    // Start server on port 8080
    fmt.Println("Server starting on http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
```

**Test**: Run and visit `http://localhost:8080` (expect 404 for now).

---

### **Phase 3: Data Models** 📊

#### Step 5: Define Go Structs
In `models/structs.go`, create structs that match the API responses:

**Think about what each struct needs:**

```go
package models

// Artist represents a band/artist
type Artist struct {
    // Add fields based on API response
    // Use json tags: json:"fieldName"
}

// Locations represents concert locations
type Locations struct {
    // Add fields
}

// Dates represents concert dates
type Dates struct {
    // Add fields
}

// Relation links locations and dates
type Relation struct {
    // Add fields
    // This will have a map: location -> []dates
}

// Index holds all locations data
type Index struct {
    // Array of locations
}
```

**Key Points**:
- Field names must be capitalized (exported)
- Use `json:"fieldname"` tags to match API
- Use correct types: int, string, []string, map[string][]string

**Manual Test**: Print the struct definitions to verify structure.

---

### **Phase 4: API Fetching** 🌐

#### Step 6: Create HTTP Request Function
In `api/fetch.go`, create a reusable function:

```go
package api

import (
    "encoding/json"
    "net/http"
)

// FetchData gets data from a URL and unmarshals into target
func FetchData(url string, target interface{}) error {
    // 1. Make HTTP GET request using http.Get()
    // 2. Check for errors
    // 3. Defer close response body
    // 4. Check status code (should be 200)
    // 5. Use json.NewDecoder to decode response into target
    // 6. Return any errors
}
```

**Implementation Steps**:
1. Use `http.Get(url)` to make request
2. Check `err != nil`
3. Defer `resp.Body.Close()`
4. Check if `resp.StatusCode != 200`
5. Use `json.NewDecoder(resp.Body).Decode(target)`
6. Handle errors at each step

**Test**: Create a test function that fetches artists and prints them.

---

#### Step 7: Fetch All API Data
Create functions to fetch each endpoint:

```go
// GetArtists fetches all artists
func GetArtists() ([]models.Artist, error) {
    var artists []models.Artist
    err := FetchData("https://groupietrackers.herokuapp.com/api/artists", &artists)
    return artists, err
}

// GetLocations fetches locations for an artist
func GetLocations(url string) (models.Locations, error) {
    // Similar pattern
}

// GetDates fetches dates for an artist
func GetDates(url string) (models.Dates, error) {
    // Similar pattern
}

// GetRelations fetches relations for an artist
func GetRelations(url string) (models.Relation, error) {
    // Similar pattern
}
```

**Test Each Function**:
```go
artists, err := GetArtists()
if err != nil {
    fmt.Println("Error:", err)
    return
}
fmt.Printf("Fetched %d artists\n", len(artists))
for _, artist := range artists {
    fmt.Println(artist.Name)
}
```

---

### **Phase 5: Web Server Setup** 🌍

#### Step 8: Create Route Handlers
In `handlers/home.go`:

```go
package handlers

import (
    "html/template"
    "net/http"
)

// HomeHandler displays all artists
func HomeHandler(w http.ResponseWriter, r *http.Request) {
    // 1. Check if path is exactly "/"
    // 2. If not, return 404 error
    // 3. Fetch artists from API
    // 4. Handle any errors (return 500)
    // 5. Parse HTML template
    // 6. Execute template with artists data
}
```

**Key Points**:
- Use `r.URL.Path` to check the route
- Use `http.Error()` for error responses
- Parse templates: `template.ParseFiles("templates/index.html")`
- Execute: `tmpl.Execute(w, artists)`

---

#### Step 9: Create Artist Detail Handler
In `handlers/artist.go`:

```go
package handlers

// ArtistHandler displays individual artist details
func ArtistHandler(w http.ResponseWriter, r *http.Request) {
    // 1. Get artist ID from URL (query parameter or path)
    // 2. Fetch all artists
    // 3. Find the artist with matching ID
    // 4. Fetch additional data (locations, dates, relations)
    // 5. Combine data into a struct for template
    // 6. Render artist.html template
}
```

**URL Pattern Options**:
- Query parameter: `/artist?id=1`
- Path parameter: `/artist/1` (requires parsing)

**Implementation Strategy**:
```go
// Get ID from query: r.URL.Query().Get("id")
// Convert to int: strconv.Atoi(idStr)
// Find artist: loop through artists, match ID
// Fetch related data using artist's location/date/relation URLs
```

---

#### Step 10: Setup Routes in Main
In `main.go`:

```go
func main() {
    // Serve static files (CSS, JS)
    fs := http.FileServer(http.Dir("static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))
    
    // Register handlers
    http.HandleFunc("/", handlers.HomeHandler)
    http.HandleFunc("/artist", handlers.ArtistHandler)
    
    // Start server
    fmt.Println("Server running on http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
```

**Test**: Visit each route and ensure no crashes.

---

### **Phase 6: HTML Templates** 🎨

#### Step 11: Create Home Page Template
In `templates/index.html`:

```html
<!DOCTYPE html>
<html>
<head>
    <title>Groupie Tracker - Artists</title>
    <link rel="stylesheet" href="/static/css/style.css">
</head>
<body>
    <h1>Music Artists</h1>
    
    <div class="artists-grid">
        {{range .}}
        <div class="artist-card">
            <img src="{{.Image}}" alt="{{.Name}}">
            <h2>{{.Name}}</h2>
            <p>Created: {{.CreationDate}}</p>
            <a href="/artist?id={{.ID}}">View Details</a>
        </div>
        {{end}}
    </div>
</body>
</html>
```

**Template Syntax**:
- `{{range .}}` - Loop through slice
- `{{.FieldName}}` - Access struct field
- `{{end}}` - End block

---

#### Step 12: Create Artist Detail Template
In `templates/artist.html`:

```html
<!DOCTYPE html>
<html>
<head>
    <title>{{.Artist.Name}} - Details</title>
    <link rel="stylesheet" href="/static/css/style.css">
</head>
<body>
    <a href="/">← Back to Artists</a>
    
    <div class="artist-detail">
        <img src="{{.Artist.Image}}" alt="{{.Artist.Name}}">
        <h1>{{.Artist.Name}}</h1>
        
        <div class="info">
            <p><strong>Members:</strong></p>
            <ul>
                {{range .Artist.Members}}
                <li>{{.}}</li>
                {{end}}
            </ul>
            
            <p><strong>Creation Date:</strong> {{.Artist.CreationDate}}</p>
            <p><strong>First Album:</strong> {{.Artist.FirstAlbum}}</p>
        </div>
        
        <div class="concerts">
            <h2>Concert Dates & Locations</h2>
            {{range $location, $dates := .Relations.DatesLocations}}
            <div class="concert">
                <h3>{{$location}}</h3>
                <ul>
                    {{range $dates}}
                    <li>{{.}}</li>
                    {{end}}
                </ul>
            </div>
            {{end}}
        </div>
    </div>
</body>
</html>
```

**Template Data Structure**:
You need to pass a struct containing all related data:
```go
type ArtistPageData struct {
    Artist    models.Artist
    Relations models.Relation
}
```

---

#### Step 13: Create Error Page Template
In `templates/error.html`:

```html
<!DOCTYPE html>
<html>
<head>
    <title>Error</title>
    <link rel="stylesheet" href="/static/css/style.css">
</head>
<body>
    <div class="error-page">
        <h1>{{.StatusCode}} - {{.StatusText}}</h1>
        <p>{{.Message}}</p>
        <a href="/">Go Home</a>
    </div>
</body>
</html>
```

**Create Error Handler Function**:
```go
type ErrorData struct {
    StatusCode int
    StatusText string
    Message    string
}

func RenderError(w http.ResponseWriter, statusCode int, message string) {
    // Set status code
    // Parse error template
    // Execute with ErrorData
}
```

---

### **Phase 7: Styling** 💅

#### Step 14: Create CSS File
In `static/css/style.css`:

**Basic Structure**:
```css
/* Reset and base styles */
* {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
}

body {
    font-family: Arial, sans-serif;
    /* Add styling */
}

/* Artists grid layout */
.artists-grid {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(250px, 1fr));
    gap: 20px;
    /* Add styling */
}

.artist-card {
    /* Card styling */
    /* Border, shadow, padding */
}

.artist-card img {
    width: 100%;
    /* Image styling */
}

/* Artist detail page */
.artist-detail {
    /* Detail page layout */
}

/* Concert information */
.concerts {
    /* Concert section styling */
}

/* Error page */
.error-page {
    text-align: center;
    /* Error styling */
}
```

**Design Considerations**:
- Mobile-responsive (use media queries)
- Good contrast and readability
- Hover effects on cards
- Professional color scheme

---

### **Phase 8: Error Handling** ⚠️

#### Step 15: Implement Comprehensive Error Handling

**Server Errors to Handle**:
1. API fetch failures
2. Invalid artist ID
3. Template parsing errors
4. Invalid routes (404)
5. Server errors (500)

**Implementation Pattern**:
```go
func HomeHandler(w http.ResponseWriter, r *http.Request) {
    // Check route
    if r.URL.Path != "/" {
        RenderError(w, 404, "Page not found")
        return
    }
    
    // Fetch data
    artists, err := api.GetArtists()
    if err != nil {
        RenderError(w, 500, "Failed to fetch artists data")
        return
    }
    
    // Parse template
    tmpl, err := template.ParseFiles("templates/index.html")
    if err != nil {
        RenderError(w, 500, "Template error")
        return
    }
    
    // Execute template
    err = tmpl.Execute(w, artists)
    if err != nil {
        RenderError(w, 500, "Failed to render page")
        return
    }
}
```

**Test Error Cases**:
- Visit `/invalid-route` → Should show 404
- Stop API (use wrong URL) → Should show 500
- Delete template file → Should show error
- Use invalid artist ID → Should show error

---

### **Phase 9: Client-Server Events** 🔄

#### Step 16: Implement Interactive Feature
Add a feature that requires server communication (one of these):

**Option A: Search/Filter Artists**
```go
// SearchHandler filters artists by name
func SearchHandler(w http.ResponseWriter, r *http.Request) {
    // Get search query from form: r.FormValue("query")
    // Fetch all artists
    // Filter artists by name containing query
    // Return filtered results (JSON or HTML)
}
```

**Option B: Artist Location Map**
- Add button to view locations on a map
- Make AJAX request to server for location data
- Server returns location coordinates as JSON

**Option C: Favorite Artists**
- Add button to mark artists as favorites
- Use POST request to send favorite to server
- Store in memory (or file)
- Display favorites on separate page

**Example: Search Implementation**

**HTML Form** (in index.html):
```html
<form action="/search" method="GET">
    <input type="text" name="query" placeholder="Search artists...">
    <button type="submit">Search</button>
</form>
```

**Handler**:
```go
func SearchHandler(w http.ResponseWriter, r *http.Request) {
    query := r.URL.Query().Get("query")
    
    artists, err := api.GetArtists()
    if err != nil {
        RenderError(w, 500, "Error fetching data")
        return
    }
    
    // Filter artists
    filtered := []models.Artist{}
    for _, artist := range artists {
        if strings.Contains(strings.ToLower(artist.Name), strings.ToLower(query)) {
            filtered = append(filtered, artist)
        }
    }
    
    // Render with filtered results
    tmpl, _ := template.ParseFiles("templates/index.html")
    tmpl.Execute(w, filtered)
}
```

**Register route**: `http.HandleFunc("/search", handlers.SearchHandler)`

---

### **Phase 10: Testing** 🧪

#### Step 17: Create Test Cases

**Manual Testing Checklist**:
- [ ] Home page loads and displays all artists
- [ ] Artist cards show image, name, creation date
- [ ] Clicking artist card opens detail page
- [ ] Detail page shows all artist information
- [ ] Concert locations and dates display correctly
- [ ] Back button returns to home
- [ ] 404 page shows for invalid routes
- [ ] Error page shows when API is unreachable
- [ ] Search/filter feature works
- [ ] CSS styling displays correctly
- [ ] Site works on mobile devices
- [ ] No server crashes under any condition

**Create Unit Tests**:
In `api/fetch_test.go`:
```go
func TestFetchArtists(t *testing.T) {
    artists, err := GetArtists()
    
    if err != nil {
        t.Errorf("Expected no error, got %v", err)
    }
    
    if len(artists) == 0 {
        t.Error("Expected artists, got empty slice")
    }
    
    // Check first artist has required fields
    if artists[0].Name == "" {
        t.Error("Artist name is empty")
    }
}
```

**Load Testing**:
- Open site in multiple browser tabs
- Click rapidly between pages
- Ensure no crashes or slow responses

---

### **Phase 11: Enhancements** ✨

#### Step 18: Add Advanced Features (Optional)

**Feature Ideas**:
1. **Pagination**: Show 10 artists per page with next/previous
2. **Sorting**: Sort by name, creation date, etc.
3. **Advanced Search**: Filter by creation year, member count
4. **Concert Calendar**: Display upcoming concerts in calendar view
5. **Dark Mode**: Toggle between light/dark themes
6. **Export Data**: Download artist info as JSON/CSV
7. **Statistics**: Show charts (most concerts, oldest bands, etc.)

**Example: Pagination Implementation**
```go
func HomeHandler(w http.ResponseWriter, r *http.Request) {
    page := 1
    perPage := 10
    
    // Get page number from query
    if p := r.URL.Query().Get("page"); p != "" {
        page, _ = strconv.Atoi(p)
    }
    
    artists, _ := api.GetArtists()
    
    // Calculate pagination
    start := (page - 1) * perPage
    end := start + perPage
    
    if end > len(artists) {
        end = len(artists)
    }
    
    pageData := struct {
        Artists []models.Artist
        Page    int
        HasNext bool
        HasPrev bool
    }{
        Artists: artists[start:end],
        Page:    page,
        HasNext: end < len(artists),
        HasPrev: page > 1,
    }
    
    // Render template
}
```

---

## 🐛 Common Issues and Solutions

### Issue 1: Template Not Found
**Error**: `template: pattern matches no files`
**Solution**: Check file paths are relative to where you run the program

### Issue 2: CORS Errors
**Error**: Cross-origin request blocked
**Solution**: Not applicable here (same-origin), but good to know for future

### Issue 3: Empty Struct Fields
**Error**: Struct fields are empty after unmarshaling
**Solution**: 
- Check json tags match API response exactly
- Ensure fields are capitalized (exported)
- Print raw JSON to see actual field names

### Issue 4: "404 page not found" for Static Files
**Error**: CSS not loading
**Solution**: 
- Check `http.Handle` for /static/ route
- Verify file paths in HTML
- Ensure FileServer is configured correctly

### Issue 5: Server Crashes on Invalid Input
**Error**: Panic on invalid artist ID
**Solution**: Always validate input and handle errors
```go
idStr := r.URL.Query().Get("id")
id, err := strconv.Atoi(idStr)
if err != nil || id < 1 {
    RenderError(w, 400, "Invalid artist ID")
    return
}
```

---

## ✅ Submission Checklist

**Functionality**:
- [ ] Website displays all artists from API
- [ ] Individual artist pages work
- [ ] All artist information displays correctly
- [ ] Concert locations and dates are linked
- [ ] Client-server event/feature implemented
- [ ] Error handling prevents crashes
- [ ] 404 page for invalid routes
- [ ] Professional design and styling

**Code Quality**:
- [ ] Code follows Go best practices
- [ ] Proper package organization
- [ ] Error handling at every step
- [ ] No hardcoded values (use constants)
- [ ] Comments explain complex logic
- [ ] No unused imports or variables

**Testing**:
- [ ] All routes tested manually
- [ ] Error cases tested
- [ ] Mobile responsive design
- [ ] Unit tests written
- [ ] No memory leaks or goroutine leaks

---

## 📖 Key Go Concepts Used

| Concept | Package/Function | Purpose |
|---------|-----------------|---------|
| HTTP Server | `http.ListenAndServe()` | Start web server |
| Route Handling | `http.HandleFunc()` | Register URL handlers |
| HTTP Requests | `http.Get()` | Fetch API data |
| JSON Parsing | `json.Unmarshal()` | Parse JSON to structs |
| Templates | `html/template` | Dynamic HTML rendering |
| File Server | `http.FileServer()` | Serve static files |
| Error Handling | `http.Error()` | Send error responses |
| Query Parameters | `r.URL.Query().Get()` | Get URL parameters |
| Struct Tags | `json:"field"` | Map JSON fields |

---

## 🎓 Learning Path

**Week 1**: Understand API, create structs, fetch data
**Week 2**: Build basic server, create templates
**Week 3**: Add styling, error handling, client-server feature
**Week 4**: Testing, refinement, documentation

---

## 🚀 Pro Tips

1. **Start with Data**: Understand the API before coding
2. **Test API Calls First**: Make sure you can fetch data before building UI
3. **Use Browser DevTools**: Inspect network requests and responses
4. **Handle Errors Early**: Don't wait until the end
5. **Keep It Simple First**: Get basic functionality working, then enhance
6. **Use Templates Wisely**: Separate logic from presentation
7. **Print Debug Info**: `fmt.Println()` is your friend during development
8. **Check API Status**: API might be down, handle gracefully
9. **Mobile First**: Test on mobile viewport from the start
10. **Read Documentation**: Go docs for net/http and html/template are excellent

---

## 📚 Additional Resources

- [Go Web Development](https://gowebexamples.com/)
- [Go HTTP Package](https://pkg.go.dev/net/http)
- [Go Templates](https://pkg.go.dev/html/template)
- [JSON in Go](https://gobyexample.com/json)
- [REST API Tutorial](https://restfulapi.net/)
- [HTTP Status Codes](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status)

---

## 💡 Extension Ideas

After completing basic requirements:
1. Add user authentication
2. Implement caching to reduce API calls
3. Create an admin panel to add custom artists
4. Build a mobile app using the same backend
5. Add GraphQL layer over REST API
6. Implement WebSocket for real-time updates
7. Create data visualizations (charts, graphs)
8. Add multilingual support

---

**Remember**: This project teaches you full-stack web development fundamentals. Take time to understand each piece, and don't hesitate to experiment! 🎸🎤