package main

import (
	"net/http"
	"os"

	gintrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gin-gonic/gin"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"

	"github.com/gin-gonic/gin"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type article struct {
	ID          string  `json:"id"`
	Label       string  `json:"label"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Cover       string  `json:"cover"`
}

var articles = []article{
	{
		ID:          "64d6597fc8f26b6a23b34e65",
		Label:       "Ut dolore nisi aliqua consectetur ullamco id ut officia et adipisicing eiusmod ut.",
		Description: "Pariatur velit officia culpa aliqua velit. Sit ex dolor mollit elit nulla cupidatat tempor eu.",
		Price:       400,
		Cover:       "https://images.unsplash.com/photo-1531722569936-825d3dd91b15?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=1740&q=80",
	},
	{
		ID:          "64d6597fb41c69997515eb26",
		Label:       "Non enim nisi aliquip adipisicing.",
		Description: "Adipisicing exercitation sunt velit Lorem sunt adipisicing ex sint veniam non ipsum ullamco aliqua eu. Non pariatur eu ipsum cupidatat voluptate duis aliquip anim nulla.",
		Price:       1,
		Cover:       "https://images.unsplash.com/photo-1531722569936-825d3dd91b15?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=1740&q=80",
	},
	{
		ID:          "64d6597f483bb7578e5495dc",
		Label:       "In eiusmod id ullamco reprehenderit nulla nulla esse ex non deserunt qui.",
		Description: "Deserunt id do tempor duis dolore exercitation cupidatat ut. Voluptate fugiat laborum minim proident commodo amet excepteur laborum eiusmod culpa officia duis pariatur.",
		Price:       8,
		Cover:       "https://images.unsplash.com/photo-1531722569936-825d3dd91b15?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=1740&q=80",
	},
	{
		ID:          "64d6597ff9c91e3c96fb20b5",
		Label:       "Incididunt esse sit non cupidatat reprehenderit enim aliqua dolore enim culpa est ad qui.",
		Description: "Ut adipisicing elit velit aute. Consequat id labore culpa ipsum ullamco proident.",
		Price:       0,
		Cover:       "https://images.unsplash.com/photo-1531722569936-825d3dd91b15?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=1740&q=80",
	},
	{
		ID:          "64d6597f2bd59b583e2e995c",
		Label:       "Dolore deserunt nostrud est mollit est laboris exercitation anim occaecat velit consequat.",
		Description: "Minim aute proident sunt nulla laboris cupidatat quis quis dolor. Incididunt officia officia qui excepteur velit.",
		Price:       7,
		Cover:       "https://images.unsplash.com/photo-1531722569936-825d3dd91b15?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=1740&q=80",
	},
	{
		ID:          "64d6597fe46f352f100f2233",
		Label:       "Veniam mollit consequat ut veniam.",
		Description: "Laborum sunt deserunt quis esse reprehenderit. Nisi adipisicing velit elit ullamco sit Lorem non nisi laboris eu laborum adipisicing sint duis.",
		Price:       1,
		Cover:       "https://images.unsplash.com/photo-1531722569936-825d3dd91b15?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=1740&q=80",
	},
	{
		ID:          "64d6597f971ffba4c42d84c4",
		Label:       "Nisi aute qui amet excepteur do fugiat ut sit culpa consequat laborum excepteur ullamco.",
		Description: "Ullamco esse qui ipsum culpa voluptate cupidatat labore dolor adipisicing Lorem fugiat. Anim ut dolor commodo aliqua cupidatat fugiat adipisicing dolore aliqua enim qui.",
		Price:       5,
		Cover:       "https://images.unsplash.com/photo-1531722569936-825d3dd91b15?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=1740&q=80",
	},
	{
		ID:          "64d6597fdc3239d1bc06692e",
		Label:       "Reprehenderit in ea in id Lorem Lorem proident occaecat dolore eu do.",
		Description: "Ex laborum aliqua ad in ut sint aute eu. Culpa elit nostrud labore ut consectetur aliqua sunt cupidatat officia duis nostrud.",
		Price:       2,
		Cover:       "https://images.unsplash.com/photo-1531722569936-825d3dd91b15?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=1740&q=80",
	},
	{
		ID:          "64d6597fb26399e9e83b476a",
		Label:       "Officia tempor voluptate anim laboris mollit anim ea Lorem ipsum eu anim.",
		Description: "Do consequat eu et cupidatat dolore fugiat reprehenderit aute ea consectetur dolor fugiat ipsum dolor. Sit sunt ea anim nulla amet et ex dolor elit tempor.",
		Price:       10,
		Cover:       "https://images.unsplash.com/photo-1531722569936-825d3dd91b15?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=1740&q=80",
	},
	{
		ID:          "64d6597f793ccb3d0a5e8d49",
		Label:       "Consectetur veniam non minim deserunt est mollit.",
		Description: "Sit sint aute laboris occaecat ad eu id incididunt. Tempor nisi ipsum sint sunt sit ut voluptate et consequat laborum in ipsum culpa elit.",
		Price:       6,
		Cover:       "https://images.unsplash.com/photo-1531722569936-825d3dd91b15?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=1740&q=80",
	},
	{
		ID:          "64d6597f9583469b6382889d",
		Label:       "Duis tempor qui elit aliquip aliquip.",
		Description: "Amet in laborum nostrud quis. Est consequat dolore nulla cillum duis aliqua cillum.",
		Price:       0,
		Cover:       "https://images.unsplash.com/photo-1531722569936-825d3dd91b15?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=1740&q=80",
	},
	{
		ID:          "64d6597f791f7135edf6ca35",
		Label:       "Anim consequat amet laboris id duis officia quis esse eiusmod duis consequat in est.",
		Description: "Velit duis velit do in. Exercitation aliqua Lorem ea veniam excepteur ad sint Lorem in.",
		Price:       4,
		Cover:       "https://images.unsplash.com/photo-1531722569936-825d3dd91b15?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=1740&q=80",
	},
	{
		ID:          "64d6597f4a1e42bc5a7f9f46",
		Label:       "Fugiat ea reprehenderit cillum sunt deserunt qui nulla.",
		Description: "Nulla ullamco amet eiusmod cillum non ullamco eiusmod proident laborum elit. Aute ex officia occaecat velit nisi do officia sint.",
		Price:       9,
		Cover:       "https://images.unsplash.com/photo-1531722569936-825d3dd91b15?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=1740&q=80",
	},
	{
		ID:          "64d6597f91919b83b221872c",
		Label:       "Duis velit sint sit minim sunt labore quis aute.",
		Description: "Nisi sit laborum laboris enim ex aliqua. Quis deserunt pariatur anim qui id eiusmod nulla elit deserunt do.",
		Price:       8,
		Cover:       "https://images.unsplash.com/photo-1531722569936-825d3dd91b15?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=1740&q=80",
	},
	{
		ID:          "64d6597fe28669bd5032ac2e",
		Label:       "Aliquip occaecat laborum incididunt velit magna tempor.",
		Description: "Veniam cupidatat ea duis officia anim dolor cillum labore magna quis non reprehenderit esse. Tempor nisi cupidatat do proident esse esse consectetur laboris veniam ex.",
		Price:       0,
		Cover:       "https://images.unsplash.com/photo-1531722569936-825d3dd91b15?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=1740&q=80",
	},
	{
		ID:          "64d6597fe15e21425098c180",
		Label:       "In est voluptate adipisicing commodo minim.",
		Description: "Sunt Lorem elit proident velit occaecat anim sint cillum sint in aute voluptate. Ut excepteur irure nulla sunt amet aliqua consequat ullamco laborum.",
		Price:       2,
		Cover:       "https://images.unsplash.com/photo-1531722569936-825d3dd91b15?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=1740&q=80",
	},
	{
		ID:          "64d6597fb2df0b8e78d13c3c",
		Label:       "Elit ea proident incididunt est.",
		Description: "Dolor occaecat cillum ipsum do minim anim est duis id magna in. Nisi enim sint ad exercitation.",
		Price:       7,
		Cover:       "https://images.unsplash.com/photo-1531722569936-825d3dd91b15?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=1740&q=80",
	},
	{
		ID:          "64d6597f7df023f0651f21ba",
		Label:       "Elit ullamco aliqua aliquip ex ad esse sint enim duis ea laborum.",
		Description: "Velit eiusmod ad mollit enim id fugiat. Nostrud consectetur excepteur do Lorem consectetur laboris non.",
		Price:       7,
		Cover:       "https://images.unsplash.com/photo-1531722569936-825d3dd91b15?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=1740&q=80",
	},
	{
		ID:          "64d6597f29b7154a67f55d18",
		Label:       "Cupidatat aute proident esse culpa incididunt voluptate.",
		Description: "Consequat aliqua enim deserunt dolor ex ex magna et. Esse ullamco do do elit aliquip consectetur voluptate aute occaecat cupidatat exercitation irure.",
		Price:       10,
		Cover:       "https://images.unsplash.com/photo-1531722569936-825d3dd91b15?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=1740&q=80",
	},
	{
		ID:          "64d6597f59e455a237783dc6",
		Label:       "Sunt sit irure veniam in voluptate occaecat ad id.",
		Description: "Irure voluptate do dolore fugiat ullamco elit Lorem irure veniam veniam magna aute enim. Ipsum excepteur fugiat consectetur dolore officia culpa in est deserunt proident labore ullamco.",
		Price:       3,
		Cover:       "https://images.unsplash.com/photo-1531722569936-825d3dd91b15?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=1740&q=80",
	},
	{
		ID:          "64d6597f419f0231c110d29b",
		Label:       "Veniam consectetur officia esse deserunt nulla exercitation Lorem nostrud sit non incididunt sit ut.",
		Description: "Irure irure sint velit elit cupidatat enim do ex ipsum excepteur et eiusmod eiusmod magna. Sit sit laboris aute ullamco reprehenderit culpa in aliquip adipisicing occaecat.",
		Price:       1,
		Cover:       "https://images.unsplash.com/photo-1531722569936-825d3dd91b15?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=1740&q=80",
	},
	{
		ID:          "64d6597f1751af4a73b6b2c6",
		Label:       "Id aute reprehenderit laborum non sint nostrud in.",
		Description: "Anim culpa commodo ipsum minim ullamco officia dolore dolor culpa culpa culpa ex. Est irure ipsum fugiat quis quis.",
		Price:       4,
		Cover:       "https://images.unsplash.com/photo-1531722569936-825d3dd91b15?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=1740&q=80",
	},
	{
		ID:          "64d6597f3a3f5e064e3209bc",
		Label:       "Magna deserunt nisi anim et voluptate Lorem est sint mollit nulla.",
		Description: "Dolore id dolore irure magna sint fugiat cillum laboris magna quis laborum voluptate. Minim tempor consectetur duis enim proident excepteur.",
		Price:       1,
		Cover:       "https://images.unsplash.com/photo-1531722569936-825d3dd91b15?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=1740&q=80",
	},
	{
		ID:          "64d6597f5da5bad1d6d8cd25",
		Label:       "Aliqua ex ullamco sunt quis aliqua non sunt culpa adipisicing exercitation eu.",
		Description: "Qui tempor voluptate aute laborum duis laborum et. Ut culpa aute duis qui officia esse.",
		Price:       8,
		Cover:       "https://images.unsplash.com/photo-1531722569936-825d3dd91b15?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=1740&q=80",
	},
	{
		ID:          "64d6597f1232118e565b5da8",
		Label:       "Ipsum ad laborum mollit veniam dolor consequat laboris adipisicing ullamco in excepteur.",
		Description: "Aute sit ullamco excepteur sint pariatur Lorem fugiat do. Eu dolore consequat nisi cillum incididunt velit nostrud aliquip et enim laboris.",
		Price:       1,
		Cover:       "https://images.unsplash.com/photo-1531722569936-825d3dd91b15?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=1740&q=80",
	},
	{
		ID:          "64d6597f7307c933fedd0e63",
		Label:       "Commodo esse voluptate amet magna deserunt nisi aliquip non deserunt veniam esse ullamco.",
		Description: "Ad officia sint proident laboris velit ad amet. Voluptate exercitation aliquip proident laboris magna cillum consectetur labore tempor et non pariatur.",
		Price:       1,
		Cover:       "https://images.unsplash.com/photo-1531722569936-825d3dd91b15?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=1740&q=80",
	},
	{
		ID:          "64d6597fca90f37f22e7f47c",
		Label:       "Ullamco dolor deserunt aliqua elit labore aliquip anim excepteur voluptate minim.",
		Description: "Ad nostrud esse ullamco ut dolor eiusmod mollit. Lorem elit minim cillum nulla excepteur do elit.",
		Price:       0,
		Cover:       "https://images.unsplash.com/photo-1531722569936-825d3dd91b15?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=1740&q=80",
	},
	{
		ID:          "64d6597fa3719372700728df",
		Label:       "Proident labore sit ea velit officia fugiat.",
		Description: "Sint cillum excepteur quis mollit veniam elit est cillum. Excepteur commodo labore elit velit laboris consequat cillum irure minim sint voluptate quis consectetur.",
		Price:       5,
		Cover:       "https://images.unsplash.com/photo-1531722569936-825d3dd91b15?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=1740&q=80",
	},
	{
		ID:          "64d6597fb48441864ffa802f",
		Label:       "Est voluptate amet ipsum reprehenderit in est excepteur nostrud minim amet laborum.",
		Description: "Anim mollit aliquip aute et ea ex exercitation ipsum reprehenderit duis aute proident irure. Consectetur enim non voluptate enim do excepteur ipsum aliqua.",
		Price:       7,
		Cover:       "https://images.unsplash.com/photo-1531722569936-825d3dd91b15?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=1740&q=80",
	},
	{
		ID:          "64d6597fe18a6388e9b47c30",
		Label:       "Fugiat id ad do incididunt labore quis proident anim.",
		Description: "Excepteur commodo sunt in qui ut mollit sunt enim reprehenderit reprehenderit. Labore occaecat quis cillum do cupidatat ut nostrud.",
		Price:       9,
		Cover:       "https://images.unsplash.com/photo-1531722569936-825d3dd91b15?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=1740&q=80",
	},
}

func main() {
	gin.SetMode(gin.ReleaseMode)

	tracer.Start()
	defer tracer.Stop()

	router := gin.Default()

	// Send app traces to Datadog
	router.Use(gintrace.Middleware("e-glide-backend"))

	// Setup logger

	// UNIX Time is faster and smaller than most timestamps
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	// Send logs to file
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logger := zerolog.New(os.Stderr).With().Str("env", gin.Mode()).Timestamp().Logger()
		logger.Fatal().Err(err).Msg("Failed to open log file")
		return
	}
	defer file.Close()

	log.Logger = zerolog.New(file).With().Str("env", gin.Mode()).Timestamp().Logger()

	// End setup logger

	router.GET("/articles", getArticles)
	router.GET("/articles/:id", getArticleById)

	router.Run("localhost:8080")
}

func getArticles(c *gin.Context) {
	log.Debug().Msg("Getting list of articles")
	c.IndentedJSON(http.StatusOK, articles)
}

func getArticleById(c *gin.Context) {
	id := c.Param("id")

	for _, article := range articles {
		if article.ID == id {
			c.IndentedJSON(http.StatusOK, article)
			log.Info().Str("article-id", id).Msg("Retrieved article")
			return
		}
	}

	log.Error().Str("article-id", id).Msg("Article not found")
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Could not find an article with this id"})
}
