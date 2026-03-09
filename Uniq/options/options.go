package options

type Options struct {
	Count      bool // -c flag
	Repeated   bool // -d flag
	Unique     bool // -u flag
	IgnoreCase bool // -i flag
	SkipFields int  // -f flag
	SkipChars  int  // -s flag
}
