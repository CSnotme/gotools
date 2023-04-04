package internal

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"time"
)

func KafkaTest() {
	conn, err := kafka.Dial("tcp", "101.43.215.103:9092")
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	// testReadWrite()
}

func testReadWrite() {
	conn, err := kafka.Dial("tcp", "101.43.215.103:9092")
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	topicConfigs := []kafka.TopicConfig{}
	for _, topic := range topicMap {
		tp := kafka.TopicConfig{Topic: topic, NumPartitions: 1, ReplicationFactor: 1}
		topicConfigs = append(topicConfigs, tp)
	}

	err = conn.CreateTopics(topicConfigs...)
	if err != nil {
		panic(err.Error())
	}

	t := time.NewTimer(200 * time.Second)

	ctx, cancel := context.WithCancel(context.Background())

	go producer(ctx, "p1", pMap["p1"], topicMap["tp1"], -1, 20)
	go producer(ctx, "p2", pMap["p2"], topicMap["tp2"], 1, 10)
	go producer(ctx, "p3", pMap["p3"], topicMap["tp3"], 5, -1)

	go consumer(ctx, "c1", cMap["c1"])
	go consumer(ctx, "c2", cMap["c2"])
	go consumer(ctx, "c3", cMap["c3"])

	select {
	case <-t.C:
		cancel()
	}

	time.Sleep(2 * time.Second)
	fmt.Printf("shutdown")
}

var topicMap = map[string]string{
	"tp1": "test-p-1",
	"tp2": "test-p-2",
	"tp3": "test-p-3",
}

var pMap = map[string]*kafka.Writer{
	"p1": {Addr: kafka.TCP("101.43.215.103:9092"), AllowAutoTopicCreation: true, Balancer: &kafka.LeastBytes{}},
	"p2": {Addr: kafka.TCP("101.43.215.103:9092"), AllowAutoTopicCreation: true, Balancer: &kafka.LeastBytes{}},
	"p3": {Addr: kafka.TCP("101.43.215.103:9092"), AllowAutoTopicCreation: true, Balancer: &kafka.LeastBytes{}},
}

var cMap = map[string]kafka.ReaderConfig{
	"c1": {Brokers: []string{"101.43.215.103:9092"}, Topic: topicMap["tp1"], GroupID: "group-1"},
	"c2": {Brokers: []string{"101.43.215.103:9092"}, Topic: topicMap["tp2"], GroupID: "group-2"},
	"c3": {Brokers: []string{"101.43.215.103:9092"}, Topic: topicMap["tp3"], GroupID: "group-3"},
}

func producer(ctx context.Context, name string, w *kafka.Writer, topic string, interval int, loopNum int) {
	fmt.Printf("producer[%v] start!!!\n", name)
	i := 0
	for {
		msg := kafka.Message{
			Topic: topic,
			Key:   []byte(fmt.Sprintf("key-%v", i)),
			Value: []byte(fmt.Sprintf("我是%v", i)),
		}

		err := w.WriteMessages(ctx, msg)
		if err != nil {
			fmt.Printf("producer[%v] write msg:[%v]-[%v] to topic:%v, err:%v\n", name, msg.Key, msg.Value, msg.Topic, err)
			break
		}

		i++
		if loopNum > 0 && i >= loopNum {
			break
		}

		if interval > 0 {
			time.Sleep(time.Second * time.Duration(interval))
		}
	}

	fmt.Printf("producer[%v] stop!!!\n", name)
}

func consumer(ctx context.Context, name string, cCfg kafka.ReaderConfig) {
	fmt.Printf("consumer[%v] start!!!\n", name)
	r := kafka.NewReader(cCfg)
	defer r.Close()

	for {
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			fmt.Printf("consumer[%v] readMsg err:%v\n", name, err)
			break
		}

		fmt.Printf("consumer[%v] readMsg %+v\n", name, msg)

		err = r.CommitMessages(ctx, msg)
		if err != nil {
			fmt.Printf("consumer[%v] commitMsg err:%v\n", name, err)
			break
		}
	}

	fmt.Printf("consumer[%v] stop!!!\n", name)
}
