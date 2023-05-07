package infra

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type ReferenceRepository struct{}

func (ReferenceRepository) Delete(topicId string) error {

	_, err := ComproMysql.NamedExec(`DELETE FROM reference WHERE topic_id = :topicId`,
		map[string]interface{}{
			"topicId": topicId,
		})
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (ReferenceRepository) Insert(topicId string, referenceDto *ReferenceDto) error {

	_, err := ComproMysql.NamedExec(`INSERT INTO reference (reference_id, topic_id, url, reference_display_name) 
                                 VALUES (:referenceId, :topic_id, :url, :referenceDisplayName)                                 
                                     `,
		map[string]interface{}{
			"referenceId":          referenceDto.ReferenceId,
			"topicId":              topicId,
			"url":                  referenceDto.Url,
			"referenceDisplayName": referenceDto.ReferenceDisplayName,
		})
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
