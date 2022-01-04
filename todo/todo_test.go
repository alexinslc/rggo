package todo_test

import (
	"io/ioutil"
	"os"
	"testing"

	"pragprog.com/rggo/todo"
)

// TestAdd tests to Add method of the List type
func TestAdd(t *testing.T) {
	// Create a list
	l := todo.List{}
	// Create a new task and add it to the list
	taskName := "New Task"
	l.Add(taskName)
	// Check the task was added successfully as the first task in the list
	if l[0].Task != taskName {
		t.Errorf("Expected %q, got %q instead.", taskName, l[0].Task)
	}
}

// TestComplete tests the Complete method of the List type
func TestComplete(t *testing.T) {
	l := todo.List{}

	taskName := "New Task"
	l.Add(taskName)

	if l[0].Task != taskName {
		t.Errorf("Expected %q, but %q instead.", taskName, l[0].Task)
	}

	if l[0].Done {
		t.Errorf("New task should not be completed.")
	}

	// Actually call the method Complete on the task
	l.Complete(1)

	if !l[0].Done {
		t.Errorf("New task should be completed.")
	}
}

// TestDelete tests the Delete method of the List type
func TestDelete(t *testing.T) {
	l := todo.List{}
	// create some dummy tasks
	tasks := []string{
		"New Task 1",
		"New Task 2",
		"New Task 3",
	}

	// add them to the list
	for _, v := range tasks {
		l.Add(v)
	}

	// Check if the first task matches up in our list
	if l[0].Task != tasks[0] {
		t.Errorf("Expected %q, got %q instead.", tasks[0], l[0].Task)
	}

	// Delete a task
	l.Delete(2)

	// Check if the length of the list is now 2
	if len(l) != 2 {
		t.Errorf("Expected list length %d, got %d instead.", 2, len(l))
	}

	// Check the task matches up
	if l[1].Task != tasks[2] {
		t.Errorf("Expected %q, got %q instead.", tasks[2], l[1].Task)
	}
}

// TestSaveGet tests the Dave and Get methods of the List type
func TestSaveGet(t *testing.T) {

	// Create some lists
	l1 := todo.List{}
	l2 := todo.List{}

	// Add a task to 1st list
	taskName := "New Task"
	l1.Add(taskName)

	// Check our task was created
	if l1[0].Task != taskName {
		t.Errorf("Expected %q, got %q instead.", taskName, l1[0].Task)
	}

	// Create a tempfile
	tf, err := ioutil.TempFile("", "")

	// Check we can create a tempfile
	if err != nil {
		t.Fatalf("Error creating temp file: %s", err)
	}
	// remove the tempfile after we done
	defer os.Remove(tf.Name())

	// check that we can actually save the file
	if err := l1.Save(tf.Name()); err != nil {
		t.Fatalf("Error saving list to file: %s", err)
	}

	// check kthat we can actually get the file
	if err := l2.Get(tf.Name()); err != nil {
		t.Fatalf("Error getting list from file: %s", err)
	}

	// Check if the tasks in our lists are the same
	if l1[0].Task != l2[0].Task {
		t.Errorf("Task %q should match %q task.", l1[0].Task, l2[0].Task)
	}
}
