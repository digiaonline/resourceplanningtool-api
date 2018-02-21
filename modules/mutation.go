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
				"description": &graphql.ArgumentConfig{
					Description: "New project description",
					Type:        graphql.NewNonNull(graphql.String),
				},
				"picture": &graphql.ArgumentConfig{
					Description: "New project picture",
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
				name := p.Args["name"].(string)
				url := p.Args["url"].(string)
				industry := p.Args["industry"].(string)
				customer := &Customer{
					Name: name,
					URL: url,
					Industry: industry,
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
	},
})
