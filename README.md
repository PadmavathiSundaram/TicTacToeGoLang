#TicTacToe:
-----------

A Go Lang version of Tic Tac Toe 2 player game.
This app aims at achieving 100% testable code.

The concept of interface is used to achieve this test coverage.

#Setup
---------

The project can be cloned into any wprkspace as it uses go modules.
Just clone the project and use the make command to get started.

// Setting up the terminal to support go mod
export GOPath=$HOME/go
export GO111MODULE=on

Installing Golint:
---------
GO111MODULE=on go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.16.0

#To run:
----------

run the below command in terminal from the root directory:

make local

This will:
 
 1. Perform the lint checks
 2. Build the project
 3. Run the tests
 4. Generate the coverage reports
 5. Run the Application
 
 #About The Game:
 ---------------
 
 https://en.wikipedia.org/wiki/Tic-tac-toe
 
 #To run tests:
 -----------
 
 make test
 
 Go version used - go1.11.5 
 
 
 
 
 
 