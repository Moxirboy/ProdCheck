package browser

import configs "engine/internal/configs"
import "context"
import "github.com/chromedp/chromedp"

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


func (client *ChromeClient) CaptureScreenshot(url string) ([]byte, error) {
    var screenshot []byte
    err := chromedp.Run(client.ctx,
        chromedp.Navigate(url),
        chromedp.FullScreenshot(&screenshot,90),
    )
    if err != nil {
        return []byte{}, err
    }
    return screenshot, nil
}