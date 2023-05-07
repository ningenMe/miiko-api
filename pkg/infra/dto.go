package infra

import "crypto/rand"

const (
	chars      = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	chars_size = len(chars)
)

type CategoryDto struct {
	CategoryId          string `db:"category_id"`
	CategoryDisplayName string `db:"category_display_name"`
	CategorySystemName  string `db:"category_system_name"`
	CategoryOrder       int32  `db:"category_order"`
	TopicSize           int32  `db:"topic_size"`
	ProblemSize         int32  `db:"problem_size"`
}

type TopicDto struct {
	TopicId          string `db:"topic_id"`
	CategoryId       string `db:"category_id"`
	TopicDisplayName string `db:"topic_display_name"`
	TopicOrder       int32  `db:"topic_order"`
	TopicText        string `db:"topic_text"`
	ProblemList      []*ProblemDto
	ReferenceList    []*ReferenceDto
}

type ProblemDto struct {
	ProblemId          string `db:"problem_id"`
	Url                string `db:"url"`
	ProblemDisplayName string `db:"problem_display_name"`
	Estimation         int32  `db:"estimation"`
	TagList            []*TagDto
}

type TagDto struct {
	ProblemId        string `db:"problem_id"`
	TopicId          string `db:"topic_id"`
	CategoryId       string `db:"category_id"`
	TopicDisplayName string `db:"topic_display_name"`
}

type ReferenceDto struct {
	ReferenceId          string `db:"reference_id"`
	Url                  string `db:"url"`
	ReferenceDisplayName string `db:"reference_display_name"`
}

func GetNewCategoryId() string {
	return "category_" + getRandomString(6)
}

func GetNewTopicId() string {
	return "topic_" + getRandomString(6)
}

func GetNewProblemId() string {
	return "problem_" + getRandomString(6)
}

// https://qiita.com/nakaryooo/items/7d269525a288c4b3ecda
func getRandomString(digit uint32) string {

	b := make([]byte, digit)
	if _, err := rand.Read(b); err != nil {
		return ""
	}

	var result string
	for _, v := range b {
		result += string(chars[int(v)%chars_size])
	}
	return result
}
