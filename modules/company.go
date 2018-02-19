package modules

type Company struct {
	ID		int
	Name		string
	URL		string
	Industry	string
}

func InsertCompany(company *Company) error {
	var id int
	err := db.QueryRow(`INSERT INTO company (name, url, industry)
			    VALUES ($1, $2, $3)
			    RETURNING id`,
			    company.Name, company.URL, company.Industry).Scan(&id)
	if err != nil {
		return err
	}
	company.ID = id
	return nil
}

func GetCompanyByID(id int) (*Company, error) {
	var name, url, industry string
	err := db.QueryRow(`SELECT name, url, industry FROM company where id=$1`, id).Scan(&name, &url, &industry)
	if err != nil {
		return nil, err
	}
	return &Company{
		ID:		id,
		Name:		name,
		URL:		url,
		Industry:	industry,
	}, nil
}

func RemoveCompanyByID(id int) error {
	_, err := db.Exec(`DELETE FROM company WHERE id=$1`, id)
	return err
}

func GetProjectsOfCompanyByID(id int) ([]*Project, error) {
	rows, err := db.Queryx(`SELECT proj.id, proj.name, proj.description
			       FROM project AS proj, projectscompany AS pc
			       WHERE proj.id = pc.project_id
			       AND pc.company_id=$1`, id)
	if err != nil {
		return nil, err
	}

	var projects = []*Project{}

	for rows.Next() {
		project := Project{}
		if err = rows.StructScan(&project); err != nil {
			return nil, err
		}
		projects = append(projects, &project)
	}
	return projects, nil
}

func GetCompaniesList() ([]*Company, error) {
	rows, err := db.Queryx(`SELECT * FROM company`)
	if err != nil {
		return nil, err
	}

	var companies = []*Company{}

	for rows.Next() {
		company := Company{}
		if err = rows.StructScan(&company); err != nil {
			return nil, err
		}
		companies = append(companies, &company)
	}
	return companies, nil
}
