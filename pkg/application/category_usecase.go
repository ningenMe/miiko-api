package application

import (
	"github.com/ningenMe/miiko-api/pkg/infra"
	miikov1 "github.com/ningenMe/miiko-api/proto/gen_go/v1"
	"net/http"
)

type CategoryUsecase struct{}

func (CategoryUsecase) CategoryListGet(isRequiredTopic bool) (*miikov1.CategoryListGetResponse, error) {

	//データ取得
	categoryDtoList := categoryRepository.GetList()

	//データ整形
	var categoryViewList []*miikov1.Category
	for _, categoryDto := range categoryDtoList {

		var topicViewList []*miikov1.Topic
		if isRequiredTopic {
			topicList := topicRepository.GetListByCategoryId(categoryDto.CategoryId, false)
			for _, topic := range topicList {
				topicViewList = append(topicViewList, &miikov1.Topic{
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
			TopicList:           topicViewList,
		})
	}

	return &miikov1.CategoryListGetResponse{
		CategoryList: categoryViewList,
	}, nil
}

func (CategoryUsecase) CategoryPost(header http.Header, categoryId string, category *miikov1.Category) (*miikov1.CategoryPostResponse, error) {

	if err := authorizationService.CheckLoggedIn(header); err != nil {
		return &miikov1.CategoryPostResponse{}, err
	}

	if category != nil {
		//作成 or 更新

		//idがないときは新規作成
		if categoryId == "" {
			categoryId = infra.GetNewCategoryId()
		}

		categoryRepository.Upsert(
			&infra.CategoryDto{
				CategoryId:          categoryId,
				CategoryDisplayName: category.GetCategoryDisplayName(),
				CategorySystemName:  category.GetCategorySystemName(),
				CategoryOrder:       category.GetCategoryOrder(),
			})

	} else {
		//削除
		//TODO 削除は別エンドポイントにしてこのロジックはどこかでなくす
		categoryRepository.Delete(categoryId)
	}

	return &miikov1.CategoryPostResponse{CategoryId: categoryId}, nil
}
