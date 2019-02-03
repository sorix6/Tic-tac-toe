package controllers

import (
	"strconv"
	"encoding/json"
	"fmt"
    "io/ioutil"
    "os"
	"net/http"
	"github.com/gorilla/mux"
	"github/sorix6/tic-tac-toe/structures"
)

func Reset(w http.ResponseWriter, r *http.Request){
	resetGame()

	params := mux.Vars(r)
	
	var game structures.Game = readFile(params["gameId"])

	json.NewEncoder(w).Encode(game)
}

func AddPlay(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)

	if (params["player"] != "O" && params["player"] != "X"){
		json.NewEncoder(w).Encode(structures.Message{Msg: "You need and X and an O for this game"})
	}
	
	decoder := json.NewDecoder(r.Body)
	
	var payload structures.Payload
	err := decoder.Decode(&payload)
	if err != nil {
		panic(err)
	}

	row := payload.Row
	column := payload.Column


	var game structures.Game = readFile(params["gameId"])

	if game.Status == "closed" {
		json.NewEncoder(w).Encode(structures.Message{Msg: "This game has ended"})
	} else {
		if game.LastMove != params["player"] {
			
			i, err := strconv.Atoi(row)
			j, err := strconv.Atoi(column)
			check(err)
			
			if (i > 2 || i < 0 || j > 2 || j < 0 || game.GameTable[i][j] != 0) {
				json.NewEncoder(w).Encode(structures.Message{Msg: "Illegal move"})

			} else {

				if params["player"] == "O" {
					game.GameTable[i][j] =  -1
				} else {
					game.GameTable[i][j] =  1
				}
				
				game.LastMove = params["player"]

				result := checkIfGameOver(game.GameTable, params["player"], game.TotalMovesMade)

				if result != "" {
					game.Winner = result
					game.Status = "closed"

					saveData(game)

					json.NewEncoder(w).Encode(structures.Message{Msg: result + " wins!"})
				} else {
					saveData(game)
					json.NewEncoder(w).Encode(game)
				}

				
			} 
			
		} else {
			json.NewEncoder(w).Encode(structures.Message{Msg: "Wait for your turn"})
		}
	}	
	
}

func GetStatus(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	
	var game structures.Game = readFile(params["gameId"])

	json.NewEncoder(w).Encode(game)
}



// helper functions


func check(e error) {
    if e != nil {
        panic(e)
    }
}

func saveData(game structures.Game) {
	jsonModel := structures.JsonModel{ID: game.ID, 
		M00: fmt.Sprint(game.GameTable[0][0]), M01: fmt.Sprint(game.GameTable[0][1]), M02: fmt.Sprint(game.GameTable[0][2]),
	    M10: fmt.Sprint(game.GameTable[1][0]), M11: fmt.Sprint(game.GameTable[1][1]), M12: fmt.Sprint(game.GameTable[1][2]),
		M20: fmt.Sprint(game.GameTable[2][0]), M21: fmt.Sprint(game.GameTable[2][1]), M22: fmt.Sprint(game.GameTable[2][2]),
		LastMove: game.LastMove,
		Status: game.Status,
		Winner: game.Winner}

	fileContent, err := json.Marshal(jsonModel)
	check(err)
   
	
	file, err := os.OpenFile("game_data/game_" + game.ID, os.O_WRONLY|os.O_CREATE, 0666)
    if err != nil {
        fmt.Println("File does not exists or cannot be created")
        os.Exit(1)
    }
	defer file.Close()
	
	file.Truncate(0)
    
    err = ioutil.WriteFile("game_data/game_" + game.ID, fileContent, 0644)

}

func resetGame() {
	
	jsonModel := structures.JsonModel{ID: "7d38148c-6526-4cd7-9b21-56b498b93b12", 
		M00: "0", M01: "0", M02: "0",
		M10: "0", M11: "0", M12: "0",
		M20: "0", M21: "0", M22: "0",
		LastMove: "",
		Status: "",
		Winner: ""}

	fileContent, err := json.Marshal(jsonModel)
	check(err)
	
	
	file, err := os.OpenFile("game_data/game_7d38148c-6526-4cd7-9b21-56b498b93b12", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("File does not exists or cannot be created")
		os.Exit(1)
	}
	defer file.Close()
	
	file.Truncate(0)
	
	err = ioutil.WriteFile("game_data/game_7d38148c-6526-4cd7-9b21-56b498b93b12", fileContent, 0644)
	
	
}

func checkIfGameOver(table [3][3]int, player string, totalMovesMade int) string{

	// the game is over if any line has the same value
	
	var diagonal1 = table[0][0] + table[1][1] + table[2][2]
	var diagonal2 = table[0][2] + table[1][1] + table[2][0]

	var line1 = table[0][0] + table[0][1] + table[0][2]
	var line2 = table[1][0] + table[1][1] + table[1][2]
	var line3 = table[2][0] + table[2][1] + table[2][2]

	var column1 = table[0][0] + table[1][0] + table[2][0]
	var column2 = table[0][1] + table[1][1] + table[2][1]
	var column3 = table[0][2] + table[1][2] + table[2][2]

	var result = "";

	if (diagonal1 == 3 || diagonal2 == 3 || line1 == 3 || line2 == 3 || line3 == 3 || column1 == 3 || column2 == 3 || column3 == 3){
		if player == "X" {
			result =  "X"
		}		

	} else if (diagonal1 == -3 || diagonal2 == -3 || line1 == -3 || line2 == -3 || line3 == -3 || column1 == -3 || column2 == -3 || column3 == -3){
		if player == "O" {
			result =  "O"
		} 
	} else if (totalMovesMade == 8){
		result = "Draw"
	}

	return result
}

func readFile(gameID string) structures.Game {
	dat, err := ioutil.ReadFile("game_data/game_" + gameID)
	data := []byte(dat)

	var dat1 map[string]interface{}

	if err := json.Unmarshal(data, &dat1); err != nil {
        panic(err)
	}	

	check(err)
	
	game := structures.Game{ID: dat1["ID"].(string), LastMove: dat1["LastMove"].(string), Status: dat1["Status"].(string), Winner: dat1["Winner"].(string)}
	totalMovesMade  := 0

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			//
			value := dat1[fmt.Sprintf("%s%d%d", "M", i, j)].(string)
			s, err := strconv.Atoi(value); 

			if s != 0 {
				totalMovesMade++
			}

			check(err)
			game.GameTable[i][j] = s
		}
	}

	game.TotalMovesMade = totalMovesMade

	return game
}