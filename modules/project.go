package modules

type Project struct {
	ID			int
	Name			string
	ShortDescription	string
	Description		string
	ContactEmail		string
	Picture			string
	Ongoing			bool
	StartTime		int
	EndTime			int
}

func InsertProject(project *Project) error {
	var id int
	err := db.QueryRow(`INSERT INTO project (name, shortdescription, description, contactemail,
			    picture, ongoing, starttime, endtime)
			    VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
			    RETURNING id`,
			    project.Name, project.ShortDescription, project.Description, project.ContactEmail, project.Picture, project.Ongoing, project.StartTime, project.EndTime).Scan(&id)
	if err != nil {
		return err
	}
	project.ID = id
	return nil
}

func GetProjectByID(id int) (*Project, error) {
	project := Project{}
	err := db.QueryRowx(`SELECT * FROM project where id=$1`, id).Scan(&project)
	if err != nil {
		return nil, err
	}
	return &project, nil
}

func RemoveProjectByID(id int) error {
	_, err := db.Exec(`DELETE FROM project WHERE id=$1`, id)
	return err
}

func UpdateProject(project *Project) error {
	_, err := db.Exec(`UPDATE project SET name=$1, shortdescription=$2, description=$3, contactemail=$4,
			   picture=$5, ongoing=$6, starttime=$7, endtime=$8 WHERE id=$9`, project.Name,
			   project.ShortDescription, project.Description, project.ContactEmail, project.Picture,
			   project.Ongoing, project.StartTime, project.EndTime, project.ID)
	return err
}

func AddPersonToProject(project_id, person_id int) error {
	_, err := db.Exec(`INSERT INTO worksinproject (project_id, person_id) VALUES ($1, $2)`, project_id, person_id)
	return err
}

func RemovePersonFromProject(project_id, person_id int) error {
	_, err := db.Exec(`DELETE FROM worksinproject WHERE project_id=$1 AND person_id=$2`, project_id, person_id)
	return err
}

func AddTechnologyToProject(project_id, technology_id int) error {
	_, err := db.Exec(`INSERT INTO usestechnology (project_id, technology_id) VALUES ($1, $2)`, project_id, technology_id)
	return err
}

func RemoveTechnologyFromProject(project_id, technology_id int) error {
	_, err := db.Exec(`DELETE FROM usestechnology WHERE project_id=$1 AND person_id=$2`, project_id, technology_id)
	return err
}

func GetPersonsInProjectByID(id int) ([]*Person, error) {
	rows, err := db.Queryx(`SELECT pers.*
			        FROM person AS pers, worksinproject AS wproj
			        WHERE pers.id = wproj.person_id
			        AND wproj.project_id=$1`, id)
	if err != nil {
		return nil, err
	}

	var persons = []*Person{}

	for rows.Next() {
		person := Person{}
		if err = rows.StructScan(&person); err != nil {
			return nil, err
		}
		persons = append(persons, &person)
	}
	return persons, nil
}

func GetProjectsCustomerByID(id int) (*Customer, error) {
	customer := Customer{}
	err := db.QueryRowx(`SELECT cust.*
			     FROM projectscustomer as pc, customer as cust
			     WHERE pc.project_id=$1
			     AND cust.id=pc.customer_id`, id).StructScan(&customer)
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

func GetTechnologiesInProjectByID(id int) ([]*Technology, error) {
	technologies := []*Technology{}
	rows, err := db.Queryx(`SELECT tech.*
			       FROM usestechnology as usestech, technology as tech
			       WHERE usestech.project_id=$1
			       AND tech.id=usestech.technology_id`, id)

	for rows.Next() {
		technology := Technology{}
		if err = rows.StructScan(&technology); err != nil {
			return nil, err
		}
		technologies = append(technologies, &technology)
	}
	return technologies, nil
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
