package utils

import "testing"

func TestCleanChirp(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "No Profanity",
			input:    "I had something interesting for breakfast",
			expected: "I had something interesting for breakfast",
		},
		{
			name:     "Single Replace",
			input:    "I hear Mastodon is better than Chirpy. sharbert I need to migrate",
			expected: "I hear Mastodon is better than Chirpy. **** I need to migrate",
		},
		{
			name:     "Multi Replace",
			input:    "I really need a kerfuffle to go to bed sooner, Fornax !",
			expected: "I really need a **** to go to bed sooner, **** !",
		},
		{
			name:     "Multi Replace with Punctuation",
			input:    "I really need a kerfuffle to go to bed sooner, Fornax!",
			expected: "I really need a **** to go to bed sooner, ****!",
		},
	}
	for _, c := range cases {
		profanity := []string{"kerfuffle", "sharbert", "fornax"}
		actual := CleanChirp(c.input, profanity)
		if len(actual) != len(c.expected) {
			t.Errorf("\nLengths do not match in test [%s]:\n\tExpected: %d\n\tActual: %d", c.name, len(c.expected), len(actual))
		}
		if actual != c.expected {
			t.Errorf("\nStrings do not match in test [%s]:\n\tExpected: %s\n\tActual: %s", c.name, c.expected, actual)
		}
	}
}
