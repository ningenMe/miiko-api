package infra

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type TopicRepository struct{}

var problemRepository = ProblemRepository{}

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
		c.ProblemList = problemRepository.GetProblemListByTopicId(c.TopicId, false)
		list = append(list, c)
	}

	return list
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
