package modules

type Skill struct {
	ID	int
	Name	string
	Level	int
}

func InsertSkill(skill *Skill) error {
	var id int
	err := db.QueryRow(`INSERT INTO skill (name, level)
			    VALUES ($1, $2) RETURNING id`,
			    skill.Name, skill.Level).Scan(&id)
	if err != nil {
		return err
	}
	skill.ID = id
	return nil
}

func GetSkillByID(id int) (*Skill, error) {
	skill := Skill{}
	err := db.QueryRowx(`SELECT * FROM skill where id=$1`, id).StructScan(&skill)
	if err != nil {
		return nil, err
	}
	return &skill, nil
}

func RemoveSkillByID(id int) error {
	_, err := db.Exec(`DELETE FROM skill WHERE id=$1`, id)
	return err
}

func GetSkillsList() ([]*Skill, error) {
	rows, err := db.Queryx(`SELECT * FROM skill`)
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
