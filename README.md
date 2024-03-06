# Advent of Code 2023 in Go
I used the 2023 Advent of Code challenges to learn how to write Go! As a first-timer for both Go and Advent of Code, this project has been a great learning experience. 


## Running the code
To run any of the challenges, navigate to the challenge directory and use the Go command:

```
cd dayX # Replace X with the challenge day
go run main.go
```
This will compile and execute the solution for that day's challenge, printing the output to the console.

[I have gitignored all of my data files](https://adventofcode.com/2023/about). If you want to test code with other inputs you can create a `data.txt` file in the day directory and uncomment the runner code in the main function.

## The setup.sh Script
I've included a setup.sh script. This script automates the creation of the directory structure and basic files for a new day's challenge. To use it, simply run the following from the root of the directory:

```
./setup.sh dayX
```
This will create a new directory for the day with a template main.go and an empty data.txt file.

## Reflections and shoutouts
I've had a really great time getting more familiar with Go, and I'm eager to use it more in the future. Shout out to [Derat's Go solutions](https://codeberg.org/derat/advent-of-code) for helping me when I got stuck, and [the subreddit](https://www.reddit.com/r/adventofcode/) for its brilliant and unhinged solutions. 

Happy to answer any questions.

