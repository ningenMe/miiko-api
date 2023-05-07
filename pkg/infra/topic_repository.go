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

func (TopicRepository) Get(topicId string) (*TopicDto, error) {
	var dto *TopicDto

	rows, err := ComproMysql.NamedQuery(
		`SELECT topic_id, category_id, topic_display_name, topic_order FROM topic WHERE topic_id = :topicId`,
		map[string]interface{}{
			"topicId": topicId,
		})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		c := &TopicDto{}
		if err = rows.StructScan(c); err != nil {
			fmt.Println(err)
		}
		dto = c
	}

	return dto, nil
}

func (TopicRepository) Upsert(topicDto *TopicDto) error {

	_, err := ComproMysql.NamedExec(`INSERT INTO topic (topic_id, category_id, topic_display_name, topic_order, topic_text) 
                                 VALUES (:topic_id, :category_id, :topic_display_name, :topic_order, :topic_text) 
                                 ON DUPLICATE KEY UPDATE 
                                     category_id=VALUES(category_id), topic_display_name=VALUES(topic_display_name), topic_order=VALUES(topic_order), topic_text=VALUES(topic_text)                                    
                                     `, topicDto)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
