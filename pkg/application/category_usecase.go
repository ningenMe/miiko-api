package application

import miikov1 "github.com/ningenMe/miiko-api/proto/gen_go/v1"

type CategoryUsecase struct{}

func (CategoryUsecase) CategoryListGet(isRequiredTopic bool) (*miikov1.CategoryListGetResponse, error) {

	categoryDtoList := categoryRepository.GetList()

	var categoryViewList []*miikov1.Category
	for _, categoryDto := range categoryDtoList {

		var viewTopicList []*miikov1.Topic
		if isRequiredTopic {
			topicList := topicRepository.GetListByCategoryId(categoryDto.CategoryId, false)
			for _, topic := range topicList {
				viewTopicList = append(viewTopicList, &miikov1.Topic{
					TopicId:          topic.TopicId,
					TopicDisplayName: topic.TopicDisplayName,
					TopicOrder:       topic.TopicOrder,
				})
			}
		}

		categoryViewList = append(categoryViewList, &miikov1.Category{
			CategoryId:          categoryDto.CategoryId,
			CategoryDisplayName: categoryDto.CategoryDisplayName,
			CategorySystemName:  categoryDto.CategorySystemName,
			CategoryOrder:       categoryDto.CategoryOrder,
			TopicSize:           categoryDto.TopicSize,
			ProblemSize:         categoryDto.ProblemSize,
			TopicList:           viewTopicList,
		})
	}

	return &miikov1.CategoryListGetResponse{
		CategoryList: categoryViewList,
	}, nil
}
