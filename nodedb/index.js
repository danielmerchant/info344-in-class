"use strict";

const mongodb = require("mongodb");
const MongoStore = require("./taskstore");
const express = require("express");
const mongoAddr = process.env.DBADDR || "localhost:27017";
const mongoURL = `mongodb://${mongoAddr}/tasks`;
const app = express();

const addr = process.env.ADDR || "localhost:4000";
const [host, port] = addr.split(":");

app.use(express.json());

mongodb.MongoClient.connect(mongoURL)
    .then(db => {
        let taskStore = new MongoStore(db, "tasks");
        app.post("/v1/tasks", (req, res) => {
            //insert new
            let task = {
                title: req.body.title,
                completed: false
            }
            taskStore.insert(task)
                .then(task => {
                    res.json(task);
                })
                .catch(err => {
                    throw err;
                });
        });
        
        app.get("/v1/tasks", (req, res) => {
            //return all not completed tasks in the database
        });
        
        app.patch("/v1/tasks/:taskID", (req, res) => {
            let taskIDToFatch = req.params.taskID;
            //update single task by id and send to client
        });
        
        app.delete("/v1/tasks/:taskID", (req, res) => {
            //delete
        })
        
        app.listen(port, host, () => {
            console.log(`server is listening at http://${addr}....`)
        });
    })
    .catch(err => {
        throw err;
});


