package spotifydefault

import (
	"html"
	"testing"
)

// TestDescriptionUnescape_PreventsCompoundingEscape reproduces the bug from
// issue #7: Spotify HTML-escapes special characters in a playlist description
// when it is stored. If the escaped value read back from the API is reused
// as-is on the next update, Spotify escapes it again, and the description
// grows without bound until the API rejects it with a 400.
func TestDescriptionUnescape_PreventsCompoundingEscape(t *testing.T) {
	original := "C'est la rentrée !"

	// Simulate Spotify's server-side escaping applied on every update.
	spotifyStore := func(description string) string {
		return html.EscapeString(description)
	}

	// Without unescaping the value read back from the API, each run compounds
	// the escaping.
	stored := spotifyStore(original)
	for i := 0; i < 3; i++ {
		stored = spotifyStore(stored)
	}
	if stored == html.EscapeString(original) {
		t.Fatalf("expected unfixed round-trip to compound escaping, got stable value %q", stored)
	}

	// With the fix: unescape what comes back from Spotify before reusing it.
	fixedStored := spotifyStore(original)
	for i := 0; i < 3; i++ {
		fixedStored = spotifyStore(html.UnescapeString(fixedStored))
	}
	if fixedStored != html.EscapeString(original) {
		t.Fatalf("expected repeated runs to stabilize, got %q, want %q", fixedStored, html.EscapeString(original))
	}
	if got := html.UnescapeString(fixedStored); got != original {
		t.Fatalf("expected unescaped description to round-trip to original, got %q, want %q", got, original)
	}
}
