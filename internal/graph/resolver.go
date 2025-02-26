package graph

import (
	"encoding/base64"
	"github.com/Cococtel/Cococtel_Gagateway/internal/services/postservice"
	"strings"

	"github.com/Cococtel/Cococtel_Gagateway/internal/domain/dtos"
	"github.com/Cococtel/Cococtel_Gagateway/internal/services/authservice"
	"github.com/Cococtel/Cococtel_Gagateway/internal/services/catalogservice"
	"github.com/graphql-go/graphql"
)

func NewSchema(
	catalogService catalogservice.ICatalog,
	authService authservice.IAuth,
	scrappingService catalogservice.IScrapping,
	aiService catalogservice.IAI,
	postsService postservice.PostsService,
) graphql.SchemaConfig {
	// Tipos ya definidos
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
	ingredientType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Ingredient",
		Fields: graphql.Fields{
			"_id":      &graphql.Field{Type: graphql.String},
			"name":     &graphql.Field{Type: graphql.String},
			"quantity": &graphql.Field{Type: graphql.String},
		},
	})

	ratingType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Rating",
		Fields: graphql.Fields{
			"user_id": &graphql.Field{Type: graphql.String},
			"rating":  &graphql.Field{Type: graphql.Int},
		},
	})

	recipeType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Recipe",
		Fields: graphql.Fields{
			"_id":           &graphql.Field{Type: graphql.String},
			"name":          &graphql.Field{Type: graphql.String},
			"category":      &graphql.Field{Type: graphql.String},
			"ingredients":   &graphql.Field{Type: graphql.NewList(ingredientType)},
			"instructions":  &graphql.Field{Type: graphql.NewList(graphql.String)},
			"creatorId":     &graphql.Field{Type: graphql.String},
			"rating":        &graphql.Field{Type: graphql.Int},
			"likes":         &graphql.Field{Type: graphql.Int},
			"liquors":       &graphql.Field{Type: graphql.NewList(graphql.String)},
			"createdAt":     &graphql.Field{Type: graphql.String},
			"ratings":       &graphql.Field{Type: graphql.NewList(ratingType)},
			"description":   &graphql.Field{Type: graphql.String},
			"averageRating": &graphql.Field{Type: graphql.Float},
		},
	})

	userType := graphql.NewObject(graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"user_id":  &graphql.Field{Type: graphql.String},
			"name":     &graphql.Field{Type: graphql.String},
			"lastname": &graphql.Field{Type: graphql.String},
			"email":    &graphql.Field{Type: graphql.String},
			"phone":    &graphql.Field{Type: graphql.String},
			"image":    &graphql.Field{Type: graphql.String},
			"username": &graphql.Field{Type: graphql.String},
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
	aiRecipeType := graphql.NewObject(graphql.ObjectConfig{
		Name: "AIRecipe",
		Fields: graphql.Fields{
			"cocktailName": &graphql.Field{Type: graphql.String},
			"ingredients":  &graphql.Field{Type: graphql.NewList(ingredientType)},
			"steps":        &graphql.Field{Type: graphql.NewList(graphql.String)},
			"observations": &graphql.Field{Type: graphql.String},
		},
	})
	productType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Product",
		Fields: graphql.Fields{
			"name":                  &graphql.Field{Type: graphql.String},
			"photo_link":            &graphql.Field{Type: graphql.String},
			"description":           &graphql.Field{Type: graphql.String},
			"additional_attributes": &graphql.Field{Type: graphql.String},
			"isbn":                  &graphql.Field{Type: graphql.String},
		},
	})
	errorType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Error",
		Fields: graphql.Fields{
			"message": &graphql.Field{Type: graphql.String},
			"status":  &graphql.Field{Type: graphql.Int},
		},
	})

	// Nuevos tipos para posts
	interactionType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Interaction",
		Fields: graphql.Fields{
			"type":      &graphql.Field{Type: graphql.Int},
			"value":     &graphql.Field{Type: graphql.String},
			"userId":    &graphql.Field{Type: graphql.String},
			"createdAt": &graphql.Field{Type: graphql.String},
		},
	})
	postType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Post",
		Fields: graphql.Fields{
			"_id":          &graphql.Field{Type: graphql.String},
			"urlImage":     &graphql.Field{Type: graphql.String},
			"title":        &graphql.Field{Type: graphql.String},
			"content":      &graphql.Field{Type: graphql.String},
			"author":       &graphql.Field{Type: graphql.String},
			"createdAt":    &graphql.Field{Type: graphql.String},
			"interactions": &graphql.Field{Type: graphql.NewList(interactionType)},
		},
	})

	// Tipos de respuesta
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
	stringProcessResponseType := graphql.NewObject(graphql.ObjectConfig{
		Name: "StringProcessResponse",
		Fields: graphql.Fields{
			"data":  &graphql.Field{Type: graphql.String},
			"error": &graphql.Field{Type: errorType},
		},
	})
	aiRecipeResponseType := graphql.NewObject(graphql.ObjectConfig{
		Name: "AIRecipeResponse",
		Fields: graphql.Fields{
			"data":  &graphql.Field{Type: aiRecipeType},
			"error": &graphql.Field{Type: errorType},
		},
	})
	productResponseType := graphql.NewObject(graphql.ObjectConfig{
		Name: "ProductResponse",
		Fields: graphql.Fields{
			"data":  &graphql.Field{Type: productType},
			"error": &graphql.Field{Type: errorType},
		},
	})
	postResponseType := graphql.NewObject(graphql.ObjectConfig{
		Name: "PostResponse",
		Fields: graphql.Fields{
			"data":  &graphql.Field{Type: postType},
			"error": &graphql.Field{Type: errorType},
		},
	})
	postsResponseType := graphql.NewObject(graphql.ObjectConfig{
		Name: "PostsResponse",
		Fields: graphql.Fields{
			"data":  &graphql.Field{Type: graphql.NewList(postType)},
			"error": &graphql.Field{Type: errorType},
		},
	})

	// Nuevas funcionalidades de autorización
	userInputType := graphql.NewInputObject(graphql.InputObjectConfig{
		Name: "UserInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"name":     &graphql.InputObjectFieldConfig{Type: graphql.String},
			"lastname": &graphql.InputObjectFieldConfig{Type: graphql.String},
			"phone":    &graphql.InputObjectFieldConfig{Type: graphql.String},
			"email":    &graphql.InputObjectFieldConfig{Type: graphql.String},
			"username": &graphql.InputObjectFieldConfig{Type: graphql.String},
			"image":    &graphql.InputObjectFieldConfig{Type: graphql.String},
		},
	})
	editProfileResponseType := graphql.NewObject(graphql.ObjectConfig{
		Name: "EditProfileResponse",
		Fields: graphql.Fields{
			"data":  &graphql.Field{Type: graphql.String}, // mensaje de éxito
			"error": &graphql.Field{Type: errorType},
		},
	})

	var imageTextResponseType = graphql.NewObject(graphql.ObjectConfig{
		Name: "ImageTextResponse",
		Fields: graphql.Fields{
			"data":  &graphql.Field{Type: graphql.NewList(graphql.String)},
			"error": &graphql.Field{Type: errorType},
		},
	})

	getUserField := &graphql.Field{
		Type: userResponseType,
		Args: graphql.FieldConfigArgument{
			"id":    &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			"token": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			id := params.Args["id"].(string)
			token := params.Args["token"].(string)
			user, apiErr := authService.GetUser(id, token)
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
				"data":  user,
				"error": nil,
			}, nil
		},
	}

	editProfileField := &graphql.Field{
		Type: editProfileResponseType,
		Args: graphql.FieldConfigArgument{
			"user":  &graphql.ArgumentConfig{Type: userInputType},
			"token": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			token := params.Args["token"].(string)
			var userInput dtos.User
			if rawUser, ok := params.Args["user"].(map[string]interface{}); ok {
				if val, exists := rawUser["name"]; exists {
					strVal := val.(string)
					userInput.Name = &strVal
				}
				if val, exists := rawUser["lastname"]; exists {
					strVal := val.(string)
					userInput.Lastname = &strVal
				}
				if val, exists := rawUser["phone"]; exists {
					strVal := val.(string)
					userInput.Phone = &strVal
				}
				if val, exists := rawUser["email"]; exists {
					strVal := val.(string)
					userInput.Email = &strVal
				}
				if val, exists := rawUser["image"]; exists {
					strVal := val.(string)
					userInput.Image = &strVal
				}
				if val, exists := rawUser["username"]; exists {
					strVal := val.(string)
					userInput.Username = &strVal
				}
			}
			apiErr := authService.EditProfile(userInput, token)
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
				"data":  "Profile updated successfully",
				"error": nil,
			}, nil
		},
	}

	// Campos de Query
	queryFields := graphql.Fields{
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
		"posts": &graphql.Field{
			Type: postsResponseType,
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				posts, apiErr := postsService.GetPosts()
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
					"data":  posts,
					"error": nil,
				}, nil
			},
		},
		"post": &graphql.Field{
			Type: postResponseType,
			Args: graphql.FieldConfigArgument{
				"_id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				id := params.Args["_id"].(string)
				post, apiErr := postsService.GetPostByID(id)
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
					"data":  post,
					"error": nil,
				}, nil
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
		"getUser": getUserField,
		"getProductByCode": &graphql.Field{
			Type: productResponseType,
			Args: graphql.FieldConfigArgument{
				"code": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				code := params.Args["code"].(string)
				product, apiErr := scrappingService.GetProductByCode(code)
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
					"data":  product,
					"error": nil,
				}, nil
			},
		},
	}

	mutationFields := graphql.Fields{
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
				"_id":                   &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
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
			Type: recipeType,
			Args: graphql.FieldConfigArgument{
				"name":     &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"category": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"ingredients": &graphql.ArgumentConfig{Type: graphql.NewList(graphql.NewInputObject(graphql.InputObjectConfig{
					Name: "IngredientInput",
					Fields: graphql.InputObjectConfigFieldMap{
						"name":     &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
						"quantity": &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
					},
				}))},
				"instructions": &graphql.ArgumentConfig{Type: graphql.NewList(graphql.String)},
				"creatorId":    &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"description":  &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				ingredientsInterface, _ := params.Args["ingredients"].([]interface{})
				var ingredients []dtos.Ingredient
				for _, ingredient := range ingredientsInterface {
					ingMap := ingredient.(map[string]interface{})
					ingredients = append(ingredients, dtos.Ingredient{
						Name:     ingMap["name"].(string),
						Quantity: ingMap["quantity"].(string),
					})
				}

				instructionsInterface, _ := params.Args["instructions"].([]interface{})
				var instructions []string
				for _, instruction := range instructionsInterface {
					instructions = append(instructions, instruction.(string))
				}

				recipe := dtos.Recipe{
					Name:         params.Args["name"].(string),
					Category:     params.Args["category"].(string),
					Ingredients:  ingredients,
					Instructions: instructions,
					CreatorId:    params.Args["creatorId"].(string),
					Description:  params.Args["description"].(string),
				}
				newRecipe, apiErr := catalogService.CreateRecipe(recipe)
				if apiErr != nil {
					return nil, apiErr.Message()
				}
				return newRecipe, nil
			},
		},
		"updateRecipe": &graphql.Field{
			Type: recipeType,
			Args: graphql.FieldConfigArgument{
				"_id":         &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"name":        &graphql.ArgumentConfig{Type: graphql.String},
				"category":    &graphql.ArgumentConfig{Type: graphql.String},
				"description": &graphql.ArgumentConfig{Type: graphql.String},
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
					return nil, apiErr.Message()
				}
				return updatedRecipe, nil
			},
		},
		"deleteRecipe": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Boolean),
			Args: graphql.FieldConfigArgument{
				"_id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				id := params.Args["_id"].(string)
				err := catalogService.DeleteRecipe(id)
				if err != nil {
					return false, err.Message()
				}
				return true, nil
			},
		},
		"createPost": &graphql.Field{
			Type: postResponseType,
			Args: graphql.FieldConfigArgument{
				"urlImage": &graphql.ArgumentConfig{Type: graphql.String},
				"title":    &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"content":  &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"author":   &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				post := dtos.Post{
					UrlImage: params.Args["urlImage"].(string),
					Title:    params.Args["title"].(string),
					Content:  params.Args["content"].(string),
					Author:   params.Args["author"].(string),
				}
				newPost, apiErr := postsService.CreatePost(post)
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
					"data":  newPost,
					"error": nil,
				}, nil
			},
		},
		"updatePost": &graphql.Field{
			Type: postResponseType,
			Args: graphql.FieldConfigArgument{
				"_id":      &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"urlImage": &graphql.ArgumentConfig{Type: graphql.String},
				"title":    &graphql.ArgumentConfig{Type: graphql.String},
				"content":  &graphql.ArgumentConfig{Type: graphql.String},
				"author":   &graphql.ArgumentConfig{Type: graphql.String},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				id := params.Args["_id"].(string)
				updates := make(map[string]interface{})
				for key, value := range params.Args {
					if key != "_id" {
						updates[key] = value
					}
				}
				updatedPost, apiErr := postsService.UpdatePost(id, updates)
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
					"data":  updatedPost,
					"error": nil,
				}, nil
			},
		},
		"deletePost": &graphql.Field{
			Type: graphql.NewObject(graphql.ObjectConfig{
				Name: "DeletePostResponse",
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
				apiErr := postsService.DeletePost(id)
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
					"data":  "post deleted successfully",
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
					Name:     stringToPointer(params.Args["name"].(string)),
					Lastname: stringToPointer(params.Args["lastname"].(string)),
					Phone:    stringToPointer(params.Args["phone"].(string)),
					Email:    stringToPointer(params.Args["email"].(string)),
					Image:    stringToPointer(params.Args["image"].(string)),
					Username: stringToPointer(params.Args["username"].(string)),
					Password: stringToPointer(params.Args["password"].(string)),
					Type:     stringToPointer(params.Args["type"].(string)),
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
				user := stringToPointer(params.Args["user"].(string))
				password := stringToPointer(params.Args["password"].(string))
				var accountType *string
				if params.Args["type"] != nil {
					accountType = stringToPointer(params.Args["type"].(string))
				}
				credentials := dtos.Login{
					User:     user,
					Password: password,
					Type:     accountType,
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
		"processStrings": &graphql.Field{
			Type: stringProcessResponseType,
			Args: graphql.FieldConfigArgument{
				"input": &graphql.ArgumentConfig{Type: graphql.NewList(graphql.String)},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				rawInput, ok := params.Args["input"].([]interface{})
				if !ok {
					return map[string]interface{}{
						"data": nil,
						"error": map[string]interface{}{
							"message": "Invalid input format",
							"status":  400,
						},
					}, nil
				}
				var input []string
				for _, item := range rawInput {
					if str, isString := item.(string); isString {
						input = append(input, str)
					}
				}
				result, apiErr := aiService.ProcessStrings(input)
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
					"data":  result,
					"error": nil,
				}, nil
			},
		},
		"createAIRecipe": &graphql.Field{
			Type: aiRecipeResponseType,
			Args: graphql.FieldConfigArgument{
				"liquor": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				liquor := params.Args["liquor"].(string)
				recipe, apiErr := aiService.CreateRecipe(liquor)
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
					"data":  recipe,
					"error": nil,
				}, nil
			},
		},
		"extractTextFromImageBytes": &graphql.Field{
			Type: imageTextResponseType,
			Args: graphql.FieldConfigArgument{
				"imageBase64": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				imageBase64 := params.Args["imageBase64"].(string)
				imageBytes, err := decodeBase64(imageBase64)
				if err != nil {
					return map[string]interface{}{
						"data": nil,
						"error": map[string]interface{}{
							"message": "Invalid base64 string",
							"status":  400,
						},
					}, nil
				}
				texts, apiErr := aiService.ExtractTextFromImage(imageBytes)
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
					"data":  texts,
					"error": nil,
				}, nil
			},
		},
		"editProfile": editProfileField,
	}

	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name:   "Query",
		Fields: queryFields,
	})

	mutationType := graphql.NewObject(graphql.ObjectConfig{
		Name:   "Mutation",
		Fields: mutationFields,
	})

	return graphql.SchemaConfig{Query: queryType, Mutation: mutationType}
}

func stringToPointer(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func decodeBase64(s string) ([]byte, error) {
	// Si la cadena contiene un prefijo (data URI), elimínalo
	if idx := strings.Index(s, ","); idx != -1 {
		s = s[idx+1:]
	}
	return base64.StdEncoding.DecodeString(s)
}
