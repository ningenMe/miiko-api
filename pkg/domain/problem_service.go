package domain

import (
	"fmt"
	"github.com/ningenMe/miiko-api/pkg/infra"
)

type ProblemService struct{}

var problemRepository = infra.ProblemRepository{}

func (ProblemService) GetProblemId(problemId string, url string) (string, error) {
	//新規でidを払い出す
	if problemId == "" {
		problemId = infra.GetNewProblemId()
		//urlが同じなのに新規作成してたらおかしいのでエラーにする
		problemWithUrl, _ := problemRepository.GetProblemByUrl(url)
		if problemWithUrl != nil {
			return "", fmt.Errorf("url is duplicated")
		}
	}
	return problemId, nil
}
