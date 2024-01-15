package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/dirham/pgxrest/requests"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
)

func main() {
	conn, err := pgx.Connect(context.Background(), "postgresql://someroot:secret_12@localhost:5432/test")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())
	// test handler
	router := chi.NewRouter()

	// select query handler
	router.Get("/{tableName}", func(w http.ResponseWriter, r *http.Request) {
		// create instance of requests query_string
		q := requests.NewQuery()
		q.SetTable(chi.URLParam(r, "tableName"))
		err := q.ParseUrl(r.URL.RequestURI())

		if err != nil {
			log.Printf("got error: %s", err.Error())
		}

		query, args, err := q.SelectQuery()
		if err != nil {
			log.Printf("got error when build query %s", err.Error())
			return
		}

		rows, err := conn.Query(r.Context(), query, args...)
		if err != nil {
			log.Printf("got error on query db: %s", err.Error())
			return
		}

		defer rows.Close()

		data, err := pgx.CollectRows(rows, pgx.RowToMap)
		if err != nil {
			log.Printf("failed to collect row: %s", err.Error())
		}
		resj, err := json.Marshal(data)
		if err != nil {
			log.Printf("error on marshal: %s", err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(resj)

		// render.Render(w, r, resj)
	})

	// add handlers for POST|PUT|DELETE

	http.ListenAndServe(":9000", router)

}
