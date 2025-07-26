package server

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
)

type Article struct {
	Title   string
	Snippet string
	Body    string
}
type Review struct {
	Title   string
	Snippet string
	Score   int
	Body    string
}

type IndexData struct {
	Reviews  []Review
	Articles []Article
}

var reviews = []Review{
	{
		Title:   "Breath of the Wild is Amazing",
		Snippet: "This game is the bomb",
		Score:   85,
		Body: `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Integer venenatis neque at odio ultricies porttitor eu at nisl. In nulla ex, molestie ac rutrum sed, auctor at justo. Sed ut vehicula dui. Phasellus sollicitudin blandit ex at molestie. Pellentesque malesuada ipsum leo, et placerat felis rutrum ac. Vivamus fermentum laoreet luctus. Morbi fringilla ante sed molestie pretium. Cras vehicula dolor sed lorem vehicula, quis varius tellus laoreet. Nulla a dignissim augue. Suspendisse pulvinar pulvinar velit eget bibendum. Mauris suscipit enim et luctus pulvinar. Aenean id rutrum massa, non finibus ipsum.

Aliquam metus neque, tincidunt et nulla sit amet, dignissim sodales turpis. Nullam vel pretium felis. Phasellus in pharetra metus. Duis feugiat elit vel diam sollicitudin rhoncus. Quisque non condimentum metus. Mauris sollicitudin varius ligula, id pellentesque orci sagittis ut. Suspendisse tempor vitae erat mattis finibus. Maecenas pellentesque, ex ac pulvinar auctor, mi libero sagittis neque, at convallis mi erat ac lectus. Pellentesque tincidunt imperdiet nunc, eget vulputate felis sollicitudin ac. Vivamus porttitor tincidunt rutrum.

Nulla nec augue eget justo finibus tristique. Morbi ac est massa. Donec pulvinar, odio eu condimentum condimentum, eros nisl malesuada nisi, nec bibendum augue nisi vitae mauris. In pulvinar dignissim vehicula. Sed at urna dictum, molestie mauris quis, ultrices velit. Maecenas auctor rhoncus nulla quis facilisis. Suspendisse potenti. Phasellus augue tortor, mattis in diam ac, ultrices dignissim nisi. Nullam tincidunt interdum tempor.

Ut auctor orci molestie, pellentesque mi eu, malesuada dolor. Phasellus tempor quis purus eget ornare. Pellentesque bibendum tempor fringilla. Integer purus neque, sodales quis efficitur sed, elementum vel lectus. Donec eget nisl tempus, viverra neque eu, rutrum metus. Duis mattis placerat sapien, sed pulvinar diam porta vitae. Nam posuere facilisis nunc, id gravida nibh pretium faucibus.

Pellentesque dapibus ante quis ipsum tincidunt, in vulputate massa consectetur. Fusce elementum urna malesuada neque facilisis scelerisque. Nullam rhoncus dui vel lacus volutpat fermentum. Cras rutrum ut nunc in hendrerit. Phasellus dignissim nibh et lacus varius, id eleifend ipsum consectetur. Donec laoreet felis erat, vulputate ullamcorper nisl interdum vel. In ullamcorper turpis in ipsum fringilla sodales. Fusce iaculis diam non erat vestibulum eleifend. Etiam viverra faucibus euismod.

`,
	},
	{
		Title:   "Breath of the Wild 2 is even MORE Amazing",
		Snippet: "This game is a BIGGER bomb",
		Score:   90,
		Body: `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Integer venenatis neque at odio ultricies porttitor eu at nisl. In nulla ex, molestie ac rutrum sed, auctor at justo. Sed ut vehicula dui. Phasellus sollicitudin blandit ex at molestie. Pellentesque malesuada ipsum leo, et placerat felis rutrum ac. Vivamus fermentum laoreet luctus. Morbi fringilla ante sed molestie pretium. Cras vehicula dolor sed lorem vehicula, quis varius tellus laoreet. Nulla a dignissim augue. Suspendisse pulvinar pulvinar velit eget bibendum. Mauris suscipit enim et luctus pulvinar. Aenean id rutrum massa, non finibus ipsum.

Aliquam metus neque, tincidunt et nulla sit amet, dignissim sodales turpis. Nullam vel pretium felis. Phasellus in pharetra metus. Duis feugiat elit vel diam sollicitudin rhoncus. Quisque non condimentum metus. Mauris sollicitudin varius ligula, id pellentesque orci sagittis ut. Suspendisse tempor vitae erat mattis finibus. Maecenas pellentesque, ex ac pulvinar auctor, mi libero sagittis neque, at convallis mi erat ac lectus. Pellentesque tincidunt imperdiet nunc, eget vulputate felis sollicitudin ac. Vivamus porttitor tincidunt rutrum.

Nulla nec augue eget justo finibus tristique. Morbi ac est massa. Donec pulvinar, odio eu condimentum condimentum, eros nisl malesuada nisi, nec bibendum augue nisi vitae mauris. In pulvinar dignissim vehicula. Sed at urna dictum, molestie mauris quis, ultrices velit. Maecenas auctor rhoncus nulla quis facilisis. Suspendisse potenti. Phasellus augue tortor, mattis in diam ac, ultrices dignissim nisi. Nullam tincidunt interdum tempor.

Ut auctor orci molestie, pellentesque mi eu, malesuada dolor. Phasellus tempor quis purus eget ornare. Pellentesque bibendum tempor fringilla. Integer purus neque, sodales quis efficitur sed, elementum vel lectus. Donec eget nisl tempus, viverra neque eu, rutrum metus. Duis mattis placerat sapien, sed pulvinar diam porta vitae. Nam posuere facilisis nunc, id gravida nibh pretium faucibus.

Pellentesque dapibus ante quis ipsum tincidunt, in vulputate massa consectetur. Fusce elementum urna malesuada neque facilisis scelerisque. Nullam rhoncus dui vel lacus volutpat fermentum. Cras rutrum ut nunc in hendrerit. Phasellus dignissim nibh et lacus varius, id eleifend ipsum consectetur. Donec laoreet felis erat, vulputate ullamcorper nisl interdum vel. In ullamcorper turpis in ipsum fringilla sodales. Fusce iaculis diam non erat vestibulum eleifend. Etiam viverra faucibus euismod.

`,
	},
}

var articles = []Article{
	{
		Title:   "I loved the Nintendo Direct",
		Snippet: "It made me more happy than Hello Kitty",
		Body: `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Integer venenatis neque at odio ultricies porttitor eu at nisl. In nulla ex, molestie ac rutrum sed, auctor at justo. Sed ut vehicula dui. Phasellus sollicitudin blandit ex at molestie. Pellentesque malesuada ipsum leo, et placerat felis rutrum ac. Vivamus fermentum laoreet luctus. Morbi fringilla ante sed molestie pretium. Cras vehicula dolor sed lorem vehicula, quis varius tellus laoreet. Nulla a dignissim augue. Suspendisse pulvinar pulvinar velit eget bibendum. Mauris suscipit enim et luctus pulvinar. Aenean id rutrum massa, non finibus ipsum.

Aliquam metus neque, tincidunt et nulla sit amet, dignissim sodales turpis. Nullam vel pretium felis. Phasellus in pharetra metus. Duis feugiat elit vel diam sollicitudin rhoncus. Quisque non condimentum metus. Mauris sollicitudin varius ligula, id pellentesque orci sagittis ut. Suspendisse tempor vitae erat mattis finibus. Maecenas pellentesque, ex ac pulvinar auctor, mi libero sagittis neque, at convallis mi erat ac lectus. Pellentesque tincidunt imperdiet nunc, eget vulputate felis sollicitudin ac. Vivamus porttitor tincidunt rutrum.

Nulla nec augue eget justo finibus tristique. Morbi ac est massa. Donec pulvinar, odio eu condimentum condimentum, eros nisl malesuada nisi, nec bibendum augue nisi vitae mauris. In pulvinar dignissim vehicula. Sed at urna dictum, molestie mauris quis, ultrices velit. Maecenas auctor rhoncus nulla quis facilisis. Suspendisse potenti. Phasellus augue tortor, mattis in diam ac, ultrices dignissim nisi. Nullam tincidunt interdum tempor.

Ut auctor orci molestie, pellentesque mi eu, malesuada dolor. Phasellus tempor quis purus eget ornare. Pellentesque bibendum tempor fringilla. Integer purus neque, sodales quis efficitur sed, elementum vel lectus. Donec eget nisl tempus, viverra neque eu, rutrum metus. Duis mattis placerat sapien, sed pulvinar diam porta vitae. Nam posuere facilisis nunc, id gravida nibh pretium faucibus.

Pellentesque dapibus ante quis ipsum tincidunt, in vulputate massa consectetur. Fusce elementum urna malesuada neque facilisis scelerisque. Nullam rhoncus dui vel lacus volutpat fermentum. Cras rutrum ut nunc in hendrerit. Phasellus dignissim nibh et lacus varius, id eleifend ipsum consectetur. Donec laoreet felis erat, vulputate ullamcorper nisl interdum vel. In ullamcorper turpis in ipsum fringilla sodales. Fusce iaculis diam non erat vestibulum eleifend. Etiam viverra faucibus euismod.

`,
	},
	{
		Title:   "What I think about the Switch 2",
		Snippet: "Why is it so effing expensive",
		Body: `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Integer venenatis neque at odio ultricies porttitor eu at nisl. In nulla ex, molestie ac rutrum sed, auctor at justo. Sed ut vehicula dui. Phasellus sollicitudin blandit ex at molestie. Pellentesque malesuada ipsum leo, et placerat felis rutrum ac. Vivamus fermentum laoreet luctus. Morbi fringilla ante sed molestie pretium. Cras vehicula dolor sed lorem vehicula, quis varius tellus laoreet. Nulla a dignissim augue. Suspendisse pulvinar pulvinar velit eget bibendum. Mauris suscipit enim et luctus pulvinar. Aenean id rutrum massa, non finibus ipsum.

Aliquam metus neque, tincidunt et nulla sit amet, dignissim sodales turpis. Nullam vel pretium felis. Phasellus in pharetra metus. Duis feugiat elit vel diam sollicitudin rhoncus. Quisque non condimentum metus. Mauris sollicitudin varius ligula, id pellentesque orci sagittis ut. Suspendisse tempor vitae erat mattis finibus. Maecenas pellentesque, ex ac pulvinar auctor, mi libero sagittis neque, at convallis mi erat ac lectus. Pellentesque tincidunt imperdiet nunc, eget vulputate felis sollicitudin ac. Vivamus porttitor tincidunt rutrum.

Nulla nec augue eget justo finibus tristique. Morbi ac est massa. Donec pulvinar, odio eu condimentum condimentum, eros nisl malesuada nisi, nec bibendum augue nisi vitae mauris. In pulvinar dignissim vehicula. Sed at urna dictum, molestie mauris quis, ultrices velit. Maecenas auctor rhoncus nulla quis facilisis. Suspendisse potenti. Phasellus augue tortor, mattis in diam ac, ultrices dignissim nisi. Nullam tincidunt interdum tempor.

Ut auctor orci molestie, pellentesque mi eu, malesuada dolor. Phasellus tempor quis purus eget ornare. Pellentesque bibendum tempor fringilla. Integer purus neque, sodales quis efficitur sed, elementum vel lectus. Donec eget nisl tempus, viverra neque eu, rutrum metus. Duis mattis placerat sapien, sed pulvinar diam porta vitae. Nam posuere facilisis nunc, id gravida nibh pretium faucibus.

Pellentesque dapibus ante quis ipsum tincidunt, in vulputate massa consectetur. Fusce elementum urna malesuada neque facilisis scelerisque. Nullam rhoncus dui vel lacus volutpat fermentum. Cras rutrum ut nunc in hendrerit. Phasellus dignissim nibh et lacus varius, id eleifend ipsum consectetur. Donec laoreet felis erat, vulputate ullamcorper nisl interdum vel. In ullamcorper turpis in ipsum fringilla sodales. Fusce iaculis diam non erat vestibulum eleifend. Etiam viverra faucibus euismod.

`,
	},
}

var indexInput = IndexData{
	Reviews:  reviews,
	Articles: articles,
}

func index(w http.ResponseWriter, r *http.Request) {
	base := path.Join("html", "base.html")
	index := path.Join("html", "index.html")

	t, _ := template.ParseFiles(base, index)
	t.Execute(w, indexInput)
}

func review(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	base := path.Join("html", "base.html")
	review := path.Join("html", "review.html")

	fmt.Printf("You are looking for a review with the id of: %s", id)

	t, _ := template.ParseFiles(base, review)
	t.Execute(w, reviews[0])
}

func article(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	base := path.Join("html", "base.html")
	review := path.Join("html", "article.html")

	fmt.Printf("You are looking for an article with the id of: %s", id)

	t, _ := template.ParseFiles(base, review)
	t.Execute(w, articles[0])
}

func StartServer() {
	// Removed http for mux to handle path params easier
	mux := http.NewServeMux()

	mux.HandleFunc("/", index)
	mux.HandleFunc("/review/{id}", review)
	mux.HandleFunc("/article/{id}", article)

	log.Println("Starting the server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
