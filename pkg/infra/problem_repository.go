package infra

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
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
		c.TagList = getTagList(c.ProblemId)
		list = append(list, c)
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
		c.TagList = getTagList(c.ProblemId)
		list = append(list, c)
	}

	return list
}

func getTagList(problemId string) []*TagDto {
	var list []*TagDto

	rows, err := ComproMysql.NamedQuery(
		`SELECT t.topic_id, category_id, topic_display_name FROM topic AS t JOIN relation_topic_problem AS rtp ON t.topic_id = rtp.topic_id WHERE problem_id = :problemId ORDER BY topic_order ASC`,
		map[string]interface{}{
			"problemId": problemId,
		})
	if err != nil {
		fmt.Println(err)
		return list
	}

	for rows.Next() {
		c := &TagDto{}
		if err = rows.StructScan(c); err != nil {
			fmt.Println(err)
		}
		list = append(list, c)
	}

	return list
}
