package modules

import (
	"github.com/graphql-go/graphql"
)

var CustomerType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Customer",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.ID),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if customer, ok := p.Source.(*Customer); ok == true {
					return customer.ID, nil
				}
				return nil, nil
			},
		},
		"name": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if customer, ok := p.Source.(*Customer); ok == true {
					return customer.Name, nil
				}
				return nil, nil
			},
		},
		"url": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if customer, ok := p.Source.(*Customer); ok == true {
					return customer.URL, nil
				}
				return nil, nil
			},
		},
		"industry": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if customer, ok := p.Source.(*Customer); ok == true {
					return customer.Industry, nil
				}
				return nil, nil
			},
		},
		"logo": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if customer, ok := p.Source.(*Customer); ok == true {
					return customer.Logo, nil
				}
				return nil, nil
			},
		},
	},
})

func init() {
	CustomerType.AddFieldConfig("customer", &graphql.Field{
		Type: CustomerType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Description: "Customer ID",
				Type:	     graphql.NewNonNull(graphql.ID),
			},
		},
	})
	CustomerType.AddFieldConfig("projects", &graphql.Field{
		Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(ProjectType))),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if customer, ok := p.Source.(*Customer); ok == true {
				return GetProjectsOfCustomerByID(customer.ID)
			}
			return nil, nil
		},
	})
}
