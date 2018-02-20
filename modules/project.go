package modules

type Project struct {
	ID		int
	Name		string
	Description	string
}

func InsertProject(project *Project) error {
	var id int
	err := db.QueryRow(`INSERT INTO project (name, description)
			    VALUES ($1, $2)
			    RETURNING id`,
			    project.Name, project.Description).Scan(&id)
	if err != nil {
		return err
	}
	project.ID = id
	return nil
}

func GetProjectByID(id int) (*Project, error) {
	var name, description string
	err := db.QueryRow(`SELECT name, description FROM project where id=$1`, id).Scan(&name, &description)
	if err != nil {
		return nil, err
	}
	return &Project{
		ID:		id,
		Name:		name,
		Description:	description,
	}, nil
	return nil, nil
}

func RemoveProjectByID(id int) error {
	_, err := db.Exec(`DELETE FROM project WHERE id=$1`, id)
	return err
}

func AddToProject(project_id, person_id int) error {
	_, err := db.Exec(`INSERT INTO worksinproject (project_id, person_id) VALUES ($1, $2)`, project_id, person_id)
	return err
}

func RemoveFromProject(project_id, person_id int) error {
	_, err := db.Exec(`DELETE FROM worksinproject WHERE project_id=$1 AND person_id=$2`, project_id, person_id)
	return err
}

func GetUsersInProjectByID(id int) ([]*Person, error) {
	rows, err := db.Query(`SELECT pers.id, pers.name, pers.email, pers.description
			       FROM person AS pers, worksinproject AS wproj
			       WHERE pers.id = wproje.person_id
			       AND wproj.project_id=$1`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var (
		persons		= []*Person{}
		person_id	int
		name		string
		email		string
		description	string
	)
	for rows.Next() {
		if err = rows.Scan(&person_id, &name, &email, &description); err != nil {
			return nil, err
		}
		persons = append(persons, &Person{ID: person_id, Name: name, Email: email, Description: description})
	}
	return persons, nil
}

func GetProjectsCustomerByID(id int) (*Customer, error) {
	customer := Customer{}
	err := db.QueryRowx(`SELECT cust.id, cust.name, cust.url, cust.industry
			     FROM projectscustomer as pc, customer as cust
			     WHERE projectscustomer.project_id=$1
			     AND customer.id=projectscustomer.customer_id
			     LIMIT 1`, id).StructScan(&customer)
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

func GetProjectsList() ([]*Project, error) {
	rows, err := db.Queryx(`SELECT * FROM project`)
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
