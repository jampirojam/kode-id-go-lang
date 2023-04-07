package model

type ElementValue struct {
	Wind  	int 	`json:"wind"`
	Water 	int 	`json:"water"`
}

type PostStatus struct {
	UserId int    	`json:"userId"`
	Id     int    	`json:"id"`
	Title  string 	`json:"title"`
	Body   string 	`json:"body"`
}
