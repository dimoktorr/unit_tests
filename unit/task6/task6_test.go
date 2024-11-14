package task6

import "testing"

//имеются методы b.ResetTimer(), b.StopTimer() и b.StartTimer()

func byCopy() Chat {
	return Chat{
		ID:          "1",
		OrderID:     "1",
		OrderNumber: "1",
		Source:      "1",
		Topic:       "1",
		Messages: []Message{
			{
				Sender:         "1",
				SystemName:     "1",
				SystemVersion:  "1",
				AppVersion:     "1",
				Device:         "1",
				MtsID:          "1",
				ClientDatetime: 1,
				CreatedAt:      1,
				DeliveredAt:    1,
			},
		},
	}
}

func messageCopy() Message {
	return Message{
		Sender:         "1",
		SystemName:     "1",
		SystemVersion:  "1",
		AppVersion:     "1",
		Device:         "1",
		MtsID:          "1",
		ClientDatetime: 1,
		CreatedAt:      1,
		DeliveredAt:    1,
	}
}
func messagePointer() *Message {
	return &Message{
		Sender:         "1",
		SystemName:     "1",
		SystemVersion:  "1",
		AppVersion:     "1",
		Device:         "1",
		MtsID:          "1",
		ClientDatetime: 1,
		CreatedAt:      1,
		DeliveredAt:    1,
	}
}

func (s Chat) stack(chat Chat) Chat {
	return chat
}

func (s Chat) stackAddMessage(message Message) Chat {
	s.Messages = append(s.Messages, message)
	return s
}

func (s *Chat) heap(chat *Chat) *Chat {
	return chat
}

func (s *Chat) heapAddMessage(message Message) *Chat {
	s.Messages = append(s.Messages, message)
	return s
}

func (s *Chat) stackHeap(chat Chat) *Chat {
	return &chat
}

func (s ChatMessagesPointer) stack(chat ChatMessagesPointer) ChatMessagesPointer {
	return chat
}

func (s ChatMessagesPointer) stackAddMessage(message *Message) ChatMessagesPointer {
	s.Messages = append(s.Messages, message)
	return s
}

func (s *ChatMessagesPointer) heap(chat *ChatMessagesPointer) *ChatMessagesPointer {
	return chat
}

func (s *ChatMessagesPointer) heapAddMessage(message *Message) *ChatMessagesPointer {
	s.Messages = append(s.Messages, message)
	return s
}

func (s *ChatMessagesPointer) stackHeap(chat ChatMessagesPointer) *ChatMessagesPointer {
	return &chat
}

func byChatPointer() *Chat {
	return &Chat{
		ID:          "1",
		OrderID:     "1",
		OrderNumber: "1",
		Source:      "1",
		Topic:       "1",
		Messages: []Message{
			{
				Sender:         "1",
				SystemName:     "1",
				SystemVersion:  "1",
				AppVersion:     "1",
				Device:         "1",
				MtsID:          "1",
				ClientDatetime: 1,
				CreatedAt:      1,
				DeliveredAt:    1,
			},
		},
	}
}

func byChatMessagesPointer() *ChatMessagesPointer {
	return &ChatMessagesPointer{
		ID:          "1",
		OrderID:     "1",
		OrderNumber: "1",
		MtsID:       "1",
		UserEmail:   "1",
		Source:      "1",
		Topic:       "1",
		Messages: []*Message{
			{
				Sender:         "1",
				SystemName:     "1",
				SystemVersion:  "1",
				AppVersion:     "1",
				Device:         "1",
				MtsID:          "1",
				ClientDatetime: 1,
				CreatedAt:      1,
				DeliveredAt:    1,
			},
		},
	}
}

func byChatMessagesCopy() ChatMessagesPointer {
	return ChatMessagesPointer{
		ID:          "1",
		OrderID:     "1",
		OrderNumber: "1",
		MtsID:       "1",
		UserEmail:   "1",
		Source:      "1",
		Topic:       "1",
		Messages: []*Message{
			{
				Sender:         "1",
				SystemName:     "1",
				SystemVersion:  "1",
				AppVersion:     "1",
				Device:         "1",
				MtsID:          "1",
				ClientDatetime: 1,
				CreatedAt:      1,
				DeliveredAt:    1,
			},
		},
	}
}

func BenchmarkMemoryStack(b *testing.B) {
	var chat Chat
	var chat1 Chat

	chat = byCopy()
	chat1 = byCopy()

	for i := 0; i < b.N; i++ {
		for i := 0; i < 1000000; i++ {
			chat = chat.stack(chat1)
		}
	}

	b.StopTimer()
}

func BenchmarkMemoryStack_AddMessage(b *testing.B) {
	var chat Chat

	chat = byCopy()

	for i := 0; i < b.N; i++ {
		for i := 0; i < 1000000; i++ {
			chat = chat.stackAddMessage(messageCopy())
		}
	}

	b.StopTimer()
}

func BenchmarkMemoryHeap(b *testing.B) {
	var chat *Chat
	var chat1 *Chat

	chat = byChatPointer()
	chat1 = byChatPointer()

	for i := 0; i < b.N; i++ {
		for i := 0; i < 1000000; i++ {
			chat = chat.heap(chat1)
		}
	}

	b.StopTimer()
}

func BenchmarkMemoryHeap_AddMessage(b *testing.B) {
	var chat *Chat

	chat = byChatPointer()

	for i := 0; i < b.N; i++ {
		for i := 0; i < 1000000; i++ {
			chat = chat.heapAddMessage(messageCopy())
		}
	}

	b.StopTimer()
}

func BenchmarkMemoryStackHeap(b *testing.B) {
	var chat *Chat
	var chat1 Chat

	chat = byChatPointer()
	chat1 = byCopy()

	for i := 0; i < b.N; i++ {
		for i := 0; i < 1000000; i++ {
			chat = chat.stackHeap(chat1)
		}
	}

	b.StopTimer()
}

func BenchmarkMemoryHeap_ChatMessagesPointer_AddMessage(b *testing.B) {
	var chat *ChatMessagesPointer

	chat = byChatMessagesPointer()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for i := 0; i < 1000000; i++ {
			chat = chat.heapAddMessage(messagePointer())
		}
	}

	b.StopTimer()
}

func BenchmarkMemoryStack_ChatMessagesPointer_AddMessage(b *testing.B) {
	var chat ChatMessagesPointer

	chat = byChatMessagesCopy()

	for i := 0; i < b.N; i++ {
		for i := 0; i < 1000000; i++ {
			chat = chat.stackAddMessage(messagePointer())
		}
	}

	b.StopTimer()
}
