package app

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("HomeHandler (Unit Test)", func() {

	It("returns redirect after POST request", func() {
		store := NewStore()

		tmpl := template.Must(
			template.New("test").Parse("body"),
		)
		handler := HomeHandler(store, tmpl)

		form := url.Values{}
		form.Add("item", "Learn E2E")

		req := httptest.NewRequest(
			http.MethodPost,
			"/",
			strings.NewReader(form.Encode()),
		)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		rr := httptest.NewRecorder()

		handler.ServeHTTP(rr, req)

		Expect(rr.Code).To(Equal(http.StatusSeeOther))
		Expect(rr.Header().Get("Location")).To(Equal("/"))
	})

	It("renders page on GET request", func() {
		store := NewStore()

		tmpl := template.Must(
			template.New("test").Parse("body"),
		)
		handler := HomeHandler(store, tmpl)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rr := httptest.NewRecorder()

		handler.ServeHTTP(rr, req)

		Expect(rr.Code).To(Equal(http.StatusOK))
		Expect(rr.Body.String()).To(ContainSubstring("Todo"))
	})
})
