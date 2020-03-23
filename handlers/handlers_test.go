package handlers

import (
    "net/http"
    "net/http/httptest"
    "github.com/joho/godotenv"
     "strings"
    "testing"
)

func TestGetNewsHandler(t *testing.T) {
     r, err := http.NewRequest("GET", "/news", nil)
     if err != nil {
          t.Fatal(err)
     }
     w := httptest.NewRecorder()
     handle := http.HandlerFunc(GetNews)
     handle.ServeHTTP(w, r)

     if err := godotenv.Load("../.env"); err != nil {
		t.Error("xxx .env file")
     }

     resp := w.Result()
     if resp.StatusCode != http.StatusOK {
          t.Errorf("Unexpected status code %d", resp.StatusCode)
     }
}

func TestPostNewsHandlerSuccess(t *testing.T) {
     r, err := http.NewRequest("POST", "/news", strings.NewReader("author=TestAuthor&body=TestContent"))
     r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
     if err != nil {
          t.Fatal(err)
     }

     if author := r.FormValue("body"); author == "" {
		t.Errorf(`req.FormValue("body") = %s, can not be empty`, author)
     }

     if body := r.FormValue("author"); body == "" {
		t.Errorf(`req.FormValue("author") = %s, can not be empty`, body)
     }
}

func TestPostNewsHandlerAuthorFailed(t *testing.T) {
     // body set to empty string
     r, err := http.NewRequest("POST", "/news", strings.NewReader("author=&body=TestContent"))
     r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
     if err != nil {
          t.Fatal(err)
     }

     w := httptest.NewRecorder()
     handle := http.HandlerFunc(PostNews)
     handle.ServeHTTP(w, r)

     resp := w.Result()
     if resp.StatusCode != http.StatusBadRequest {
          t.Errorf("Unexpected status code %d", resp.StatusCode)
     }
}

func TestPostNewsHandlerBodyFailed(t *testing.T) {
     r, err := http.NewRequest("POST", "/news", strings.NewReader("author=TestAuthor&body="))
     r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
     if err != nil {
          t.Fatal(err)
     }

     w := httptest.NewRecorder()
     handle := http.HandlerFunc(PostNews)
     handle.ServeHTTP(w, r)

     resp := w.Result()
     if resp.StatusCode != http.StatusBadRequest {
          t.Errorf("Unexpected status code %d", resp.StatusCode)
     }
}