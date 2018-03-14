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

func UpdateCustomer(customer *Customer) error {
	old_customer := Customer{}
	err := db.QueryRowx(`SELECT * FROM customer WHERE customer.id = $1`, customer.ID).StructScan(&old_customer)
	if err != nil {
		return err
	}

	err = AggregateStructs(customer, &old_customer)
	if err != nil {
		return err
	}

	_, err = db.Exec(`UPDATE customer SET name=$1, url=$2, industry=$3, logo=$4 WHERE id=$5`, customer.Name,
			  customer.URL, customer.Industry, customer.Logo, customer.ID)

	return err
}

func AddProjectToCustomer(customer_id, project_id int) error {
	_, err := db.Exec(`INSERT INTO projectscustomer (project_id, customer_id) VALUES ($1, $2)`, project_id, customer_id)
	return err
}

func RemoveProjectFromCustomer(project_id, customer_id int) error {
	_, err := db.Exec(`DELETE FROM projectscustomer WHERE project_id=$1 AND customer_id=$2`, project_id, customer_id)
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
