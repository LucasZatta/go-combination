package server

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type Score struct {
	Score string `json:"score"`
}

func (s *Server) RegisterRoutes() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/verify", s.VeirfyHandler).Methods(http.MethodPost)

	return r
}

func (s *Server) VeirfyHandler(w http.ResponseWriter, r *http.Request) {

	var score Score

	err := json.NewDecoder(r.Body).Decode(&score)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	teamScores, err := teamScore(score.Score)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	possibleCombinations := possibleCombinations(teamScores)

	resp := make(map[string]string)
	resp["combinations"] = strconv.Itoa(possibleCombinations)

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	_, _ = w.Write(jsonResp)
}

func teamScore(s string) (teamScores []int, err error) {
	teamScores = make([]int, 2)
	scores := strings.Split(s, "x")

	for index, scoreString := range scores {
		scoreInt, err := strconv.Atoi(scoreString)
		if err != nil {
			return teamScores, err
		}
		teamScores[index] = scoreInt
	}

	return

}

func possibleCombinations(scores []int) int {
	possiblePoints := []int{3, 6, 7, 8} //fg, td, td+1 and td+2
	return verifyScore(possiblePoints, len(possiblePoints), scores[0]) * verifyScore(possiblePoints, len(possiblePoints), scores[1])
}

func verifyScore(points []int, n int, sum int) int {
	if sum == 0 {
		return 1
	}
	if sum < 0 {
		return 0
	}

	if n <= 0 {
		return 0
	}

	return verifyScore(points, n-1, sum) + verifyScore(points, n, sum-points[n-1])
}
