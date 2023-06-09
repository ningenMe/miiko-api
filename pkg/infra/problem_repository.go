package infra

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type ProblemRepository struct{}

// TODO ハンドリングを直す
func (ProblemRepository) GetProblemListByTopicId(topicId string, isRequiredTag bool) []*ProblemDto {
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

	if isRequiredTag {
		tagMap := getTagMap(problemIdlist)
		for _, problem := range list {
			problem.TagList = tagMap[problem.ProblemId]
		}
	}

	return list
}

func (ProblemRepository) GetProblemWithTagList(limit int32, offset int32, isDesc bool) []*ProblemDto {
	var list []*ProblemDto

	sortOrder := "DESC"
	if !isDesc {
		sortOrder = "ASC"
	}

	rows, err := ComproMysql.NamedQuery(
		`SELECT problem_id, url, problem_display_name, estimation FROM problem ORDER BY updated_time `+sortOrder+` LIMIT :limit OFFSET :offset`,
		map[string]interface{}{
			"limit":  limit,
			"offset": offset,
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

func (ProblemRepository) GetProblemWithTag(problemId string) (*ProblemDto, error) {
	var dto *ProblemDto

	rows, err := ComproMysql.NamedQuery(
		`SELECT problem_id, url, problem_display_name, estimation FROM problem WHERE problem_id = :problemId`,
		map[string]interface{}{
			"problemId": problemId,
		})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for rows.Next() {
		c := &ProblemDto{}
		if err = rows.StructScan(c); err != nil {
			fmt.Println(err)
		}
		dto = c
	}

	tagMap := getTagMap([]string{dto.ProblemId})
	dto.TagList = tagMap[dto.ProblemId]

	return dto, nil
}

func (ProblemRepository) GetProblemByUrl(url string) (*ProblemDto, error) {
	var dto *ProblemDto

	rows, err := ComproMysql.NamedQuery(
		`SELECT problem_id, url, problem_display_name, estimation FROM problem WHERE url = :url`,
		map[string]interface{}{
			"url": url,
		})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for rows.Next() {
		c := &ProblemDto{}
		if err = rows.StructScan(c); err != nil {
			fmt.Println(err)
		}
		dto = c
	}

	return dto, nil
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

	for idx, _ := range list {
		tmpList := mp[list[idx].ProblemId]
		tmpList = append(tmpList, &list[idx])
		mp[list[idx].ProblemId] = tmpList
	}

	return mp
}

func (ProblemRepository) Upsert(tx *sqlx.Tx, problem *ProblemDto) error {

	_, err := tx.NamedExec(`INSERT INTO problem (problem_id, url, problem_display_name, estimation) 
                                 VALUES (:problem_id, :url, :problem_display_name, :estimation) 
                                 ON DUPLICATE KEY UPDATE 
                                    url=VALUES(url), problem_display_name=VALUES(problem_display_name), estimation=VALUES(estimation)                                     
                                     `, problem)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (ProblemRepository) DeleteTag(tx *sqlx.Tx, problemId string) error {

	_, err := tx.NamedExec(`DELETE FROM relation_topic_problem WHERE problem_id = :problemId`,
		map[string]interface{}{
			"problemId": problemId,
		})
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (ProblemRepository) InsertTag(tx *sqlx.Tx, problemId string, topicId string) error {
	_, err := tx.NamedExec(`INSERT INTO relation_topic_problem (problem_id, topic_id) VALUES (:problemId, :topicId) `,
		map[string]interface{}{
			"problemId": problemId,
			"topicId":   topicId,
		})
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
