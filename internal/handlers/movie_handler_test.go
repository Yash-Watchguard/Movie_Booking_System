package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Yash-Watchguard/MovieTicketBooking/internal/mocks/servicemock"
	"github.com/Yash-Watchguard/MovieTicketBooking/internal/models/contextkey"
	role "github.com/Yash-Watchguard/MovieTicketBooking/internal/models/roles"
)

func TestAddmovie(t *testing.T) {
     tests:=[]struct{
		name string
		ctx context.Context
        requestBody string
		expectedStatus int
		expectedMessage string
		shouldSucced bool
		MockMovieService *servicemock.MockMovieService
	 }{
        {  
			name: "unaut",
		   ctx: context.WithValue(context.Background(),contextkey.UserRole,role.Customer),
           requestBody: `{"sdsh":"sjdsjd"}`,
		   expectedStatus: http.StatusForbidden,
		   expectedMessage: "unauthorized",
		   shouldSucced: false,
		   MockMovieService: &servicemock.MockMovieService{
			Shoulderr: false,
		   },
		},
		{
			name: "body",
			ctx: context.WithValue(context.Background(),contextkey.UserRole,role.Admin),
           requestBody: `{"name": "sjdsjd", "movie_type": "some type", "duration": "not an int"}`,
		   expectedStatus: http.StatusBadRequest,
		   expectedMessage: "Invalid request body",
		   shouldSucced: false,
		   MockMovieService: &servicemock.MockMovieService{
			Shoulderr: false,
		   },
		},
		{
		   name: "duration",
		   ctx: context.WithValue(context.Background(),contextkey.UserRole,role.Admin),
           requestBody: `{"name":"sjdjs","movie_type":"jdjsd","duration":0}`,
		   expectedStatus: http.StatusBadRequest,
		   expectedMessage: "invalid movie duration",
		   shouldSucced: false,
		   MockMovieService: &servicemock.MockMovieService{
		   Shoulderr: false,
		   },
	    },
		{
		   name: "fail",
		   ctx: context.WithValue(context.Background(),contextkey.UserRole,role.Admin),
           requestBody: `{"name":"sjdjs","movie_type":"jdjsd","duration":7}`,
		   expectedStatus: http.StatusInternalServerError,
		   expectedMessage: "Movie add fail",
		   shouldSucced: false,
		   MockMovieService: &servicemock.MockMovieService{
		   Shoulderr: true,
		   },
	    },
		{
			name: "success",
		   ctx: context.WithValue(context.Background(),contextkey.UserRole,role.Admin),
           requestBody: `{"name":"sjdjs","movie_type":"jdjsd","duration":7}`,
		   expectedStatus: http.StatusCreated,
		   expectedMessage: "movie added successfully",
		   shouldSucced: false,
		   MockMovieService: &servicemock.MockMovieService{
		   Shoulderr: false,
		   },

		},
	}

	 for _,test:=range tests{
		t.Run(test.name,func(t *testing.T) {
			 req:=httptest.NewRequest(http.MethodPost,"/v1/movies/addmovie",strings.NewReader(test.requestBody))
			 req=req.WithContext(test.ctx)

			 w:=httptest.NewRecorder()

			 MovieHandler:=NewMovieHandler(test.MockMovieService)

			 MovieHandler.AddMovie(w,req)

			 if w.Code!=test.expectedStatus{
				t.Errorf("sjdjsd")
			 }

			 if test.shouldSucced{
				var got Successresponse
				err:=json.Unmarshal(w.Body.Bytes(),&got)
				if err!=nil{
					t.Errorf("sjdjsd")
				}
				if test.expectedMessage!=got.Message{
					t.Errorf("sjdj")
				}
			 }else{
				var got ErrorResponse
				err:=json.Unmarshal(w.Body.Bytes(),&got)
				if err!=nil{
					t.Errorf("sjdjsd")
				}
				if test.expectedMessage!=got.Message{
					t.Errorf("sjdj")
				}
			 }
		})
       
	 }
}

func TestViewMovie(t *testing.T){
	 tests:=[]struct{
		name string
	
		expectedStatus int
		expectedMessage string
		shouldSucced bool
		MockMovieService *servicemock.MockMovieService
	 }{
		{
		   name: "fail",
		   
		   expectedStatus: http.StatusInternalServerError,
		   expectedMessage: "Fail view all movie",
		   shouldSucced: false,
		   MockMovieService: &servicemock.MockMovieService{
		   Shoulderr: true,
		   },
	    },
		{
			name: "success",
		  
		   expectedStatus: http.StatusOK,
		   expectedMessage: "movies retrived successfully",
		   shouldSucced: false,
		   MockMovieService: &servicemock.MockMovieService{
		   Shoulderr: false,
		   },

		},
	 }
	  for _,test:=range tests{
		t.Run(test.name,func(t *testing.T) {
			 req:=httptest.NewRequest(http.MethodPost,"/v1/movies/addmovie",nil)
			

			 w:=httptest.NewRecorder()

			 MovieHandler:=NewMovieHandler(test.MockMovieService)

			 MovieHandler.ViewAllMovies(w,req)

			 if w.Code!=test.expectedStatus{
				t.Errorf("sjdjsd")
			 }

			 if test.shouldSucced{
				var got Successresponse
				err:=json.Unmarshal(w.Body.Bytes(),&got)
				if err!=nil{
					t.Errorf("sjdjsd")
				}
				if test.expectedMessage!=got.Message{
					t.Errorf("sjdj")
				}
			 }else{
				var got ErrorResponse
				err:=json.Unmarshal(w.Body.Bytes(),&got)
				if err!=nil{
					t.Errorf("sjdjsd")
				}
				if test.expectedMessage!=got.Message{
					t.Errorf("sjdj")
				}
			 }
		})
	}
}
