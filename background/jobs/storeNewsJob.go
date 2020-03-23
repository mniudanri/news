package jobs

import(
    "fmt"
    "time"
	"news-app/models"
    "database/sql"
	"net/http"
	"os"
	"io/ioutil"
	"errors"
    _ "github.com/lib/pq"
	"strings"
)

type NewsStoreJob struct { // job class
    Id int64
    Author string
    Body string
    CreatedTime string
}

var (
	db *sql.DB
	news NewsStoreJob
)

func InitDataNewsStoreJob(Author string, Body string) (NewsStoreJob) {
	news.Author = Author
	news.Body = Body

	var err error
    db, err = sql.Open("postgres", "host=localhost port=5432 user=postgres "+
            "password=12345678 dbname=masterdata sslmode=disable")

    if err != nil {
        panic(err)
    }

    if err = db.Ping(); err != nil {
        panic(err)
    }

	return news
}

func (job *NewsStoreJob) Process() {
	t := time.Now()
	data := model.NewsPostResponse{}

    rows := db.QueryRow("INSERT INTO news (author, body, created_time) VALUES ($1, $2, $3) RETURNING id, created_time", news.Author, news.Body, t)
	_ = rows.Scan(&data.Id, &data.CreatedTime)

	// reformat date
    formatted := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
	t.Year(), t.Month(), t.Day(),
	t.Hour(), t.Minute(), t.Second())
	data.CreatedTime = formatted
	data.Body = news.Body

    _ = PostNewsToElasticSearch(data)
}

func (job *NewsStoreJob) Dispatch() {
	jobQueue := InitJobQueue(10)
	jobQueue.Start()

	jobQueue.Push(job)

	// NOTE: Queue set timeout 1 seconds
    time.AfterFunc(time.Second * 1, func() {
		jobQueue.Stop()
    })
    time.Sleep(time.Second * 6)
}

// NOTE: process upsert to elasticsearch
func PostNewsToElasticSearch(news model.NewsPostResponse) []byte {
	jsonElastic := fmt.Sprintf(`
	{
		"doc" : {
			"id" : "%d",
			"created_time": "%s"
		},
		"doc_as_upsert" : true
	}`, news.Id, news.CreatedTime);

	payload := strings.NewReader(jsonElastic)
	req, _ := http.NewRequest("POST", fmt.Sprintf("http://localhost:9200/news-detail/_update/%d", news.Id), payload)
	req.Header.Set("Content-Type", "application/json")

	req.SetBasicAuth(os.Getenv("ELASTIC_USER"), os.Getenv("ELASTIC_PWD"))
	elasticRes, _ := http.DefaultClient.Do(req)
	fmt.Println("Response Elastic", elasticRes)

	defer elasticRes.Body.Close()
	body, err := ioutil.ReadAll(elasticRes.Body)

	if err != nil {
		panic(errors.New("Can not get data from elasticsearch!"))
	}

	fmt.Println(string(body))
	return body
}