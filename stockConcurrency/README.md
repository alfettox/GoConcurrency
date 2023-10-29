# Concurrency stock example

The program simulates two financial institutions, Morgan Stanley and NBC, managing and tracking stock options for two different companies, Ericsson and Huawei. The task involves tracking the stock options and leaving messages about stock purchases.

There are a few issues in the original program:

1. **Deadlock Issue**: The original program faces a deadlock issue. It occurs when the program waits for the completion of two goroutines while also trying to read from a channel. Since the channel is not closed, the program becomes stuck in a deadlock.

2. **Message Handling**: Messages left by the financial institutions are not properly handled. The program should receive and display these messages to ensure proper communication.

## Code

The program consists of three main goroutines:

1. **Morgan Stanley Goroutine**: Simulates Morgan Stanley's activity in managing Ericsson stock options. It uses a mutex for synchronization.

2. **NBC Goroutine**: Simulates NBC's activity in managing Huawei stock options. It also uses a mutex for synchronization.

3. **Channel Reader Goroutine**: A dedicated goroutine that reads messages from a channel and prints them.
This goroutine ensures that messages are properly displayed.

The main functions in the program include:

- `trackStockOptions`: This function simulates the stock option tracking process for a financial institution.
It uses a mutex to protect the shared stock option variable.

- `leaveNote`: This function is used to leave messages about stock purchases. It sends messages to a channel.

- `main`: The main function initializes the program, creates goroutines, waits for their completion, and handles message reading.

## Solutions

To address the issues in the original program, the following solutions are implemented:

1. **Deadlock Issue**: A dedicated goroutine is added to read and print messages from the channel. Additionally, the `leaveNote` function now closes the channel after sending all messages. This ensures that the channel reader can complete its task and prevents a deadlock.

2. **Message Handling**: The channel reader goroutine is responsible for reading and displaying messages, ensuring effective communication.

## Running the Program

To execute the program, you can use the following command:

```bash
go run concurrency.go