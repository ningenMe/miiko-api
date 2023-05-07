package application

import (
	miikov1 "github.com/ningenMe/miiko-api/proto/gen_go/v1"
)

type ProblemUsecase struct{}

func (ProblemUsecase) ProblemListGet(offset int32, limit int32) (*miikov1.ProblemListGetResponse, error) {

	//データ取得
	problemDtoList := problemRepository.GetProblemWithTagList(offset, limit)

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
