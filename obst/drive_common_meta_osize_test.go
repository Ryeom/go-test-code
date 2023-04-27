package test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"testing"
)

type CommonMeta struct {
	FileSize    int64  `json:"fileSize"`
	UserNo      string `json:"userNo"`
	ServiceCode string `json:"serviceCode"`
}
type SizeMeta struct {
	Used        string `json:"used"`
	UserNo      string `json:"userNo"`
	ServiceCode string `json:"serviceCode"`
	Class       string `json:"class"`
	BucketName  string `json:"bucketName"`
	Number      string `json:"number"`
	UserId      string `json:"userId"`
	WasteUsed   string `json:"wasteUsed"`
	ImageUsed   string `json:"imageUsed"`
	VideoUsed   string `json:"videoUsed""`
	AudioUsed   string `json:"audioUsed"`
	EtcUsed     string `json:"etcUsed""`
}

func TestCommonMetaToSizeMeta(t *testing.T) {
	file, _ := ioutil.ReadFile("n_7651.json")
	var targetList []CommonMeta
	jerr := json.Unmarshal(file, &targetList)

	if jerr != nil {
		fmt.Println(jerr)
		return
	}

	resultRow := map[string]*CommonMeta{}
	for _, v := range targetList {
		if resultRow[v.UserNo+"-"+v.ServiceCode] != nil {
			resultRow[v.UserNo+"-"+v.ServiceCode].FileSize = resultRow[v.UserNo+"-"+v.ServiceCode].FileSize + v.FileSize
		} else {
			resultRow[v.UserNo+"-"+v.ServiceCode] = &CommonMeta{
				UserNo:      v.UserNo,
				ServiceCode: v.ServiceCode,
				FileSize:    v.FileSize,
			}

		}
	}
	number := "7651"
	bucketName := "n:" + number
	const class = "com.example.hello.storage.dto.FileSizeMeta"
	//
	//ctx, _ := context.WithTimeout(context.Background(), 80000*time.Second)
	//clientOptions := options.Client().ApplyURI("mongodb://10.20.30.40:26017").SetMaxPoolSize(100)
	//client, err := mongo.Connect(ctx, clientOptions)
	//if err != nil {
	//	fmt.Println("1 mongodb err", err)
	//	panic(err)
	//}
	//client.Ping(ctx, nil)
	//if err != nil {
	//	fmt.Println("2 mongodb err", err)
	//	panic(err)
	//}
	//
	//collection := client.Database("storage").Collection("SizeMeta")
	for _, v := range resultRow {
		doc := SizeMeta{
			Used:        "NumberLong(" + strconv.Itoa(int(v.FileSize)) + ")",
			UserNo:      v.UserNo,
			ServiceCode: v.ServiceCode,
			Class:       class,
			BucketName:  bucketName,
			Number:      number,
			UserId:      v.UserNo,
			WasteUsed:   "NumberLong(0)",
			ImageUsed:   "NumberLong(0)",
			VideoUsed:   "NumberLong(0)",
		}
		//result, err := collection.InsertOne(context.TODO(), doc)
		//if err != nil {
		//	fmt.Println("insert err", err)
		//}
		//fmt.Println(result)
		b, _ := json.Marshal(doc)
		fmt.Println(string(b), ",")
	}

}
