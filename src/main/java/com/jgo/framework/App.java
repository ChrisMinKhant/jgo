package com.jgo.framework;

import com.jgo.framework.gorunner.GoRunner;
import com.jgo.framework.gorunner.golib.GoTest;
import com.jgo.framework.gorunner.golib.GoTestSecond;
import com.jgo.framework.gorunner.implementation.GoRunnerImplementation;

public class App {
    public static void main(String[] args) {
        GoRunner<GoTest> runner = new GoRunnerImplementation<GoTest>("GoTest", GoTest.class);
        GoRunner<GoTestSecond> runnerTwo = new GoRunnerImplementation<GoTestSecond>("GoTestSecond", GoTestSecond.class);

        GoTest goTest = (GoTest) runner.getRunner();
        GoTestSecond goTestSecond = (GoTestSecond) runnerTwo.getRunner();

        goTest.Test(1, 2, 3);
        goTestSecond.TestSecond();
    }
}