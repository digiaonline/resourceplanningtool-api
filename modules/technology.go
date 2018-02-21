package modules

type Technology struct {
	ID		int
	Name		string
	Description	string
}

func InsertTechnology(technology *Technology) error {
	var id int
	err := db.QueryRow(`INSERT INTO technology (name, description)
			    VALUES ($1, $2) RETURNING id`,
			    technology.Name, technology.Description).Scan(&id)
	if err != nil {
		return err
	}
	technology.ID = id
	return nil
}

func GetTechnologyByID(id int) (*Technology, error) {
	technology := Technology{}
	err := db.QueryRowx(`SELECT * FROM technology where id=$1`, id).StructScan(&technology)
	if err != nil {
		return nil, err
	}
	return &technology, nil
}

func RemoveTechnologyByID(id int) error {
	_, err := db.Exec(`DELETE FROM technology WHERE id=$1`, id)
	return err
}

func GetTechnologiesList() ([]*Technology, error) {
	rows, err := db.Queryx(`SELECT * FROM technology`)
	if err != nil {
		return nil, err
	}

	var technologies = []*Technology{}

	for rows.Next() {
		technology := Technology{}
		if err = rows.StructScan(&technology); err != nil {
			return nil, err
		}
		technologies = append(technologies, &technology)
	}
	return technologies, nil
}
