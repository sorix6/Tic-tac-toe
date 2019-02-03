# Tic-Tac-Toe Game

The purpose of this project is to allow users to play a game of Tic-Tac-Toe.
The project only consists of an API, no interface is available for now.

### Technologies and implementation

The project has bee written in Go and uses a file to store the state of the game.
No database has been used

There are 3 exposed end-points:
- /{gameId}/reset - allows the users to reset the game
- /{gameId}/add/{player} - where player can be either X or O, allows users to add X or O to the game
- /{gameId}/get-status - returns the status of the game

The gameId is a UUID V4. The idea behind this choice was to be able to allow multiple games to be played at the same time, but this feature has not yet been implemented. The only available game for the moment has the ID **7d38148c-6526-4cd7-9b21-56b498b93b12**

The state of the game is stored in a 3 X 3 matrix.

The first player to get 3 of her marks in a row (up, down, across, or diagonally) is the winner.
When all 9 squares are full, the game is over. If no player has 3 marks in a row, the game ends in a tie.

### Playing the game

In order to play the game, the **./tic-tac-toe** file has to be executed.

Since the end-points for resetting and playing use the CRUD methods DELETE and POST, respectively, it is recommended to use an API development environment, like Postman

* To get the state of the game:
  - end-point: /7d38148c-6526-4cd7-9b21-56b498b93b12/get-status
  - method: GET

* To play the game:
    - As O:
        - end-point: 7d38148c-6526-4cd7-9b21-56b498b93b12/add/O
    - AS X:
        - end-point: 7d38148c-6526-4cd7-9b21-56b498b93b12/add/X

    - method: POST
    - payload: {"row":"1", "column":"1"} - the payload must be updated with the coordinates of the desired position, in the 3 X 3 matrix, as explained below

    | R: 0, C: 0 | R: 0, C: 1 | R: 0, C: 2 |
    | --- | --- | --- |
    | R: 1, C: 0 | R: 1, C: 1 | R: 1, C: 2 |
    | --- | --- | --- |
    | R: 2, C: 0 | R: 2, C: 1 | R: 2, C: 2 |

* To reset the game:
    - end-point: 7d38148c-6526-4cd7-9b21-56b498b93b12/reset
    - method: DELETE