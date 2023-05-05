package infra

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type ProblemRepository struct{}

func (ProblemRepository) GetProblemListByTopicId(topicId string) []*ProblemDto {
	var list []*ProblemDto

	rows, err := ComproMysql.NamedQuery(
		`SELECT p.problem_id, url, problem_display_name, estimation FROM problem AS p JOIN relation_topic_problem AS rtp ON p.problem_id = rtp.problem_id WHERE rtp.topic_id = :topicId ORDER BY estimation`,
		map[string]interface{}{
			"topicId": topicId,
		})
	if err != nil {
		fmt.Println(err)
		return list
	}

	for rows.Next() {
		c := &ProblemDto{}
		if err = rows.StructScan(c); err != nil {
			fmt.Println(err)
		}
		list = append(list, c)
	}
	var problemIdlist []string
	for _, problem := range list {
		problemIdlist = append(problemIdlist, problem.ProblemId)
	}
	tagMap := getTagMap(problemIdlist)
	for _, problem := range list {
		problem.TagList = tagMap[problem.ProblemId]
	}

	return list
}

func (ProblemRepository) GetProblemList(sortType string, offset int32, limit int32) []*ProblemDto {
	var list []*ProblemDto

	rows, err := ComproMysql.NamedQuery(
		`SELECT problem_id, url, problem_display_name, estimation FROM problem ORDER BY :sortType DESC LIMIT :limit OFFSET :offset`,
		map[string]interface{}{
			"sortType": sortType,
			"offset":   offset,
			"limit":    limit,
		})
	if err != nil {
		fmt.Println(err)
		return list
	}

	for rows.Next() {
		c := &ProblemDto{}
		if err = rows.StructScan(c); err != nil {
			fmt.Println(err)
		}
		list = append(list, c)
	}

	var problemIdlist []string
	for _, problem := range list {
		problemIdlist = append(problemIdlist, problem.ProblemId)
	}
	tagMap := getTagMap(problemIdlist)
	for _, problem := range list {
		problem.TagList = tagMap[problem.ProblemId]
	}

	return list
}

func getTagMap(problemIdList []string) map[string][]*TagDto {
	mp := make(map[string][]*TagDto)

	sql := `SELECT t.topic_id, category_id, topic_display_name, problem_id FROM topic AS t JOIN relation_topic_problem AS rtp ON t.topic_id = rtp.topic_id WHERE problem_id IN (?) ORDER BY topic_order ASC`

	sql, params, err := sqlx.In(sql, problemIdList)
	if err != nil {
		fmt.Println(err)
		return mp
	}

	var list []TagDto
	err = ComproMysql.Select(&list, ComproMysql.Rebind(sql), params...)
	if err != nil {
		fmt.Println(err)
		return mp
	}

	for _, tag := range list {
		tmpList := mp[tag.ProblemId]
		tmpList = append(tmpList, &tag)
		mp[tag.ProblemId] = tmpList
	}

	return mp
}
