package integration

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"

	"e2e-demo/app"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("HomeHandler Integration Test", func() {

	It("accepts POST request and redirects correctly", func() {
		// Arrange
		store := app.NewStore()

		// Minimal template (no filesystem dependency)
		tmpl := template.Must(
			template.New("test").Parse("<html><body>{{.}}</body></html>"),
		)

		handler := app.HomeHandler(store, tmpl)

		form := url.Values{}
		form.Add("item", "Learn E2E")

		req := httptest.NewRequest(
			http.MethodPost,
			"/",
			strings.NewReader(form.Encode()),
		)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		rr := httptest.NewRecorder()

		// Act
		handler.ServeHTTP(rr, req)

		// Assert
		Expect(rr.Code).To(Equal(http.StatusSeeOther))
		Expect(rr.Header().Get("Location")).To(Equal("/"))
	})

})
