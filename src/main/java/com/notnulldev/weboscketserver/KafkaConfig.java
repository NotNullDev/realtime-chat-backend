package com.notnulldev.weboscketserver;

import org.apache.kafka.clients.admin.NewTopic;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.kafka.config.TopicBuilder;

@Configuration
public class KafkaConfig {

    @Value("{app.incomingMessageTopic:incomingMessage}")
    public static final String INCOMING_MESSAGE_TOPIC = "incomingMessage";

    @Bean
    NewTopic incomingMessage() {
        return TopicBuilder
                .name(INCOMING_MESSAGE_TOPIC)
                .replicas(1)
                .build();
    }

    @KafkaListener(topics = INCOMING_MESSAGE_TOPIC, groupId = KafkaConfig.INCOMING_MESSAGE_TOPIC)
    void newMessageTopic(IncomingMessage message) {
        System.out.println(message);
    }
}
