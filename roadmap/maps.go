package roadmap

import "fmt"

func maps() {
	var studentScores = make(map[string]int)
	studentScores["Jada"] = 50
	studentScores["Chris"] = 60
	studentScores["Will"] = 70

	fmt.Println(studentScores)
}
