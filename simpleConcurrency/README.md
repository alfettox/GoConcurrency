# Easy Go Concurrency Example

This Go program demonstrates how to use concurrency to make your code faster by allowing multiple threads (goroutines) to perform tasks simultaneously.

## How It Works

1. **Initialization**: We start by initializing our threads (goroutines) and set up a `sync.WaitGroup` to keep track of their completion.

2. **Communication Channels**: We establish two channels for communication among the threads. These channels allow threads to send and receive messages.

3. **Thread Jobs**: There are two types of threads, each with specific tasks that involve performing calculations. After completing their tasks, they send the results through channels for sharing.

4. **Synchronization**: We ensure that we wait for all threads to complete their work before proceeding to the next step.

5. **Aggregating Results**: Finally, we collect and aggregate all the results from the threads, calculating the final outcome.

## Execution

To run the program, use the following command in your terminal:

```shell
go run main.go