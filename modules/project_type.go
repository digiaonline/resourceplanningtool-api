package modules

import (
	"github.com/graphql-go/graphql"
)

var ProjectType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Project",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.ID),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if project, ok := p.Source.(*Project); ok == true {
					return project.ID, nil
				}
				return nil, nil
			},
		},
		"name": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if project, ok := p.Source.(*Project); ok == true {
					return project.Name, nil
				}
				return nil, nil
			},
		},
		"description": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if project, ok := p.Source.(*Project); ok == true {
					return project.Description, nil
				}
				return nil, nil
			},
		},
	},
})

func init() {
	ProjectType.AddFieldConfig("project", &graphql.Field{
		Type: ProjectType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Description: "Project ID",
				Type:	     graphql.NewNonNull(graphql.ID),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if project, ok := p.Source.(*Project); ok == true {
				return GetProjectsCompanyByID(project.ID)
			}
			return nil, nil
		},
	})
}
