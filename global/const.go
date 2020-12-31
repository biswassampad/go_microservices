package global

const (
	dburi       = "mongodb://127.0.0.1:27017"
	dbname      = "go-blog-application"
	performance = 100
)

var (
	jwtsecret = []byte("blogSecret")
)
