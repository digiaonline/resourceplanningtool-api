package modules

type Person struct {
	ID		int
	Name		string
	Email		string
	Title		string
	Description	string
	Location	string
	Picture		string
	GithubURL	string
	LinkedInURL	string
	StartDate	int
}

func InsertPerson(person *Person) error {
	var id int
	err := db.QueryRow(`INSERT INTO person (name, email, title, description, location,
			    picture, githuburl, linkedinurl, startdate)
			    VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
			    RETURNING id`,
			    person.Name, person.Email, person.Title, person.Description,
			    person.Location, person.Picture, person.GithubURL,
			    person.LinkedInURL, person.StartDate).Scan(&id)
	if err != nil {
		return err
	}
	person.ID = id
	return nil
}

func GetPersonByID(id int) (*Person, error) {
	person := Person{}
	person.ID = id
	err := db.QueryRowx(`SELECT* FROM person where id=$1 LIMIT 1`, id).StructScan(&person)
	if err != nil {
		return nil, err
	}
	return &person, nil
}

func RemovePersonByID(id int) error {
	_, err := db.Exec(`DELETE FROM person WHERE id=$1`, id)
	return err
}

func HasSkill(person_id, skill_id int) error {
	_, err := db.Exec(`INSERT INTO hasskill (person_id, skill_id) VALUES ($1, $2)`, person_id, skill_id)
	return err
}

func RemoveHasSkillByID(id int) error {
	_, err := db.Exec(`DELETE FROM hasskill WHERE id=$1`, id)
	return err
}

func GetPersonsSkillsByID(id int) ([]*Skill, error) {
	rows, err := db.Queryx(`SELECT sk.id, sk.name, sk.level
			        FROM skill AS sk, hasskill AS hskill
			        WHERE hskill.person_id=$1`, id)
	if err != nil {
		return nil, err
	}

	var skills = []*Skill{}

	for rows.Next() {
		skill := Skill{}
		if err = rows.StructScan(&skill); err != nil {
			return nil, err
		}
		skills = append(skills, &skill)
	}
	return skills, nil
}

func GetPersonsList() ([]*Person, error) {
	rows, err := db.Queryx(`SELECT * FROM person`)
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
