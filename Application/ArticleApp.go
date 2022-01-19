package Application

import repo "Go-BarisAkdas-Blog/Domain/Repository"

type ArticleApplication struct{}

var articleRepo = repo.NewArticleRepositoory()

func (a ArticleApplication) GetAllArticles() ([]repo.ArticleDTO, error) {
	return articleRepo.GetAllArticles()
}
func (a ArticleApplication) GetAllActiveArticles() ([]repo.ArticleDTO, error) {
	return articleRepo.GetAllActiveArticles()
}
func (a ArticleApplication) GetArticleById(id int) (repo.ArticleDTO, error) {
	return articleRepo.GetArticleById(id)
}
func (a ArticleApplication) GetArticleBySlug(slug string) (repo.ArticleDTO, error) {
	return articleRepo.GetArticleBySlug(slug)
}
func (a ArticleApplication) CreateArticle(article *repo.ArticleDTO) (repo.ArticleDTO, error) {
	return articleRepo.CreateArticle(article)
}
func (a ArticleApplication) UpdateArticle(article *repo.ArticleDTO) (repo.ArticleDTO, error) {
	return articleRepo.UpdateArticle(article)
}
func (a ArticleApplication) DeleteArticle(id int) (repo.ArticleDTO, error) {
	return articleRepo.DeleteArticle(id)
}
