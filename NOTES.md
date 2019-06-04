# NOTES:

### Potential Use Case:

1. A single task does all the reads
2. It then creates and returns a new task with the results of those reads under the `Next()` call.

The idea behind the above is that the first task is a "read task" while a subsequent task created through `Next()` is the "write task".
