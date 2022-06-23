
# Go_STUDY

### A bunch of random exercise, functions, logics, tests, algorithms, and tests while studying the Go programming language. This is my journey of understanding the GoLang documented.

###

### I will be writing most of the logics, functions or algorithms in this repo in both Go and Python. I find that I learn even better when I'm studying a language while comparing it with a language I can already strongly type in. It helps me understand the language I am studying in contrast to the one I already know. I choose Python for this mainly because Python can be run almost anywhere on any project structure without having to setup anything. 

### For example, if the project root folder is opened in visual studio code, you can run the python files directly by just typing in windows

``` 
python file.py
```
### While for Linux or Mac
``` 
python3 file.py
```
### To run the Go files, you run the main.go file
```
go run main.go
```
### The python codes on this repo are written on version 3.10, so there are some features used that may not be available for previous versions of python. You can edit these though if you are running a version lower than 3.10

### According to the generic speed test I conducted on 3 popular programming languages (Python, Go, and Ruby), It showed that my preferred Language of choice, Python was the slowest by a mile. But would it sway me away though, never!. Below are the test results conducted. The programs looped through 10000000 appended a random number between 1 and 100 in an array/list and sorted this array before printing out the time it took to complete the task. The test was run on a laptop with a corei3 10th Gen @1.20GHz 8GB ram(I don't think this matters though) on power saver mood.

```
python main.py
```
#### Result: 7116.7362ms


```
ruby main.rb
```
#### Result: 1917.7393ms

```
go run main.go
```
#### Result: 960.0633ms

### From the result, you can see Go was 2X faster than Ruby, and 8X faster than Python. On the other hand, Ruby was almost 4X faster than Python and was actually faster when I switched the loops from for to while loops. Ruby came in at 1534.7823ms while the other two maintained the same result as their respective For Loops.


### This test was merely for practice and is by no means a conclusion. Programming isn't always about speed, there are so many deciding factors. The first languages I ever looked into were Java & PHP, but I would forever be the Python guy as I love Python irrespective of it's shortcomings. Just learning how much fun Ruby and Go are to write with though, but still not as much fun it is writing Python.


#### Will continually update this readme as I progress.
