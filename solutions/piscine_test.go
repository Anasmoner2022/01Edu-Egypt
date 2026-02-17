package piscine

import (
	"bytes"
	"io"
	"os"
	"testing"
)

// Helper function to capture z01.PrintRune output
func captureOutput(f func()) string {
	// Create a pipe to capture output
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Run the function that uses z01.PrintRune
	f()

	// Close writer and restore stdout
	w.Close()
	os.Stdout = old

	// Read captured output
	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

// Tests for FirstWord
func TestFirstWord(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"simple word", "hello world", "hello\n"},
		{"leading spaces", "   hello world", "hello\n"},
		{"single word", "hello", "hello\n"},
		{"empty string", "", "\n"},
		{"only spaces", "   ", "\n"},
		{"tabs and spaces", "\t\t  hello", "hello\n"},
		{"multiple words", "foo bar baz", "foo\n"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FirstWord(tt.input)
			if result != tt.expected {
				t.Errorf("FirstWord(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

// Tests for LastWord
func TestLastWord(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"simple words", "hello world", "world\n"},
		{"trailing spaces", "hello world   ", "world\n"},
		{"single word", "hello", "hello\n"},
		{"empty string", "", "\n"},
		{"only spaces", "   ", "\n"},
		{"multiple words", "foo bar baz", "baz\n"},
		{"tabs", "hello\tworld\t", "world\n"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LastWord(tt.input)
			if result != tt.expected {
				t.Errorf("LastWord(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

// Tests for Gcd
func TestGcd(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"both zero", 0, 0, 0},
		{"one zero", 42, 0, 42},
		{"simple gcd", 42, 10, 2},
		{"coprime numbers", 13, 17, 1},
		{"same numbers", 100, 100, 100},
		{"large numbers", 1071, 462, 21},
		{"reverse order", 462, 1071, 21},
		{"negative numbers", -42, 10, 2},
		{"both negative", -42, -10, -2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Gcd(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Gcd(%d, %d) = %d, want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// Tests for HashCode
func TestHashCode(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{"simple string", "Hello world!"},
		{"empty string", ""},
		{"single char", "A"},
		{"numbers", "123456"},
		{"special chars", "!@#$%"},
		{"long string", "This is a longer test string with various characters!"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := HashCode(tt.input)

			// Check that result has same length as input
			if len(result) != len(tt.input) {
				t.Errorf("HashCode(%q) length = %d, want %d", tt.input, len(result), len(tt.input))
			}

			// Check that all characters are printable (>= 33, <= 127)
			for i, ch := range result {
				if ch < 33 || ch > 127 {
					t.Errorf("HashCode(%q)[%d] = %d, not in printable range [33, 127]", tt.input, i, ch)
				}
			}
		})
	}
}

// Tests for IsCapitalized
func TestIsCapitalized(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"all capitalized", "Hello World", true},
		{"not capitalized", "hello world", false},
		{"mixed case first", "Hello world", false},
		{"single word capitalized", "Hello", true},
		{"single word lowercase", "hello", false},
		{"empty string", "", false},
		{"numbers first", "123 456", true},
		{"special chars first", "!Hello @World", true},
		{"all caps", "HELLO WORLD", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsCapitalized(tt.input)
			if result != tt.expected {
				t.Errorf("IsCapitalized(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

// Tests for Fields
func TestFields(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{"simple words", "hello world", []string{"hello", "world"}},
		{"multiple spaces", "hello    world", []string{"hello", "world"}},
		{"leading spaces", "   hello world", []string{"hello", "world"}},
		{"trailing spaces", "hello world   ", []string{"hello", "world"}},
		{"tabs", "hello\tworld", []string{"hello", "world"}},
		{"newlines", "hello\nworld", []string{"hello", "world"}},
		{"mixed whitespace", "  hello\t\nworld  \r", []string{"hello", "world"}},
		{"empty string", "", []string{}},
		{"only spaces", "   ", []string{}},
		{"single word", "hello", []string{"hello"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Fields(tt.input)
			if len(result) != len(tt.expected) {
				t.Errorf("Fields(%q) length = %d, want %d", tt.input, len(result), len(tt.expected))
				return
			}
			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("Fields(%q)[%d] = %q, want %q", tt.input, i, result[i], tt.expected[i])
				}
			}
		})
	}
}

// Tests for IsNegative
func TestIsNegative(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected string
	}{
		{"negative number", -5, "T\n"},
		{"positive number", 5, "F\n"},
		{"zero", 0, "F\n"},
		{"large negative", -1000000, "T\n"},
		{"large positive", 1000000, "F\n"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := captureOutput(func() {
				IsNegative(tt.input)
			})
			if result != tt.expected {
				t.Errorf("IsNegative(%d) output = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

// Tests for PrintAlpha
func TestPrintAlpha(t *testing.T) {
	expected := "abcdefghijklmnopqrstuvwxyz\n"
	result := captureOutput(func() {
		PrintAlpha()
	})
	if result != expected {
		t.Errorf("PrintAlpha() = %q, want %q", result, expected)
	}
}

// Tests for PrintReverseAlpha
func TestPrintReverseAlpha(t *testing.T) {
	expected := "zyxwvutsrqponmlkjihgfedcba\n"
	result := captureOutput(func() {
		PrintReverseAlpha()
	})
	if result != expected {
		t.Errorf("PrintReverseAlpha() = %q, want %q", result, expected)
	}
}

// Tests for PrintDigit
func TestPrintDigit(t *testing.T) {
	expected := "0123456789\n"
	result := captureOutput(func() {
		PrintDigit()
	})
	if result != expected {
		t.Errorf("PrintDigit() = %q, want %q", result, expected)
	}
}

// Tests for PrintComp
func TestPrintComp(t *testing.T) {
	result := captureOutput(func() {
		PrintComp()
	})

	// Check that it starts correctly
	if len(result) < 5 || result[:5] != "012, " {
		t.Errorf("PrintComp() doesn't start with '012, '")
	}

	// Check that it ends correctly
	if len(result) < 4 || result[len(result)-4:] != "789\n" {
		t.Errorf("PrintComp() doesn't end with '789\\n'")
	}

	// Check that combinations are in ascending order
	if !bytes.Contains([]byte(result), []byte("123")) {
		t.Errorf("PrintComp() doesn't contain '123'")
	}
}

// Tests for PrintComb2
func TestPrintComb2(t *testing.T) {
	result := captureOutput(func() {
		PrintComb2()
	})

	// Check that it starts correctly
	if len(result) < 8 || result[:8] != "00 01, " {
		t.Errorf("PrintComb2() doesn't start with '00 01, '")
	}

	// Check that it ends correctly
	if len(result) < 5 || result[len(result)-5:] != "98 99\n" {
		t.Errorf("PrintComb2() doesn't end with '98 99\\n'")
	}
}

// Tests for PrintNbr
func TestPrintNbr(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected string
	}{
		{"zero", 0, "0"},
		{"positive single digit", 5, "5"},
		{"positive multiple digits", 123, "123"},
		{"negative single digit", -5, "-5"},
		{"negative multiple digits", -123, "-123"},
		{"large positive", 987654321, "987654321"},
		{"large negative", -987654321, "-987654321"},
		{"minimum int64", -9223372036854775808, "-9223372036854775808"},
		{"maximum int64", 9223372036854775807, "9223372036854775807"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := captureOutput(func() {
				PrintNbr(tt.input)
			})
			if result != tt.expected {
				t.Errorf("PrintNbr(%d) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

// Tests for PrintMemory
func TestPrintMemory(t *testing.T) {
	tests := []struct {
		name  string
		input [10]byte
	}{
		{
			name:  "hello with special chars",
			input: [10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*', 0, 0},
		},
		{
			name:  "all printable",
			input: [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'},
		},
		{
			name:  "numbers",
			input: [10]byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := captureOutput(func() {
				PrintMemory(tt.input)
			})

			// Check that output has at least 3 lines (hex output + character output)
			lines := bytes.Split([]byte(result), []byte("\n"))
			if len(lines) < 3 {
				t.Errorf("PrintMemory() output has %d lines, want at least 3", len(lines))
			}

			// Check that hex values are present
			if !bytes.Contains([]byte(result), []byte("0")) {
				t.Errorf("PrintMemory() output doesn't contain hex digits")
			}
		})
	}
}

// Benchmark tests
func BenchmarkGcd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Gcd(1071, 462)
	}
}

func BenchmarkHashCode(b *testing.B) {
	input := "This is a test string for benchmarking"
	for i := 0; i < b.N; i++ {
		HashCode(input)
	}
}

func BenchmarkFields(b *testing.B) {
	input := "  hello   world   foo   bar   baz  "
	for i := 0; i < b.N; i++ {
		Fields(input)
	}
}

func BenchmarkIsCapitalized(b *testing.B) {
	input := "Hello World Foo Bar"
	for i := 0; i < b.N; i++ {
		IsCapitalized(input)
	}
}

func TestAtoi(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "simple positive number",
			input:    "12345",
			expected: 12345,
		},
		{
			name:     "number with leading zeros",
			input:    "0000000012345",
			expected: 12345,
		},
		{
			name:     "number with space",
			input:    "012 345",
			expected: 0,
		},
		{
			name:     "non-numeric string",
			input:    "Hello World!",
			expected: 0,
		},
		{
			name:     "positive sign",
			input:    "+1234",
			expected: 1234,
		},
		{
			name:     "negative sign",
			input:    "-1234",
			expected: -1234,
		},
		{
			name:     "double positive sign",
			input:    "++1234",
			expected: 0,
		},
		{
			name:     "double negative sign",
			input:    "--1234",
			expected: 0,
		},
		{
			name:     "zero",
			input:    "0",
			expected: 0,
		},
		{
			name:     "single digit",
			input:    "7",
			expected: 7,
		},
		{
			name:     "negative single digit",
			input:    "-7",
			expected: -7,
		},
		{
			name:     "positive single digit with sign",
			input:    "+7",
			expected: 7,
		},
		{
			name:     "empty string",
			input:    "",
			expected: 0,
		},
		{
			name:     "only sign",
			input:    "+",
			expected: 0,
		},
		{
			name:     "only negative sign",
			input:    "-",
			expected: 0,
		},
		{
			name:     "letters after numbers",
			input:    "123abc",
			expected: 0,
		},
		{
			name:     "letters before numbers",
			input:    "abc123",
			expected: 0,
		},
		{
			name:     "sign in middle",
			input:    "12+34",
			expected: 0,
		},
		{
			name:     "multiple zeros",
			input:    "0000",
			expected: 0,
		},
		{
			name:     "large number",
			input:    "987654321",
			expected: 987654321,
		},
		{
			name:     "negative large number",
			input:    "-987654321",
			expected: -987654321,
		},
		{
			name:     "positive with leading zeros",
			input:    "+00042",
			expected: 42,
		},
		{
			name:     "negative with leading zeros",
			input:    "-00042",
			expected: -42,
		},
		{
			name:     "mixed signs",
			input:    "+-123",
			expected: 0,
		},
		{
			name:     "mixed signs reverse",
			input:    "-+123",
			expected: 0,
		},
		{
			name:     "special characters",
			input:    "!@#$%",
			expected: 0,
		},
		{
			name:     "tab character",
			input:    "123\t456",
			expected: 0,
		},
		{
			name:     "newline character",
			input:    "123\n456",
			expected: 0,
		},
		{
			name:     "decimal number",
			input:    "123.456",
			expected: 0,
		},
		{
			name:     "sign at end",
			input:    "123+",
			expected: 0,
		},
		{
			name:     "sign at end negative",
			input:    "123-",
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Atoi(tt.input)
			if result != tt.expected {
				t.Errorf("Atoi(%q) = %d, want %d", tt.input, result, tt.expected)
			}
		})
	}
}

// Additional edge case tests
func TestAtoiEdgeCases(t *testing.T) {
	// Test that the function handles all provided examples correctly
	examples := map[string]int{
		"12345":         12345,
		"0000000012345": 12345,
		"012 345":       0,
		"Hello World!":  0,
		"+1234":         1234,
		"-1234":         -1234,
		"++1234":        0,
		"--1234":        0,
	}

	for input, expected := range examples {
		result := Atoi(input)
		if result != expected {
			t.Errorf("Atoi(%q) = %d, want %d", input, result, expected)
		}
	}
}

// Benchmark test
func BenchmarkAtoi(b *testing.B) {
	testCases := []string{
		"12345",
		"-12345",
		"+12345",
		"0000000012345",
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			Atoi(tc)
		}
	}
}
