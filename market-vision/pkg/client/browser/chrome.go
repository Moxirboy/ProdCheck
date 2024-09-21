package browser

import configs "market-vision/internal/config"
import "context"
import "github.com/chromedp/chromedp"
import "time"

type ChromeClient struct {
    ctx    context.Context
    cancel context.CancelFunc
}

// InitializeChrome sets up and returns a ChromeClient
func InitializeChrome(cfg configs.Config) *ChromeClient {
    opts := []chromedp.ExecAllocatorOption{
        chromedp.Flag("headless", true),
    }

    allocCtx, cancelAlloc := chromedp.NewExecAllocator(context.Background(), opts...)
    ctx, ctxCancel := chromedp.NewContext(allocCtx)
    ctx, cancelWithTimeout := context.WithTimeout(ctx, cfg.Browser.Timeout)

    return &ChromeClient{
        ctx: ctx,
        cancel: func() {
            cancelWithTimeout()
            ctxCancel()
            cancelAlloc()
        },
    }
}


func (client *ChromeClient) FetchHTML(url string, waitTime time.Duration) (string, error) {
    var htmlContent string
    err := chromedp.Run(client.ctx,
        chromedp.Navigate(url),
        chromedp.WaitReady("html"),
		
        chromedp.OuterHTML("html", &htmlContent),
    )
    if err != nil {
        return "", err
    }
    return htmlContent, nil
}