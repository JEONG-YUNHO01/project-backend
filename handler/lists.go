package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/JEONG-YUNHO01/project-backend/models"

	"github.com/labstack/echo/v4"
)

const format string = "json"
const numOfRows string = "10"

// 의약품 검색 결과 리스트를 반환
func GetSearchResultLists(c echo.Context) error {
	itemName := c.QueryParam("itemName")
	pageNo := c.QueryParam("pageNo")
	entpName := c.QueryParam("entpName")
	efcyQesitm := c.QueryParam("efcyQesitm")
	apiKey := os.Getenv("apiKey")
	params := "serviceKey=" + url.QueryEscape(apiKey) + "&" +
		"pageNo=" + url.QueryEscape(pageNo) + "&" +
		"numOfRows=" + url.QueryEscape(numOfRows) + "&" +
		"itemName=" + url.QueryEscape(itemName) + "&" +
		"entpName=" + url.QueryEscape(entpName) + "&" +
		"efcyQesitm=" + url.QueryEscape(efcyQesitm) + "&" +
		"type=" + url.QueryEscape(format)

	path := fmt.Sprintf("%s?%s", os.Getenv("apiUrl"), params)

	resp, err := http.Get(path)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	resData := models.ResData{}
	json.Unmarshal(body, &resData)

	return c.JSON(http.StatusOK, resData)
}

// 의약품 1개에 대한 상세페이지용 정보를 반환
func GetOneDrugDetail(c echo.Context) error {
	apiKey := os.Getenv("apiKey")
	itemSeq := c.QueryParam("itemSeq")
	params := "serviceKey=" + url.QueryEscape(apiKey) + "&" +
		"itemSeq=" + url.QueryEscape(itemSeq) + "&" +
		"type=" + url.QueryEscape(format)

	path := fmt.Sprintf("%s?%s", os.Getenv("apiUrl"), params)

	resp, err := http.Get(path)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	resData := models.ResData{}
	json.Unmarshal(body, &resData)
	resData.Body.Items[0].OpenDe = resData.Body.Items[0].OpenDe[0:10]
	resData.Body.Items[0].UpdateDe = resData.Body.Items[0].UpdateDe[0:4] + "-" + resData.Body.Items[0].UpdateDe[4:6] + "-" + resData.Body.Items[0].UpdateDe[6:8]
	return c.JSON(http.StatusOK, resData)
}
