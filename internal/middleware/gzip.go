package middleware

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"
)

/* ---------- RESPONSE (compress) ---------- */

type compressWriter struct {
	w          http.ResponseWriter
	zw         *gzip.Writer
	compressed bool
}

func newCompressWriter(w http.ResponseWriter) *compressWriter {
	return &compressWriter{
		w:  w,
		zw: gzip.NewWriter(w),
	}
}

func (c *compressWriter) Header() http.Header {
	return c.w.Header()
}

func (c *compressWriter) Write(b []byte) (int, error) {

	ct := c.w.Header().Get("Content-Type")

	if strings.Contains(ct, "application/json") ||
		strings.Contains(ct, "text/html") {

		c.compressed = true
		c.w.Header().Set("Content-Encoding", "gzip")

		return c.zw.Write(b)
	}

	return c.w.Write(b)
}

func (c *compressWriter) WriteHeader(statusCode int) {
	ct := c.w.Header().Get("Content-Type")

	if strings.Contains(ct, "application/json") || strings.Contains(ct, "text/html") {
		c.w.Header().Set("Content-Encoding", "gzip")
	}

	c.w.WriteHeader(statusCode)
}

func (c *compressWriter) Close() error {
	if c.compressed {
		return c.zw.Close()
	}

	return nil
}

/* ---------- REQUEST (decompress) ---------- */

type compressReader struct {
	r  io.ReadCloser
	zr *gzip.Reader
}

func newCompressReader(r io.ReadCloser) (*compressReader, error) {
	zr, err := gzip.NewReader(r)
	if err != nil {
		return nil, err
	}

	return &compressReader{r: r, zr: zr}, nil
}

func (c *compressReader) Read(p []byte) (int, error) {
	return c.zr.Read(p)
}

func (c *compressReader) Close() error {
	if err := c.zr.Close(); err != nil {
		return err
	}
	return c.r.Close()
}

/* ---------- MIDDLEWARE ---------- */

func GzipMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// 1. decompress request
		if r.Header.Get("Content-Encoding") == "gzip" {
			cr, err := newCompressReader(r.Body)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			r.Body = cr
			defer cr.Close()
		}

		// 2. compress response
		if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			cw := newCompressWriter(w)
			defer cw.Close()
			w = cw
		}

		// 3. call next handler
		next.ServeHTTP(w, r)
	})
}
