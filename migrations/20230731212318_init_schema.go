package migrations

import "database/sql"

func init() {
	migrator.AddMigrations(&Migration{
		Version: "20230731212318",
		Up:      mig_20230731212318_init_schema_up,
		Down:    mig_20230731212318_init_schema_down,
	})
}

func mig_20230731212318_init_schema_up(tx *sql.Tx) error {
	_, err := tx.Exec("CREATE TABLE users ( name varchar(255) );")
	if err != nil {
		return err
	}
	return nil
}

func mig_20230731212318_init_schema_down(tx *sql.Tx) error {
	return nil
}
