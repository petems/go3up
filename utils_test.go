package main

import (
	"errors"
	"testing"
)

func TestIsRecoverable(t *testing.T) {
	if msg := "broken pipes all over"; isRecoverable(errors.New(msg)) {
		t.Errorf("Expected %s to NOT be recoverable", msg)
	}

	if msg := "Oh noes, I broken pipe"; !isRecoverable(errors.New(msg)) {
		t.Errorf("Expected %s to BE recoverable", msg)
	}
}

func TestMsg(t *testing.T) {
	actual := msg()
	if expected := "Error, no message available"; actual != expected {
		t.Error("Expected an error, got", actual)
	}

	verbose := opts.verbose
	opts.verbose = true
	if actual := msg("Foo", "bar", "baz"); actual != "Foo\n" {
		t.Error("Expected Foo\\n got", actual)
	}

	quiet := opts.quiet
	opts.quiet = true
	opts.verbose = false
	if actual := msg("Foo", "bar", "baz"); actual != "baz" {
		t.Error("Expected baz got", actual)
	}

	opts.quiet = false
	if actual := msg("Foo", "bar", "baz"); actual != "bar" {
		t.Error("Expected bar got", actual)
	}

	if actual := msg("Foo"); actual != "Foo\n" {
		t.Error("Expected Foo\\n got", actual)
	}

	opts.verbose = verbose
	opts.quiet = quiet
}

func TestQuit(t *testing.T) {
	t.Skip("No idea how to test this one yet")
}

func TestBetterMime(t *testing.T) {
	assertions := map[string]string{
		".html": "text/html; charset=utf-8",
		".jpg":  "image/jpeg",
		".JPG":  "image/jpeg",
		".ttf":  "binary/octet-stream",
	}

	for ext, mime := range assertions {
		if expected := betterMime(ext); expected != mime {
			t.Error("Expected", mime, "got", expected)
		}
	}
}