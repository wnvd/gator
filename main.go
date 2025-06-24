package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/wnvd/gator/internal/config"
	"github.com/wnvd/gator/internal/database"
)

const (
	USERNAME = "naveed"
)

func main() {

	var st state
	var err error
	st.cfg, err = config.Read()
	if err != nil {
		fmt.Println("unable to read file %w", err)
	}

	// database connection
	db, err := sql.Open("postgres", st.cfg.DBURL)
	if err != nil {
		fmt.Println("Failed to connect to the database ", err)
	}

	dbQueries := database.New(db)
	st.db = dbQueries

	var c commands
	c.cmdsReg = make(map[string]func(*state, command) error)

	// Registering commands
	c.register("login", handlerLogin)
	c.register("register", handlerRegister)
	c.register("users", handlerGetUsers)
	c.register("agg", handlerAggregate)
	c.register("addfeed", handlerAddFeed)

	// NOTE: This command is for development purpose only
	c.register("reset", handlerReset)

	args := os.Args

	if len(args) < 2 {
		fmt.Println("bad command")
		fmt.Println("Format: <command> <args>")
		os.Exit(1)
	}

	cmd := command{
		name:	args[1],
		args:	args[2:],
	}

	if err := c.run(&st, cmd); err != nil {
		fmt.Printf("Unable to run the command %v\n", err)
	}
}
