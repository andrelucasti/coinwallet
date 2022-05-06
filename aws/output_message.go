package aws

type Result struct {
	OutputMessage OutputMessage
}

type OutputMessage struct {
	MessageId string
	Body      string
}
