# Word Count Tool 

Write Your Own wc Tool
1st challenge in the Coding Challenges by John Crickett: 

# Steps 
While John's site has details on what to do, I want to keep a log on my approach. 
Link: https://codingchallenges.fyi/challenges/challenge-wc/ 

## Step 0
IDE - VSCode 
- This is alredy installed on my machine at the time of the project.

Language - Go 
- I do not know Go but have been wanting to try it out for a long time now.


## Step 1
What should be done?
- [X] Read the number of bytes in a file
- [X] Output the number of bytes in the file
- [X] Add a "-c" CLI option
- [X] Read the input from the CLI ("test.txt")

What did I learn?
- Assignment operator, for loop, packages, opening/closing files in Go
- CLI arguments
- CLI flags

## Step 2
What should be done?
- [X] Add a "-l" CLI option
- [X] Output the number of lines in the file

What did I learn?
- How to use bufio to scan a file line by line

## Step 3
What should be done?
- [X] Add a "-w" CLI option
- [X] Output the number of words in the file

What did I learn?
- ScanWords in bufio can be used to read a file word by word

## Step 4
What should be done?
- [X] Add a "-m" CLI option
- [X] Output the number of characters in the file

What did I learn?
- ScanRunes in bufio can be used to read a file character by character


