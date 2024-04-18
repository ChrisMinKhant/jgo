package com.jgo.framework;

import java.io.File;
import java.util.UUID;

import com.jgo.framework.javalib.jgocommute.JgoCommute;
import com.jgo.framework.javalib.jgocommute.implementation.JgoCommuteImplementation;

public class App {
    public static void main(String[] args) {

        JgoCommute jgoCommute = new JgoCommuteImplementation(UUID.randomUUID().toString(),"GoTest","Hello this is testing string.");

        new File("/mnt/edisk/jgo/.jgo");

        jgoCommute.push();
    }
}