package handlers

import(
    "news-app/background/jobs"
    "net/http"
    "database/sql"
    // "html/template"
    _ "github.com/lib/pq"
    "strconv"
    "encoding/json"
    "errors"
	"news-app/models"
    "news-app/connection"
)

var (
	// tpl *template.Template
	db *sql.DB
)

func init(){
   db = connection.Connect()
}

func GetNews(w http.ResponseWriter, r *http.Request) {
    pPage, _ := r.URL.Query()["page"]
    pLimit, _ := r.URL.Query()["limit"]
    var page int64 = 1
    var limit int64 = 10

    if len(pPage) != 0 {
        page, _ = strconv.ParseInt(pPage[0], 0, 32)
    }
    if len(pLimit) != 0 {
        limit, _ = strconv.ParseInt(pLimit[0], 0, 32)
    }
    from := (page*limit) + 1

    if page == 1 {
        from = 0
    }
    listNews := GetNewsFromElastic(limit, from)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(listNews)

    // tpl.ExecuteTemplate(w, "news.gohtml", listNews)
}

func PostNews(w http.ResponseWriter, r *http.Request) {
    res := model.NewsPostResponse{}

    json.NewDecoder(r.Body).Decode(&res)
    err := validate(res)
    if err != nil {
        http.Error(w, err.Error()+"\n\n"+http.StatusText(400), http.StatusBadRequest)
        return
    }

    job := jobs.InitDataNewsStoreJob(res.Author, res.Body)
    job.Dispatch()

    w.WriteHeader(http.StatusCreated)
    w.Write([]byte("{\"status\":true, \"message\":\"News created\"}"))
    json.NewEncoder(w)
    return
}

func HandleNews(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        GetNews(w, r)
        return
    }else if (r.Method == "POST"){
        PostNews(w, r)
        return
    }else {
        http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
        return
    }
}

func validate(data model.NewsPostResponse) error {
  if data.Author == "" {
    return errors.New("Author is required")
  }else if data.Body == "" {
    return errors.New("Body is required")
  }

  return nil
}