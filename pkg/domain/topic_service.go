package domain

import (
	"github.com/ningenMe/miiko-api/pkg/infra"
)

type TopicService struct{}

var topicRepository = infra.TopicRepository{}
var referenceRepository = infra.ReferenceRepository{}

func (TopicService) Upsert(topicDto *infra.TopicDto) error {
	if err := topicRepository.Upsert(topicDto); err != nil {
		return err
	}
	if err := referenceRepository.Delete(topicDto.TopicId); err != nil {
		return err
	}
	for _, referenceDto := range topicDto.ReferenceList {
		if err := referenceRepository.Insert(topicDto.TopicId, referenceDto); err != nil {
			return err
		}
	}
	return nil
}

func (TopicService) GetWithProblemWithTag(topicId string) (*infra.TopicDto, error) {
	topicDto, err := topicRepository.Get(topicId)
	if err != nil {
		return nil, err
	}
	//TODO ハンドリングを直す
	topicDto.ProblemList = problemRepository.GetProblemListByTopicId(topicId, true)
	//TODO エラー拾うと空の時おかしくなる気がする、後で修正
	topicDto.ReferenceList, _ = referenceRepository.Get(topicId)

	return topicDto, nil
}

func (TopicService) GetListWithProblemWithTagByCategoryId(categoryId string) ([]*infra.TopicDto, error) {
	topicDtoList, err := topicRepository.GetListByCategoryId(categoryId)
	if err != nil {
		return nil, err
	}
	for _, topicDto := range topicDtoList {
		//TODO ハンドリングを直す
		topicDto.ProblemList = problemRepository.GetProblemListByTopicId(topicDto.TopicId, false)
		//TODO エラー拾うと空の時おかしくなる気がする、後で修正
		topicDto.ReferenceList, _ = referenceRepository.Get(topicDto.TopicId)
	}

	return topicDtoList, nil
}
