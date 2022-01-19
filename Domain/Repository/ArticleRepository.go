package Repository

import (
	model "Go-BarisAkdas-Blog/Domain/Model"
	"time"
)

type IArticleRepository interface {
	GetAllArticles() ([]ArticleDTO, error)
	GetAllActiveArticles() ([]ArticleDTO, error)
	GetArticleById(id int) (ArticleDTO, error)
	GetArticleBySlug(slug string) (ArticleDTO, error)
	CreateArticle(article *ArticleDTO) (ArticleDTO, error)
	UpdateArticle(article *ArticleDTO) (ArticleDTO, error)
	DeleteArticle(id int) (ArticleDTO, error)
}

type ArticleRepository struct{}

type ArticleDTO struct {
	ID          uint
	Title       string
	Slug        string
	Description string
	Content     string
	PictureUrl  string
	CreatedDate time.Time
	IsDeleted   bool
}

func NewArticleRepositoory() IArticleRepository {
	return &ArticleRepository{}
}

func (a ArticleRepository) GetAllArticles() ([]ArticleDTO, error) {
	data, err := model.Article{}.GetAll()
	if err != nil {
		return nil, err
	}

	var dataDtos []ArticleDTO

	for _, value := range data {
		newData := ArticleDTO{
			ID:          value.ID,
			Title:       value.Title,
			Slug:        value.Slug,
			Description: value.Description,
			Content:     value.Content,
			PictureUrl:  value.PictureUrl,
			CreatedDate: value.CreatedAt,
			IsDeleted:   value.IsDeleted,
		}
		dataDtos = append(dataDtos, newData)
	}

	return dataDtos, nil
}

func (a ArticleRepository) GetAllActiveArticles() ([]ArticleDTO, error) {
	data, err := model.Article{}.GetAll("is_deleted = false")
	if err != nil {
		return nil, err
	}

	var dataDtos []ArticleDTO

	for _, value := range data {
		newData := ArticleDTO{
			ID:          value.ID,
			Title:       value.Title,
			Slug:        value.Slug,
			Description: value.Description,
			Content:     value.Content,
			PictureUrl:  value.PictureUrl,
			CreatedDate: value.CreatedAt,
			IsDeleted:   value.IsDeleted,
		}
		dataDtos = append(dataDtos, newData)
	}

	return dataDtos, nil
}

func (a ArticleRepository) GetArticleById(id int) (ArticleDTO, error) {
	data, err := model.Article{}.Get(id)
	if err != nil {
		return ArticleDTO{}, err
	}

	result := ArticleDTO{
		ID:          data.ID,
		Title:       data.Title,
		Slug:        data.Slug,
		Description: data.Description,
		Content:     data.Content,
		PictureUrl:  data.PictureUrl,
		CreatedDate: data.CreatedAt,
		IsDeleted:   data.IsDeleted,
	}
	return result, nil
}

func (a ArticleRepository) GetArticleBySlug(slug string) (ArticleDTO, error) {
	data, err := model.Article{}.Get("slug = ?", slug)
	if err != nil {
		return ArticleDTO{}, err
	}

	result := ArticleDTO{
		ID:          data.ID,
		Title:       data.Title,
		Slug:        data.Slug,
		Description: data.Description,
		Content:     data.Content,
		PictureUrl:  data.PictureUrl,
		CreatedDate: data.CreatedAt,
		IsDeleted:   data.IsDeleted,
	}
	return result, nil
}

func (a ArticleRepository) CreateArticle(article *ArticleDTO) (ArticleDTO, error) {
	data, err := model.Article{
		Title:       article.Title,
		Slug:        article.Slug,
		Description: article.Description,
		Content:     article.Content,
		PictureUrl:  article.PictureUrl,
		IsDeleted:   false,
	}.Add()
	if err != nil {
		return ArticleDTO{}, err
	}

	return ArticleDTO{
		ID:          data.ID,
		Title:       data.Title,
		Slug:        data.Slug,
		Description: data.Description,
		Content:     data.Content,
		PictureUrl:  data.PictureUrl,
		CreatedDate: data.CreatedAt,
		IsDeleted:   data.IsDeleted}, nil
}

func (a ArticleRepository) UpdateArticle(article *ArticleDTO) (ArticleDTO, error) {
	data, err := model.Article{}.Get(article.ID)
	if err != nil {
		return ArticleDTO{}, err
	}

	result, err := data.Update(model.Article{
		Title:       article.Title,
		Slug:        article.Slug,
		Description: article.Description,
		Content:     article.Content,
		PictureUrl:  article.PictureUrl,
		IsDeleted:   false,
	})
	if err != nil {
		return ArticleDTO{}, err
	}

	return ArticleDTO{
		ID:          result.ID,
		Title:       result.Title,
		Slug:        result.Slug,
		Description: result.Description,
		Content:     result.Content,
		PictureUrl:  result.PictureUrl,
		CreatedDate: result.CreatedAt,
		IsDeleted:   result.IsDeleted}, nil
}

func (a ArticleRepository) DeleteArticle(id int) (ArticleDTO, error) {
	data, err := model.Article{}.Get(id)
	if err != nil {
		return ArticleDTO{}, err
	}

	result, err := data.Update(model.Article{
		IsDeleted: true,
	})
	if err != nil {
		return ArticleDTO{}, err
	}

	return ArticleDTO{
		ID:          result.ID,
		Title:       result.Title,
		Slug:        result.Slug,
		Description: result.Description,
		Content:     result.Content,
		PictureUrl:  result.PictureUrl,
		CreatedDate: result.CreatedAt,
		IsDeleted:   result.IsDeleted}, nil
}
