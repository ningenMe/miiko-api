package infra

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type StatisticsRepository struct{}

func (StatisticsRepository) Get() (*StatisticsDto, error) {
	var dto *StatisticsDto

	rows, err := ComproMysql.Queryx(
		`SELECT
       (select count(category_id) from category) AS category_size,
       (select count(topic_id) from topic) AS topic_size,
       (select count(problem_id) from problem) AS problem_size,
       (select count(*) from relation_topic_problem) AS tag_size,
       (select count(reference_id) from reference) AS reference_size,
       (select max(updated_time) from category) AS category_timestamp,
       (select max(updated_time) from topic) AS topic_timestamp,
       (select max(updated_time) from problem) AS problem_timestamp,
       (select max(created_time) from relation_topic_problem) AS tag_timestamp,
       (select max(updated_time) from reference) AS reference_timestamp
       `)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		c := &StatisticsDto{}
		if err = rows.StructScan(c); err != nil {
			fmt.Println(err)
		}
		dto = c
	}

	return dto, nil
}
