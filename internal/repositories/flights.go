package repositories

import (
	"context"
	"go-grpc-api/internal/databases"
	"go-grpc-api/schemas"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type FlightsRepository struct {
	db *pgxpool.Pool
}

func NewFlightsRepository() *FlightsRepository {
	repository := &FlightsRepository{
		db: databases.GetDB(),
	}
	return repository
}

func (repo *FlightsRepository) List(skip int, limit int) ([]schemas.Flight, int, error) {

	query_total := `SELECT COUNT(1) FROM public.flights`
	var total int

	err := repo.db.QueryRow(context.Background(), query_total).Scan(&total)
	if err != nil {
		log.Printf("Failed to fetch flights total: %s\n", err.Error())
		return nil, 0, err
	}

	query := `SELECT year, month, dayofmonth, dayofweek, deptime, crsdeptime, coalesce(arrtime,0) arrtime, crsarrtime, uniquecarrier
			  FROM public.flights order by ctid offset $1 limit $2`

	rows, err := repo.db.Query(context.Background(), query, skip, limit)
	if err != nil {
		log.Printf("Failed to query flights: %s\n", err.Error())
		return nil, 0, err
	}

	flights := []schemas.Flight{}
	for rows.Next() {
		var f schemas.Flight
		err = rows.Scan(&f.Year, &f.Month, &f.DayofMonth, &f.DayOfWeek, &f.DepTime, &f.CRSDepTime, &f.ArrTime, &f.CRSArrTime, &f.UniqueCarrier)
		if err != nil {
			log.Printf("Failed to fetch flights: %s\n", err.Error())
			return nil, 0, err
		}
		flights = append(flights, f)
	}

	return flights, total, nil
}
