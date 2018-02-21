package modules

import (
	"github.com/graphql-go/graphql"
)

var PersonType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Person",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.ID),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if person, ok := p.Source.(*Person); ok == true {
					return person.ID, nil
				}
				return nil, nil
			},
		},
		"name": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if person, ok := p.Source.(*Person); ok == true {
					return person.Name, nil
				}
				return nil, nil
			},
		},
		"email": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if person, ok := p.Source.(*Person); ok == true {
					return person.Email, nil
				}
				return nil, nil
			},
		},
		"title": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if person, ok := p.Source.(*Person); ok == true {
					return person.Title, nil
				}
				return nil, nil
			},
		},
		"description": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if person, ok := p.Source.(*Person); ok == true {
					return person.Description, nil
				}
				return nil, nil
			},
		},
		"location": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if person, ok := p.Source.(*Person); ok == true {
					return person.Location, nil
				}
				return nil, nil
			},
		},
		"picture": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if person, ok := p.Source.(*Person); ok == true {
					return person.Picture, nil
				}
				return nil, nil
			},
		},
		"githuburl": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if person, ok := p.Source.(*Person); ok == true {
					return person.GithubURL, nil
				}
				return nil, nil
			},
		},
		"linkedinurl": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if person, ok := p.Source.(*Person); ok == true {
					return person.LinkedInURL, nil
				}
				return nil, nil
			},
		},
		"startdate": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if person, ok := p.Source.(*Person); ok == true {
					return person.StartDate, nil
				}
				return nil, nil
			},
		},
	},
})

func init() {
	PersonType.AddFieldConfig("person", &graphql.Field{
		Type: PersonType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Description: "Person ID",
				Type:	     graphql.NewNonNull(graphql.ID),
			},
		},
	})
	PersonType.AddFieldConfig("skills", &graphql.Field{
		Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(SkillType))),
		Resolve: func(p graphql.ResolveParams) (interface {}, error) {
			if person, ok := p.Source.(*Person); ok == true {
				return GetPersonsSkillsByID(person.ID)
			}
			return nil, nil
		},
	})
}
