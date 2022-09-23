package dto

import "pdf-server/macro"

type WKConvertReq struct {
	Type          macro.PDFConvertType   `json:"type" binding:"required"`                    // Source type [1] url [2] text
	Download      bool                   `json:"download" binding:"omitempty"`               // Is download
	Filename      *string                `json:"filename" binding:"omitempty,min=1,max=100"` // Download file name
	Text          []byte                 `json:"text" binding:"omitempty"`                   // Text
	Url           *string                `json:"url" binding:"omitempty,uri"`                // Url
	GlobalOptions *WKConvertGlobalOption `json:"global_options" binding:"omitempty,dive"`    // Global options
	PageOptions   *WKConvertPageOption   `json:"page_options" binding:"omitempty,dive"`      // Page options
}

type WKConvertGlobalOption struct {
	CookieJar        *string `json:"cookie_jar" binding:"omitempty"`         // Read and write cookies from and to the supplied cookie jar file
	Copies           *uint   `json:"copies" binding:"omitempty"`             // Number of copies to print into the pdf file (default 1)
	Dpi              *uint   `json:"dpi" binding:"omitempty"`                // Change the dpi explicitly (this has no effect on X11 based systems)
	ExtendedHelp     *bool   `json:"extended_help" binding:"omitempty"`      // Display more extensive help, detailing less common command switches
	Grayscale        *bool   `json:"grayscale" binding:"omitempty"`          // PDF will be generated in grayscale
	Help             *bool   `json:"help" binding:"omitempty"`               // Display help
	HTMLDoc          *bool   `json:"html_doc" binding:"omitempty"`           // Output program html help
	ImageDpi         *uint   `json:"image_dpi" binding:"omitempty"`          // When embedding images scale them down to this dpi (default 600)
	ImageQuality     *uint   `json:"image_quality" binding:"omitempty"`      // When jpeg compressing images use this quality (default 94)
	License          *bool   `json:"license" binding:"omitempty"`            // Output license information and exit
	LowQuality       *bool   `json:"low_quality" binding:"omitempty"`        // Generates lower quality pdf/ps. Useful to shrink the result document space
	ManPage          *bool   `json:"man_page" binding:"omitempty"`           // Output program man page
	MarginBottom     *uint   `json:"margin_bottom" binding:"omitempty"`      // Set the page bottom margin
	MarginLeft       *uint   `json:"margin_left" binding:"omitempty"`        // Set the page left margin (default 10mm)
	MarginRight      *uint   `json:"margin_right" binding:"omitempty"`       // Set the page right margin (default 10mm)
	MarginTop        *uint   `json:"margin_top" binding:"omitempty"`         // Set the page top margin
	NoCollate        *bool   `json:"no_collate" binding:"omitempty"`         // Do not collate when printing multiple copies (default collate)
	NoPdfCompression *bool   `json:"no_pdf_compression" binding:"omitempty"` // Do not use lossless compression on pdf objects
	Orientation      *string `json:"orientation" binding:"omitempty"`        // Set orientation to Landscape or Portrait (default Portrait)
	PageHeight       *uint   `json:"page_height" binding:"omitempty"`        // Page height
	PageSize         *string `json:"page_size" binding:"omitempty"`          // Set paper size to: A4, Letter, etc. (default A4)
	PageWidth        *uint   `json:"page_width" binding:"omitempty"`         // Page width
	Quiet            *bool   `json:"quiet" binding:"omitempty"`              // Be less verbose
	Readme           *bool   `json:"readme" binding:"omitempty"`             // Output program readme
	Title            *string `json:"title" binding:"omitempty"`              // The title of the generated pdf file (The title of the first document is used if not specified)
	Version          *bool   `json:"version" binding:"omitempty"`            // Output version information and exit
}

type WKConvertPageOption struct {
	Allow                     []string          `json:"allow" binding:"omitempty"`                        // Allow the file or files from the specified folder to be loaded (repeatable)
	BypassProxyFor            []string          `json:"bypass_proxy_for" binding:"omitempty"`             // Bypass proxy for host
	CacheDir                  *string           `json:"cache_dir" binding:"omitempty"`                    // Web cache directory
	CheckboxCheckedSvg        *string           `json:"checkbox_checked_svg" binding:"omitempty"`         // Use this SVG file when rendering checked checkboxes
	CheckboxSvg               *string           `json:"checkbox_svg" binding:"omitempty"`                 // Use this SVG file when rendering unchecked checkboxes
	Cookie                    map[string]string `json:"cookie" binding:"omitempty"`                       // Set an additional cookie (repeatable), value should be url encoded
	CustomHeader              map[string]string `json:"custom_header" binding:"omitempty"`                // Set an additional HTTP header (repeatable)
	CustomHeaderPropagation   *bool             `json:"custom_header_propagation" binding:"omitempty"`    // Add HTTP headers specified by --custom-header for each resource request
	DebugJavascript           *bool             `json:"debug_javascript" binding:"omitempty"`             // Show javascript debugging output
	DefaultHeader             *bool             `json:"default_header" binding:"omitempty"`               // Add a default header, with the name of the page to the left, and the page number to the right, this is short for: --header-left='[webpage]' --header-right='[page]/[toPage]' --top 2cm --header-line
	DisableExternalLinks      *bool             `json:"disable_external_links" binding:"omitempty"`       // Do not make links to remote web pages
	DisableInternalLinks      *bool             `json:"disable_internal_links" binding:"omitempty"`       // Do not make local links
	DisableJavascript         *bool             `json:"disable_javascript" binding:"omitempty"`           // Do not allow web pages to run javascript
	DisableLocalFileAccess    *bool             `json:"disable_local_file_access" binding:"omitempty"`    // Do not allowed conversion of a local file to read in other local files, unless explicitly allowed with --allow
	DisableSmartShrinking     *bool             `json:"disable_smart_shrinking" binding:"omitempty"`      // Disable the intelligent shrinking strategy used by WebKit that makes the pixel/dpi ratio none constant
	EnableForms               *bool             `json:"enable_forms" binding:"omitempty"`                 // Turn HTML form fields into pdf form fields
	EnableLocalFileAccess     *bool             `json:"enable_local_file_access" binding:"omitempty"`     // Allowed conversion of a local file to read in other local files
	EnablePlugins             *bool             `json:"enable_plugins" binding:"omitempty"`               // Enable installed plugins (plugins will likely not work)
	EnableTocBackLinks        *bool             `json:"enable_toc_back_links" binding:"omitempty"`        // Link from section header to toc
	Encoding                  *string           `json:"encoding" binding:"omitempty"`                     // Set the default text encoding, for input
	ExcludeFromOutline        *bool             `json:"exclude_from_outline" binding:"omitempty"`         // Do not include the page in the table of contents and outlines
	JavascriptDelay           *uint             `json:"javascript_delay" binding:"omitempty"`             // Wait some milliseconds for javascript finish (default 200)
	KeepRelativeLinks         *bool             `json:"keep_relative_links" binding:"omitempty"`          // Keep relative external links as relative external links
	LoadErrorHandling         *string           `json:"load_error_handling" binding:"omitempty"`          // Specify how to handle pages that fail to load: abort, ignore or skip (default abort)
	LoadMediaErrorHandling    *string           `json:"load_media_error_handling" binding:"omitempty"`    // Specify how to handle media files that fail to load: abort, ignore or skip (default ignore)
	MinimumFontSize           *uint             `json:"minimum_font_size" binding:"omitempty"`            // Minimum font size
	NoBackground              *bool             `json:"no_background" binding:"omitempty"`                // Do not print background
	NoCustomHeaderPropagation *bool             `json:"no_custom_header_propagation" binding:"omitempty"` // Do not add HTTP headers specified by --custom-header for each resource request
	NoImages                  *bool             `json:"no_images" binding:"omitempty"`                    // Do not load or print images
	NoStopSlowScripts         *bool             `json:"no_stop_slow_scripts" binding:"omitempty"`         // Do not Stop slow running javascripts
	PageOffset                *uint             `json:"page_offset" binding:"omitempty"`                  // Set the starting page number (default 0)
	Password                  *string           `json:"password" binding:"omitempty"`                     // HTTP Authentication password
	Post                      map[string]string `json:"post" binding:"omitempty"`                         // Add an additional post field (repeatable)
	PostFile                  map[string]string `json:"post_file" binding:"omitempty"`                    // Post an additional file (repeatable)
	PrintMediaType            *bool             `json:"print_media_type" binding:"omitempty"`             // Use print media-type instead of screen
	Proxy                     *string           `json:"proxy" binding:"omitempty"`                        // Use a proxy
	ProxyHostnameLookup       *bool             `json:"proxy_hostname_lookup" binding:"omitempty"`        // Use the proxy for resolving hostnames
	RadiobuttonCheckedSvg     *string           `json:"radiobutton_checked_svg" binding:"omitempty"`      // Use this SVG file when rendering checked radiobuttons
	RadiobuttonSvg            *string           `json:"radiobutton_svg" binding:"omitempty"`              // Use this SVG file when rendering unchecked radiobuttons
	RunScript                 []string          `json:"run_script" binding:"omitempty"`                   // Run this additional javascript after the page is done loading (repeatable)
	SslCrtPath                *string           `json:"ssl_crt_path" binding:"omitempty"`                 // Path to the ssl client cert public key in OpenSSL PEM format, optionally followed by intermediate ca and trusted certs
	SslKeyPassword            *string           `json:"ssl_key_password" binding:"omitempty"`             // Password to ssl client cert private key
	SslKeyPath                *string           `json:"ssl_key_path" binding:"omitempty"`                 // Path to ssl client cert private key in OpenSSL PEM format
	Username                  *string           `json:"username" binding:"omitempty"`                     // HTTP Authentication username
	UserStyleSheet            *string           `json:"user_style_sheet" binding:"omitempty"`             // Specify a user style sheet, to load with every page
	ViewportSize              *string           `json:"viewport_size" binding:"omitempty"`                // Set viewport size if you have custom scrollbars or css attribute overflow to emulate window size
	WindowStatus              *string           `json:"window_status" binding:"omitempty"`                // Wait until window.status is equal to this string before rendering page
	Zoom                      *float64          `json:"zoom" binding:"omitempty"`                         // Use this zoom factor (default 1)
	FooterCenter              *string           `json:"footer_center" binding:"omitempty"`                // Centered footer text
	FooterFontName            *string           `json:"footer_font_name" binding:"omitempty"`             // Set footer font name (default Arial)
	FooterFontSize            *uint             `json:"footer_font_size" binding:"omitempty"`             // Set footer font size (default 12)
	FooterHTML                *string           `json:"footer_html" binding:"omitempty"`                  // Adds a html footer
	FooterLeft                *string           `json:"footer_left" binding:"omitempty"`                  // Left aligned footer text
	FooterLine                *bool             `json:"footer_line" binding:"omitempty"`                  // Display line above the footer
	FooterRight               *string           `json:"footer_right" binding:"omitempty"`                 // Right aligned footer text
	FooterSpacing             *float64          `json:"footer_spacing" binding:"omitempty"`               // Spacing between footer and content in mm (default 0)
	HeaderCenter              *string           `json:"header_center" binding:"omitempty"`                // Centered header text
	HeaderFontName            *string           `json:"header_font_name" binding:"omitempty"`             // Set header font name (default Arial)
	HeaderFontSize            *uint             `json:"header_font_size" binding:"omitempty"`             // Set header font size (default 12)
	HeaderHTML                *string           `json:"header_html" binding:"omitempty"`                  // Adds a html header
	HeaderLeft                *string           `json:"header_left" binding:"omitempty"`                  // Left aligned header text
	HeaderLine                *bool             `json:"header_line" binding:"omitempty"`                  // Display line below the header
	HeaderRight               *string           `json:"header_right" binding:"omitempty"`                 // Right aligned header text
	HeaderSpacing             *float64          `json:"header_spacing" binding:"omitempty"`               // Spacing between header and content in mm (default 0)
	Replace                   map[string]string `json:"replace" binding:"omitempty"`
}
