package todo

import "testing"

type fakeRepo struct {
	createCalled bool
}

func (f *fakeRepo) Create(title string) (Task, error) {
	f.createCalled = true
	return Task{
		ID:    1,
		Title: title,
	}, nil
}

func TestCreateTask_EmptyTitle(t *testing.T) {
	repo := &fakeRepo{}
	service := NewService(repo)

	_, err := service.CreateTask("")

	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestCreateTask_ValidTitle(t *testing.T) {
	repo := &fakeRepo{}
	service := NewService(repo)

	task, err := service.CreateTask("learn go")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if task.Title != "learn go" {
		t.Fatalf("expected title 'learn go', got %s", task.Title)
	}

	if !repo.createCalled {
		t.Fatal("expected repository Create to be called")
	}
}
