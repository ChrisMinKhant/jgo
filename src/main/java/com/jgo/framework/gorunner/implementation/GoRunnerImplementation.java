package com.jgo.framework.gorunner.implementation;

import com.jgo.framework.gorunner.GoRunner;
import com.sun.jna.Library;
import com.sun.jna.Native;

public class GoRunnerImplementation<T extends Library> implements GoRunner<T> {

    private T runner;

    public GoRunnerImplementation(String path, Class<T> goLibraryClass) {
        this.runner = (T) Native.load("src/lib/" + path + "/" + path + ".so", goLibraryClass);
    }

    @Override
    public T getRunner() {
        return this.runner;
    }
}