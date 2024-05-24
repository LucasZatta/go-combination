
# Go-Combination

## Project Structure:

```
├── cmd
│   └──api  
|      └──main.go      \\ Server startup
├── internal
│   └──server
|      └──routes.go     \\ Where routes and handlers are setup
|      └──server.go     \\ Server config and setups
├── Dockerfile
├── go.mod              \\ Go dependencies file
├── go.sum
├── makefile            
```

## Considerations

### Structure
Simple api structure implemented to handle a single POST route that solves the score combination problem.

### Score Combinations
Given an American Football match score, find how many unique combinations there are given the following score rules:
  Touchdow: 6 points -> followed by an extra 1 or 2 points (or zero)
  Field Goal: 3 points with no extras

## Rest API
### POST    /verify
Body Example: {"score":"3x15"}

Response Example: {"combinations" : "4"}

## Running the project
`docker run -p 8080:8080 lucasszatta/go-sol-combinations:latest`
