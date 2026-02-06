package e2e

import (
	"context"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/chromedp/chromedp"
)

var _ = Describe("Todo App E2E", func() {
	It("should persist item after refresh", func() {

		opts := append(chromedp.DefaultExecAllocatorOptions[:],
			chromedp.Flag("headless", "new"),
			chromedp.Flag("disable-gpu", true),
			chromedp.Flag("no-sandbox", true),
			chromedp.Flag("disable-dev-shm-usage", true),
			chromedp.Flag("remote-debugging-port", "9222"),
		)

		allocCtx, cancel := chromedp.NewExecAllocator(
			context.Background(),
			opts...,
		)
		defer cancel()

		ctx, cancel := chromedp.NewContext(
			allocCtx,
			chromedp.WithLogf(GinkgoWriter.Printf),
		)
		defer cancel()

		ctx, cancel = context.WithTimeout(ctx, 60*time.Second)
		defer cancel()

		var (
			runErr error
		)

		// Initial actions
		runErr = chromedp.Run(ctx,
			chromedp.Navigate("http://localhost:8080"),
			chromedp.WaitVisible("#item"),
			chromedp.SendKeys("#item", "test-data"),
			chromedp.Click("button"),
		)

		Expect(runErr).ToNot(HaveOccurred())

		// Eventually with error handling
		Eventually(func() string {
			var text string
			chromedp.Run(ctx,
				chromedp.Text("body", &text),
			)
			return text
		}, 15*time.Second, 500*time.Millisecond).
			Should(ContainSubstring("test-data"))

	})
})
