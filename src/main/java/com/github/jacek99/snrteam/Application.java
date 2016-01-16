package com.github.jacek99.snrteam;

import static spark.Spark.*;

public class Application {

    public static void main(String[] args) {
        get("/hello", (req, res) -> "Hello World");
    }

}