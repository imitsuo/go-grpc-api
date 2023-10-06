package schemas

type Flight struct {
	Year          int32   `json:"year"`
	Month         int32   `json:"month"`
	DayofMonth    int32   `json:"day_of_month"`
	DayOfWeek     int32   `json:"day_of_week"`
	DepTime       float32 `json:"dep_time"`
	CRSDepTime    int32   `json:"crs_dep_time"`
	ArrTime       float32 `json:"arr_time"`
	CRSArrTime    int32   `json:"crs_arr_time"`
	UniqueCarrier string  `json:"unique_carrier"`
	// FlightNum         int     `json:"flight_num"`
	// TailNum           string  `json:"tail_num"`
	// ActualElapsedTime float64 `json:"actual_elapsed_time"`
	// CRSElapsedTime    float64 `json:"crs_elapsed_time"`
	// AirTime           float64 `json:"air_time"`
	// ArrDelay          float64 `json:"arr_delay"`
	// DepDelay          float64 `json:"dep_delay"`
	// Origin            string  `json:"origin"`
	// Dest              string  `json:"dest"`
	// Distance          int     `json:"distance"`
	// TaxiIn            float64 `json:"taxi_in"`
	// TaxiOut           float64 `json:"taxi_out"`
	// Cancelled         int     `json:"cancelled"`
	// CancellationCode  string  `json:"cancellation_code"`
	// Diverted          int     `json:"diverted"`
	// CarrierDelay      float64 `json:"carrier_delay"`
	// WeatherDelay      float64 `json:"weather_delay"`
	// NASDelay          float64 `json:"nas_delay"`
	// SecurityDelay     float64 `json:"security_delay"`
	// LateAircraftDelay float64 `json:"late_aircraft_delay"`
}
