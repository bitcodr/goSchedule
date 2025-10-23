package model

import "testing"

func TestEmailModel(t *testing.T) {
	email := Email{
		Attachments: []string{"file1.pdf", "file2.pdf"},
		Body:        "This is a test email body",
		Receiver:    "receiver@example.com",
		Subject:     "Test Email Subject",
	}

	if email.Body != "This is a test email body" {
		t.Errorf("Expected body to be 'This is a test email body', got '%s'", email.Body)
	}

	if email.Receiver != "receiver@example.com" {
		t.Errorf("Expected receiver to be 'receiver@example.com', got '%s'", email.Receiver)
	}

	if email.Subject != "Test Email Subject" {
		t.Errorf("Expected subject to be 'Test Email Subject', got '%s'", email.Subject)
	}

	if len(email.Attachments) != 2 {
		t.Errorf("Expected 2 attachments, got %d", len(email.Attachments))
	}

	if email.Attachments[0] != "file1.pdf" {
		t.Errorf("Expected first attachment to be 'file1.pdf', got '%s'", email.Attachments[0])
	}
}

func TestEmptyEmailModel(t *testing.T) {
	email := Email{}

	if email.Body != "" {
		t.Errorf("Expected empty body, got '%s'", email.Body)
	}

	if email.Receiver != "" {
		t.Errorf("Expected empty receiver, got '%s'", email.Receiver)
	}

	if email.Subject != "" {
		t.Errorf("Expected empty subject, got '%s'", email.Subject)
	}

	if email.Attachments != nil {
		t.Errorf("Expected nil attachments, got %v", email.Attachments)
	}
}
