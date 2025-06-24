/*
Package ulog provides beautiful styled console logging for Go applications.

ULog is a lightweight Go logging library that provides elegant, colorful, and
boxed console output for various message types with a simple and intuitive API.
It enhances your application logs with visually distinct message categories and
formatting to improve readability and debugging experience.

# Features

  - Colorful boxed messages with customizable tags
  - Multiple message types (Error, Warning, Info, Success, Ongoing, Message)
  - Timestamp support
  - Structured data display utilities
  - Simple API with both global functions and configurable logger instances

# Installation

To install ulog, use the following command:

	go get github.com/utsav-56/ulog

# Basic Usage

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

# Creating Custom Loggers

You can create custom logger instances with specific settings:

	// Create a custom logger without timestamps and with 2 spaces of padding
	customLogger := ulog.NewLogger(false, 2)
	customLogger.Success("This is a custom formatted message", "CUSTOM")

# Data Structure Utilities

The package also provides utilities for working with data structures:

	// Print a map
	data := map[string]any{
	    "key1": "value1",
	    "key2": map[string]any{"nested": "value"}
	}
	ulog.PrintMap(data)

	// Print a list
	items := []string{"apple", "banana", "cherry"}
	ulog.PrintList(items)

For more examples and detailed documentation, see the README.md file.
*/
package ulog
