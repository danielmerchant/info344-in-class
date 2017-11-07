"use strict";

const MongoStore = require("./taskStore.js");
const mongodb = require("mongodb")

const mongoAddr = process.env.DBADDR || "localhost:27017";
const mongoURL = `mongodb://${mongoAddr}/tasks`;

describe("mongo task store", () => {
    test("crud cycle", () => {
        return mongodb.MongoClient.connect(mongoURL)
            .then(db => {
                let store = new MongoStore(db, "tasks");
                let task = {
                    title: "learn node in mongo",
                    tags: ["mongodb", "node.js", "info344"]
                };                 
                return store.insert(task)
                    .then(task => {
                        expect(task._id).toBeDefined();
                        return task._id;
                    })
                    .then(taskId => {
                        return store.get(taskId);
                    })
                    .then(fetchedTask => {
                        expect(fetchedTask).toEqual(task);
                        return store.update(task._id, {completed: true});
                    })
                    .then(updatedTask => {
                        expect(updatedTask.completed).toBe(true);
                        return store.delete(task._id);
                    })
                    .then(() => {
                        return store.get(task._id);
                    })
                    .then(fetchedTask => {
                        expect(fetchedTask).toBeFalsy();
                    })
                    .then(() => {
                        db.close();
                    })
                    .catch(err => {
                        db.close();
                        throw err;
                    });
            });
    });
});