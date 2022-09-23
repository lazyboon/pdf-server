package dto

import (
	"github.com/chromedp/cdproto/page"
	"pdf-server/macro"
)

type ChromeConvertReq struct {
	Type     macro.PDFConvertType  `json:"type" binding:"required"`                    // Source type [1] url [2] text
	Download bool                  `json:"download" binding:"omitempty"`               // Is download
	Filename *string               `json:"filename" binding:"omitempty,min=1,max=100"` // Download file name
	Text     []byte                `json:"text" binding:"omitempty"`
	Url      *string               `json:"url" binding:"omitempty,uri"`
	Options  *ChromeConvertOptions `json:"options" binding:"omitempty,dive"`
}

type ChromeConvertOptions struct {
	Landscape           *bool                        `json:"landscape" binding:"omitempty"`             // Paper orientation. Defaults to false.
	DisplayHeaderFooter *bool                        `json:"display_header_footer" binding:"omitempty"` // Display header and footer. Defaults to false.
	PrintBackground     *bool                        `json:"print_background" binding:"omitempty"`      // Print background graphics. Defaults to false.
	Scale               *float64                     `json:"scale" binding:"omitempty"`                 // Scale of the webpage rendering. Defaults to 1.
	PaperWidth          *float64                     `json:"paper_width" binding:"omitempty"`           // Paper width in inches. Defaults to 8.5 inches.
	PaperHeight         *float64                     `json:"paper_height" binding:"omitempty"`          // Paper height in inches. Defaults to 11 inches.
	MarginTop           *float64                     `json:"margin_top" binding:"omitempty"`            // Top margin in inches. Defaults to 1cm (~0.4 inches).
	MarginBottom        *float64                     `json:"margin_bottom" binding:"omitempty"`         // Bottom margin in inches. Defaults to 1cm (~0.4 inches).
	MarginLeft          *float64                     `json:"margin_left" binding:"omitempty"`           // Left margin in inches. Defaults to 1cm (~0.4 inches).
	MarginRight         *float64                     `json:"margin_right" binding:"omitempty"`          // Right margin in inches. Defaults to 1cm (~0.4 inches).
	PageRanges          *string                      `json:"page_ranges" binding:"omitempty"`           // Paper ranges to print, one based, e.g., '1-5, 8, 11-13'. Pages are printed in the document order, not in the order specified, and no more than once. Defaults to empty string, which implies the entire document is printed. The page numbers are quietly capped to actual page count of the document, and ranges beyond the end of the document are ignored. If this results in no pages to print, an error is reported. It is an error to specify a range with start greater than end.
	HeaderTemplate      *string                      `json:"header_template" binding:"omitempty"`       // HTML template for the print header. Should be valid HTML markup with following classes used to inject printing values into them: - date: formatted print date - title: document title - url: document location - pageNumber: current page number - totalPages: total pages in the document  For example, <span class=title></span> would generate span containing the title.
	FooterTemplate      *string                      `json:"footer_template" binding:"omitempty"`       // HTML template for the print footer. Should use the same format as the headerTemplate.
	PreferCSSPageSize   *bool                        `json:"prefer_css_page_size" binding:"omitempty"`  // Whether or not to prefer page size as defined by css. Defaults to false, in which case the content will be scaled to fit the paper size.
	TransferMode        *page.PrintToPDFTransferMode `json:"transfer_mode" binding:"omitempty"`         // return as stream
}
