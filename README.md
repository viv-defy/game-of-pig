# game-of-pig
This program simulates game-of-pig as defined in [one2n go-bootcamp](https://playbook.one2n.in/go-bootcamp/go-projects/a-game-of-pig/a-game-of-pig-exercise). 

**game-of-pig** is a two player turn-based game, played with a six-sided dice. This game involves each player rolling a dice, summing all the outcomes and adding it to the total from previous turns. But with two caveats
- on receiving 1, the player forfeits his turn without adding the current sum to his previous total
- the player forfeits his turn voluntarily

These caveats gives rise to various strategies that players adopt to win the game as they balance the risk(getting 1) and reward(maximising the turn total)

To understand the game in detail read [one2n go-bootcamp](https://playbook.one2n.in/go-bootcamp/go-projects/a-game-of-pig/a-game-of-pig-exercise).

## Assumptions
- **Each player plays to win**
- **The input range for strategy is restricted from 1 to 100**. An input less than 1 as the beginning of the range raises an error. An input over 100 as the end of the range is replaced with 100

## Functions and Types
The code contains a main file and a player file. 
### main.go
This file contains the main logic of the game
#### functions
- parseStrategies - parses the input arguments
- parseRange - parses the range from a string
- fixStrategies - loops through the ranges for all the valid player-holds
- simulateGames - simulates 10 games with given holds
- game - simulates a single game
- findMin - a helper function to find min of two functions
### player.go
- player - a struct with variables
    - id
    - hold
    - score

    and methods
    - turn - play till 1, hold or win
    - reset - set score to 0

## Running the code
- Clone the project
    ```
    git clone git@github.com:viv-defy/game-of-pig.git
    cd game-of-pig
    ```
- Build the project
    ```
    make build
    ```
- Run the code by entering two strategies
    ```
    ./main 20 25
    ```
    Strategies can either be a range `start-end` or `hold` 
    For simulating the stories given in the go-bootcamp, run the following

    Story 1
    ```
    ./main 10 15
    ```
    Story 2
    ```
    ./main 21 1-100
    ```
    Story 3
    ```
    ./main 1-100 1-100
    ```
