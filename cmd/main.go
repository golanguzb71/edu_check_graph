package main

import (
	"context"
	"edu_test_graph/graph"
	"edu_test_graph/internal/config"
	"edu_test_graph/internal/repository"
	"edu_test_graph/internal/service"
	"github.com/gorilla/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	loadEnv()
	port := os.Getenv("PORT")
	database.ConnectPostgres()
	database.ConnectRedis()
	db := database.DB
	groupRepo := repository.NewGroupRepository(db)
	answerRepo := repository.NewAnswerRepository(db)
	collectionRepo := repository.NewCollectionRepository(db)
	studentRepo := repository.NewStudentRepository(db)
	questionRepo := repository.NewQuestionRepository(db)

	groupService := service.NewGroupService(groupRepo)
	answerService := service.NewAnswerService(answerRepo)
	collectionService := service.NewCollectionService(collectionRepo)
	studentService := service.NewStudentService(studentRepo)
	questionService := service.NewQuestionService(questionRepo)

	gqlServer := startGraphQLServer(port, groupService, answerService, collectionService, studentService, questionService)

	waitForShutDown(gqlServer)
}

func waitForShutDown(server *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down servers...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}

	log.Println("Servers exiting")
}

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func startGraphQLServer(port string, groupService *service.GroupService, answerService *service.AnswerService, collectionService *service.CollectionService, studentService *service.StudentService, questionService *service.QuestionService) *http.Server {
	gqlMux := http.NewServeMux()

	gqlMux.Handle("/", playground.Handler("GraphQL playground", "/query"))

	gqlMux.Handle("/query", handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: &graph.Resolver{
			GroupService:      groupService,
			AnswerService:     answerService,
			CollectionService: collectionService,
			StudentService:    studentService,
			QuestionService:   questionService,
		},
	})))

	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Origin", "Content-Type", "Authorization"}),
	)

	gqlSrv := &http.Server{
		Addr:    ":" + port,
		Handler: corsHandler(gqlMux),
	}

	go func() {
		if err := gqlSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting GraphQL server: %v", err)
		}
	}()
	log.Println("Server starting ...")
	return gqlSrv
}
