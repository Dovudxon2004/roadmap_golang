package roadmap

type Student struct {
	name  string
	age   int
	score []int
}

var Tolib = Student{"Tolib", 20, []int{21, 23, 54, 23, 34, 46, 25}}

func (s Student) getAverageScore() float32 {

	sum := 0
	for _, v := range s.score {
		sum += v
	}
	return float32(sum) / float32(len(s.score))
}

func GetAverageScore1() float32 {
	n := Tolib.getAverageScore()
	return n
}
