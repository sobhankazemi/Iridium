package handlers

import (
	"api/models"
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

type Repository struct {
	DB *sql.DB
}

func New(db *sql.DB) *Repository {
	return &Repository{
		DB: db,
	}
}

func (s *Repository) PostInfo(w http.ResponseWriter, r *http.Request) {
	infoReq := &models.Information{}
	postData, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err.Error())
	}
	err = json.Unmarshal(postData, infoReq)
	if err != nil {
		log.Println(err.Error())
	}

	retchan := make(chan error)

	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
		defer cancel()
		query := `insert into information (Os,
		KernelName,
		HostName,
		KernelRelease,
		KernelVersion,
		Machine,
		Processor,
		HwPlatform,
		UsedSpace,
		DateTime) values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)`
		_, err := s.DB.ExecContext(ctx, query,
			infoReq.Os,
			infoReq.KernelName,
			infoReq.HostName,
			infoReq.KernelRelease,
			infoReq.KernelVersion,
			infoReq.Machine,
			infoReq.Processor,
			infoReq.HwPlatform,
			infoReq.UsedSpace,
			infoReq.DateTime)

		retchan <- err
	}()

	err = <-retchan
	if err != nil {
		log.Println(err.Error())
	}
}
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
func (s *Repository) GetInfo(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	infoResp := []models.Information{}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	query := `select Os,KernelName,HostName,KernelRelease,KernelVersion,Machine,Processor,HwPlatform,UsedSpace,DateTime from information`
	rows, err := s.DB.QueryContext(ctx, query)
	if err != nil {
		log.Println(err.Error())
	}
	for rows.Next() {
		info := models.Information{}
		err := rows.Scan(
			&info.Os,
			&info.KernelName,
			&info.HostName,
			&info.KernelRelease,
			&info.KernelVersion,
			&info.Machine,
			&info.Processor,
			&info.HwPlatform,
			&info.UsedSpace,
			&info.DateTime)
		if err != nil {
			log.Println(err.Error())
		}
		infoResp = append(infoResp, info)
	}

	jsonResp, _ := json.Marshal(infoResp)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}
