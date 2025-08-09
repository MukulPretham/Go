# Go Module Dependencies: Complete Guide

This guide covers the three main approaches to handling module dependencies in Go, with practical examples for each.

## Overview

Go modules can be imported in three ways:
1. **Published/Internet Modules** - Fully automatic
2. **Local Modules with Replace Directives** - Manual setup, then automatic
3. **Local Modules with Go Workspaces** - Modern approach (Go 1.18+)

---

## 1. Published/Internet Modules (Fully Automatic)

### How it works:
- Import modules published on GitHub, GitLab, or other repositories
- Go automatically downloads and manages dependencies
- Zero manual configuration required

### Example:

**Project Structure:**
```
my-project/
├── go.mod
└── main.go
```

**go.mod:**
```go
module my-project

go 1.24.3
```

**main.go:**
```go
package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/gorilla/mux"
)

func main() {
    // Using gin
    r := gin.Default()
    
    // Using gorilla/mux
    router := mux.NewRouter()
    
    fmt.Println("Dependencies loaded automatically!")
}
```

**Commands:**
```bash
# Just run go mod tidy - dependencies are automatically added
go mod tidy

# Run your code
go run main.go
```

**Result - go.mod after `go mod tidy`:**
```go
module my-project

go 1.24.3

require (
    github.com/gin-gonic/gin v1.9.1
    github.com/gorilla/mux v1.8.0
)

require (
    // ... indirect dependencies added automatically
)
```

### Advantages:
- ✅ Completely automatic
- ✅ Version management handled by Go
- ✅ Works everywhere (CI/CD, team members)
- ✅ Semantic versioning support

---

## 2. Local Modules - Replace Directives

### How it works:
- Use `replace` directive to redirect module paths to local directories
- Manually add the replace directive, then `go mod tidy` handles the rest
- Traditional approach, works with all Go versions

### Example:

**Project Structure:**
```
my-project/
├── auth/
│   ├── go.mod
│   └── auth.go
├── api/
│   ├── go.mod
│   └── main.go
└── database/
    ├── go.mod
    └── db.go
```

**auth/go.mod:**
```go
module company.com/auth

go 1.24.3
```

**auth/auth.go:**
```go
package auth

import "fmt"

func Login(username string) string {
    return fmt.Sprintf("User %s logged in", username)
}

func Logout(username string) string {
    return fmt.Sprintf("User %s logged out", username)
}
```

**database/go.mod:**
```go
module company.com/database

go 1.24.3
```

**database/db.go:**
```go
package database

import "fmt"

func Connect() string {
    return "Database connected"
}

func Query(sql string) string {
    return fmt.Sprintf("Executing: %s", sql)
}
```

**api/go.mod (before setup):**
```go
module company.com/api

go 1.24.3
```

**api/main.go:**
```go
package main

import (
    "fmt"
    "company.com/auth"
    "company.com/database"
)

func main() {
    // Use auth module
    result := auth.Login("john")
    fmt.Println(result)
    
    // Use database module
    conn := database.Connect()
    fmt.Println(conn)
    
    query := database.Query("SELECT * FROM users")
    fmt.Println(query)
}
```

**Setup Commands:**
```bash
cd api

# Add replace directives
go mod edit -replace company.com/auth=../auth
go mod edit -replace company.com/database=../database

# Let Go find and add the dependencies
go mod tidy
```

**api/go.mod (after setup):**
```go
module company.com/api

go 1.24.3

replace company.com/auth => ../auth
replace company.com/database => ../database

require (
    company.com/auth v0.0.0-00010101000000-000000000000
    company.com/database v0.0.0-00010101000000-000000000000
)
```

**Run the code:**
```bash
go run main.go
```

### Advantages:
- ✅ Works with all Go versions
- ✅ Explicit dependency mapping
- ✅ Good for CI/CD pipelines
- ✅ Clear in go.mod what's local vs remote

### Disadvantages:
- ❌ Manual setup required
- ❌ Clutters go.mod files
- ❌ Must manage replace directives for each module

---

## 3. Local Modules - Go Workspaces (Modern Approach)

### How it works:
- Create a `go.work` file that lists all local modules
- Go automatically discovers and uses local modules in the workspace
- No replace directives needed
- Available since Go 1.18

### Example:

**Project Structure:**
```
my-workspace/
├── go.work
├── frontend/
│   ├── go.mod
│   └── main.go
├── backend/
│   ├── api/
│   │   ├── go.mod
│   │   └── api.go
│   └── auth/
│       ├── go.mod
│       └── auth.go
└── shared/
    ├── utils/
    │   ├── go.mod
    │   └── utils.go
    └── models/
        ├── go.mod
        └── models.go
```

**Setup the workspace:**
```bash
cd my-workspace

# Initialize workspace with all modules
go work init \
    ./frontend \
    ./backend/api \
    ./backend/auth \
    ./shared/utils \
    ./shared/models
```

**go.work (created automatically):**
```go
go 1.24.3

use (
    ./frontend
    ./backend/api
    ./backend/auth
    ./shared/utils
    ./shared/models
)
```

**shared/utils/go.mod:**
```go
module company.com/utils

go 1.24.3
```

**shared/utils/utils.go:**
```go
package utils

import "strings"

func FormatName(name string) string {
    return strings.Title(strings.ToLower(name))
}

func ValidateEmail(email string) bool {
    return strings.Contains(email, "@")
}
```

**backend/auth/go.mod:**
```go
module company.com/auth

go 1.24.3
```

**backend/auth/auth.go:**
```go
package auth

import (
    "fmt"
    "company.com/utils"
)

func CreateUser(name, email string) string {
    if !utils.ValidateEmail(email) {
        return "Invalid email"
    }
    
    formattedName := utils.FormatName(name)
    return fmt.Sprintf("User created: %s (%s)", formattedName, email)
}
```

**frontend/go.mod:**
```go
module company.com/frontend

go 1.24.3
```

**frontend/main.go:**
```go
package main

import (
    "fmt"
    "company.com/auth"
    "company.com/utils"
)

func main() {
    // Use utils directly
    formatted := utils.FormatName("john DOE")
    fmt.Println("Formatted name:", formatted)
    
    // Use auth (which also uses utils internally)
    result := auth.CreateUser("jane smith", "jane@example.com")
    fmt.Println(result)
    
    // Test invalid email
    invalid := auth.CreateUser("bob", "invalid-email")
    fmt.Println(invalid)
}
```

**Run from any module:**
```bash
# From frontend directory
cd frontend
go run main.go

# Or from workspace root
cd my-workspace
go run ./frontend

# Dependencies are automatically resolved!
```

**Adding more modules to workspace:**
```bash
# Add a new module to existing workspace
go work use ./new-module

# Or add modules from different locations
go work use /Users/username/other-project/module
go work use ../external-library
```

### Advantages:
- ✅ Clean go.mod files (no replace directives)
- ✅ Automatic local module discovery
- ✅ Easy to add/remove modules
- ✅ Supports modules anywhere on filesystem
- ✅ Modern, recommended approach
- ✅ Great for team development

### Disadvantages:
- ❌ Requires Go 1.18+
- ❌ One more file to manage (go.work)

---

## Comparison Summary

| Aspect | Published Modules | Replace Directives | Go Workspaces |
|--------|------------------|-------------------|---------------|
| **Setup Complexity** | None | Manual per module | One-time setup |
| **go.mod Cleanliness** | Clean | Cluttered | Clean |
| **Go Version** | Any | Any | 1.18+ |
| **Automatic Discovery** | ✅ | ❌ | ✅ |
| **Team Sharing** | ✅ | ✅ | ✅ |
| **CI/CD Friendly** | ✅ | ✅ | ✅ |
| **Multi-location Support** | N/A | ✅ | ✅ |

---

## Best Practices

### Use Published Modules when:
- Working with third-party libraries
- Your modules are stable and versioned
- Sharing code across different projects

### Use Replace Directives when:
- Working with older Go versions (< 1.18)
- Need explicit control over module resolution
- Working in environments that don't support workspaces

### Use Go Workspaces when:
- Developing multiple related modules locally
- Using Go 1.18+
- Want clean, maintainable setup
- Modules may be in different filesystem locations

---

## Quick Reference Commands

### Published Modules:
```bash
go mod tidy              # Auto-add dependencies
go get package@version   # Add specific version
go mod download          # Download dependencies
```

### Replace Directives:
```bash
go mod edit -replace old=new    # Add replace directive
go mod edit -dropreplace old    # Remove replace directive
go mod tidy                     # Sync dependencies
```

### Go Workspaces:
```bash
go work init module1 module2    # Create workspace
go work use ./module            # Add module to workspace
go work use -r ./dir           # Add all modules in directory
go work edit -dropuse ./module  # Remove module from workspace
go work sync                    # Sync workspace dependencies
```

---

## Conclusion

All three approaches are valid and serve different use cases:

- **Published modules** are the gold standard for production dependencies
- **Replace directives** provide explicit control and work universally
- **Go workspaces** offer the cleanest developer experience for local multi-module projects

Choose the approach that best fits your project's needs, Go version constraints, and team preferences.