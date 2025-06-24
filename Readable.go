package ulog

import (
	"fmt"
	"net"
	"strings"
	"time"
)

func ReadableFileSize(size int64) string {
	if size < 1024 {
		return fmt.Sprintf("%d B", size)
	} else if size < 1024*1024 {
		return fmt.Sprintf("%.2f KB", float64(size)/1024)
	} else if size < 1024*1024*1024 {
		return fmt.Sprintf("%.2f MB", float64(size)/(1024*1024))
	} else if size < 1024*1024*1024*1024 {
		return fmt.Sprintf("%.2f GB", float64(size)/(1024*1024*1024))
	} else {
		return fmt.Sprintf("%.2f TB", float64(size)/(1024*1024*1024*1024))
	}

}

func ReadableTime(seconds int64) string {
	if seconds < 60 {
		return fmt.Sprintf("%d seconds", seconds)
	} else if seconds < 3600 {
		minutes := seconds / 60
		return fmt.Sprintf("%d minutes", minutes)
	} else if seconds < 86400 {
		hours := seconds / 3600
		minutes := (seconds % 3600) / 60
		return fmt.Sprintf("%d hours %d minutes", hours, minutes)
	} else {
		days := seconds / 86400
		hours := (seconds % 86400) / 3600
		minutes := (seconds % 3600) / 60
		return fmt.Sprintf("%d days %d hours %d minutes", days, hours, minutes)
	}
}

func InMBPS(speed float64) string {
	if speed <= 0 {
		return "0.00 MB/s" // Avoid division by zero
	}
	return fmt.Sprintf("%.2f MB/s", speed/(1024*1024)) // Convert bytes to megabytes
}

func ReadablePercentage(percentage float64) string {
	if percentage < 0 {
		return "0.00%"
	} else if percentage > 100 {
		return "100.00%"
	}
	return fmt.Sprintf("%.2f%%", percentage)
}

// ReadableDuration formats a time.Duration into a human-readable string
func ReadableDuration(duration time.Duration) string {
	if duration < time.Millisecond {
		return fmt.Sprintf("%d ns", duration.Nanoseconds())
	} else if duration < time.Second {
		return fmt.Sprintf("%.2f ms", float64(duration.Nanoseconds())/float64(time.Millisecond))
	} else if duration < time.Minute {
		return fmt.Sprintf("%.2f s", float64(duration.Nanoseconds())/float64(time.Second))
	} else {
		return ReadableTime(int64(duration.Seconds()))
	}
}

// ReadableTimestamp formats a time.Time into a human-readable string
func ReadableTimestamp(t time.Time) string {
	return t.Format("2006-01-02 15:04:05.000")
}

// ReadableMemoryUsage formats memory usage in a human-readable format
func ReadableMemoryUsage(bytes uint64) string {
	return ReadableFileSize(int64(bytes))
}

// ReadableIP formats an IP address to a human-readable string
func ReadableIP(ip net.IP) string {
	if ip == nil {
		return "unknown"
	}
	return ip.String()
}

// ReadableLatency formats latency in a human-readable way with appropriate units
func ReadableLatency(latencyMs float64) string {
	if latencyMs < 1 {
		return fmt.Sprintf("%.2f μs", latencyMs*1000)
	} else if latencyMs < 1000 {
		return fmt.Sprintf("%.2f ms", latencyMs)
	} else {
		return fmt.Sprintf("%.2f s", latencyMs/1000)
	}
}

// ReadableCount formats large numbers with comma separators
func ReadableCount(count int64) string {
	str := fmt.Sprintf("%d", count)
	if count < 1000 {
		return str
	}

	result := ""
	for i, char := range str {
		if i > 0 && (len(str)-i)%3 == 0 {
			result += ","
		}
		result += string(char)
	}
	return result
}

// ReadableStatus formats a status code with description
func ReadableStatus(code int) string {
	var status string
	switch {
	case code >= 100 && code < 200:
		status = "Informational"
	case code >= 200 && code < 300:
		status = "Success"
	case code >= 300 && code < 400:
		status = "Redirection"
	case code >= 400 && code < 500:
		status = "Client Error"
	case code >= 500:
		status = "Server Error"
	default:
		status = "Unknown"
	}
	return fmt.Sprintf("%d (%s)", code, status)
}

// ReadableBool formats a boolean value as "Yes" or "No"
func ReadableBool(value bool) string {
	if value {
		return "Yes"
	}
	return "No"
}

// ReadableLevel formats a log level in a consistent way
func ReadableLevel(level string) string {
	level = strings.ToUpper(level)
	switch level {
	case "DEBUG", "INFO", "WARN", "ERROR", "FATAL":
		return level
	default:
		return "UNKNOWN"
	}
}
