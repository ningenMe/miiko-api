package application

import (
	"fmt"
	miikov1 "github.com/ningenMe/miiko-api/proto/gen_go/v1"
)

type TopicUsecase struct{}

func (TopicUsecase) TopicListGet(categorySystemName string) (*miikov1.TopicListGetResponse, error) {

	//データ取得
	categoryDto := categoryRepository.Get(categorySystemName)
	if categoryDto == nil {
		return &miikov1.TopicListGetResponse{}, fmt.Errorf("category not found")
	}
	topicDtoList := topicRepository.GetListByCategoryId(categoryDto.CategoryId, true)

	//データ整形
	categoryView := &miikov1.Category{
		CategoryId:          categoryDto.CategoryId,
		CategorySystemName:  categoryDto.CategorySystemName,
		CategoryDisplayName: categoryDto.CategoryDisplayName,
		CategoryOrder:       categoryDto.CategoryOrder,
	}

	var topicViewList []*miikov1.Topic
	for _, topicDto := range topicDtoList {

		var problemViewList []*miikov1.Problem
		for _, problemDto := range topicDto.ProblemList {

			var tagViewList []*miikov1.Tag
			for _, tagDto := range problemDto.TagList {
				tagViewList = append(tagViewList, &miikov1.Tag{
					TopicId:          tagDto.TopicId,
					CategoryId:       tagDto.CategoryId,
					TopicDisplayName: tagDto.TopicDisplayName,
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

		topicViewList = append(topicViewList, &miikov1.Topic{
			TopicId:          topicDto.TopicId,
			TopicDisplayName: topicDto.TopicDisplayName,
			TopicOrder:       topicDto.TopicOrder,
			ProblemList:      problemViewList,
		})
	}

	return &miikov1.TopicListGetResponse{
		Category:  categoryView,
		TopicList: topicViewList,
	}, nil
}
