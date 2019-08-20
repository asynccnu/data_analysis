package insight

import (
	"fmt"
	"github.com/asynccnu/data_analysis_service_v1/handler"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"strconv"
)

func Error(c *gin.Context) {
	ErrMap := make(map[string]string) //:= only use in function
	duration := c.DefaultQuery("duration", "1d")
	errorSql := "q=SELECT count(*) FROM data1 WHERE time > now()- " + duration + " AND \"mianCat\"='apiEvent' AND \"value\"='error' tz('Asia/Shanghai');"
	allSql := "q=SELECT count(*) FROM data1 WHERE time > now()- " + duration + " AND \"mianCat\"='apiEvent' tz('Asia/Shanghai');"
	errorJson := sql2json(errorSql)
	allJson := sql2json(allSql)

	time := gjson.Get(allJson, "results.0.series.0.values.0.0").String()
	ErrMap["time"] = time

	errorNum := json2num(errorJson)
	allNum := json2num(allJson)
	pr := fmt.Sprintf("%f%%", (errorNum/allNum)*100)
	ErrMap["proportion"] = pr
	ErrMap["error_num"] = strconv.FormatFloat(errorNum, 'f', -1, 64)
	ErrMap["all_num"] = strconv.FormatFloat(allNum, 'f', -1, 64)
	handler.SendResponse(c, nil, ErrMap)
}

func json2num(json string) float64 {
	value := gjson.Get(json, "results.0.series.0.values.0.2").Float()
	return value
}
