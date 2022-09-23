package router

import (
	"bytes"
	"fmt"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"pdf-server/core"
	"pdf-server/dto"
	"pdf-server/macro"
)

type WKRouter struct{}

func NewWKRouter() *WKRouter {
	return &WKRouter{}
}

func (w *WKRouter) Convert(ctx *core.Context) {
	req := &dto.WKConvertReq{}
	err := ctx.ShouldBindJSON(req)
	if err != nil {
		ctx.AbortBadRequest()
		return
	}
	switch req.Type {
	case macro.ConvertFromText:
		if len(req.Text) == 0 {
			ctx.AbortBadRequest()
			return
		}
	case macro.ConvertFromUrl:
		if req.Url == nil {
			ctx.AbortBadRequest()
			return
		}
	}
	// generate pdf
	generator, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		ctx.AbortWithError(err)
		return
	}
	wkConvertSetGlobalOptions(req.GlobalOptions, generator)
	var provider wkhtmltopdf.PageProvider
	switch req.Type {
	case macro.ConvertFromUrl:
		pg := wkhtmltopdf.NewPage(*req.Url)
		wkConvertPageOptions(req.PageOptions, &pg.PageOptions)
		provider = pg
	case macro.ConvertFromText:
		pg := wkhtmltopdf.NewPageReader(bytes.NewReader(req.Text))
		wkConvertPageOptions(req.PageOptions, &pg.PageOptions)
		provider = pg
	}
	generator.AddPage(provider)
	err = generator.Create()
	if err != nil {
		ctx.AbortWithError(err)
		return
	}
	// response
	if req.Download {
		var filename string
		if req.Filename != nil {
			filename = *req.Filename
		} else {
			filename = fmt.Sprintf("%s.pdf", uuid.New().String())
		}
		ctx.Header("Content-Disposition", "attachment; filename="+filename)
		ctx.Data(http.StatusOK, "application/octet-stream", generator.Bytes())
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"data": generator.Bytes(),
		})
	}
}

func wkConvertSetGlobalOptions(opts *dto.WKConvertGlobalOption, generator *wkhtmltopdf.PDFGenerator) {
	if opts == nil {
		return
	}
	if opts.CookieJar != nil {
		generator.CookieJar.Set(*opts.CookieJar)
	}
	if opts.Copies != nil {
		generator.Copies.Set(*opts.Copies)
	}
	if opts.Dpi != nil {
		generator.Dpi.Set(*opts.Dpi)
	}
	if opts.ExtendedHelp != nil {
		generator.ExtendedHelp.Set(*opts.ExtendedHelp)
	}
	if opts.Grayscale != nil {
		generator.Grayscale.Set(*opts.Grayscale)
	}
	if opts.Help != nil {
		generator.Help.Set(*opts.Help)
	}
	if opts.HTMLDoc != nil {
		generator.HTMLDoc.Set(*opts.HTMLDoc)
	}
	if opts.ImageDpi != nil {
		generator.ImageDpi.Set(*opts.ImageDpi)
	}
	if opts.ImageQuality != nil {
		generator.ImageQuality.Set(*opts.ImageQuality)
	}
	if opts.License != nil {
		generator.License.Set(*opts.License)
	}
	if opts.LowQuality != nil {
		generator.LowQuality.Set(*opts.LowQuality)
	}
	if opts.ManPage != nil {
		generator.ManPage.Set(*opts.ManPage)
	}
	if opts.MarginBottom != nil {
		generator.MarginBottom.Set(*opts.MarginBottom)
	}
	if opts.MarginLeft != nil {
		generator.MarginLeft.Set(*opts.MarginLeft)
	}
	if opts.MarginRight != nil {
		generator.MarginRight.Set(*opts.MarginRight)
	}
	if opts.MarginTop != nil {
		generator.MarginTop.Set(*opts.MarginTop)
	}
	if opts.NoCollate != nil {
		generator.NoCollate.Set(*opts.NoCollate)
	}
	if opts.NoPdfCompression != nil {
		generator.NoPdfCompression.Set(*opts.NoPdfCompression)
	}
	if opts.Orientation != nil {
		generator.Orientation.Set(*opts.Orientation)
	}
	if opts.PageHeight != nil {
		generator.PageHeight.Set(*opts.PageHeight)
	}
	if opts.PageSize != nil {
		generator.PageSize.Set(*opts.PageSize)
	}
	if opts.PageWidth != nil {
		generator.PageWidth.Set(*opts.PageWidth)
	}
	if opts.Quiet != nil {
		generator.Quiet.Set(*opts.Quiet)
	}
	if opts.Readme != nil {
		generator.Readme.Set(*opts.Readme)
	}
	if opts.Title != nil {
		generator.Title.Set(*opts.Title)
	}
	if opts.Version != nil {
		generator.Version.Set(*opts.Version)
	}
}

func wkConvertPageOptions(opts *dto.WKConvertPageOption, ans *wkhtmltopdf.PageOptions) {
	if opts == nil {
		return
	}
	for _, item := range opts.Allow {
		ans.Allow.Set(item)
	}
	for _, item := range opts.BypassProxyFor {
		ans.BypassProxyFor.Set(item)
	}
	if opts.CacheDir != nil {
		ans.CacheDir.Set(*opts.CacheDir)
	}
	if opts.CheckboxCheckedSvg != nil {
		ans.CheckboxCheckedSvg.Set(*opts.CheckboxCheckedSvg)
	}
	if opts.CheckboxSvg != nil {
		ans.CheckboxSvg.Set(*opts.CheckboxSvg)
	}
	for key, val := range opts.Cookie {
		ans.Cookie.Set(key, val)
	}
	for key, val := range opts.CustomHeader {
		ans.CustomHeader.Set(key, val)
	}
	if opts.CustomHeaderPropagation != nil {
		ans.CustomHeaderPropagation.Set(*opts.CustomHeaderPropagation)
	}
	if opts.DebugJavascript != nil {
		ans.DebugJavascript.Set(*opts.DebugJavascript)
	}
	if opts.DefaultHeader != nil {
		ans.DefaultHeader.Set(*opts.DefaultHeader)
	}
	if opts.DisableExternalLinks != nil {
		ans.DisableExternalLinks.Set(*opts.DisableExternalLinks)
	}
	if opts.DisableInternalLinks != nil {
		ans.DisableInternalLinks.Set(*opts.DisableInternalLinks)
	}
	if opts.DisableJavascript != nil {
		ans.DisableJavascript.Set(*opts.DisableJavascript)
	}
	if opts.DisableLocalFileAccess != nil {
		ans.DisableLocalFileAccess.Set(*opts.DisableLocalFileAccess)
	}
	if opts.DisableSmartShrinking != nil {
		ans.DisableSmartShrinking.Set(*opts.DisableSmartShrinking)
	}
	if opts.EnableForms != nil {
		ans.EnableForms.Set(*opts.EnableForms)
	}
	if opts.EnableLocalFileAccess != nil {
		ans.EnableLocalFileAccess.Set(*opts.EnableLocalFileAccess)
	}
	if opts.EnablePlugins != nil {
		ans.EnablePlugins.Set(*opts.EnablePlugins)
	}
	if opts.EnableTocBackLinks != nil {
		ans.EnableTocBackLinks.Set(*opts.EnableTocBackLinks)
	}
	if opts.Encoding != nil {
		ans.Encoding.Set(*opts.Encoding)
	}
	if opts.ExcludeFromOutline != nil {
		ans.ExcludeFromOutline.Set(*opts.ExcludeFromOutline)
	}
	if opts.JavascriptDelay != nil {
		ans.JavascriptDelay.Set(*opts.JavascriptDelay)
	}
	if opts.KeepRelativeLinks != nil {
		ans.KeepRelativeLinks.Set(*opts.KeepRelativeLinks)
	}
	if opts.LoadErrorHandling != nil {
		ans.LoadErrorHandling.Set(*opts.LoadErrorHandling)
	}
	if opts.LoadMediaErrorHandling != nil {
		ans.LoadMediaErrorHandling.Set(*opts.LoadMediaErrorHandling)
	}
	if opts.MinimumFontSize != nil {
		ans.MinimumFontSize.Set(*opts.MinimumFontSize)
	}
	if opts.NoBackground != nil {
		ans.NoBackground.Set(*opts.NoBackground)
	}
	if opts.NoCustomHeaderPropagation != nil {
		ans.NoCustomHeaderPropagation.Set(*opts.NoCustomHeaderPropagation)
	}
	if opts.NoImages != nil {
		ans.NoImages.Set(*opts.NoImages)
	}
	if opts.NoStopSlowScripts != nil {
		ans.NoStopSlowScripts.Set(*opts.NoStopSlowScripts)
	}
	if opts.PageOffset != nil {
		ans.PageOffset.Set(*opts.PageOffset)
	}
	if opts.Password != nil {
		ans.Password.Set(*opts.Password)
	}
	for key, val := range opts.Post {
		ans.Post.Set(key, val)
	}
	for key, val := range opts.PostFile {
		ans.PostFile.Set(key, val)
	}
	if opts.PrintMediaType != nil {
		ans.PrintMediaType.Set(*opts.PrintMediaType)
	}
	if opts.Proxy != nil {
		ans.Proxy.Set(*opts.Proxy)
	}
	if opts.ProxyHostnameLookup != nil {
		ans.ProxyHostnameLookup.Set(*opts.ProxyHostnameLookup)
	}
	if opts.RadiobuttonCheckedSvg != nil {
		ans.RadiobuttonCheckedSvg.Set(*opts.RadiobuttonCheckedSvg)
	}
	if opts.RadiobuttonSvg != nil {
		ans.RadiobuttonSvg.Set(*opts.RadiobuttonSvg)
	}
	for _, item := range opts.RunScript {
		ans.RunScript.Set(item)
	}
	if opts.SslCrtPath != nil {
		ans.SslCrtPath.Set(*opts.SslCrtPath)
	}
	if opts.SslKeyPassword != nil {
		ans.SslKeyPassword.Set(*opts.SslKeyPassword)
	}
	if opts.SslKeyPath != nil {
		ans.SslKeyPath.Set(*opts.SslKeyPath)
	}
	if opts.Username != nil {
		ans.Username.Set(*opts.Username)
	}
	if opts.UserStyleSheet != nil {
		ans.UserStyleSheet.Set(*opts.UserStyleSheet)
	}
	if opts.ViewportSize != nil {
		ans.ViewportSize.Set(*opts.ViewportSize)
	}
	if opts.WindowStatus != nil {
		ans.WindowStatus.Set(*opts.WindowStatus)
	}
	if opts.Zoom != nil {
		ans.Zoom.Set(*opts.Zoom)
	}
	if opts.FooterCenter != nil {
		ans.FooterCenter.Set(*opts.FooterCenter)
	}
	if opts.FooterFontName != nil {
		ans.FooterFontName.Set(*opts.FooterFontName)
	}
	if opts.FooterFontSize != nil {
		ans.FooterFontSize.Set(*opts.FooterFontSize)
	}
	if opts.FooterHTML != nil {
		ans.FooterHTML.Set(*opts.FooterHTML)
	}
	if opts.FooterLeft != nil {
		ans.FooterLeft.Set(*opts.FooterLeft)
	}
	if opts.FooterLine != nil {
		ans.FooterLine.Set(*opts.FooterLine)
	}
	if opts.FooterRight != nil {
		ans.FooterRight.Set(*opts.FooterRight)
	}
	if opts.FooterSpacing != nil {
		ans.FooterSpacing.Set(*opts.FooterSpacing)
	}
	if opts.HeaderCenter != nil {
		ans.HeaderCenter.Set(*opts.HeaderCenter)
	}
	if opts.HeaderFontName != nil {
		ans.HeaderFontName.Set(*opts.HeaderFontName)
	}
	if opts.HeaderFontSize != nil {
		ans.HeaderFontSize.Set(*opts.HeaderFontSize)
	}
	if opts.HeaderHTML != nil {
		ans.HeaderHTML.Set(*opts.HeaderHTML)
	}
	if opts.HeaderLeft != nil {
		ans.HeaderLeft.Set(*opts.HeaderLeft)
	}
	if opts.HeaderLine != nil {
		ans.HeaderLine.Set(*opts.HeaderLine)
	}
	if opts.HeaderRight != nil {
		ans.HeaderRight.Set(*opts.HeaderRight)
	}
	if opts.HeaderSpacing != nil {
		ans.HeaderSpacing.Set(*opts.HeaderSpacing)
	}
	for key, val := range opts.Replace {
		ans.Replace.Set(key, val)
	}
}
