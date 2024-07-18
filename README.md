# Todo list
I made this project in 3 days just to practice and improve my skills. Thanks for [that project](https://github.com/GolangLessons/url-shortener/tree/main) for giving motivation and discovering some features that I tried to use in my project.

## How to run it
Before you run it you need to set environment variable `CONFIG_PATH`. \
I was using Git Bash, here how you can set it there:
```bash
export CONFIG_PATH=./config/config.yaml
```
And then just run the project:
```bash
go run ./cmd/todo_list/main.go
```

## What does it do
### Add new task
You can add new task by POST method to `/add`. \
Request example:
```json
{
    "Title": "Play baseball",
    "Done": true
}
```

### View all tasks
You can view all tasks you've created by GET method to `/`.

### Delete task
You can't delete tasks lol. 🫡

## Some info about project
- I've used [go-chi](https://github.com/go-chi/chi) as web framework of project
- I've used [cleanenv](https://github.com/ilyakaznacheev/cleanenv) to read configs of project
- I've used [sqlite](https://github.com/mattn/go-sqlite3) as sqlite3 driver

## Conclusion
I think this project improved me as a developer a bit, but I'm not sure. And I'm not going to continue it.

That's all.
