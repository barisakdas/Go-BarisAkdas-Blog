package cmd

import (
	model "Go-BarisAkdas-Blog/Domain/Model"
	. "Go-BarisAkdas-Blog/Infrastructure/Auth"
	. "Go-BarisAkdas-Blog/Presentation"
	"net/http"
)

func Init() {
	model.CreateDb()
	model.Migration()

	http.HandleFunc("/", GetAllEndpoints)                // For endpoints
	http.HandleFunc("/get-access-token", GetAccessToken) // For access token

	// About
	http.Handle("/getAllAbouts", IsAuthorized(GetAllAbouts))
	http.Handle("/getAllActiveAbouts", IsAuthorized(GetAllActiveAbouts))
	http.Handle("/getAboutById", IsAuthorized(GetAboutById))
	http.Handle("/createAbout", IsAuthorized(CreateAbout))
	http.Handle("/updateAbout", IsAuthorized(UpdateAbout))
	http.Handle("/deleteAbout", IsAuthorized(DeleteAbout))

	// Article
	http.Handle("/getAllArticles", IsAuthorized(GetAllArticles))
	http.Handle("/getAllActiveArticles", IsAuthorized(GetAllActiveArticles))
	http.Handle("/getArticleById", IsAuthorized(GetArticleById))
	http.Handle("/getArticleBySlug", IsAuthorized(GetArticleBySlug))
	http.Handle("/createArticle", IsAuthorized(CreateArticle))
	http.Handle("/updateArticle", IsAuthorized(UpdateArticle))
	http.Handle("/deleteArticle", IsAuthorized(DeleteArticle))

	// Category
	http.Handle("/getAllCategories", IsAuthorized(GetAllCategories))
	http.Handle("/getAllActiveCategories", IsAuthorized(GetAllActiveCategories))
	http.Handle("/getCategoryById", IsAuthorized(GetCategoryById))
	http.Handle("/createCategory", IsAuthorized(CreateCategory))
	http.Handle("/updateCategory", IsAuthorized(UpdateCategory))
	http.Handle("/deleteCategory", IsAuthorized(DeleteCategory))

	// Message
	http.Handle("/getAllMessages", IsAuthorized(GetAllMessages))
	http.Handle("/getAllActiveMessages", IsAuthorized(GetAllActiveMessages))
	http.Handle("/getMessageById", IsAuthorized(GetMessageById))
	http.Handle("/createMessage", IsAuthorized(CreateMessage))
	http.Handle("/updateMessage", IsAuthorized(UpdateMessage))
	http.Handle("/replyMessage", IsAuthorized(ReplyMessage))
	http.Handle("/deleteMessage", IsAuthorized(DeleteMessage))

	// Role
	http.Handle("/getAllRoles", IsAuthorized(GetAllRoles))
	http.Handle("/getAllActiveRoles", IsAuthorized(GetAllActiveRoles))
	http.Handle("/getRoleById", IsAuthorized(GetRoleById))
	http.Handle("/createRole", IsAuthorized(CreateRole))
	http.Handle("/updateRole", IsAuthorized(UpdateRole))
	http.Handle("/deleteRole", IsAuthorized(DeleteRole))

	// User
	http.Handle("/getAllUsers", IsAuthorized(GetAllUsers))
	http.Handle("/getAllActiveUsers", IsAuthorized(GetAllActiveUsers))
	http.Handle("/getUserById", IsAuthorized(GetUserById))
	http.Handle("/createUser", IsAuthorized(CreateUser))
	http.Handle("/updateUser", IsAuthorized(UpdateUser))
	http.Handle("/deleteUser", IsAuthorized(DeleteUser))

	http.ListenAndServe(":8080", nil)
}
