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
					Description: "New Project Name",
					Type:        graphql.NewNonNull(graphql.String),
				},
				"description": &graphql.ArgumentConfig{
					Description: "New Project Description",
					Type:        graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				name := p.Args["name"].(string)
				description := p.Args["description"].(string)
				project := &Project{
					Name: name,
					Description: description,
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
		"createCompany": &graphql.Field{
			Type: CompanyType,
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Description: "New company name",
					Type:        graphql.NewNonNull(graphql.String),
				},
				"url": &graphql.ArgumentConfig{
					Description: "New company URL",
					Type:        graphql.NewNonNull(graphql.String),
				},
				"industry": &graphql.ArgumentConfig{
					Description: "New company industry",
					Type:        graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				name := p.Args["name"].(string)
				url := p.Args["url"].(string)
				industry := p.Args["industry"].(string)
				company := &Company{
					Name: name,
					URL: url,
					Industry: industry,
				}
				err := InsertCompany(company)
				return company, err
			},
		},
		"removeCompany": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Description: "Company ID to remove",
					Type:        graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				i := p.Args["id"].(string)
				id, err := strconv.Atoi(i)
				if err != nil {
					return nil, err
				}
				err = RemoveCompanyByID(id)
				return (err == nil), err
			},
		},
	},
})
