package repository

func (r *Repository) CreateScheduler() error {

	query := `
        CREATE TABLE scheduler (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            date CHAR(8) NOT NULL DEFAULT "",
            title VARCHAR(256) NOT NULL DEFAULT "",
            comment TEXT NOT NULL DEFAULT "",
            repeat CHAR(128) NOT NULL DEFAULT ""
        );
        CREATE INDEX date_scheduler ON scheduler (date);`
	if _, err := r.db.Exec(query); err != nil {
		return err
	}

	return nil
}
