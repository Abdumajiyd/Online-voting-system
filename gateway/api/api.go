package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/online_voting_service/gateway/api/handler"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"google.golang.org/grpc"
)

func NewGin(connPublic, connVote *grpc.ClientConn) *gin.Engine {

	handler := handler.NewHandler(connPublic, connVote)

	router := gin.Default()
	router.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, 
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	election := router.Group("api/v1/election")
	{
		election.POST("", handler.CreateElectionHandler)           
		election.PUT("", handler.UpdateElectionHandler)            
		election.DELETE("", handler.DeleteElectionHandler)         
		election.GET("/id", handler.GetElectionByIdHandler)        
		election.GET("/all", handler.GetAllElectionsHandler)       
		election.GET("/results", handler.GetCandidateVotesHandler) 
	}

	candidate := router.Group("/api/v1/candidate")
	{
		candidate.POST("", handler.CreateCandidateHandler)     
		candidate.PUT("", handler.UpdateCandidateHandler)      
		candidate.DELETE("", handler.DeleteCandidateHandler)   
		candidate.GET("/id", handler.GetCandidateByIdHandler)  
		candidate.GET("/all", handler.GetAllCandidatesHandler) 
	}
	publicVote := router.Group("/api/v1/public_vote")
	{
		publicVote.POST("", handler.CreatePublicVoteHandler)                 
		publicVote.GET("/public/id", handler.GetPublicVoteByPublicIdHandler) 
		publicVote.GET("/vote/id", handler.GetPublicVoteByVoteIdHandler)     
		publicVote.GET("/public/all", handler.GetAllPublicVotesHandler)      
		publicVote.GET("/vote/all", handler.GetAllVotesHandler)              
	}
	party := router.Group("api/v1/party")
	{
		party.POST("", handler.CreatePartyHandler)      
		party.PUT("", handler.UpdatePartyHandler)      
		party.DELETE("", handler.DeletePartyHandler)    
		party.GET("/id", handler.GetPartyByIdHandler)   
		party.GET("/all", handler.GetAllPartiesHandler) 
	}
	public := router.Group("api/v1/public")
	{
		public.POST("", handler.CreatePublicHandler)
		public.PUT("", handler.UpdatePublicHandler)
		public.DELETE("", handler.DeletePublicHandler)
		public.GET("/id", handler.GetPublicByIdHandler)  
		public.GET("/all", handler.GetAllPublicsHandler) 
	}
	return router
}
