package infra

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type TopicRepository struct{}

var problemRepository = ProblemRepository{}

// TODO メソッドを分解する
func (TopicRepository) GetListByCategoryId(categoryId string, isRequiredProblem bool) []*TopicDto {
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
		if isRequiredProblem {
			c.ProblemList = problemRepository.GetProblemListByTopicId(c.TopicId, false)
		}
		list = append(list, c)
	}

	return list
}

// TODO TopicProblemWithTagの取得
func (TopicRepository) Get(topicId string) *TopicDto {
	var dto *TopicDto

	rows, err := ComproMysql.NamedQuery(
		`SELECT topic_id, category_id, topic_display_name, topic_order FROM topic WHERE topic_id = :topicId`,
		map[string]interface{}{
			"topicId": topicId,
		})
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		c := &TopicDto{}
		if err = rows.StructScan(c); err != nil {
			fmt.Println(err)
		}
		c.ProblemList = problemRepository.GetProblemListByTopicId(c.TopicId, true)
		dto = c
	}

	return dto
}

func (TopicRepository) Upsert(topic *TopicDto) {

	_, err := ComproMysql.NamedExec(`INSERT INTO topic (topic_id, category_id, topic_display_name, topic_order) 
                                 VALUES (:topic_id, :category_id, :topic_display_name, :topic_order) 
                                 ON DUPLICATE KEY UPDATE 
                                     category_id=VALUES(category_id), topic_display_name=VALUES(topic_display_name), topic_order=VALUES(topic_order)                                     
                                     `, topic)
	if err != nil {
		fmt.Println(err)
	}
}
