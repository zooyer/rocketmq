package rocketmq

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/zooyer/rocketmq/common/mix"
	"strings"
	"testing"
)

func TestCheckTopic(t *testing.T) {
	tests := []string{
		"Hello",
		"%RETRY%Hello",
		"_%RETRY%Hello",
		"-%RETRY%Hello",
		"223-%RETRY%Hello",
	}
	for _, name := range tests {
		t.Run(name, func(t *testing.T) {
			if err := CheckTopic(name); err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestCheckTopic_HasIllegalCharacters(t *testing.T) {
	illegalTopic := "TOPIC&*^"
	if err := CheckTopic(illegalTopic); err != nil {
		diff := cmp.Diff(err.Error(), fmt.Sprintf("the specified topic[%s] contains illegal characters, allowing only %s", illegalTopic, ValidRegexpString))
		if diff != "" {
			t.Fatal(diff)
		}
	}
}

func TestCheckTopic_UseDefaultTopic(t *testing.T) {
	defaultTopic := mix.AutoCreateTopicKeyTopic
	if err := CheckTopic(defaultTopic); err != nil {
		diff := cmp.Diff(err.Error(), fmt.Sprintf("the topic[%s] is conflict with AUTO_CREATE_TOPIC_KEY_TOPIC", defaultTopic))
		if diff != "" {
			t.Fatal(diff)
		}
	}
}

func TestCheckTopic_BlankTopic(t *testing.T) {
	blankTopic := ""
	if err := CheckTopic(blankTopic); err != nil {
		diff := cmp.Diff(err.Error(), "the specified topic is blank")
		if diff != "" {
			t.Fatal(diff)
		}
	}
}

func TestCheckTopic_TooLongTopic(t *testing.T) {
	var buf strings.Builder
	tooLongTopic := "TooLongTopic"
	buf.WriteString(tooLongTopic)
	for i := 0; i < CharacterMaxLength+1-len(tooLongTopic); i++ {
		buf.WriteByte('_')
	}
	tooLongTopic = buf.String()
	if err := CheckTopic(tooLongTopic); err != nil {
		diff := cmp.Diff(err.Error(), "the specified topic is longer than topic max length 255")
		if diff != "" {
			t.Fatal(diff)
		}
	}
}
