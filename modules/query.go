package modules

import (
	"strconv"

	"github.com/graphql-go/graphql"
)

var QueryType = graphql.NewObject(graphql.ObjectConfig{
	Name:	"Query",
	Fields:	graphql.Fields{
		"project": &graphql.Field{
			Type:	ProjectType,
			Args:	graphql.FieldConfigArgument{
				"id":	&graphql.ArgumentConfig{
					Description:	"Project ID",
					Type:		graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				i := p.Args["id"].(string)
				id, err := strconv.Atoi(i)
				if err != nil {
					return nil, err
				}
				return GetProjectByID(id)
			},
		},
		"list_projects": &graphql.Field{
			Type:	graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(ProjectType))),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return GetProjectsList()
			},
		},
		"person": &graphql.Field{
			Type:	PersonType,
			Args:	graphql.FieldConfigArgument{
				"id":	&graphql.ArgumentConfig{
					Description:	"Person ID",
					Type:		graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				i := p.Args["id"].(string)
				id, err := strconv.Atoi(i)
				if err != nil {
					return nil, err
				}
				return GetPersonByID(id)
			},
		},
		"list_persons": &graphql.Field{
			Type:	graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(PersonType))),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return GetPersonsList()
			},
		},
		"skill": &graphql.Field{
			Type:	SkillType,
			Args:	graphql.FieldConfigArgument{
				"id":	&graphql.ArgumentConfig{
					Description:	"Skill ID",
					Type:		graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				i := p.Args["id"].(string)
				id, err := strconv.Atoi(i)
				if err != nil {
					return nil, err
				}
				return GetSkillByID(id)
			},
		},
		"list_skills": &graphql.Field{
			Type:	graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(SkillType))),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return GetSkillsList()
			},
		},
		"customer": &graphql.Field{
			Type:	CustomerType,
			Args:	graphql.FieldConfigArgument{
				"id":	&graphql.ArgumentConfig{
					Description:	"Customer ID",
					Type:		graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				i := p.Args["id"].(string)
				id, err := strconv.Atoi(i)
				if err != nil {
					return nil, err
				}
				return GetCustomerByID(id)
			},
		},
		"list_customers": &graphql.Field{
			Type:	graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(CustomerType))),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return GetCustomersList()
			},
		},
	},
})
