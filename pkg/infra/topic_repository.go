package infra

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type TopicRepository struct{}

func (TopicRepository) GetListByCategoryId(categoryId string) []*TopicDto {
	var list []*TopicDto

	rows, err := ComproMysql.NamedQuery(
		`SELECT topic_id, category_id, topic_display_name, topic_order FROM topic WHERE category_id = :categoryId ORDER BY topic_order ASC`,
		map[string]interface{}{
			"categoryId": categoryId,
		})
	if err != nil {
		fmt.Println(err)
		return list
	}
	defer rows.Close()

	for rows.Next() {
		c := &TopicDto{}
		if err = rows.StructScan(c); err != nil {
			fmt.Println(err)
		}
		c.ProblemList = getProblemWithListByTopicId(c.TopicId)
		list = append(list, c)
	}

	return list
}

func getProblemWithListByTopicId(topicId string) []*ProblemDto {
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

	return list
}

func getProblemWithListByTopicId(topicId string) []*ProblemDto {
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

	return list
}
