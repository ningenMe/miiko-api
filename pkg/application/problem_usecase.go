package application

import (
	"github.com/ningenMe/miiko-api/pkg/domain"
	"github.com/ningenMe/miiko-api/pkg/infra"
	miikov1 "github.com/ningenMe/miiko-api/proto/gen_go/v1"
	"net/http"
)

const todoTopic = "topic_000266"

type ProblemUsecase struct{}

var problemService = domain.ProblemService{}

func (ProblemUsecase) ProblemListGet(limit int32, offset int32, isDesc bool) (*miikov1.ProblemListGetResponse, error) {

	//データ取得
	problemDtoList := problemRepository.GetProblemWithTagList(limit, offset, isDesc)

	//データ整形
	var problemViewList []*miikov1.Problem
	for _, problemDto := range problemDtoList {

		var tagViewList []*miikov1.Tag
		for _, tagView := range problemDto.TagList {
			tagViewList = append(tagViewList, &miikov1.Tag{
				TopicId:          tagView.TopicId,
				CategoryId:       tagView.CategoryId,
				TopicDisplayName: tagView.TopicDisplayName,
			})
		}

		problemViewList = append(problemViewList, &miikov1.Problem{
			ProblemId:          problemDto.ProblemId,
			Url:                problemDto.Url,
			ProblemDisplayName: problemDto.ProblemDisplayName,
			Estimation:         problemDto.Estimation,
			TagList:            tagViewList,
		})
	}

	return &miikov1.ProblemListGetResponse{
		ProblemList: problemViewList,
	}, nil
}

func (ProblemUsecase) ProblemGet(problemId string) (*miikov1.ProblemGetResponse, error) {

	problemDto, err := problemRepository.GetProblemWithTag(problemId)
	if err != nil {
		return &miikov1.ProblemGetResponse{}, err
	}

	var tagViewList []*miikov1.Tag
	for _, tagDto := range problemDto.TagList {
		tagViewList = append(tagViewList, &miikov1.Tag{
			TopicId:          tagDto.TopicId,
			CategoryId:       tagDto.CategoryId,
			TopicDisplayName: tagDto.TopicDisplayName,
		})
	}

	return &miikov1.ProblemGetResponse{
		Problem: &miikov1.Problem{
			ProblemId:          problemDto.ProblemId,
			Url:                problemDto.Url,
			ProblemDisplayName: problemDto.ProblemDisplayName,
			Estimation:         problemDto.Estimation,
			TagList:            tagViewList,
		},
	}, nil
}

func (ProblemUsecase) ProblemPost(header http.Header, problemId string, problem *miikov1.Problem) (*miikov1.ProblemPostResponse, error) {
	if err := authorizationService.CheckLoggedIn(header); err != nil {
		return &miikov1.ProblemPostResponse{}, err
	}

	//新規IDの払い出しやurlチェックなど
	problemId, err := problemService.GetProblemId(problemId, problem.Url)
	if err != nil {
		return &miikov1.ProblemPostResponse{}, err
	}

	var tagDtoList []*infra.TagDto
	for _, tag := range problem.TagList {
		tagDtoList = append(tagDtoList, &infra.TagDto{
			TopicId: tag.TopicId,
		})
	}
	//NOTE タグがない場合は未登録タグに入れる
	if len(tagDtoList) == 0 {
		tagDtoList = append(tagDtoList, &infra.TagDto{
			TopicId: todoTopic,
		})
	}

	problemDto := &infra.ProblemDto{
		ProblemId:          problemId,
		Url:                problem.Url,
		ProblemDisplayName: problem.ProblemDisplayName,
		Estimation:         problem.Estimation,
		TagList:            tagDtoList,
	}

	problemRepository.Upsert(problemDto)
	problemRepository.DeleteTag(problemId)
	for _, tag := range problem.TagList {
		problemRepository.InsertTag(problemId, tag.TopicId)
	}

	return &miikov1.ProblemPostResponse{
		ProblemId: problemId,
	}, nil
}
