package common

import (
	"strings"
	"testing"
)

func TestTest(t *testing.T) {
	var str = `
StringHashSet[PROPERTY_TRACE_SWITCH] = PROPERTY_TRACE_SWITCH
	StringHashSet[PROPERTY_MSG_REGION] = PROPERTY_MSG_REGION
	StringHashSet[PROPERTY_KEYS] = PROPERTY_KEYS
	StringHashSet[PROPERTY_TAGS] = PROPERTY_TAGS
	StringHashSet[PROPERTY_WAIT_STORE_MSG_OK] = PROPERTY_WAIT_STORE_MSG_OK
	StringHashSet[PROPERTY_DELAY_TIME_LEVEL] = PROPERTY_DELAY_TIME_LEVEL
	StringHashSet[PROPERTY_RETRY_TOPIC] = PROPERTY_RETRY_TOPIC
	StringHashSet[PROPERTY_REAL_TOPIC] = PROPERTY_REAL_TOPIC
	StringHashSet[PROPERTY_REAL_QUEUE_ID] = PROPERTY_REAL_QUEUE_ID
	StringHashSet[PROPERTY_TRANSACTION_PREPARED] = PROPERTY_TRANSACTION_PREPARED
	StringHashSet[PROPERTY_PRODUCER_GROUP] = PROPERTY_PRODUCER_GROUP
	StringHashSet[PROPERTY_MIN_OFFSET] = PROPERTY_MIN_OFFSET
	StringHashSet[PROPERTY_MAX_OFFSET] = PROPERTY_MAX_OFFSET
	StringHashSet[PROPERTY_BUYER_ID] = PROPERTY_BUYER_ID
	StringHashSet[PROPERTY_ORIGIN_MESSAGE_ID] = PROPERTY_ORIGIN_MESSAGE_ID
	StringHashSet[PROPERTY_TRANSFER_FLAG] = PROPERTY_TRANSFER_FLAG
	StringHashSet[PROPERTY_CORRECTION_FLAG] = PROPERTY_CORRECTION_FLAG
	StringHashSet[PROPERTY_MQ2_FLAG] = PROPERTY_MQ2_FLAG
	StringHashSet[PROPERTY_RECONSUME_TIME] = PROPERTY_RECONSUME_TIME
	StringHashSet[PROPERTY_UNIQ_CLIENT_MESSAGE_ID_KEYIDX] = PROPERTY_UNIQ_CLIENT_MESSAGE_ID_KEYIDX
	StringHashSet[PROPERTY_MAX_RECONSUME_TIMES] = PROPERTY_MAX_RECONSUME_TIMES
	StringHashSet[PROPERTY_CONSUME_START_TIMESTAMP] = PROPERTY_CONSUME_START_TIMESTAMP
	StringHashSet[PROPERTY_INSTANCE_ID] = PROPERTY_INSTANCE_ID
`
	var buf strings.Builder
	bytes := []byte(str)
	for i := 0; i < len(bytes); i++ {
		c := bytes[i]
		switch {
		case 'A' <= c && c <= 'Z':

		default:
			buf.WriteByte(c)
		}
	}
}
