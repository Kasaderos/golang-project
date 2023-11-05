package kafka

import (
	"github.com/Shopify/sarama"
)

func BuildMessage(topic string, key string, message []byte) *sarama.ProducerMessage {
	// if you want to send headers ([]string)
	// do smth like this

	// if len(headersKV)%2 != 0 {
	// 	return nil, errors.New("wrong number of headersKV")
	// }

	// headers := make([]sarama.RecordHeader, 0, len(headersKV)/2)
	// for i := 0; i < len(headersKV); i += 2 {
	// 	headers = append(headers, sarama.RecordHeader{
	// 		Key:   []byte(headersKV[i]),
	// 		Value: []byte(headersKV[i+1]),
	// 	})
	// }
	// add headers to Headers

	return &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.ByteEncoder(message),
	}
}
