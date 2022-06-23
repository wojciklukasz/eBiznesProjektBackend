package controllers

import (
	"ProjektBackend/api/v1/models"
	"context"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/google"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

var googleConfig = &oauth2.Config{
	ClientID:     os.Getenv("GOOGLE_CLIENTID"),
	ClientSecret: os.Getenv("GOOGLE_CLIENTSECRET"),
	RedirectURL:  "http://localhost:8000/api/v1/oauth/google/callback",
	Scopes: []string{
		"https://www.googleapis.com/auth/userinfo.email",
		"https://www.googleapis.com/auth/userinfo.profile",
		"openid",
	},
	Endpoint: google.Endpoint,
}

var githubConfig = &oauth2.Config{
	ClientID:     os.Getenv("GITH_CLIENTID"),
	ClientSecret: os.Getenv("GITH_CLIENTSECRET"),
	Scopes: []string{
		"user:email",
	},
	Endpoint: github.Endpoint,
}

type GoogleUserInfo struct {
	Sub           string `json:"sub"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Profile       string `json:"profile"`
	Picture       string `json:"picture"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Gender        string `json:"gender"`
}

type GithubUserInfo struct {
	Email      string `json:"email"`
	Primary    bool   `json:"primary"`
	Verified   bool   `json:"verified"`
	Visibility string `json:"visibility"`
}

func GetOauthRouting(e *echo.Group) {
	g := e.Group("/oauth")

	g.GET("/google", func(c echo.Context) error {
		url := GetLoginURL("google")
		return c.JSON(http.StatusOK, map[string]string{"url": url})
	})

	g.GET("/github", func(c echo.Context) error {
		url := GetLoginURL("github")
		return c.JSON(http.StatusOK, map[string]string{"url": url})
	})

	g.GET("/google/callback", func(c echo.Context) error {
		err := HandleCallback("google", c)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, map[string]string{"message": "login successful"})
	})

	g.GET("/github/callback", func(c echo.Context) error {
		err := HandleCallback("github", c)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, map[string]string{"message": "login successful"})
	})
}

func GetLoginURL(provider string) string {
	var config *oauth2.Config
	if provider == "google" {
		config = googleConfig
	} else if provider == "github" {
		config = githubConfig
	} else {
		return "Invalid provider"
	}

	url := config.AuthCodeURL("state")
	return url
}

func GetTokenFromWeb(provider string, code string) *oauth2.Token {
	var config *oauth2.Config
	if provider == "google" {
		config = googleConfig
	} else if provider == "github" {
		config = githubConfig
	} else {
		panic("INVALID OAUTH PROVIDER")
	}

	tok, err := config.Exchange(context.Background(), code)
	if err != nil {
		panic(err)
	}

	return tok
}

func FetchGoogleUserInfo(token oauth2.Token) (*GoogleUserInfo, error) {
	client := googleConfig.Client(context.Background(), &token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Error closing body")
		}
	}(resp.Body)

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result GoogleUserInfo
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func FetchGithubUserInfo(token oauth2.Token) (*GithubUserInfo, error) {
	client := githubConfig.Client(context.Background(), &token)
	resp, err := client.Get("https://api.github.com/user/emails")
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Error closing body")
		}
	}(resp.Body)

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var results []GithubUserInfo
	if err := json.Unmarshal(data, &results); err != nil {
		return nil, err
	}

	for _, result := range results {
		if result.Primary == true {
			return &result, nil
		}
	}

	return nil, nil
}

func HandleCallback(provider string, c echo.Context) error {
	var token *oauth2.Token
	var err error
	var u models.User

	if provider == "google" {
		var email *GoogleUserInfo
		token = GetTokenFromWeb("google", c.QueryParam("code"))
		email, err = FetchGoogleUserInfo(*token)
		found := FindUser(email.Email, "google")
		if found == false {
			AddUser(email.Email, "google")
		}
		u = GetUser(email.Email, "google")
	} else if provider == "github" {
		var email *GithubUserInfo
		token = GetTokenFromWeb("github", c.QueryParam("code"))
		email, err = FetchGithubUserInfo(*token)
		found := FindUser(email.Email, "github")
		if found == false {
			AddUser(email.Email, "github")
		}
		u = GetUser(email.Email, "github")
	} else {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "invalid oauth provider"})
	}

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Failed to fetch user info"})
	}

	if provider == "google" {
		err = c.Redirect(http.StatusFound, "https://ebiznesprojekt.azurewebsites.net/login/google/"+u.GoToken+"&"+u.Email)
		if err != nil {
			return err
		}
	} else {
		err = c.Redirect(http.StatusFound, "https://ebiznesprojekt.azurewebsites.net/login/github/"+u.GoToken+"&"+u.Email)
		if err != nil {
			return err
		}
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Logged in successfully"})
}
