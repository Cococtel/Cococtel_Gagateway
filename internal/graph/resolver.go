package graph

import (
	"github.com/Cococtel/Cococtel_Gagateway/internal/domain/dtos"
	"github.com/Cococtel/Cococtel_Gagateway/internal/services/authservice"
	"github.com/Cococtel/Cococtel_Gagateway/internal/services/catalogservice"
	"github.com/graphql-go/graphql"
)

func NewSchema(catalogService catalogservice.ICatalog, authService authservice.IAuth) graphql.SchemaConfig {
	liquorType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Liquor",
		Fields: graphql.Fields{
			"_id":                   &graphql.Field{Type: graphql.String},
			"name":                  &graphql.Field{Type: graphql.String},
			"EAN":                   &graphql.Field{Type: graphql.Int},
			"category":              &graphql.Field{Type: graphql.String},
			"description":           &graphql.Field{Type: graphql.String},
			"additional_attributes": &graphql.Field{Type: graphql.String},
		},
	})
	recipeType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Recipe",
		Fields: graphql.Fields{
			"_id":          &graphql.Field{Type: graphql.String},
			"name":         &graphql.Field{Type: graphql.String},
			"ingredients":  &graphql.Field{Type: graphql.NewList(graphql.String)},
			"instructions": &graphql.Field{Type: graphql.String},
			"category":     &graphql.Field{Type: graphql.String},
			"createdAt":    &graphql.Field{Type: graphql.String},
			"liquors":      &graphql.Field{Type: graphql.NewList(graphql.String)},
		},
	})

	userType := graphql.NewObject(graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"user_id":  &graphql.Field{Type: graphql.String},
			"name":     &graphql.Field{Type: graphql.String},
			"lastname": &graphql.Field{Type: graphql.String},
			"email":    &graphql.Field{Type: graphql.String},
			"country":  &graphql.Field{Type: graphql.String},
			"phone":    &graphql.Field{Type: graphql.String},
			"image":    &graphql.Field{Type: graphql.String},
		},
	})

	successfulLoginType := graphql.NewObject(graphql.ObjectConfig{
		Name: "SuccessfulLogin",
		Fields: graphql.Fields{
			"user_id":      &graphql.Field{Type: graphql.String},
			"name":         &graphql.Field{Type: graphql.String},
			"double_auth":  &graphql.Field{Type: graphql.Boolean},
			"expiration":   &graphql.Field{Type: graphql.String},
			"token":        &graphql.Field{Type: graphql.String},
			"account_type": &graphql.Field{Type: graphql.String},
		},
	})

	errorType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Error",
		Fields: graphql.Fields{
			"message": &graphql.Field{Type: graphql.String},
			"status":  &graphql.Field{Type: graphql.Int},
		},
	})

	liquorResponseType := graphql.NewObject(graphql.ObjectConfig{
		Name: "LiquorResponse",
		Fields: graphql.Fields{
			"data":  &graphql.Field{Type: liquorType},
			"error": &graphql.Field{Type: errorType},
		},
	})

	liquorsResponseType := graphql.NewObject(graphql.ObjectConfig{
		Name: "LiquorsResponse",
		Fields: graphql.Fields{
			"data":  &graphql.Field{Type: graphql.NewList(liquorType)},
			"error": &graphql.Field{Type: errorType},
		},
	})

	recipeResponseType := graphql.NewObject(graphql.ObjectConfig{
		Name: "RecipeResponse",
		Fields: graphql.Fields{
			"data":  &graphql.Field{Type: recipeType},
			"error": &graphql.Field{Type: errorType},
		},
	})

	recipesResponseType := graphql.NewObject(graphql.ObjectConfig{
		Name: "RecipesResponse",
		Fields: graphql.Fields{
			"data":  &graphql.Field{Type: graphql.NewList(recipeType)},
			"error": &graphql.Field{Type: errorType},
		},
	})

	userResponseType := graphql.NewObject(graphql.ObjectConfig{
		Name: "UserResponse",
		Fields: graphql.Fields{
			"data":  &graphql.Field{Type: userType},
			"error": &graphql.Field{Type: errorType},
		},
	})

	loginResponseType := graphql.NewObject(graphql.ObjectConfig{
		Name: "LoginResponse",
		Fields: graphql.Fields{
			"data":  &graphql.Field{Type: successfulLoginType},
			"error": &graphql.Field{Type: errorType},
		},
	})

	verifyResponseType := graphql.NewObject(graphql.ObjectConfig{
		Name: "VerifyResponse",
		Fields: graphql.Fields{
			"data":  &graphql.Field{Type: graphql.String},
			"error": &graphql.Field{Type: errorType},
		},
	})

	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"liquors": &graphql.Field{
				Type: liquorsResponseType,
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					liquors, apiErr := catalogService.GetLiquors()
					if apiErr != nil {
						return map[string]interface{}{
							"data": nil,
							"error": map[string]interface{}{
								"message": apiErr.Message().Error(),
								"status":  apiErr.Status(),
							},
						}, nil
					}
					return map[string]interface{}{
						"data":  liquors,
						"error": nil,
					}, nil
				},
			},
			"liquor": &graphql.Field{
				Type: liquorResponseType,
				Args: graphql.FieldConfigArgument{
					"_id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id := params.Args["_id"].(string)
					liquor, apiErr := catalogService.GetLiquorByID(id)
					if apiErr != nil {
						return map[string]interface{}{
							"data": nil,
							"error": map[string]interface{}{
								"message": apiErr.Message().Error(),
								"status":  apiErr.Status(),
							},
						}, nil
					}
					return map[string]interface{}{
						"data":  liquor,
						"error": nil,
					}, nil
				},
			},
			"recipes": &graphql.Field{
				Type: recipesResponseType,
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					recipes, apiErr := catalogService.GetRecipes()
					if apiErr != nil {
						return map[string]interface{}{"data": nil, "error": apiErr}, nil
					}
					return map[string]interface{}{"data": recipes, "error": nil}, nil
				},
			},
			"recipe": &graphql.Field{
				Type: recipeResponseType,
				Args: graphql.FieldConfigArgument{
					"_id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id := params.Args["_id"].(string)
					recipe, apiErr := catalogService.GetRecipeByID(id)
					if apiErr != nil {
						return map[string]interface{}{"data": nil, "error": apiErr}, nil
					}
					return map[string]interface{}{"data": recipe, "error": nil}, nil
				},
			},
			"verify": &graphql.Field{
				Type: verifyResponseType,
				Args: graphql.FieldConfigArgument{
					"token": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					token := params.Args["token"].(string)
					apiErr := authService.Verify(token)
					if apiErr != nil {
						return map[string]interface{}{
							"data": nil,
							"error": map[string]interface{}{
								"message": apiErr.Message().Error(),
								"status":  apiErr.Status(),
							},
						}, nil
					}
					return map[string]interface{}{
						"data":  "ok",
						"error": nil,
					}, nil
				},
			},
		},
	})

	mutationType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createLiquor": &graphql.Field{
				Type: liquorResponseType,
				Args: graphql.FieldConfigArgument{
					"name":                  &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
					"EAN":                   &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
					"category":              &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
					"description":           &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
					"additional_attributes": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					liquor := dtos.Liquor{
						Name:                 params.Args["name"].(string),
						EAN:                  params.Args["EAN"].(int),
						Category:             params.Args["category"].(string),
						Description:          params.Args["description"].(string),
						AdditionalAttributes: params.Args["additional_attributes"].(string),
					}
					newLiquor, apiErr := catalogService.CreateLiquor(liquor)
					if apiErr != nil {
						return map[string]interface{}{
							"data": nil,
							"error": map[string]interface{}{
								"message": apiErr.Message().Error(),
								"status":  apiErr.Status(),
							},
						}, nil
					}
					return map[string]interface{}{
						"data":  newLiquor,
						"error": nil,
					}, nil
				},
			},
			"updateLiquor": &graphql.Field{
				Type: liquorResponseType,
				Args: graphql.FieldConfigArgument{
					"name":                  &graphql.ArgumentConfig{Type: graphql.String},
					"EAN":                   &graphql.ArgumentConfig{Type: graphql.Int},
					"category":              &graphql.ArgumentConfig{Type: graphql.String},
					"description":           &graphql.ArgumentConfig{Type: graphql.String},
					"additional_attributes": &graphql.ArgumentConfig{Type: graphql.String},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id := params.Args["_id"].(string)
					updates := make(map[string]interface{})
					for key, value := range params.Args {
						if key != "_id" {
							updates[key] = value
						}
					}

					updatedLiquor, apiErr := catalogService.UpdateLiquor(id, updates)
					if apiErr != nil {
						return map[string]interface{}{
							"data": nil,
							"error": map[string]interface{}{
								"message": apiErr.Message().Error(),
								"status":  apiErr.Status(),
							},
						}, nil
					}
					return map[string]interface{}{
						"data":  updatedLiquor,
						"error": nil,
					}, nil
				},
			},
			"deleteLiquor": &graphql.Field{
				Type: graphql.NewObject(graphql.ObjectConfig{
					Name: "DeleteLiquorResponse",
					Fields: graphql.Fields{
						"data":  &graphql.Field{Type: graphql.String},
						"error": &graphql.Field{Type: errorType},
					},
				}),
				Args: graphql.FieldConfigArgument{
					"_id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id := params.Args["_id"].(string)
					apiErr := catalogService.DeleteLiquor(id)

					if apiErr != nil {
						return map[string]interface{}{
							"data": nil,
							"error": map[string]interface{}{
								"message": apiErr.Message().Error(),
								"status":  apiErr.Status(),
							},
						}, nil
					}

					return map[string]interface{}{
						"data":  "liquor deleted successfully",
						"error": nil,
					}, nil
				},
			},
			"createRecipe": &graphql.Field{
				Type: recipeResponseType,
				Args: graphql.FieldConfigArgument{
					"name":         &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
					"ingredients":  &graphql.ArgumentConfig{Type: graphql.NewList(graphql.String)},
					"instructions": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
					"category":     &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
					"liquors":      &graphql.ArgumentConfig{Type: graphql.NewList(graphql.String)},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					ingredients := []string{}
					if rawIngredients, ok := params.Args["ingredients"].([]interface{}); ok {
						for _, item := range rawIngredients {
							if str, isString := item.(string); isString {
								ingredients = append(ingredients, str)
							}
						}
					}
					liquors := []string{}
					if rawLiquors, ok := params.Args["liquors"].([]interface{}); ok {
						for _, item := range rawLiquors {
							if str, isString := item.(string); isString {
								liquors = append(liquors, str)
							}
						}
					}

					recipe := dtos.Recipe{
						Name:         params.Args["name"].(string),
						Ingredients:  ingredients,
						Instructions: params.Args["instructions"].(string),
						Category:     params.Args["category"].(string),
						Liquors:      liquors,
					}

					newRecipe, apiErr := catalogService.CreateRecipe(recipe)
					if apiErr != nil {
						return map[string]interface{}{
							"data": nil,
							"error": map[string]interface{}{
								"message": apiErr.Message().Error(),
								"status":  apiErr.Status(),
							},
						}, nil
					}

					return map[string]interface{}{
						"data":  newRecipe,
						"error": nil,
					}, nil
				},
			},
			"updateRecipe": &graphql.Field{
				Type: recipeResponseType,
				Args: graphql.FieldConfigArgument{
					"_id":          &graphql.ArgumentConfig{Type: graphql.String},
					"name":         &graphql.ArgumentConfig{Type: graphql.String},
					"ingredients":  &graphql.ArgumentConfig{Type: graphql.NewList(graphql.String)},
					"instructions": &graphql.ArgumentConfig{Type: graphql.String},
					"category":     &graphql.ArgumentConfig{Type: graphql.String},
					"liquors":      &graphql.ArgumentConfig{Type: graphql.NewList(graphql.String)},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id := params.Args["_id"].(string)
					updates := make(map[string]interface{})
					for key, value := range params.Args {
						if key != "_id" {
							updates[key] = value
						}
					}

					updatedRecipe, apiErr := catalogService.UpdateRecipe(id, updates)
					if apiErr != nil {
						return map[string]interface{}{
							"data": nil,
							"error": map[string]interface{}{
								"message": apiErr.Message().Error(),
								"status":  apiErr.Status(),
							},
						}, nil
					}
					return map[string]interface{}{
						"data":  updatedRecipe,
						"error": nil,
					}, nil
				},
			},
			"deleteRecipe": &graphql.Field{
				Type: graphql.NewObject(graphql.ObjectConfig{
					Name: "DeleteRecipeResponse",
					Fields: graphql.Fields{
						"data":  &graphql.Field{Type: graphql.String},
						"error": &graphql.Field{Type: errorType},
					},
				}),
				Args: graphql.FieldConfigArgument{
					"_id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id := params.Args["_id"].(string)
					apiErr := catalogService.DeleteRecipe(id)

					if apiErr != nil {
						return map[string]interface{}{
							"data": nil,
							"error": map[string]interface{}{
								"message": apiErr.Message().Error(),
								"status":  apiErr.Status(),
							},
						}, nil
					}

					return map[string]interface{}{
						"data":  "recipe deleted successfully",
						"error": nil,
					}, nil
				},
			},
			"register": &graphql.Field{
				Type: userResponseType,
				Args: graphql.FieldConfigArgument{
					"name":     &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
					"lastname": &graphql.ArgumentConfig{Type: graphql.String},
					"phone":    &graphql.ArgumentConfig{Type: graphql.String},
					"email":    &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
					"image":    &graphql.ArgumentConfig{Type: graphql.String},
					"username": &graphql.ArgumentConfig{Type: graphql.String},
					"password": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
					"type":     &graphql.ArgumentConfig{Type: graphql.String},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					user := dtos.Register{
						Name:     params.Args["name"].(*string),
						Lastname: params.Args["lastname"].(*string),
						Phone:    params.Args["phone"].(*string),
						Email:    params.Args["email"].(*string),
						Image:    params.Args["image"].(*string),
						Username: params.Args["username"].(*string),
						Password: params.Args["password"].(*string),
						Type:     params.Args["type"].(*string),
					}
					newUser, apiErr := authService.Register(user)
					if apiErr != nil {
						return map[string]interface{}{"data": nil, "error": apiErr}, nil
					}
					return map[string]interface{}{"data": newUser, "error": nil}, nil
				},
			},
			"login": &graphql.Field{
				Type: loginResponseType,
				Args: graphql.FieldConfigArgument{
					"user":     &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
					"password": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
					"type":     &graphql.ArgumentConfig{Type: graphql.String},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					credentials := dtos.Login{
						User:     params.Args["user"].(*string),
						Password: params.Args["password"].(*string),
						Type:     params.Args["type"].(*string),
					}

					loginResponse, apiErr := authService.Login(credentials)
					if apiErr != nil {
						return map[string]interface{}{
							"data": nil,
							"error": map[string]interface{}{
								"message": apiErr.Message().Error(),
								"status":  apiErr.Status(),
							},
						}, nil
					}
					return map[string]interface{}{
						"data":  loginResponse,
						"error": nil,
					}, nil
				},
			},
		},
	})

	return graphql.SchemaConfig{Query: queryType, Mutation: mutationType}
}
