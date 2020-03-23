package handlers

import(
    "fmt"
    "net/http"
    "strconv"
    "encoding/json"
    "io/ioutil"
    "strings"
	"os"
    "news-app/models"
    "news-app/connection"
    "log"
)

func init(){
    db = connection.Connect()
}

func GetNewsFromElastic(limit,from int64) ([]model.NewsGetResponse) {
    // NOTE json elasticsearch to get news data

    listNews := make([]model.NewsGetResponse, 0)

	// Elasticsearch json
    jsonElastic := fmt.Sprintf(`
    {
        "size": %d,
        "from": %d,
        "sort":{
            "created_time": "desc"
        }
    }`, limit, from);

    payload := strings.NewReader(jsonElastic)
    req, err := http.NewRequest("POST", "http://localhost:9200/news-detail/_search/", payload)

    if err != nil {
        log.Print("Can not get data from elasticsearch!")
        return listNews
	}
	// Header use content-type json and get cache
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Cache-Control", "public, max-age=31557600")

	// Use asic Authentication
    req.SetBasicAuth(os.Getenv("ELASTIC_USER"), os.Getenv("ELASTIC_PWD"))
    elasticRes, err := http.DefaultClient.Do(req)

	if err != nil {
        log.Print("Can not get data from elasticsearch!")
        return listNews
	}

	// close curl
	defer elasticRes.Body.Close()

	// read and decode response
    body, _ := ioutil.ReadAll(elasticRes.Body)
    interfaceMap := make(map[string]interface{})
    _ = json.Unmarshal([]byte(body), &interfaceMap)

    if interfaceMap["error"] != nil{
        log.Print("Can not get data from elasticsearch!")
        return listNews
    }
	// get hits
    hits := interfaceMap["hits"].(map[string]interface{})["hits"].(interface{}).([]interface{})

	// NOTE loop data news
	// store to news interface
    for _, data := range hits {
        news := model.NewsGetResponse{}
        temp := data.(map[string]interface{})["_source"].(map[string]interface{})
        news.CreatedTime = temp["created_time"].(string)
        news.Id, _ = strconv.ParseInt(temp["id"].(string), 0, 32)
        listNews = append(listNews, news)
	}

	return listNews
}