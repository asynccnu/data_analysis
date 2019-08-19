package insight

import (
	. "github.com/asynccnu/data_analysis_service_v1/handler"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"sort"
	"strconv"
)

const (
	maxNum = 21
)

//必须要大写，否则外面不能导入
type PV struct {
	SubCat string `json:"sub_cat"`
	Value  uint64 `json:"value"`
}

func Pv(c *gin.Context) {
	duration := c.DefaultQuery("duration", "1d")
	sql := "q=SELECT COUNT(*) FROM data1 WHERE time > now()- " + duration + " AND \"mianCat\"='pageView' AND \"pid\"='ccnubox' GROUP BY subCat tz('Asia/Shanghai');"
	json := sql2json(sql)

	time := gjson.Get(json, "results.0.series.0.values.0.0").String()
	var pvMap []PV
	for i := 0; i < maxNum; i++ {
		valuePath := "results.0.series." + strconv.Itoa(i) + ".values.0.2"
		subCatPath := "results.0.series." + strconv.Itoa(i) + ".tags.subCat"
		subCat := gjson.Get(json, subCatPath).String()
		value := gjson.Get(json, valuePath).Uint()
		pvMap = append(pvMap, PV{subCat, value})
	}
	sort.Slice(pvMap, func(i, j int) bool {
		return pvMap[i].Value > pvMap[j].Value
	})
	pvMap = append(pvMap, PV{time, 0})
	SendResponse(c, nil, pvMap)
}
