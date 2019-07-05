package message

const (
	PropertyKeys                           = "KEYS"
	PropertyTags                           = "TAGS"
	PropertyWaitStoreMsgOk                 = "WAIT"
	PropertyDelayTimeLevel                 = "DELAY"
	PropertyRetryTopic                     = "RETRY_TOPIC"
	PropertyRealTopic                      = "REAL_TOPIC"
	PropertyRealQueueId                    = "REAL_QID"
	PropertyTransactionPrepared            = "TRAN_MSG"
	PropertyProducerGroup                  = "PGROUP"
	PropertyMinOffset                      = "MIN_OFFSET"
	PropertyMaxOffset                      = "MAX_OFFSET"
	PropertyBuyerId                        = "BUYER_ID"
	PropertyOriginMessageId                = "ORIGIN_MESSAGE_ID"
	PropertyTransferFlag                   = "TRANSFER_FLAG"
	PropertyCorrectionFlag                 = "CORRECTION_FLAG"
	PropertyMq2Flag                        = "MQ2_FLAG"
	PropertyReconsumeTime                  = "RECONSUME_TIME"
	PropertyMsgRegion                      = "MSG_REGION"
	PropertyTraceSwitch                    = "TRACE_ON"
	PropertyUniqClientMessageIdKeyidx      = "UNIQ_KEY"
	PropertyMaxReconsumeTimes              = "MAX_RECONSUME_TIMES"
	PropertyConsumeStartTimestamp          = "CONSUME_START_TIME"
	PropertyTransactionPreparedQueueOffset = "TRAN_PREPARED_QUEUE_OFFSET"
	PropertyTransactionCheckTimes          = "TRANSACTION_CHECK_TIMES"
	PropertyCheckImmunityTimeInSeconds     = "CHECK_IMMUNITY_TIME_IN_SECONDS"
	PropertyInstanceId                     = "INSTANCE_ID"

	KeySeparator = " "
)

var StringHashSet = make(map[string]struct{})

func init() {
	StringHashSet[PropertyTraceSwitch] = struct{}{}
	StringHashSet[PropertyMsgRegion] = struct{}{}
	StringHashSet[PropertyKeys] = struct{}{}
	StringHashSet[PropertyTags] = struct{}{}
	StringHashSet[PropertyWaitStoreMsgOk] = struct{}{}
	StringHashSet[PropertyDelayTimeLevel] = struct{}{}
	StringHashSet[PropertyRetryTopic] = struct{}{}
	StringHashSet[PropertyRealTopic] = struct{}{}
	StringHashSet[PropertyRealQueueId] = struct{}{}
	StringHashSet[PropertyTransactionPrepared] = struct{}{}
	StringHashSet[PropertyProducerGroup] = struct{}{}
	StringHashSet[PropertyMinOffset] = struct{}{}
	StringHashSet[PropertyMaxOffset] = struct{}{}
	StringHashSet[PropertyBuyerId] = struct{}{}
	StringHashSet[PropertyOriginMessageId] = struct{}{}
	StringHashSet[PropertyTransferFlag] = struct{}{}
	StringHashSet[PropertyCorrectionFlag] = struct{}{}
	StringHashSet[PropertyMq2Flag] = struct{}{}
	StringHashSet[PropertyReconsumeTime] = struct{}{}
	StringHashSet[PropertyUniqClientMessageIdKeyidx] = struct{}{}
	StringHashSet[PropertyMaxReconsumeTimes] = struct{}{}
	StringHashSet[PropertyConsumeStartTimestamp] = struct{}{}
	StringHashSet[PropertyInstanceId] = struct{}{}
}
