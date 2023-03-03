package startup

import (
	"context"
	"fmt"
	"linebot/di"
	model_com "linebot/models/commons"
	"log"
	"os"
)

var (
	LOCALDEBUG        = true
	version    string = "development"
	buildNum   string
	buildTime  string
	user       string
	branch     string
	commit     string
)

// start up the server
func Run() {
	info := &model_com.SystemInfo{
		Version:   version,
		BuildNum:  buildNum,
		Branch:    branch,
		Commit:    commit,
		BuildUser: user,
		BuildTime: buildTime,
	}
	// di
	if server, err := di.CreateLinebotServer(context.Background(), info); err != nil {
		fmt.Fprintf(os.Stderr, "Error during dependency injection: %v", err)
		os.Exit(1)
	} else if err := server.Run(); err != nil {
		log.Fatal(err)
		fmt.Fprintf(os.Stderr, "fatal error: %v", err)
		os.Exit(1)
	}
}
