package modules

type News struct {
	ID		int
	URL		string
	Description	string
}

func InsertNews(news *News) error {
	var id int
	err := db.QueryRow(`INSERT INTO news (url, description)
			    VALUES ($1, $2) RETURNING id`,
			    news.URL, news.Description).Scan(&id)
	if err != nil {
		return err
	}
	news.ID = id
	return nil
}

func GetNewsByID(id int) (*News, error) {
	news := News{}
	err := db.QueryRowx(`SELECT * FROM news where id=$1`, id).StructScan(&news)
	if err != nil {
		return nil, err
	}
	return &news, nil
}

func RemoveNewsByID(id int) error {
	_, err := db.Exec(`DELETE FROM news WHERE id=$1`, id)
	return err
}

func GetNewsList() ([]*News, error) {
	rows, err := db.Queryx(`SELECT * FROM news`)
	if err != nil {
		return nil, err
	}

	var newsList = []*News{}

	for rows.Next() {
		news := News{}
		if err = rows.StructScan(&news); err != nil {
			return nil, err
		}
		newsList = append(newsList, &news)
	}
	return newsList, nil
}
