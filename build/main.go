package main

import (
  "github.com/alpacahq/alpaca-trade-api-go/alpaca"
  "github.com/alpacahq/alpaca-trade-api-go/common"
  "github.com/joho/godotenv"
  "html/template"
  "net/http"
  "fmt"
  "log"
  "os"
)

var (
  tpl *template.Template
)

func init() {
  tpl = template.Must(template.ParseGlob("templates/*"))
  // loads values from .env into the system
  if err := godotenv.Load(); err != nil {
    log.Print("No .env file found")
  }

  common.EnvApiKeyID, exists := os.LookupEnv("APCA_API_KEY_ID")
  if exists {
	   log.Printf("Api Key ID: %s", common.Credentials().ID)
  }
  common.EnvApiSecretKey, exists := os.LookupEnv("APCA_API_SECRET_KEY")
  if exists {
	   log.Printf("API Secret Key: %s", common.Credentials().Secret)
  }
  alpacaBaseUrl, exists := os.LookupEnv("APCA_API_BASE_URL")
  if exists {
	   log.Printf("Alpaca Base Url: %s", alpacaBaseUrl)
  }
  alpaca.SetBaseUrl(alpacaBaseUrl)
}

func main() {


  alpacaClient := alpaca.NewClient(common.Credentials())
  acct, err := alpacaClient.GetAccount()
  if err != nil {
    panic(err)
  }
  log.Print(*acct)

  http.HandleFunc("/", indexHandler)

  // Serve static files out of the public directory.
	// By configuring a static handler in app.yaml, App Engine serves all the
	// static content itself. As a result, the following two lines are in
	// effect for development only.
	public := http.StripPrefix("/public", http.FileServer(http.Dir("public")))
	http.Handle("/public/", public)

  // [START setting_port]
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
	// [END setting_port]
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/", http.StatusFound)
    return
  }
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}
