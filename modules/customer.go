package modules

type Customer struct {
	ID		int
	Name		string
	URL		string
	Industry	string
	Logo		string
}

func InsertCustomer(customer *Customer) error {
	var id int
	err := db.QueryRow(`INSERT INTO customer (name, url, industry, logo)
			    VALUES ($1, $2, $3, $4)
			    RETURNING id`,
			    customer.Name, customer.URL, customer.Industry, customer.Logo).Scan(&id)
	if err != nil {
		return err
	}
	customer.ID = id
	return nil
}

func GetCustomerByID(id int) (*Customer, error) {
	customer := Customer{}
	err := db.QueryRowx(`SELECT * FROM customer where id=$1 LIMIT 1`, id).StructScan(&customer)
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

func RemoveCustomerByID(id int) error {
	_, err := db.Exec(`DELETE FROM customer WHERE id=$1`, id)
	return err
}

func GetProjectsOfCustomerByID(id int) ([]*Project, error) {
	rows, err := db.Queryx(`SELECT proj.*
			        FROM project AS proj, projectscustomer AS pc
			        WHERE proj.id = pc.project_id
			        AND pc.customer_id=$1`, id)
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

func GetCustomersList() ([]*Customer, error) {
	rows, err := db.Queryx(`SELECT * FROM customer`)
	if err != nil {
		return nil, err
	}

	var customers = []*Customer{}

	for rows.Next() {
		customer := Customer{}
		if err = rows.StructScan(&customer); err != nil {
			return nil, err
		}
		customers = append(customers, &customer)
	}
	return customers, nil
}
