# ULog - Beautiful Styled Console Logging for Go

A lightweight Go logging library that provides elegant, colorful, and boxed console output for various message types with a simple and intuitive API.

## Installation

```bash
go get github.com/utsav-56/ulog
```

## Features

-   Colorful boxed messages with customizable tags
-   Multiple message types (Error, Warning, Info, Success, Ongoing, Message)
-   Timestamp support
-   Structured data display utilities
-   Simple API with both global functions and configurable logger instances

## Quick Start

```go
package main

import (
    "github.com/utsav-56/ulog"
)

func main() {
    // Using global functions
    ulog.Info("This is an information message")
    ulog.Error("Something went wrong!", "ERROR")
    ulog.Success("Operation completed successfully", "SUCCESS")

    // Display structured data
    data := map[string]interface{}{
        "user": "john_doe",
        "role": "admin",
        "settings": map[string]interface{}{
            "notifications": true,
            "theme": "dark",
        },
    }

    // Print the map with proper formatting
    ulog.Success("User data retrieved:")
    ulog.PrintMapWithIndent(data, "  ")
}
```

## Logger Functions

### Global Functions

These functions use the default logger instance with predefined settings:

#### Error

Logs an error message in a red box.

-   **Parameters**:

    -   `message`: The error message to display
    -   `tag`: Optional tag to show in the top border of the box

-   **Returns**: void

<details>
<summary>Usage Example</summary>

```go
ulog.Error("Failed to connect to database", "DB ERROR")
```

</details>

#### Warning

Logs a warning message in a yellow box.

-   **Parameters**:

    -   `message`: The warning message to display
    -   `tag`: Optional tag to show in the top border of the box

-   **Returns**: void

<details>
<summary>Usage Example</summary>

```go
ulog.Warning("API rate limit approaching", "API WARNING")
```

</details>

#### Info

Logs an info message in the default terminal color.

-   **Parameters**:

    -   `message`: The informational message to display
    -   `tag`: Optional tag to show in the top border of the box

-   **Returns**: void

<details>
<summary>Usage Example</summary>

```go
ulog.Info("System initialized successfully", "STARTUP")
```

</details>

#### Success

Logs a success message in a green box.

-   **Parameters**:

    -   `message`: The success message to display
    -   `tag`: Optional tag to show in the top border of the box

-   **Returns**: void

<details>
<summary>Usage Example</summary>

```go
ulog.Success("User registration completed", "REGISTRATION")
```

</details>

#### Message

Logs a message in a blue box.

-   **Parameters**:

    -   `message`: The message to display
    -   `tag`: Optional tag to show in the top border of the box

-   **Returns**: void

<details>
<summary>Usage Example</summary>

```go
ulog.Message("Processing user request", "REQUEST")
```

</details>

#### Ongoing

Logs an ongoing operation message in an orange-like color.

-   **Parameters**:

    -   `message`: The ongoing operation message to display
    -   `tag`: Optional tag to show in the top border of the box

-   **Returns**: void

<details>
<summary>Usage Example</summary>

```go
ulog.Ongoing("Importing data, please wait...", "IMPORT")
```

</details>

### Creating Custom Loggers

You can create custom logger instances with specific settings:

```go
func NewLogger(showTimestamp bool, padding int) *Logger
```

-   **Parameters**:

    -   `showTimestamp`: Whether to display timestamps in the log messages
    -   `padding`: The padding inside the box (minimum 1)

-   **Returns**: A new Logger instance with the specified settings

<details>
<summary>Usage Example</summary>

```go
// Create a custom logger without timestamps and with 2 spaces of padding
customLogger := ulog.NewLogger(false, 2)
customLogger.Success("This is a custom formatted message", "CUSTOM")
```

</details>

## Data Structure Utilities

### PrintMap

Prints the contents of a map to the standard output, recursively handling nested maps.

-   **Parameters**:

    -   `m`: The map to be printed

-   **Returns**: void

<details>
<summary>Usage Example</summary>

```go
data := map[string]any{
    "key1": "value1",
    "key2": map[string]any{"nested": "value"}
}
ulog.PrintMap(data)
```

</details>

### PrintList

Prints a list of strings to the standard output with 1-based indexing.

-   **Parameters**:

    -   `list`: The string slice to be printed

-   **Returns**: void

<details>
<summary>Usage Example</summary>

```go
items := []string{"apple", "banana", "cherry"}
ulog.PrintList(items)
// Output:
// 1: apple
// 2: banana
// 3: cherry
```

</details>

### MapAsPrettyString

Converts a map to a pretty string representation.

-   **Parameters**:

    -   `m`: The map to be converted to a string
    -   `beforeMessage`: Optional message to prepend to the map string

-   **Returns**: A formatted string representation of the map

<details>
<summary>Usage Example</summary>

```go
data := map[string]interface{}{"name": "John", "age": 30}
str := ulog.MapAsPrettyString(data, "User info:")
fmt.Println(str)
// Output: User info: {age: 30, name: "John"}
```

</details>

### ValueAsString

Converts a value of any type to its string representation.

-   **Parameters**:

    -   `v`: The value to be converted to a string

-   **Returns**: A string representation of the value

<details>
<summary>Usage Example</summary>

```go
str1 := ulog.ValueAsString("hello")  // Returns: "hello"
str2 := ulog.ValueAsString(42)       // Returns: 42
str3 := ulog.ValueAsString(map[string]interface{}{"key": "value"})  // Returns: {key: "value"}
```

</details>

### ListAsPrettyString

Converts a slice of strings to a pretty string representation.

-   **Parameters**:

    -   `list`: The string slice to be converted to a string
    -   `beforeMessage`: Optional message to prepend to the list string

-   **Returns**: A formatted string representation of the list

<details>
<summary>Usage Example</summary>

```go
items := []string{"apple", "banana", "cherry"}
str := ulog.ListAsPrettyString(items, "Fruits:")
fmt.Println(str)
// Output: Fruits: ["apple", "banana", "cherry"]
```

</details>

### ListAsPrettyStringWithIndex

Converts a slice of strings to a pretty string representation with 1-based indexing.

-   **Parameters**:

    -   `list`: The string slice to be converted to a string
    -   `beforeMessage`: Optional message to prepend to the list string

-   **Returns**: A formatted string representation of the list with indexes

<details>
<summary>Usage Example</summary>

```go
items := []string{"apple", "banana", "cherry"}
str := ulog.ListAsPrettyStringWithIndex(items, "Fruits:")
fmt.Println(str)
// Output: Fruits: [1: "apple", 2: "banana", 3: "cherry"]
```

</details>

### PrintMapWithIndent

Prints a map with proper indentation for nested structures, sorting keys alphabetically.

-   **Parameters**:

    -   `m`: The map to be printed
    -   `indent`: The current indentation string to use

-   **Returns**: void

<details>
<summary>Usage Example</summary>

```go
data := map[string]interface{}{
    "user": map[string]interface{}{
        "name": "John",
        "scores": []interface{}{95, 87, 92},
    },
}
ulog.PrintMapWithIndent(data, "")
```

</details>

### PrintStruct

Prints the contents of a struct to the standard output by converting it to a map first.

-   **Parameters**:

    -   `data`: The struct to be printed

-   **Returns**: void

<details>
<summary>Usage Example</summary>

```go
type User struct {
    Name  string
    Email string
    Age   int
}

user := User{
    Name:  "John Doe",
    Email: "john@example.com",
    Age:   30,
}

ulog.PrintStruct(user)
```

</details>

### ConvertStructToMap

Converts a Go struct to a map[string]interface{} using JSON marshaling and unmarshaling.

-   **Parameters**:

    -   `data`: The input struct or value to be converted

-   **Returns**:
    -   A map representation of the struct
    -   An error if marshaling or unmarshaling fails

<details>
<summary>Usage Example</summary>

```go
type User struct {
    Name  string `json:"name"`
    Email string `json:"email"`
    Age   int    `json:"age"`
}

user := User{
    Name:  "John Doe",
    Email: "john@example.com",
    Age:   30,
}

userMap, err := ulog.ConvertStructToMap(user)
if err != nil {
    ulog.Error("Failed to convert struct: " + err.Error())
    return
}

ulog.PrintMap(userMap)
```

</details>

### FormatJSON

Returns a formatted JSON string with specified indentation.

-   **Parameters**:

    -   `data`: The Go data structure to be marshaled into JSON
    -   `indent`: The number of spaces to use for each indentation level

-   **Returns**:
    -   The formatted JSON string
    -   An error if marshaling fails

<details>
<summary>Usage Example</summary>

```go
type User struct {
    Name  string `json:"name"`
    Email string `json:"email"`
    Age   int    `json:"age"`
}

user := User{
    Name:  "John Doe",
    Email: "john@example.com",
    Age:   30,
}

jsonStr, err := ulog.FormatJSON(user, 4)
if err != nil {
    ulog.Error("Failed to format JSON: " + err.Error())
    return
}

fmt.Println(jsonStr)
```

</details>

## Complete Example

```go
package main

import (
    "github.com/utsav-56/ulog"
)

type Config struct {
    Database struct {
        Host     string `json:"host"`
        Port     int    `json:"port"`
        Username string `json:"username"`
        Password string `json:"password"`
    } `json:"database"`
    Server struct {
        Port    int    `json:"port"`
        Timeout int    `json:"timeout"`
        Debug   bool   `json:"debug"`
    } `json:"server"`
}

func main() {
    // Basic usage
    ulog.Info("Application starting...")

    // Custom logger
    logger := ulog.NewLogger(true, 2)
    logger.Warning("This uses a custom logger with extra padding", "CUSTOM")

    // Using different message types
    ulog.Success("Database connection established", "DB")
    ulog.Error("Failed to load configuration file", "CONFIG")
    ulog.Ongoing("Processing large dataset...", "DATA")
    ulog.Message("User logged in: john_doe", "AUTH")

    // Working with data structures
    items := []string{"Configure database", "Setup server", "Initialize cache"}
    ulog.Info("Pending tasks:")
    ulog.PrintList(items)

    // Create a complex data structure
    cfg := Config{}
    cfg.Database.Host = "localhost"
    cfg.Database.Port = 5432
    cfg.Database.Username = "admin"
    cfg.Database.Password = "********"
    cfg.Server.Port = 8080
    cfg.Server.Timeout = 30
    cfg.Server.Debug = true

    // Convert struct to map and print
    configMap, _ := ulog.ConvertStructToMap(cfg)
    ulog.Success("Configuration loaded:")
    ulog.PrintMapWithIndent(configMap, "  ")

    // Format as JSON
    jsonStr, _ := ulog.FormatJSON(cfg, 2)
    ulog.Info("Configuration as JSON:\n" + jsonStr)

    ulog.Success("Application initialized successfully", "DONE")
}
```

## License

This project is licensed under the MIT License - see the LICENSE file for details.
