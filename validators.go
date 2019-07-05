package rocketmq

import (
	"errors"
	"fmt"
	"github.com/zooyer/rocketmq/common/message"
	"github.com/zooyer/rocketmq/common/mix"
	"github.com/zooyer/rocketmq/common/utils"
	"regexp"
)

const (
	ValidRegexpString = "[%|a-zA-Z0-9_-]+$"
	ValidPatternString = ValidRegexpString
	CharacterMaxLength = 255
)

var Regexp = regexp.MustCompile(ValidRegexpString)

//func GetGroupWithRegularExpression(origin, patternStr string) string {
//	var pattern = regexp.MustCompile(patternStr)
//	matcher = pattern.MatchString(origin)
//}

// RegularExpressionMatcher if, and only if, the entire origin sequence matches this matcher's pattern
func RegularExpressionMatcher(origin string, regexp *regexp.Regexp) bool {
	if regexp == nil {
		return true
	}
	return regexp.MatchString(origin)
}

func CheckMessage(msg *message.Message) error {
	return nil
}

// CheckTopic validate topic format
func CheckTopic(topic string) error {
	if utils.IsBlank(topic) {
		return errors.New("the specified topic is blank")
	}
	if !RegularExpressionMatcher(topic, Regexp) {
		return fmt.Errorf("the specified topic[%s] contains illegal characters, allowing only %s", topic, ValidRegexpString)
	}
	if len(topic) > CharacterMaxLength {
		return errors.New("the specified topic is longer than topic max length 255")
	}
	if topic == mix.AutoCreateTopicKeyTopic {
		return fmt.Errorf("the topic[%s] is conflict with AUTO_CREATE_TOPIC_KEY_TOPIC", topic)
	}
	return nil
}
