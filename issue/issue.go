package issue

import (
	"fmt"
	. "gin-restapi-sample/db"
	. "gin-restapi-sample/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	rows, err := Dbm.Select(IssueData{}, "select * from issue")
	if err != nil {
		fmt.Println(err)
		c.String(http.StatusInternalServerError, "error")
	} else if len(rows) == 0 {
		c.String(http.StatusNotFound, "not data")
	} else {
		c.JSON(http.StatusOK, rows)
	}
}

func Show(c *gin.Context) {
	id := c.Param("id") //TODO:idのvalidate
	rows, err := getList("where id=" + id)

	if err != nil {
		fmt.Println(err)
		c.String(http.StatusInternalServerError, "error")
	} else if len(rows) == 0 {
		c.String(http.StatusNotFound, "not data")
	} else {
		fmt.Println("getData:", rows)
		c.JSON(http.StatusOK, rows)
	}
}

//common function for get issues
func getList(condition string) ([]IssueData, error) {
	sql := "select * from issue " + condition
	rows, err := Dbm.Select(IssueData{}, sql)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	issue_list := make([]IssueData, len(rows))
	cnt := 0
	for _, row := range rows {
		issuedata := row.(*IssueData)
		issue_list[cnt].ID = issuedata.ID
		issue_list[cnt].Title = issuedata.Title
		issue_list[cnt].Source = issuedata.Source
		issue_list[cnt].Detail = issuedata.Detail
		issue_list[cnt].Priority = issuedata.Priority
		issue_list[cnt].Status = issuedata.Status
		issue_list[cnt].LimitStr = UnixTimeToDayString(issuedata.Limit)
		issue_list[cnt].CreatedStr = UnixTimeToDateString(issuedata.Created)
		issue_list[cnt].UpdatedStr = UnixTimeToDateString(issuedata.Updated)
		cnt++
	}
	return issue_list, nil
}
