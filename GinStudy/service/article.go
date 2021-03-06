/*
* @Time    : 2020-11-24 10:33
* @Author  : CoderCharm
* @File    : article.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :
**/
package service

import (
	"gin_study/global"
	"gin_study/model"
	"gin_study/model/request"
)

/*
获取推荐文章
*/
func FetchRecommendArticleList() (RecommendArticleList []model.RecommendArticle) {
	// 执行原生SQL语句
	//global.GIN_DB.Raw("select id,title,content from article").Scan(&article)

	global.GIN_DB.Table("article_article").
		Where("is_recommend = ?", "1").
		Order("update_time DESC").Limit(5).
		Scan(&RecommendArticleList)

	return RecommendArticleList
}

/*
获取文章列表
*/
func FetchArticleIndexList(Category request.ArticleCategory) (ArticleIndexList []model.ArticleIndex, total int64) {
	limit := Category.PageSize
	offset := Category.PageSize * (Category.Page - 1)

	cateId := Category.CateId

	if cateId == 0 {
		global.GIN_DB.Table("article_article").Where("is_open = ?", "1").Count(&total)
		global.GIN_DB.Table("article_article").Where("is_open = ?", "1").
			Order("add_time DESC").Limit(limit).Offset(offset).Scan(&ArticleIndexList)
	} else {
		global.GIN_DB.Table("article_article").
			Where("is_open = ?", "1").
			Where("category_id = ?", cateId).
			Count(&total)
		global.GIN_DB.Table("article_article").
			Where("is_open = ?", "1").Where("category_id = ?", cateId).
			Order("add_time DESC").Limit(limit).
			Offset(offset).Scan(&ArticleIndexList)
	}

	return ArticleIndexList, total

}

/*
获取所有分类
*/
func FetchCategoryList() (CategoryList []model.Category) {

	global.GIN_DB.Table("article_category").
		Where("active = ?", "1").
		Scan(&CategoryList)
	return CategoryList
}

/*
获取文章详情
*/
func FetchArticleDetail(DetailHref request.ArticleDetail) (DetailArticle model.ArticleDetail) {
	href := DetailHref.Href
	global.GIN_DB.Table("article_article").
		Where("is_open = ?", "1").
		Where("article_url = ?", href).
		Scan(&DetailArticle)
	return DetailArticle
}
