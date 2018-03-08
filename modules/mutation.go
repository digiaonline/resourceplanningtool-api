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
				"liveat": &graphql.ArgumentConfig{
					Description: "New project liveat URL",
					Type:        graphql.NewNonNull(graphql.String),
				},
				"githuburl": &graphql.ArgumentConfig{
					Description: "New project githuburl",
					Type:        graphql.NewNonNull(graphql.String),
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
					LiveAt: p.Args["liveat"].(string),
					GithubURL: p.Args["githuburl"].(string),
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
		"updateProject": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Description: "New project ID",
					Type:        graphql.NewNonNull(graphql.Int),
				},
				"name": &graphql.ArgumentConfig{
					Description:	"New project name",
					Type:		graphql.String,
				},
				"shortdescription": &graphql.ArgumentConfig{
					Description: "New project short description",
					Type:        graphql.String,
				},
				"description": &graphql.ArgumentConfig{
					Description:	"New project description",
					Type:		graphql.String,
				},
				"contactemail": &graphql.ArgumentConfig{
					Description:	"New project description",
					Type:		graphql.String,
				},
				"picture": &graphql.ArgumentConfig{
					Description:	"New project picture",
					Type:		graphql.String,
				},
				"ongoing": &graphql.ArgumentConfig{
					Description:	"New project ongoing",
					Type:		graphql.Boolean,
				},
				"starttime": &graphql.ArgumentConfig{
					Description:	"New project starttime",
					Type:		graphql.Int,
				},
				"endtime": &graphql.ArgumentConfig{
					Description: "New project endtime",
					Type:        graphql.Int,
				},
				"liveat": &graphql.ArgumentConfig{
					Description:	"New project liveat URL",
					Type:		graphql.String,
				},
				"githuburl": &graphql.ArgumentConfig{
					Description:	"New project githuburl",
					Type:		graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				project := Project{}
				for field, value := range p.Args {
					if value != "" {
						if field == "id" {
							project.ID = value.(int)
						}
						if field == "name" {
							project.Name = value.(string)
						}
						if field == "shortdescription" {
							project.ShortDescription = value.(string)
						}
						if field == "description" {
							project.Description = value.(string)
						}
						if field == "contactemail" {
							project.ContactEmail = value.(string)
						}
						if field == "picture" {
							project.Picture = value.(string)
						}
						if field == "ongoing" {
							project.Ongoing = value.(bool)
						}
						if field == "starttime" {
							project.StartTime = value.(int)
						}
						if field == "endtime" {
							project.EndTime = value.(int)
						}
						if field == "liveat" {
							project.LiveAt = value.(string)
						}
						if field == "githuburl" {
							project.GithubURL = value.(string)
						}
					}

				}
				err := UpdateProject(&project)
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
		"updateCustomer": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Description: "Customer ID",
					Type:        graphql.NewNonNull(graphql.Int),
				},
				"name": &graphql.ArgumentConfig{
					Description:	"Customer name",
					Type:		graphql.String,
				},
				"url": &graphql.ArgumentConfig{
					Description:	"Customer url",
					Type:		graphql.String,
				},
				"industry": &graphql.ArgumentConfig{
					Description:	"Customer industry",
					Type:		graphql.String,
				},
				"logo": &graphql.ArgumentConfig{
					Description:	"Customer logo",
					Type:		graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				customer := Customer{}
				for field, value := range p.Args {
					if value != "" {
						if field == "id" {
							customer.ID = value.(int)
						}
						if field == "name" {
							customer.Name = value.(string)
						}
						if field == "url" {
							customer.URL = value.(string)
						}
						if field == "industry" {
							customer.Industry = value.(string)
						}
						if field == "logo" {
							customer.Logo = value.(string)
						}
					}

				}
				err := UpdateCustomer(&customer)
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
		"updatePerson": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Description: "Customer ID",
					Type:        graphql.NewNonNull(graphql.Int),
				},
				"name": &graphql.ArgumentConfig{
					Description:	"Customer name",
					Type:		graphql.String,
				},
				"email": &graphql.ArgumentConfig{
					Description:	"Person email",
					Type:		graphql.String,
				},
				"title": &graphql.ArgumentConfig{
					Description:	"Person title",
					Type:		graphql.String,
				},
				"description": &graphql.ArgumentConfig{
					Description:	"Person description",
					Type:		graphql.String,
				},
				"location": &graphql.ArgumentConfig{
					Description:	"Person location",
					Type:		graphql.String,
				},
				"picture": &graphql.ArgumentConfig{
					Description:	"Person picture",
					Type:		graphql.String,
				},
				"githuburl": &graphql.ArgumentConfig{
					Description:	"Github url",
					Type:		graphql.String,
				},
				"linkedinurl": &graphql.ArgumentConfig{
					Description:	"LinkedIn url",
					Type:		graphql.String,
				},
				"startdate": &graphql.ArgumentConfig{
					Description:	"Person start date",
					Type:		graphql.Int,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				person := Person{}
				for field, value := range p.Args {
					if value != "" {
						if field == "id" {
							person.ID = value.(int)
						}
						if field == "name" {
							person.Name = value.(string)
						}
						if field == "email" {
							person.Email = value.(string)
						}
						if field == "title" {
							person.Title = value.(string)
						}
						if field == "description" {
							person.Description = value.(string)
						}
						if field == "location" {
							person.Location = value.(string)
						}
						if field == "githuburl" {
							person.GithubURL = value.(string)
						}
						if field == "linkedinurl" {
							person.LinkedInURL = value.(string)
						}
						if field == "startdate" {
							person.StartDate = value.(int)
						}
					}

				}
				err := UpdatePerson(&person)
				return (err == nil), err
			},
		},
		"addSkillForPerson": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"person_id": &graphql.ArgumentConfig{
					Description: "Person ID",
					Type:        graphql.NewNonNull(graphql.Int),
				},
				"skill_id": &graphql.ArgumentConfig{
					Description: "Skill ID",
					Type:        graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				person_id := p.Args["person_id"].(int)
				skill_id := p.Args["skill_id"].(int)
				err := AddSkill(person_id, skill_id)
				return (err == nil), err
			},
		},
		"removeSkillForPerson": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Description: "ID",
					Type:        graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id := p.Args["id"].(int)
				err := RemovePersonsSkillByID(id)
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
		"createNews": &graphql.Field{
			Type: NewsType,
			Args: graphql.FieldConfigArgument{
				"url": &graphql.ArgumentConfig{
					Description: "News URL",
					Type:        graphql.NewNonNull(graphql.String),
				},
				"description": &graphql.ArgumentConfig{
					Description: "News description",
					Type:        graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				news := &News{
					URL: p.Args["url"].(string),
					Description: p.Args["description"].(string),
				}
				err := InsertNews(news)
				return news, err
			},
		},
		"removeNews": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Description: "News ID to remove",
					Type:        graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				i := p.Args["id"].(string)
				id, err := strconv.Atoi(i)
				if err != nil {
					return nil, err
				}
				err = RemoveNewsByID(id)
				return (err == nil), err
			},
		},
		"addNewsToProject": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"project_id": &graphql.ArgumentConfig{
					Description: "Project ID",
					Type:        graphql.NewNonNull(graphql.Int),
				},
				"news_id": &graphql.ArgumentConfig{
					Description: "News ID",
					Type:        graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				project_id := p.Args["project_id"].(int)
				news_id := p.Args["news_id"].(int)
				err := AddNewsToProject(project_id, news_id)
				return (err == nil), err
			},
		},
		"removeNewsFromProject": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"project_id": &graphql.ArgumentConfig{
					Description: "Project ID",
					Type:        graphql.NewNonNull(graphql.Int),
				},
				"news_id": &graphql.ArgumentConfig{
					Description: "News ID",
					Type:        graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				project_id := p.Args["project_id"].(int)
				news_id := p.Args["news_id"].(int)
				err := RemoveNewsFromProject(project_id, news_id)
				return (err == nil), err
			},
		},
	},
})
