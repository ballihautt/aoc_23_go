# aoc_Go

Aoc completion in Go.

## Why ?

This repository aims to show what I can do in Go, and so I chose to complete the Advent of Code.
I wanted to finish the 2023 calendar I started last year, but not in C. I start it again, now in Go.

## How ?

I will try to create a single entry point, so I will not recreate basic tools each day.
For that, they will be an option `-f`, taking a string, and which will be the input file path.
A second option, `-d`, will take an unsigned integer, which determines the subfunction to call.
I will be able to call the day solution using nearly same command for that for the other days.
