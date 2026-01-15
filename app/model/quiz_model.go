package model

type QuizContent struct {
	Id          int      `json:"id"`
	Question    string   `json:"question"`
	Options     []Option `json:"options"`
	Answer      int      `json:"answer"`
	Explanation string   `json:"explanation"`
}

type Option struct {
	Id   int    `json:"id"`
	Text string `json:"text"`
}

type Quiz struct {
	ID      string      `json:"id"`
	Content QuizContent `json:"content"`
}
