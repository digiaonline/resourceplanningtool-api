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
		"shortdescription": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if project, ok := p.Source.(*Project); ok == true {
					return project.ShortDescription, nil
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
		"contactemail": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if project, ok := p.Source.(*Project); ok == true {
					return project.ContactEmail, nil
				}
				return nil, nil
			},
		},
		"picture": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if project, ok := p.Source.(*Project); ok == true {
					return project.Picture, nil
				}
				return nil, nil
			},
		},
		"ongoing": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Boolean),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if project, ok := p.Source.(*Project); ok == true {
					return project.Ongoing, nil
				}
				return nil, nil
			},
		},
		"starttime": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if project, ok := p.Source.(*Project); ok == true {
					return project.StartTime, nil
				}
				return nil, nil
			},
		},
		"endtime": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if project, ok := p.Source.(*Project); ok == true {
					return project.EndTime, nil
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
	})
	ProjectType.AddFieldConfig("persons", &graphql.Field{
		Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(PersonType))),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if project, ok := p.Source.(*Project); ok == true {
				return GetPersonsInProjectByID(project.ID)
			}
			return nil, nil
		},
	})
	ProjectType.AddFieldConfig("customer", &graphql.Field{
		Type: CustomerType,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if project, ok := p.Source.(*Project); ok == true {
				return GetProjectsCustomerByID(project.ID)
			}
			return nil, nil
		},
	})
}
