package tokenizer

import "testing"

func TestTokenizeHostnameOnly(t *testing.T) {
	want := URI{Scheme: "http", Hostname: "example.com", Path: "/", Port: 80}
	if got := Tokenize(want.Hostname); got != want {
		t.Errorf("Tokenize() = %q, want %q", got, want)
	}
}

func TestTokenizeSchemeHostname(t *testing.T) {
	want := URI{Scheme: "http", Hostname: "example.com", Path: "/", Port: 80}
	if got := Tokenize(want.Repr()); got != want {
		t.Errorf("Tokenize() = %q, want %q", got, want)
	}
}

func TestTokenizeSchemeHostnamePath(t *testing.T) {
	want := URI{Scheme: "http", Hostname: "example.com", Path: "/foo", Port: 80}
	if got := Tokenize(want.Repr()); got != want {
		t.Errorf("Tokenize() = %q, want %q", got, want)
	}
}
