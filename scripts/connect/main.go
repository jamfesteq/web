package main

import (
	"context"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	err := run()
	if err != nil {
		fmt.Println("Failed:", err)
		os.Exit(1)
	}
}

type zone struct {
	ID        int    `db:"zoneidnumber"`
	ShortName string `db:"short_name"`
	Expansion int    `db:"expansion"`
	LongName  string `db:"long_name"`
}

func run() error {
	var err error
	var db *sqlx.DB
	db, err = sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", "peq", "peqpass", "127.0.0.1", "3306", "peq"))
	if err != nil {
		return fmt.Errorf("sql.Open: %w", err)
	}

	zones, err := zonesQuery(context.Background(), db)
	if err != nil {
		return fmt.Errorf("zones: %w", err)
	}

	query := "SELECT zone, target_zone_id, min_expansion, max_expansion FROM zone_points"
	rows, err := db.Query(query)
	if err != nil {
		return fmt.Errorf("db.Query: %w", err)
	}
	defer rows.Close()

	type zonePoint struct {
		zoneName       string
		targetZoneID   int
		targetZoneName string
		minExpansion   int
		maxExpansion   int
	}

	w, err := os.Create("zone_points.txt")
	if err != nil {
		return fmt.Errorf("os.Create: %w", err)
	}
	defer w.Close()

	for rows.Next() {
		var z zonePoint
		err = rows.Scan(&z.zoneName, &z.targetZoneID, &z.minExpansion, &z.maxExpansion)
		if err != nil {
			return fmt.Errorf("rows.Scan: %w", err)
		}
		if z.targetZoneID == 0 {
			continue
		}
		if z.targetZoneID == 33480 {
			continue
		}
		if z.targetZoneID == 1478099295 {
			continue
		}

		entry, ok := zones[z.targetZoneID]
		if !ok {
			fmt.Println("zone not found:", z.targetZoneID)
			//return fmt.Errorf("zone not found: %d", z.targetZoneID)
			continue
		}
		z.targetZoneName = entry.ShortName
		if z.targetZoneName == z.zoneName {
			continue
		}
		//w.addZoneLine('airplane', 'ecommons', 2, 'jump off island to');
		_, err = fmt.Fprintf(w, "%d w.addZoneLine('%s', '%s'); // %d\n", entry.Expansion, z.zoneName, z.targetZoneName, entry.Expansion)
		if err != nil {
			return fmt.Errorf("fmt.Fprintf: %w", err)
		}
	}

	return nil
}

func zonesQuery(ctx context.Context, db *sqlx.DB) (map[int]zone, error) {
	query := "SELECT zoneidnumber, short_name, expansion, long_name FROM zone"
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("db.Query: %w", err)
	}
	defer rows.Close()

	zones := map[int]zone{}

	for rows.Next() {
		var z zone
		err = rows.Scan(&z.ID, &z.ShortName, &z.Expansion, &z.LongName)
		if err != nil {
			return nil, fmt.Errorf("rows.Scan: %w", err)
		}
		zones[z.ID] = z
	}

	return zones, nil
}
