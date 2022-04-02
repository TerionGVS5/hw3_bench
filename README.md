# hw3_bench
There is a function that looks for something in the file. But she doesn't do it very quickly. We need to optimize it.

The task to work with the pprof profiler.

The purpose of the task is to learn how to work with pprof, find hot spots in the code, be able to build a cpu and memory consumption profile, optimize the code taking into account this information. Writing the fastest solution is not the goal of the assignment.

To generate a graph, you will need graphviz. For windows users, don't forget to add it to PATH so that the dot command is available.

I recommend that you read carefully. materials in Russian - there are many more optimization examples and explanations on how to work with the profiler. In fact, there is all the information to complete this task.

There are dozens of places where you can optimize.

To complete the task, one of the parameters ( ns/op, B/op, allocs/op ) must be faster than in BenchmarkSolution ( fast < solution ) and one more better than BenchmarkSolution + 20% ( fast < solution * 1.2), for example ( fast allocs/op < 10422*1.2=12506 ).

## My optimisation result
| function | count | ns/op | B/op | allocs/op |
| ---: | ---: | ---: | ---: | ---: |
| BenchmarkSlow-4       |         43    |       28169590 ns/op    |     18850657 B/op   |   195818 allocs/op  | 
| BenchmarkFast-4       |       490     |       2426283 ns/op     |     1741911 B/op    |    9589 allocs/op   | 
