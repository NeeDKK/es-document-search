package controller

import (
	"esDocumentSearch/config"
	"esDocumentSearch/entity"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"strconv"
)

type SearchController struct {
}

func Search(c *gin.Context) {

	searchContent := c.DefaultQuery("searchContent", "")
	size := c.DefaultQuery("size", "2")
	page := c.DefaultQuery("page", "0")
	sizeInt, _ := strconv.Atoi(size)
	pageInt, _ := strconv.Atoi(page)
	query := elastic.NewBoolQuery()
	//如果默认查询为空。分页返回所有
	if searchContent == "" {
		searchContentResult, err := config.EsClient.Search(config.RESUMEINDEX).Query(query).Size(sizeInt).From((pageInt - 1) * sizeInt).Do(c)
		if err != nil {
			fmt.Println("查询失败", err.Error())
			entity.FailWithMessage("查询失败", c)
			return
		}
		entity.OkWithDetailed(searchContentResult, "查询成功", c)
		return
	}
	higthFild := elastic.NewHighlight().Fields(elastic.NewHighlighterField("attachment.content").PreTags("<span style='color:red'>").PostTags("</span>"))
	query.Should(elastic.NewMatchQuery("name.keyword", searchContent)).
		Should(elastic.NewMatchQuery("school", searchContent)).
		Should(elastic.NewMatchQuery("attachment.content", searchContent).Analyzer("ik_max_word"))
	searchContentResult, err := config.EsClient.Search(config.RESUMEINDEX).Query(query).Highlight(higthFild).Size(sizeInt).From((pageInt - 1) * sizeInt).Do(c)
	if err != nil {
		fmt.Println("查询失败", err.Error())
		entity.FailWithMessage("查询失败", c)
		return
	}
	entity.OkWithDetailed(searchContentResult, "查询成功", c)
}
