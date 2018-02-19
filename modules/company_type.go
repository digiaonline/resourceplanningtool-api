package modules

import (
	"github.com/graphql-go/graphql"
)

var CompanyType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Company",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.ID),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if company, ok := p.Source.(*Company); ok == true {
					return company.ID, nil
				}
				return nil, nil
			},
		},
		"name": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if company, ok := p.Source.(*Company); ok == true {
					return company.Name, nil
				}
				return nil, nil
			},
		},
		"url": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if company, ok := p.Source.(*Company); ok == true {
					return company.URL, nil
				}
				return nil, nil
			},
		},
		"industry": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if company, ok := p.Source.(*Company); ok == true {
					return company.Industry, nil
				}
				return nil, nil
			},
		},
	},
})

func init() {
	CompanyType.AddFieldConfig("company", &graphql.Field{
		Type: CompanyType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Description: "Company ID",
				Type:	     graphql.NewNonNull(graphql.ID),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if company, ok := p.Source.(*Company); ok == true {
				return GetProjectsOfCompanyByID(company.ID)
			}
			return nil, nil
		},
	})
	CompanyType.AddFieldConfig("projects", &graphql.Field{
		Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(ProjectType))),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if company, ok := p.Source.(*Company); ok == true {
				return GetProjectsOfCompanyByID(company.ID)
			}
			return []Project{}, nil
		},
	})
}
