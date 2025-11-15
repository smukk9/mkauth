package health

import (
	"time"

	"github.com/smukk9/mkauth/internal/db"
)

type Service struct {
	db *db.Database
}

type DBStatus struct {
	LastCheck time.Time
}

func NewService(db *db.Database) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) CheckDatabase() (*DBStatus, error) {
	// Insert a test record
	query := `INSERT INTO health_check (checked_at) VALUES (CURRENT_TIMESTAMP)`
	_, err := s.db.Conn.Exec(query)
	if err != nil {
		return nil, err
	}

	// Retrieve the last check time
	var lastCheck time.Time
	querySelect := `SELECT checked_at FROM health_check ORDER BY id DESC LIMIT 1`
	err = s.db.Conn.QueryRow(querySelect).Scan(&lastCheck)
	if err != nil {
		return nil, err
	}

	return &DBStatus{
		LastCheck: lastCheck,
	}, nil
}
