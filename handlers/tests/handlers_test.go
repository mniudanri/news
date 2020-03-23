package handlers

import (
    "net/http"
    "net/http/httptest"
     // "news-app/handlers"
     // "fmt"
     "strings"
    "testing"
)

// func TestNewsHandler(t *testing.T) {
//      r, err := http.NewRequest("GET", "/news", nil)
//      if err != nil {
//           t.Fatal(err)
//      }
//      w := httptest.NewRecorder()
//      handler := http.HandlerFunc(MyLoginHandler)
//      handler.ServeHTTP(w, r)

//      resp := w.Result()

//      if resp.StatusCode != http.StatusOK {
//           t.Errorf("Unexpected status code %d", resp.StatusCode)
//      }
// }

func TestNewsHandler(t *testing.T) {
     r, err := http.NewRequest("GET", "/news", nil)
     if err != nil {
          t.Fatal(err)
     }
     w := httptest.NewRecorder()
     handler := http.HandlerFunc(GetNews)
     handler.ServeHTTP(w, r)

     resp := w.Result()

     if resp.StatusCode != http.StatusOK {
          t.Errorf("Unexpected status code %d", resp.StatusCode)
     }
}

func TestPostGetNewsHandler(t *testing.T) {
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

     // if z := r.Method(); z != "POST" {
     //      t.Errorf(`must accept POST method`, q)
     // }
     // w := httptest.NewRecorder()
     // handler := http.HandlerFunc(PostNews)
     // handler.ServeHTTP(w, r)

     // r.ParseForm()
     // author := r.Form.Get("author")
     // body := r.Form.Get("body")

     // if author != "" {
     //      t.Errorf("Expected request to have 'topic=meaningful-topic', got: '%s'", author)
     // }

     // if body != "" {
     //      t.Errorf("Expected request to have 'topic=meaningful-topic', got: '%s'", body)
     // }



     
     // resp := w.Result()

     // if resp.StatusCode != http.StatusOK {
     //      t.Errorf("Unexpected status code %d", resp.StatusCode)
     // }
}
// func TestPublishOK(t *testing.T) {
//   ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//      w.WriteHeader(http.StatusOK)
//      if r.Method != "POST" && r.Method != "GET" {
//           t.Errorf("Expected 'POST' and 'GET' request, got '%s'", r.Method)
//      }
//      rr := httptest.NewRecorder()

//      if r.URL.EscapedPath() != "/news" {
//           if status := rr.Code; status != http.StatusOK && status != http.StatusOK {
//                t.Errorf("Expected request to '/news', got '%s'", r.URL.EscapedPath())
//           }
//      }
//      // r.ParseForm()
//      // topic := r.Form.Get("topic")
//      // if topic != "meaningful-topic" {  
//      //      t.Errorf("Expected request to have 'topic=meaningful-topic', got: '%s'", topic)
//      // }
//   }))

//   defer ts.Close()
//   nsqdUrl := ts.URL
//   err := Publish(nsqdUrl, "hello")
//   if err != nil {
//     t.Errorf("Publish() returned an error: %s", err)
//   }
// }