package badge

import (
	"bytes"
	"html/template"
	"io"
	"sync"

	"github.com/golang/freetype/truetype"
	"github.com/hawl1/go-badge/fonts"
	"golang.org/x/image/font"
)

type badge struct {
	Subject       string
	Status        string
	Color         Color
	Bounds        bounds
	LabelLength   float64
	MessageLength float64
}

type bounds struct {
	// SubjectDx is the width of subject string of the badge.
	SubjectDx float64
	SubjectX  float64
	// StatusDx is the width of status string of the badge.
	StatusDx float64
	StatusX  float64
}

func (b bounds) Dx() float64 {
	return b.SubjectDx + b.StatusDx
}

type badgeDrawer struct {
	fd    *font.Drawer
	tmpl  *template.Template
	mutex *sync.Mutex
}

func (d *badgeDrawer) Render(subject, status string, color Color, w io.Writer) error {
	d.mutex.Lock()
	subjectDx := d.measureString(subject)
	statusDx := d.measureString(status)
	d.mutex.Unlock()

	labelWidth := float64(subjectDx) / 2.0 + 1
	glyphOffset := float64(extraDx)
	labelLength := labelWidth - glyphOffset

	messageWidth := float64(statusDx) / 2.0 + 1
	messageLength := messageWidth - glyphOffset

	bdg := badge{
		Subject:       subject,
		Status:        status,
		Color:         color,
		Bounds: bounds{
			SubjectDx: subjectDx,
			SubjectX:  (subjectDx/2.0 + 1) * 10,
			StatusDx:  statusDx,
			StatusX:   (subjectDx + statusDx/2.0 - 1) * 10,
		},
		LabelLength:   labelLength,
		MessageLength: messageLength,
	}

	return d.tmpl.Execute(w, bdg)
}

func (d *badgeDrawer) RenderBytes(subject, status string, color Color) ([]byte, error) {
	buf := &bytes.Buffer{}
	err := d.Render(subject, status, color, buf)
	return buf.Bytes(), err
}

// shield.io uses Verdana.ttf to measure text width with an extra 10px.
// As we use Vera.ttf, we have to tune this value a little.
const extraDx = 10

func (d *badgeDrawer) measureString(s string) float64 {
	return float64(d.fd.MeasureString(s)>>6) + extraDx
}

// Render renders a badge of the given color, with given subject and status to w.
func Render(subject, status string, color Color, w io.Writer) error {
	return drawer.Render(subject, status, color, w)
}

// RenderBytes renders a badge of the given color, with given subject and status to bytes.
func RenderBytes(subject, status string, color Color) ([]byte, error) {
	return drawer.RenderBytes(subject, status, color)
}

const (
	dpi      = 72
	fontsize = 11
)

var drawer *badgeDrawer

func init() {
	drawer = &badgeDrawer{
		fd:    mustNewFontDrawer(fontsize, dpi),
		tmpl:  template.Must(template.New("flat-template").Parse(flatTemplate)),
		mutex: &sync.Mutex{},
	}
}

func mustNewFontDrawer(size, dpi float64) *font.Drawer {
	ttf, err := truetype.Parse(fonts.VeraSans)
	if err != nil {
		panic(err)
	}
	return &font.Drawer{
		Face: truetype.NewFace(ttf, &truetype.Options{
			Size:    size,
			DPI:     dpi,
			Hinting: font.HintingFull,
		}),
	}
}
