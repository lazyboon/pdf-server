package macro

type PDFConvertType uint8

const (
	ConvertFromUrl  PDFConvertType = iota + 1 // From URL
	ConvertFromText                           // From HTML text
)
