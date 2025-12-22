package psolver

import (
	"bytes"
	"io"
	"testing"

	"github.com/fatih/color"
)

type MockExporter struct{}

func (m *MockExporter) Export(mat *Matrix, w io.Writer) error {
	_, err := w.Write([]byte("mock export"))
	return err
}

func TestExporterInterface(t *testing.T) {
	var _ Exporter = &MockExporter{}

	mat := NewMatrix(5, 5)
	e := &MockExporter{}
	var buf bytes.Buffer
	if err := e.Export(mat, &buf); err != nil {
		t.Fatalf("Export failed: %v", err)
	}
	if buf.String() != "mock export" {
		t.Errorf("Expected 'mock export', got '%s'", buf.String())
	}
}

func TestStringExporter(t *testing.T) {
	m := NewMatrix(3, 3)
	// Fill matrix with some data to have a predictable string output if needed,
	// or just rely on empty matrix string representation.
	// Matrix.String() for empty 3x3:
	// ...
	// ...
	// ...
	expected := "...\n...\n...\n"

	exporter := &StringExporter{}
	var buf bytes.Buffer
	if err := exporter.Export(m, &buf); err != nil {
		t.Fatalf("Export failed: %v", err)
	}

	if buf.String() != expected {
		t.Errorf("Expected:\n%q\nGot:\n%q", expected, buf.String())
	}
}

func TestColorStringExporter(t *testing.T) {
	// Initialize color system for testing, usually color detects TTY
	// We force it to enable color output
	color.NoColor = false

	m := NewMatrix(3, 3)
	// Place a piece to check color output
	m.data[0] = 'F'
	m.data[1] = 'I'
	m.data[2] = 0 // empty

	exporter := NewColorStringExporter()
	var buf bytes.Buffer
	if err := exporter.Export(m, &buf); err != nil {
		t.Fatalf("Export failed: %v", err)
	}

	// Helper to check if buffer contains substring
	contains := func(s string) bool {
		return bytes.Contains(buf.Bytes(), []byte(s))
	}

	// We expect ANSI codes.
	// Red background for F: \x1b[41;37mF\x1b[0m
	// Green background for I: \x1b[42;30mI\x1b[0m

	if !contains("F") {
		t.Error("Output should contain 'F'")
	}
	if !contains("I") {
		t.Error("Output should contain 'I'")
	}
	// Check for at least some escape code start
	if !contains("\x1b[") {
		t.Error("Output should contain ANSI escape codes")
	}
}

func TestSVGExporter(t *testing.T) {
	m := NewMatrix(3, 3)
	m.data[0] = 'F'
	m.data[1] = 'I'
	m.data[4] = 'L' // center

	exporter := NewSVGExporter()
	var buf bytes.Buffer
	if err := exporter.Export(m, &buf); err != nil {
		t.Fatalf("Export failed: %v", err)
	}

	// Check header
	if !bytes.Contains(buf.Bytes(), []byte("<svg width=\"60\" height=\"60\"")) {
		t.Error("SVG header missing or incorrect dimensions")
	}
	// Check for rects
	if !bytes.Contains(buf.Bytes(), []byte("<rect")) {
		t.Error("SVG should contain rectangles")
	}
	// Check for specific color (F is Red #E6194B)
	if !bytes.Contains(buf.Bytes(), []byte("#E6194B")) {
		t.Error("SVG should contain color for F (#E6194B)")
	}
	// Check for specific color (I is Green #3CB44B)
	if !bytes.Contains(buf.Bytes(), []byte("#3CB44B")) {
		t.Error("SVG should contain color for I (#3CB44B)")
	}
}
