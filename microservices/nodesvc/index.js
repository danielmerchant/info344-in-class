// @ts-check
"use strict";

//load express and morgan modules
const express = require("express");
const morgan = require("morgan");

const addr = process.env.ADDR;
const[host, port] = addr.split(":");
const portNum = parseInt(port);

const app = express();
app.use(morgan(process.env.LOG_FORMAT || "dev"));

app.get("/", (req, res) => {
    res.set("Content-Type", "text/plain");
    res.send("Hello, Node.js!");
});

app.get("/v1/users/me/hello", (req, res) => {
    let userJSON = req.get("X-User");
    if (!userJSON) {
        throw new Error("no x-user header provided");
    }
    let user = JSON.parse(userJSON);
    res.json({
        message: `hello ${user.firstName} ${user.lastName}`
    });
});

const handlers = require("./handlers.js");
app.use(handlers({}));

app.use((err, req, res, next) => {
    console.error(err.stack);
    res.set("Content-Type", "text/plain");
    res.send(err.message);
});

app.listen(portNum, host, () => {
    console.log(`server is listening at http://${addr}...`);
});