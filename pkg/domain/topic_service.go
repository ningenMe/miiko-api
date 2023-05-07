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
