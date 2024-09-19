module github.com/ballihautt/aoc_23_go

go 1.23.1

require (
    "github.com/ballihautt/aoc_23_go/day01" v0.0.0
    "github.com/ballihautt/aoc_23_go/day02" v0.0.0
)

replace (
    "github.com/ballihautt/aoc_23_go/day01" => "./day01"
    "github.com/ballihautt/aoc_23_go/day02" => "./day02"
)