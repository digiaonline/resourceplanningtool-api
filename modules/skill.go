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
	var name string
	var level int
	err := db.QueryRow(`SELECT name, level FROM skill where id=$1`, id).Scan(&name, &level)
	if err != nil {
		return nil, err
	}
	return &Skill{
		ID:	id,
		Name:	name,
		Level:	level,
	}, nil
}

func RemoveSkillByID(id int) error {
	_, err := db.Exec(`DELETE FROM skill WHERE id=$1`, id)
	return err
}
