package com.notnulldev.weboscketserver;

import lombok.AllArgsConstructor;
import lombok.NoArgsConstructor;
import lombok.ToString;

@AllArgsConstructor
@ToString
@NoArgsConstructor
public class IncomingMessage {
    private int Id;
    private String channelName;
    private String ownerId;
    private String messageContent;
}
