// internal/liquidationengine/liquidationengine_test.go
package liquidationengine

import (
    "testing"
)

// TestNewApp verifies the NewApp function returns a non-nil app instance
// with verbose mode enabled and initial processed count set to 0.
func TestNewApp(t *testing.T) {
    app := NewApp(true)
    if app == nil {
        t.Fatal("NewApp returned nil")
    }
    if !app.Verbose {
        t.Error("Expected verbose to be true")
    }
    if app.ProcessedCount != 0 {
        t.Errorf("Expected ProcessedCount to be 0, got %d", app.ProcessedCount)
    }
}

// TestProcess tests the Process method with sample data, checking for success
// and correct processed count increment.
func TestProcess(t *testing.T) {
    app := NewApp(false)
    result, err := app.Process("test data")
    
    if err != nil {
        t.Fatalf("Process returned error: %v", err)
    }
    
    if !result.Success {
        t.Error("Expected result.Success to be true")
    }
    
    if app.ProcessedCount != 1 {
        t.Errorf("Expected ProcessedCount to be 1, got %d", app.ProcessedCount)
    }
}

// TestRun tests the Run method with empty input, checking for no error returned.
func TestRun(t *testing.T) {
    app := NewApp(false)
    err := app.Run("", "")
    if err != nil {
        t.Errorf("Run returned unexpected error: %v", err)
    }
}