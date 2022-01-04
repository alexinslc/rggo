package main_test

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

// variables
var (
	binName  = "todo"
	fileName = ".todo.json"
)

func TestMain(m *testing.M) {
	fmt.Println("Building tool...")

	// If it's windows, append .exe for us
	if runtime.GOOS == "windows" {
		binName += ".exe"
	}

	// Create the command to build it
	build := exec.Command("go", "build", "-o", binName)

	// if there's an error building, tell us about it
	if err := build.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Cannot build tool %s: %s", binName, err)
		os.Exit(1)
	}

	fmt.Println("Running tests...")
	result := m.Run()

	fmt.Println("Cleaning up...")
	os.Remove(binName)
	os.Remove(fileName)
	os.Exit(result)
}

func TestTodoCLI(t *testing.T) {
	task := "test task number 1"
	dir, err := os.Getwd()

	if err != nil {
		t.Fatal(err)
	}

	cmdPath := filepath.Join(dir, binName)

	// For this subtest, we're setting the name to AddNewTask as the first argument to t.Run() to make it easier
	// to see the results. Also, we're executing the compiled binary with the expected argument by splitting the
	// task variable. The test fails if an error occurs while adding the task.
	t.Run("AddNewTask", func(t *testing.T) {
		cmd := exec.Command(cmdPath, strings.Split(task, " ")...)
		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})

	// For this subtest, we're setting the name to ListTasks. Then we execute the tool with no arguments
	// capturing it's output in the variable 'out'. The test fails immediately if an error occurs while executing
	// the tool. If the execution succeeds, we compare the output with the task name, failing the test if they
	// don't match.
	t.Run("ListTasks", func(t *testing.T) {
		cmd := exec.Command(cmdPath)
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}
		expected := task + "\n"
		if expected != string(out) {
			t.Errorf("Expected %q, got %q instead\n", expected, string(out))
		}
	})
}
