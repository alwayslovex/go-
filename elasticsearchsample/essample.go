package elasticsearchsample

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
)

//es中的存储结构
type DutyInfo struct {
	Title        string `json:"title"`
	Content      string `json:"content"`
	Results      string `json:"results"`
	Dutyer       string `json:"dutyer"`
	HappenedDate string `json:"happeneddate"`
}

//读取模版数据导入到es
func putEs(fileName string) {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	//elastic.SetSniff(false)
	client, err := elastic.NewSimpleClient(elastic.SetURL("http://127.0.0.1:9200"))
	if err != nil {
		panic(err)
	}
	info, code, err := client.Ping(host).Do(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	lines := strings.Split(string(content), "\n")
	bodys := make([]DutyInfo, 0)
	for _, v := range lines {
		line := strings.Split(v, "\t")
		d := line[1] //strings.Split(line[1]," ")[0]
		//只是为了打印
		s := fmt.Sprintf("POST /onduty/_doc \n{\n\"title\":\"%s\",\n\"content\":\"%s\",\n\"results\":\"%s\",\n\"dutyer\":\"linhaote\",\n\"happen_date\":\"%s\"\n}", line[0], line[0],
			line[5], strings.ReplaceAll(d, "/", "-"))
		fmt.Println(s)
		var duty DutyInfo
		duty.Title = line[0]
		duty.Content = line[0]
		duty.Results = line[5]
		duty.Dutyer = "linhaote"
		duty.HappenedDate = strings.ReplaceAll(d, "/", "-")
		bodys = append(bodys, duty)
		put1, er := client.Index().Index("onduty").BodyJson(duty).Do(ctx)
		if er != nil {
			panic(er)
		}
		fmt.Println(put1.Result)
	}
}

var host = "http://localhost:9200"

func searchEs(title string) {
	ctx := context.Background()
	//elastic.SetSniff(false)
	client, err := elastic.NewSimpleClient(elastic.SetURL("http://127.0.0.1:9200"))
	if err != nil {
		panic(err)
	}
	info, code, err := client.Ping(host).Do(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)
	boolSearch := elastic.NewBoolQuery()
	boolSearch.Must(elastic.NewMatchQuery("title", title)).Must(elastic.NewMatchQuery("content", title))
	result, err := client.Search().
		Index("onduty").
		Query(boolSearch).
		Size(100).
		Do(context.Background())

	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Println("共查询到", result.Hits.TotalHits.Value, " 条记录\n查询结果是:\n")
	var dutyinfo DutyInfo
	for _, item := range result.Each(reflect.TypeOf(dutyinfo)) {
		if t, ok := item.(DutyInfo); ok {
			fmt.Printf("%s : %s : %s\n", t.Content, t.Results, t.Happen_date)
		}
	}
}

//参考 https://olivere.github.io/elastic/
//可视化开源组件 https://github.com/360EntSecGroup-Skylar/ElasticHD
func main() {
	method := os.Args[1]
	keywd := os.Args[2]
	if method == "put" {
		putEs(keywd)
	}
	if method == "get" {
		//title := "巴西快车adyen渠道支付成功数"
		searchEs(keywd)
	}
}
