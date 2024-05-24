
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
Given an American Football match score, find how many unique combinations there are given the following score rules:</br>
  - Touchdow: 6 points -> followed by an extra 1 or 2 points (or zero) </br>
  - Field Goal: 3 points with no extras
### Solution
For each possible scoring option, we can either include it on the target sum, or exclude it. By including it, we call the verify function recursively with the new sum value and the set of possible scoring options. If we exclude it, the function is called recursively with the same sum and the number of the remaining scoring options(len({3,6,7,8}) -1).

Our base cases for this approach are:
- When the sum is 0, which means the only way to "progress" further is by selecting NO scoring option. So we return 1.
- when the sum is negative, which means that the current branch is impossible for that case. Returning 0.

The time complexity for this implementation is O(2<sup>sum</sup>)

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
