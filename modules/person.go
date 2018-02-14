package modules

type Person struct {
	ID		int
	Name		string
	Email		string
	Description	string
}

func InsertPerson(person *Person) error {
	var id int
	err := db.QueryRow(`INSERT INTO person (name, email, description)
			    VALUES ($1, $2, $3)
			    RETURNING id`,
			    person.Name, person.Email, person.Description).Scan(&id)
	if err != nil {
		return err
	}
	person.ID = id
	return nil
}

func GetPersonByID(id int) (*Person, error) {
	var name, email, description string
	err := db.QueryRow(`SELECT name, email, description FROM person where id=$1`, id).Scan(&name, &email, &description)
	if err != nil {
		return nil, err
	}
	return &Person{
		ID:		id,
		Name:		name,
		Email:		email,
		Description:	description,
	}, nil
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

func GetUsersSkillsByID(id int) ([]*Skill, error) {
	rows, err := db.Query(`SELECT sk.id, sk.name, sk.level
			       FROM skill AS sk, hasskill AS hskill
			       WHERE hskill.person_id=$1`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var (
		skills		= []*Skill{}
		skill_id	int
		name		string
		skill_level	int
	)
	for rows.Next() {
		if err = rows.Scan(&skill_id, &name, &skill_level); err != nil {
			return nil, err
		}
		skills = append(skills, &Skill{ID: skill_id, Name: name, Level: skill_level})
	}
	return skills, nil
}
