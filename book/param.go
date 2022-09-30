package book

type BookInput struct {
	Title    string `json:"title" binding:"required"`
	Price    int    `json:"price" binding:"required"`
	SubTitle string `json:"sub_title"`
}
