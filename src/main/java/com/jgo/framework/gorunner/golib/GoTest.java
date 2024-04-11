package com.jgo.framework.gorunner.golib;

import com.sun.jna.Library;

public interface GoTest extends Library {
    void Test(int firstValue, int secondValue);
}