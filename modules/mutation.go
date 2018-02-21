package modules

import (
	"strconv"

	"github.com/graphql-go/graphql"
)

var MutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"createProject": &graphql.Field{
			Type: ProjectType,
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Description: "New project name",
					Type:        graphql.NewNonNull(graphql.String),
				},
				"shortdescription": &graphql.ArgumentConfig{
					Description: "New project short description",
					Type:        graphql.NewNonNull(graphql.String),
				},
				"description": &graphql.ArgumentConfig{
					Description: "New project description",
					Type:        graphql.NewNonNull(graphql.String),
				},
				"contactemail": &graphql.ArgumentConfig{
					Description: "New project description",
					Type:        graphql.NewNonNull(graphql.String),
				},
				"picture": &graphql.ArgumentConfig{
					Description: "New project picture",
					Type:        graphql.NewNonNull(graphql.String),
				},
				"ongoing": &graphql.ArgumentConfig{
					Description: "New project ongoing",
					Type:        graphql.NewNonNull(graphql.Boolean),
				},
				"starttime": &graphql.ArgumentConfig{
					Description: "New project starttime",
					Type:        graphql.NewNonNull(graphql.Int),
				},
				"endtime": &graphql.ArgumentConfig{
					Description: "New project endtime",
					Type:        graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				project := &Project{
					Name: p.Args["name"].(string),
					ShortDescription: p.Args["shortdescription"].(string),
					Description: p.Args["description"].(string),
					ContactEmail: p.Args["contactemail"].(string),
					Picture: p.Args["picture"].(string),
					Ongoing: p.Args["ongoing"].(bool),
					StartTime: p.Args["starttime"].(int),
					EndTime: p.Args["endtime"].(int),
				}
				err := InsertProject(project)
				return project, err
			},
		},
		"removeProject": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Description: "Project ID to remove",
					Type:        graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				i := p.Args["id"].(string)
				id, err := strconv.Atoi(i)
				if err != nil {
					return nil, err
				}
				err = RemoveProjectByID(id)
				return (err == nil), err
			},
		},
		"createCustomer": &graphql.Field{
			Type: CustomerType,
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Description: "New customer name",
					Type:        graphql.NewNonNull(graphql.String),
				},
				"url": &graphql.ArgumentConfig{
					Description: "New customer URL",
					Type:        graphql.NewNonNull(graphql.String),
				},
				"industry": &graphql.ArgumentConfig{
					Description: "New customer industry",
					Type:        graphql.NewNonNull(graphql.String),
				},
				"logo": &graphql.ArgumentConfig{
					Description: "New customer logo",
					Type:        graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				customer := &Customer{
					Name: p.Args["name"].(string),
					URL: p.Args["url"].(string),
					Industry: p.Args["industry"].(string),
					Logo: p.Args["logo"].(string),
				}
				err := InsertCustomer(customer)
				return customer, err
			},
		},
		"removeCustomer": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Description: "Customer ID to remove",
					Type:        graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				i := p.Args["id"].(string)
				id, err := strconv.Atoi(i)
				if err != nil {
					return nil, err
				}
				err = RemoveCustomerByID(id)
				return (err == nil), err
			},
		},
		"createPerson": &graphql.Field{
			Type: PersonType,
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Description: "New person name",
					Type:        graphql.NewNonNull(graphql.String),
				},
				"email": &graphql.ArgumentConfig{
					Description: "New person email",
					Type:        graphql.NewNonNull(graphql.String),
				},
				"title": &graphql.ArgumentConfig{
					Description: "New person title",
					Type:        graphql.NewNonNull(graphql.String),
				},
				"description": &graphql.ArgumentConfig{
					Description: "New person description",
					Type:        graphql.NewNonNull(graphql.String),
				},
				"location": &graphql.ArgumentConfig{
					Description: "New person location",
					Type:        graphql.NewNonNull(graphql.String),
				},
				"picture": &graphql.ArgumentConfig{
					Description: "New person picture",
					Type:        graphql.NewNonNull(graphql.String),
				},
				"githuburl": &graphql.ArgumentConfig{
					Description: "New person GithubURL",
					Type:        graphql.NewNonNull(graphql.String),
				},
				"linkedinurl": &graphql.ArgumentConfig{
					Description: "New person LinkedInURL",
					Type:        graphql.NewNonNull(graphql.String),
				},
				"startdate": &graphql.ArgumentConfig{
					Description: "New person startdate",
					Type:        graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				person := &Person{
					Name: p.Args["name"].(string),
					Email: p.Args["email"].(string),
					Title: p.Args["title"].(string),
					Description: p.Args["description"].(string),
					Location: p.Args["location"].(string),
					Picture: p.Args["picture"].(string),
					GithubURL: p.Args["githuburl"].(string),
					LinkedInURL: p.Args["linkedinurl"].(string),
					StartDate: p.Args["startdate"].(int),
				}
				err := InsertPerson(person)
				return person, err
			},
		},
		"removePerson": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Description: "Person ID to remove",
					Type:        graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				i := p.Args["id"].(string)
				id, err := strconv.Atoi(i)
				if err != nil {
					return nil, err
				}
				err = RemovePersonByID(id)
				return (err == nil), err
			},
		},
		"createSkill": &graphql.Field{
			Type: SkillType,
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Description: "New skill name",
					Type:        graphql.NewNonNull(graphql.String),
				},
				"level": &graphql.ArgumentConfig{
					Description: "New skill level",
					Type:        graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				skill := &Skill{
					Name: p.Args["name"].(string),
					Level: p.Args["level"].(int),
				}
				err := InsertSkill(skill)
				return skill, err
			},
		},
		"removeSkill": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Description: "Skill ID to remove",
					Type:        graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				i := p.Args["id"].(string)
				id, err := strconv.Atoi(i)
				if err != nil {
					return nil, err
				}
				err = RemoveSkillByID(id)
				return (err == nil), err
			},
		},
		"createTechnology": &graphql.Field{
			Type: TechnologyType,
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Description: "New technology name",
					Type:        graphql.NewNonNull(graphql.String),
				},
				"description": &graphql.ArgumentConfig{
					Description: "New technology description",
					Type:        graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				technology := &Technology{
					Name: p.Args["name"].(string),
					Description: p.Args["description"].(string),
				}
				err := InsertTechnology(technology)
				return technology, err
			},
		},
		"removeTechnology": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Description: "Technology ID to remove",
					Type:        graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				i := p.Args["id"].(string)
				id, err := strconv.Atoi(i)
				if err != nil {
					return nil, err
				}
				err = RemoveTechnologyByID(id)
				return (err == nil), err
			},
		},
		"addPersonToProject": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"project_id": &graphql.ArgumentConfig{
					Description: "Project ID",
					Type:        graphql.NewNonNull(graphql.Int),
				},
				"person_id": &graphql.ArgumentConfig{
					Description: "Person ID",
					Type:        graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				project_id := p.Args["project_id"].(int)
				person_id := p.Args["person_id"].(int)
				err := AddPersonToProject(project_id, person_id)
				return (err == nil), err
			},
		},
		"removePersonFromProject": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"project_id": &graphql.ArgumentConfig{
					Description: "Project ID",
					Type:        graphql.NewNonNull(graphql.Int),
				},
				"person_id": &graphql.ArgumentConfig{
					Description: "Person ID",
					Type:        graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				project_id := p.Args["project_id"].(int)
				person_id := p.Args["person_id"].(int)
				err := RemovePersonFromProject(project_id, person_id)
				return (err == nil), err
			},
		},
		"addTechnologyToProject": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"project_id": &graphql.ArgumentConfig{
					Description: "Project ID",
					Type:        graphql.NewNonNull(graphql.Int),
				},
				"technology_id": &graphql.ArgumentConfig{
					Description: "Technology ID",
					Type:        graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				project_id := p.Args["project_id"].(int)
				technology_id := p.Args["technology_id"].(int)
				err := AddTechnologyToProject(project_id, technology_id)
				return (err == nil), err
			},
		},
		"removeTechnologyFromProject": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"project_id": &graphql.ArgumentConfig{
					Description: "Project ID",
					Type:        graphql.NewNonNull(graphql.Int),
				},
				"technology_id": &graphql.ArgumentConfig{
					Description: "Technology ID",
					Type:        graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				project_id := p.Args["project_id"].(int)
				technology_id := p.Args["technology_id"].(int)
				err := RemoveTechnologyFromProject(project_id, technology_id)
				return (err == nil), err
			},
		},
	},
})
