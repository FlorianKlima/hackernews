func main() {
port := os.Getenv("PORT")
if port == "" {
port = defaultPort
}

router := chi.NewRouter()

database.InitDB()
database.Migrate()
server := handler.NewDefaultServer(hackernews.NewExecutableSchema(hackernews.Config{Resolvers: &hackernews.Resolver{}}))
router.Handle("/", playground.Handler("GraphQL playground", "/query"))
router.Handle("/query", server)

log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
log.Fatal(http.ListenAndServe(":"+port, router))
}