package com.jgo.framework.javalib.jgocommute.implementation;

import java.io.File;
import java.io.FileWriter;
import java.io.IOException;
import java.nio.charset.Charset;
import java.nio.charset.StandardCharsets;
import java.security.InvalidAlgorithmParameterException;
import java.security.InvalidKeyException;
import java.security.NoSuchAlgorithmException;
import java.security.SecureRandom;
import java.util.Base64;

import javax.crypto.BadPaddingException;
import javax.crypto.Cipher;
import javax.crypto.IllegalBlockSizeException;
import javax.crypto.NoSuchPaddingException;
import javax.crypto.SecretKey;
import javax.crypto.spec.GCMParameterSpec;
import javax.crypto.spec.SecretKeySpec;

import com.jgo.framework.javalib.jgocommute.JgoCommute;

public class JgoCommuteImplementation implements JgoCommute {

    private final String jgoFilePath = "/mnt/edisk/jgo/.jgo";
    private final String jgoSecretPath = "/mnt/edisk/jgo/.secret";
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

        String preparedMessage = encode(this.messageId) + "|" + encode(this.functionName) + "|" + encode(this.message)
                + "|" + encode(this.commit.toString());

        String encryptedMessage = "";

        try {

            String[] returnedValue = new String[2];

            returnedValue = encryptMessage(preparedMessage);

            encryptedMessage = returnedValue[1];

            FileWriter fileWriter = new FileWriter(this.jgoSecretPath,true);

            fileWriter.write(returnedValue[0]);

            fileWriter.close();

        } catch (InvalidKeyException e) {
            // TODO Auto-generated catch block
            e.printStackTrace();
        } catch (NoSuchAlgorithmException e) {
            // TODO Auto-generated catch block
            e.printStackTrace();
        } catch (NoSuchPaddingException e) {
            // TODO Auto-generated catch block
            e.printStackTrace();
        } catch (InvalidAlgorithmParameterException e) {
            // TODO Auto-generated catch block
            e.printStackTrace();
        } catch (IllegalBlockSizeException e) {
            // TODO Auto-generated catch block
            e.printStackTrace();
        } catch (BadPaddingException e) {
            // TODO Auto-generated catch block
            e.printStackTrace();
        } catch (IOException e){
            e.printStackTrace();
        }
        return encryptedMessage + "|";
    }

    private String encode(String rawString) {
        return Base64.getEncoder().encodeToString(rawString.getBytes());
    }

    @Override
    public String[] encryptMessage(String message) throws NoSuchAlgorithmException, NoSuchPaddingException,
            InvalidKeyException, InvalidAlgorithmParameterException, IllegalBlockSizeException, BadPaddingException {

        final SecureRandom secure = new SecureRandom();
        byte[] key = new byte[32];
        secure.nextBytes(key);

        SecretKey secretKey = new SecretKeySpec(key, "AES");

        Cipher cipher = Cipher.getInstance("AES/GCM/NoPadding");

        byte[] iv = new byte[12];
        secure.nextBytes(iv);

        GCMParameterSpec gcmParameterSpec = new GCMParameterSpec(128, iv);

        cipher.init(Cipher.ENCRYPT_MODE, secretKey, gcmParameterSpec);

        byte[] cipherText = cipher.doFinal("Hello this is testing string to encrypt".getBytes(StandardCharsets.UTF_8));

        byte[] combinnedCipterText = new byte[iv.length + cipherText.length];

        System.arraycopy(iv, 0, combinnedCipterText, 0, iv.length);
        System.arraycopy(cipherText, 0, combinnedCipterText, iv.length, cipherText.length);

        String[] encryptedMessageWithSecret = new String[2];

        encryptedMessageWithSecret[0] = new String(secretKey.getEncoded());
        encryptedMessageWithSecret[1] = new String(combinnedCipterText);

        return encryptedMessageWithSecret;
    }

}
