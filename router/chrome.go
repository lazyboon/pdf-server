package router

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"net/http/httptest"
	"pdf-server/core"
	"pdf-server/dto"
	"pdf-server/macro"
	"time"
)

type ChromeRouter struct{}

func NewChromeRouter() *ChromeRouter {
	return &ChromeRouter{}
}

func (c *ChromeRouter) Convert(ctx *core.Context) {
	// request param
	req := &dto.ChromeConvertReq{}
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
	dpCtx, cancel := chromedp.NewContext(ctx, chromedp.WithBrowserOption(chromedp.WithDialTimeout(60*time.Second)))
	defer cancel()
	var url string
	switch req.Type {
	case macro.ConvertFromUrl:
		url = *req.Url
	case macro.ConvertFromText:
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/html")
			_, _ = fmt.Fprintf(w, string(req.Text))
		}))
		defer ts.Close()
		url = ts.URL
	}

	var ans []byte
	err = chromedp.Run(
		dpCtx,
		chromedp.Navigate(url),
		chromedp.ActionFunc(func(ctx context.Context) error {
			params := page.PrintToPDF()
			opts := req.Options
			if opts != nil {
				if opts.Landscape != nil {
					params = params.WithLandscape(*opts.Landscape)
				}
				if opts.DisplayHeaderFooter != nil {
					params = params.WithDisplayHeaderFooter(*opts.DisplayHeaderFooter)
				}
				if opts.PrintBackground != nil {
					params = params.WithPrintBackground(*opts.PrintBackground)
				}
				if opts.Scale != nil {
					params = params.WithScale(*opts.Scale)
				}
				if opts.PaperWidth != nil {
					params = params.WithPaperWidth(*opts.PaperWidth)
				}
				if opts.PaperHeight != nil {
					params = params.WithPaperHeight(*opts.PaperHeight)
				}
				if opts.MarginTop != nil {
					params = params.WithMarginTop(*opts.MarginTop)
				}
				if opts.MarginBottom != nil {
					params = params.WithMarginBottom(*opts.MarginBottom)
				}
				if opts.MarginLeft != nil {
					params = params.WithMarginLeft(*opts.MarginLeft)
				}
				if opts.MarginRight != nil {
					params = params.WithMarginRight(*opts.MarginRight)
				}
				if opts.PageRanges != nil {
					params = params.WithPageRanges(*opts.PageRanges)
				}
				if opts.HeaderTemplate != nil {
					params = params.WithHeaderTemplate(*opts.HeaderTemplate)
				}
				if opts.FooterTemplate != nil {
					params = params.WithFooterTemplate(*opts.FooterTemplate)
				}
				if opts.PreferCSSPageSize != nil {
					params = params.WithPreferCSSPageSize(*opts.PreferCSSPageSize)
				}
				if opts.TransferMode != nil {
					params = params.WithTransferMode(*opts.TransferMode)
				}
			}
			ans, _, err = params.Do(ctx)
			return err
		}),
	)
	if err != nil {
		ctx.AbortWithError(err)
		return
	}

	// response
	// response
	if req.Download {
		var filename string
		if req.Filename != nil {
			filename = *req.Filename
		} else {
			filename = fmt.Sprintf("%s.pdf", uuid.New().String())
		}
		ctx.Header("Content-Disposition", "attachment; filename="+filename)
		ctx.Data(http.StatusOK, "application/octet-stream", ans)
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"data": ans,
		})
	}
}
