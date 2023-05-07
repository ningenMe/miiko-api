package application

import (
	"fmt"
	"github.com/ningenMe/miiko-api/pkg/domain"
	"github.com/ningenMe/miiko-api/pkg/infra"
	miikov1 "github.com/ningenMe/miiko-api/proto/gen_go/v1"
	"net/http"
)

type TopicUsecase struct{}

var categoryServie = domain.CategoryService{}
var topicService = domain.TopicService{}

func (TopicUsecase) TopicListGet(categorySystemName string) (*miikov1.TopicListGetResponse, error) {

	//データ取得
	categoryDto, err := categoryRepository.Get(categorySystemName)
	if err != nil {
		return &miikov1.TopicListGetResponse{}, err
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

func (TopicUsecase) TopicPost(header http.Header, topicId string, topic *miikov1.Topic, categoryId string) (*miikov1.TopicPostResponse, error) {

	if err := authorizationService.CheckLoggedIn(header); err != nil {
		return &miikov1.TopicPostResponse{}, err
	}

	//idがないときは新規作成
	if topicId == "" {
		topicId = infra.GetNewTopicId()
	}

	var referenceDtoList []*infra.ReferenceDto
	for _, reference := range topic.ReferenceList {
		referenceDtoList = append(referenceDtoList, &infra.ReferenceDto{
			ReferenceId:          infra.GetNewReferenceId(), //更新のたびにIDは作り直す
			Url:                  reference.Url,
			ReferenceDisplayName: reference.ReferenceDisplayName,
		})
	}

	if err := topicService.Upsert(&infra.TopicDto{
		TopicId:          topicId,
		CategoryId:       categoryId,
		TopicDisplayName: topic.TopicDisplayName,
		TopicOrder:       topic.TopicOrder,
		TopicText:        topic.TopicText,
		ReferenceList:    referenceDtoList,
	}); err != nil {
		return nil, err
	}

	return &miikov1.TopicPostResponse{TopicId: topicId}, nil
}

func (TopicUsecase) TopicGet(topicId string) (*miikov1.TopicGetResponse, error) {

	//データ取得
	topicDto, err := topicService.GetWithProblemWithTag(topicId)
	if err != nil {
		return &miikov1.TopicGetResponse{}, fmt.Errorf("topic not found")
	}
	categoryDto, err := categoryServie.GetCategoryByCategoryId(topicDto.CategoryId)
	if err != nil {
		return &miikov1.TopicGetResponse{}, err
	}

	//データ整形
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
	var referenceViewList []*miikov1.Reference
	for _, referenceDto := range topicDto.ReferenceList {
		referenceViewList = append(referenceViewList, &miikov1.Reference{
			ReferenceId:          referenceDto.ReferenceId,
			Url:                  referenceDto.Url,
			ReferenceDisplayName: referenceDto.ReferenceDisplayName,
		})
	}

	return &miikov1.TopicGetResponse{
		Category: &miikov1.Category{
			CategoryId:          categoryDto.CategoryId,
			CategorySystemName:  categoryDto.CategorySystemName,
			CategoryDisplayName: categoryDto.CategoryDisplayName,
			CategoryOrder:       categoryDto.CategoryOrder,
			TopicSize:           categoryDto.TopicSize,
			ProblemSize:         categoryDto.ProblemSize,
		},
		Topic: &miikov1.Topic{
			TopicId:          topicDto.TopicId,
			TopicDisplayName: topicDto.TopicDisplayName,
			TopicOrder:       topicDto.TopicOrder,
			TopicText:        topicDto.TopicText,
			ProblemList:      problemViewList,
		},
	}, nil
}
