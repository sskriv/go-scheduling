# go-scheduling

for self-learning goroutines and channels

## Description

This program uses goroutines to execute tasks defined in the Job struct.  
It ensures that all running tasks are completed before exiting by utilizing sync.WaitGroup and channels.

## Features

- execute tasks using goroutines.
- graceful termination of the program.
