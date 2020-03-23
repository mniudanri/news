package handlers

import(
    "fmt"
    "net/http"
    "strconv"
    "encoding/json"
    "errors"
    "io/ioutil"
    "strings"
	"os"
	"news-app/models"
)

func GetNewsFromElastic(limit,from int64) ([]model.NewsGetResponse) {
	// NOTE json elasticsearch to get news data

	// Elasticsearch json
    jsonElastic := fmt.Sprintf(`
    {
        "size": %d,
        "from": %d,
        "sort":{
            "created_time.keyword": "desc"
        }
    }`, limit, from);

    payload := strings.NewReader(jsonElastic)
	req, _ := http.NewRequest("POST", "http://localhost:9200/news-detail/_search/", payload)

	// Header use content-type json and get cache
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Cache-Control", "public, max-age=31557600")

	// Use asic Authentication
    req.SetBasicAuth(os.Getenv("ELASTIC_USER"), os.Getenv("ELASTIC_PWD"))
    elasticRes, err := http.DefaultClient.Do(req)

	if err != nil {
		panic(errors.New("Can not get data from elasticsearch!"))
	}

	// close curl
	defer elasticRes.Body.Close()

	// read and decode response
    body, _ := ioutil.ReadAll(elasticRes.Body)
    interfaceMap := make(map[string]interface{})
    _ = json.Unmarshal([]byte(body), &interfaceMap)

	// get hits
    hits := interfaceMap["hits"].(map[string]interface{})["hits"].(interface{}).([]interface{})

	// NOTE loop data news
	// store to news interface
	listNews := make([]model.NewsGetResponse, 0)
    for _, data := range hits {
        news := model.NewsGetResponse{}
        temp := data.(map[string]interface{})["_source"].(map[string]interface{})
        news.CreatedTime = temp["created_time"].(string)
        news.Id, _ = strconv.ParseInt(temp["id"].(string), 0, 32)
        listNews = append(listNews, news)
	}

	return listNews
}