package com.jgo.framework.javalib.jgocommute.implementation;

import java.io.FileWriter;
import java.io.IOException;
import java.util.Base64;

import com.jgo.framework.javalib.jgocommute.JgoCommute;

public class JgoCommuteImplementation implements JgoCommute {

    private final String jgoFilePath = "/mnt/edisk/jgo/.jgo";
    private String messageId;
    private String functionName;
    private String message;
    private Boolean commit = false;

    public JgoCommuteImplementation(String messageId, String functionName, String message) {
        this.messageId = messageId;
        this.functionName = functionName;
        this.message = message;
    }

    @Override
    public void push() {
        writeToFile();
    }

    private void writeToFile() {
        try {
            FileWriter writer = new FileWriter(this.jgoFilePath, true);

            writer.write(validateMessage());

            writer.close();
        } catch (IOException e) {
            // TODO Auto-generated catch block
            e.printStackTrace();
        }
    }

    private String validateMessage() {

        String preparedMessage = encode(this.messageId)+ "|" + encode(this.functionName) + "|" + encode(this.message) + "|" + encode(this.commit.toString());

        return encode(preparedMessage)+"|";
    }

    private String encode(String rawString){
        return Base64.getEncoder().encodeToString(rawString.getBytes());
    }

}
