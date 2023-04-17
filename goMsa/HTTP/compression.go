package main

import (
	"compress/gzip"
	"net/http"
	"strings"
)

// HTTP의 압축 방법 Gzip괴 deflate

type GzipResponseWritter struct {
	gw *gzip.Writer
	http.ResponseWriter
}

type GZipHandler struct {
	next http.Handler
}

func (h *GZipHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	encodings := r.Header.Get("Accept-Encoding")

	if strings.Contains(encodings, "gzip") {
		h.serveGzipped(w, r)
	} else if strings.Contains(encodings, "deflate") {
		panic("Deflate not implemented")
	} else { // 클라이언트가 압축되지 않은 컨텐츠를 요구한 경우
		h.servePlain(w, r)
	}
}

func (h *GZipHandler) serveGzipped(w http.ResponseWriter, r *http.Request) {
	gzw := gzip.NewWriter(w)
	defer gzw.Close()

	w.Header().Set("Content-Encoding", "gzip")
	h.next.ServeHTTP(GzipResponseWritter{gzw, w}, r)
}

func (h *GZipHandler) servePlain(w http.ResponseWriter, r *http.Request) {
	h.next.ServeHTTP(w, r)
}

func (w GzipResponseWritter) write(b []byte) (int, error) {
	if _, ok := w.Header()["Content-Type"]; !ok {
		// 콘텐츠 타입이 설정되지 않는 경우 압축되지 않은 본문에서 콘텐츠 타입을 유추한다.
		w.Header().Set("Content-Type", http.DetectContentType(b))
	}
	return w.gw.Write(b)
}
