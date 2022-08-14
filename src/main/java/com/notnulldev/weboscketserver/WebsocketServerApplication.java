package com.notnulldev.weboscketserver;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.http.ResponseEntity;
import org.springframework.kafka.core.KafkaTemplate;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;


@SpringBootApplication
public class WebsocketServerApplication {

    private final static Logger logger = LoggerFactory.getLogger(WebsocketServerApplication.class);

    public static void main(String[] args) {
        SpringApplication.run(WebsocketServerApplication.class, args);
    }
}

@RestController
@RequestMapping(name = "api/v1")
class IncomingMessageController {
    private final KafkaTemplate<String, IncomingMessage> kafkaTemplate;
    private final Logger logger = LoggerFactory.getLogger(IncomingMessageController.class);

    @Value("{incomingTopicName:MOVED}")
    private String responseMessage;

    IncomingMessageController(KafkaTemplate<String, IncomingMessage> kafkaTemplate) {
        this.kafkaTemplate = kafkaTemplate;
        logger.info("Created IncomingMessageController");
    }

    @PostMapping
    public ResponseEntity<IncomingMessage> putMessage(IncomingMessage message) {
        kafkaTemplate.send(KafkaConfig.INCOMING_MESSAGE_TOPIC, message);
        logger.info(message.toString());
        return ResponseEntity.status(200).body(message);
    }

}