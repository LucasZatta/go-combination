
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

## Extra
Attempt at solving the score combination problem, which is essentially a number partition problem given a limited set of numbers(6,7,8,3), using recursion.
Did not manage to think of a solution using this initial approach that would return me only the UNIQUE combinations.
```
func verify(n int) int {
	if n == 0 {
		return 1
	} else if n < 0 {
		return 0
	}
	return (verify(n-6) + verify(n-7) + verify(n-8) + verify(n-3))
}
```
