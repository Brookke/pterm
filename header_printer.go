package pterm

var (
	// DefaultHeaderPrinter returns the printer for a default header text.
	// Defaults to LightWhite, Bold Text and a Gray Header background.
	DefaultHeaderPrinter = HeaderPrinter{Header: Header{
		TextStyle:       Style{FgLightWhite, Bold},
		BackgroundStyle: Style{BgGray},
		Margin:          5,
	}}

	// PrintHeader is the short form of DefaultHeaderPrinter.Println
	PrintHeader = DefaultHeaderPrinter.Println
)

// Header contains the data used to craft a header.
// A Header is printed as a big box with text in it.
// Can be used as title screens or section separator.
type Header struct {
	TextStyle       Style
	BackgroundStyle Style
	Margin          int
}

// HeaderPrinter is the printer used to print a Header.
type HeaderPrinter struct {
	Header Header
}

// Sprint formats using the default formats for its operands and returns the resulting string.
// Spaces are added between operands when neither is a string.
func (p HeaderPrinter) Sprint(a ...interface{}) string {
	text := Sprint(a...)
	textLength := len(text) + p.Header.Margin*2
	var marginString string
	for i := 0; i < p.Header.Margin; i++ {
		marginString += " "
	}
	var blankLine string
	for i := 0; i < textLength; i++ {
		blankLine += " "
	}

	var ret string

	ret += p.Header.BackgroundStyle.Sprint(blankLine) + "\n"
	ret += p.Header.BackgroundStyle.Sprint(p.Header.TextStyle.Sprint(marginString+text+marginString)) + "\n"
	ret += p.Header.BackgroundStyle.Sprint(blankLine) + "\n\n"

	return ret
}

// Sprintln formats using the default formats for its operands and returns the resulting string.
// Spaces are always added between operands and a newline is appended.
func (p HeaderPrinter) Sprintln(a ...interface{}) string {
	return Sprint(p.Sprint(a...) + "\n")
}

// Sprintf formats according to a format specifier and returns the resulting string.
func (p HeaderPrinter) Sprintf(format string, a ...interface{}) string {
	panic("implement me")
}

// Print formats using the default formats for its operands and writes to standard output.
// Spaces are added between operands when neither is a string.
// It returns the number of bytes written and any write error encountered.
func (p HeaderPrinter) Print(a ...interface{}) GenericPrinter {
	Print(p.Sprint(a...))
	return p
}

// Println formats using the default formats for its operands and writes to standard output.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
func (p HeaderPrinter) Println(a ...interface{}) GenericPrinter {
	Println(p.Sprint(a...))
	return p
}

// Printf formats according to a format specifier and writes to standard output.
// It returns the number of bytes written and any write error encountered.
func (p HeaderPrinter) Printf(format string, a ...interface{}) GenericPrinter {
	p.Print(Sprintf(format, a...))
	return p
}