package psolver

import (
	"fmt"
	"io"

	"github.com/fatih/color"
)

// Exporter is the interface for exporting a matrix to an io.Writer
type Exporter interface {
	Export(m *Matrix, w io.Writer) error
}

// StringExporter exports the matrix as a string
type StringExporter struct{}

// Export writes the string representation of the matrix to the writer
func (s *StringExporter) Export(m *Matrix, w io.Writer) error {
	_, err := io.WriteString(w, m.String())
	return err
}

// ColorStringExporter exports the matrix using ANSI colors
type ColorStringExporter struct {
	colorMap map[byte]*color.Color
}

// NewColorStringExporter creates a new ColorStringExporter with default colors
func NewColorStringExporter() *ColorStringExporter {
	return &ColorStringExporter{
		colorMap: map[byte]*color.Color{
			'F': color.New().AddBgRGB(230, 25, 75).AddRGB(230, 25, 75),     // Red
			'I': color.New().AddBgRGB(60, 180, 75).AddRGB(60, 180, 75),     // Green
			'L': color.New().AddBgRGB(255, 225, 25).AddRGB(255, 225, 25),   // Yellow
			'N': color.New().AddBgRGB(0, 130, 200).AddRGB(0, 130, 200),     // Blue
			'P': color.New().AddBgRGB(245, 130, 48).AddRGB(245, 130, 48),   // Orange
			'T': color.New().AddBgRGB(145, 30, 180).AddRGB(145, 30, 180),   // Purple
			'U': color.New().AddBgRGB(70, 240, 240).AddRGB(70, 240, 240),   // Cyan
			'V': color.New().AddBgRGB(240, 50, 230).AddRGB(240, 50, 230),   // Magenta
			'W': color.New().AddBgRGB(210, 245, 60).AddRGB(210, 245, 60),   // Lime
			'X': color.New().AddBgRGB(250, 190, 212).AddRGB(250, 190, 212), // Pink
			'Y': color.New().AddBgRGB(0, 128, 128).AddRGB(0, 128, 128),     // Teal
			'Z': color.New().AddBgRGB(220, 190, 255).AddRGB(220, 190, 255), // Lavender
			'O': color.New().AddBgRGB(0, 0, 0).AddRGB(0, 0, 0),             // Balck
		},
	}
}

// Export writes the matrix with colors to the writer
func (c *ColorStringExporter) Export(m *Matrix, w io.Writer) error {
	for j := 0; j < m.Height; j++ {
		for i := 0; i < m.Width; i++ {
			val := m.data[j*m.Width+i]
			if val == 0 {
				if _, err := io.WriteString(w, " ."); err != nil {
					return err
				}
			} else {
				if col, ok := c.colorMap[val]; ok {
					if _, err := col.Fprint(w, " "+string(val)); err != nil {
						return err
					}
				} else {
					if _, err := io.WriteString(w, " "+string(val)); err != nil {
						return err
					}
				}
			}
		}
		if _, err := io.WriteString(w, "\n"); err != nil {
			return err
		}
	}
	return nil
}

// SVGExporter exports the matrix as an SVG image
type SVGExporter struct {
	CellSize int
	colorMap map[byte]string
}

// NewSVGExporter creates a new SVGExporter with default settings
func NewSVGExporter() *SVGExporter {
	return &SVGExporter{
		CellSize: 20,
		colorMap: map[byte]string{
			'F': "#E6194B", // Red
			'I': "#3CB44B", // Green
			'L': "#FFE119", // Yellow
			'N': "#0082C8", // Blue
			'P': "#F58230", // Orange
			'T': "#911EB4", // Purple
			'U': "#46F0F0", // Cyan
			'V': "#F032E6", // Magenta
			'W': "#D2F53C", // Lime
			'X': "#FABED4", // Pink
			'Y': "#008080", // Teal
			'Z': "#DCBEFF", // Lavender
		},
	}
}

// Export writes the matrix as an SVG to the writer
func (s *SVGExporter) Export(m *Matrix, w io.Writer) error {
	width := m.Width * s.CellSize
	height := m.Height * s.CellSize

	if _, err := fmt.Fprintf(w, "<svg width=\"%d\" height=\"%d\" xmlns=\"http://www.w3.org/2000/svg\">\n", width, height); err != nil {
		return err
	}

	// Background (optional, but good for transparency handling if needed)
	// fmt.Fprintf(w, "<rect width=\"100%%\" height=\"100%%\" fill=\"white\"/>\n")

	for j := 0; j < m.Height; j++ {
		for i := 0; i < m.Width; i++ {
			val := m.data[j*m.Width+i]
			if val != 0 {
				color, ok := s.colorMap[val]
				if !ok {
					color = "black" // Fallback
				}
				x := i * s.CellSize
				y := j * s.CellSize
				if _, err := fmt.Fprintf(w, "<rect x=\"%d\" y=\"%d\" width=\"%d\" height=\"%d\" fill=\"%s\" stroke=\"black\" stroke-width=\"1\"/>\n", x, y, s.CellSize, s.CellSize, color); err != nil {
					return err
				}
			}
		}
	}

	if _, err := io.WriteString(w, "</svg>\n"); err != nil {
		return err
	}
	return nil
}
