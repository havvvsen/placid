package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

// getFreePort returns a free, available port on localhost.
func getFreePort() (string, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return "", err
	}
	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return "", err
	}
	defer l.Close()
	return fmt.Sprintf("%d", l.Addr().(*net.TCPAddr).Port), nil
}

func TestFileServer(t *testing.T) {
	port, err := getFreePort()
	if err != nil {
		t.Fatalf("Failed to get free port: %v", err)
	}

	// Create a temporary directory for the file server base path
	tempDir, err := os.MkdirTemp("", "fileserver-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create a dummy file in the temp directory
	testContent := "hello world"
	testFile := filepath.Join(tempDir, "test.txt")
	if err := os.WriteFile(testFile, []byte(testContent), 0644); err != nil {
		t.Fatalf("Failed to write test file: %v", err)
	}

	// Create a nested file
	nestedDir := filepath.Join(tempDir, "nested")
	if err := os.Mkdir(nestedDir, 0755); err != nil {
		t.Fatalf("Failed to create nested dir: %v", err)
	}
	nestedFile := filepath.Join(nestedDir, "nested.txt")
	if err := os.WriteFile(nestedFile, []byte("nested content"), 0644); err != nil {
		t.Fatalf("Failed to write nested file: %v", err)
	}

	// Set up command to run the main program
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cmd := exec.CommandContext(ctx, "go", "run", "main.go")
	// Use clean environment to ensure tests don't inherit unexpected env vars
	cmd.Env = append(os.Environ(),
		"FILE_SERVER_PORT="+port,
		"FILE_SERVER_BASE_PATH="+tempDir,
	)

	if err := cmd.Start(); err != nil {
		t.Fatalf("Failed to start server: %v", err)
	}
	defer func() {
		if cmd.Process != nil {
			cmd.Process.Kill()
		}
	}()

	// Wait for the server to be ready
	serverURL := fmt.Sprintf("http://localhost:%s", port)
	ready := false
	for i := 0; i < 20; i++ {
		time.Sleep(100 * time.Millisecond)
		resp, err := http.Get(serverURL + "/test.txt")
		if err == nil {
			resp.Body.Close()
			if resp.StatusCode == http.StatusOK {
				ready = true
				break
			}
		}
	}

	if !ready {
		t.Fatalf("Server did not become ready in time")
	}

	// Test 1: Fetch existing file
	t.Run("ServeExistingFile", func(t *testing.T) {
		resp, err := http.Get(serverURL + "/test.txt")
		if err != nil {
			t.Fatalf("Failed to make request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status 200, got %d", resp.StatusCode)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("Failed to read response body: %v", err)
		}

		if string(body) != testContent {
			t.Errorf("Expected body %q, got %q", testContent, string(body))
		}
	})

	// Test 2: Fetch nested file
	t.Run("ServeNestedFile", func(t *testing.T) {
		resp, err := http.Get(serverURL + "/nested/nested.txt")
		if err != nil {
			t.Fatalf("Failed to make request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status 200, got %d", resp.StatusCode)
		}

		body, _ := io.ReadAll(resp.Body)
		if string(body) != "nested content" {
			t.Errorf("Expected nested content, got %q", string(body))
		}
	})

	// Test 3: Fetch nonexistent file (404)
	t.Run("ServeNonexistentFile", func(t *testing.T) {
		resp, err := http.Get(serverURL + "/nonexistent.txt")
		if err != nil {
			t.Fatalf("Failed to make request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusNotFound {
			t.Errorf("Expected status 404, got %d", resp.StatusCode)
		}
	})

	// Test 4: Directory listing
	t.Run("ServeDirectoryListing", func(t *testing.T) {
		resp, err := http.Get(serverURL + "/")
		if err != nil {
			t.Fatalf("Failed to make request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status 200 for dir, got %d", resp.StatusCode)
		}

		body, _ := io.ReadAll(resp.Body)
		if !strings.Contains(string(body), "test.txt") {
			t.Errorf("Expected directory listing to contain test.txt")
		}
		if !strings.Contains(string(body), "nested") {
			t.Errorf("Expected directory listing to contain nested directory")
		}
	})
}

func TestMissingEnvVars(t *testing.T) {
	t.Run("MissingPort", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		cmd := exec.CommandContext(ctx, "go", "run", "main.go")
		
		// Create a clean environment without FILE_SERVER_PORT
		env := os.Environ()
		var cleanEnv []string
		for _, e := range env {
			if !strings.HasPrefix(e, "FILE_SERVER_PORT=") {
				cleanEnv = append(cleanEnv, e)
			}
		}
		cleanEnv = append(cleanEnv, "FILE_SERVER_BASE_PATH=/tmp")
		cmd.Env = cleanEnv

		output, err := cmd.CombinedOutput()
		if err == nil {
			t.Fatal("Expected command to fail due to missing FILE_SERVER_PORT")
		}

		if !strings.Contains(string(output), "$FILE_SERVER_PORT is required") {
			t.Errorf("Expected error message about missing port, got: %s", string(output))
		}
	})

	t.Run("MissingBasePath", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		cmd := exec.CommandContext(ctx, "go", "run", "main.go")
		
		// Create a clean environment without FILE_SERVER_BASE_PATH
		env := os.Environ()
		var cleanEnv []string
		for _, e := range env {
			if !strings.HasPrefix(e, "FILE_SERVER_BASE_PATH=") {
				cleanEnv = append(cleanEnv, e)
			}
		}
		cleanEnv = append(cleanEnv, "FILE_SERVER_PORT=8080")
		cmd.Env = cleanEnv

		output, err := cmd.CombinedOutput()
		if err == nil {
			t.Fatal("Expected command to fail due to missing FILE_SERVER_BASE_PATH")
		}

		if !strings.Contains(string(output), "$FILE_SERVER_BASE_PATH is required") {
			t.Errorf("Expected error message about missing base path, got: %s", string(output))
		}
	})
}
