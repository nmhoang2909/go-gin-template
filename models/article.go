package models

type Article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var articles = []Article{
	{
		ID:      1,
		Title:   "Title 1",
		Content: "Content 1",
	},
	{
		ID:      2,
		Title:   "Title 2",
		Content: "Content 2",
	},
}

func GetArticles() []Article {
	return articles
}

func GetArticleByID(id int) Article {
	for _, a := range articles {
		if a.ID == id {
			return a
		}
	}
	return Article{}
}
