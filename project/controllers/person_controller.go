package controllers

import (
	"e-mobile/models"
	"e-mobile/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PersonController struct {
	DB  *gorm.DB
	API *services.API
}

func NewPersonController(db *gorm.DB) *PersonController {
	return &PersonController{DB: db, API: &services.API{}}
}

// CreatePerson godoc
// @Summary Создать нового человека
// @Description Создает нового человека с обогащенными данными (возраст, пол, национальность).
// @Accept json
// @Produce json
// @Param person body models.Person true "Данные человека"
// @Success 201 {object} models.Person "Человек успешно создан"
// @Router /api/people [post]
func (pc *PersonController) CreatePerson(c *gin.Context) {
	var input models.Person

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := pc.addPerson(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add person"})
		return
	}

	if err := pc.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save person"})
		return
	}

	c.JSON(http.StatusCreated, input)
}

// GetPeople godoc
// @Summary Получить всех людей
// @Description Возвращает список всех людей из базы данных.
// @Produce json
// @Success 200 {array} models.Person "Список людей"
// @Router /api/people [get]
func (pc *PersonController) GetPeople(c *gin.Context) {
	var people []models.Person
	if err := pc.DB.Find(&people).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch people"})
		return
	}
	c.JSON(http.StatusOK, people)
}

// GetPersonByID godoc
// @Summary Получить человека по ID
// @Description Возвращает данные человека по его уникальному ID.
// @Produce json
// @Param id path int true "ID человека"
// @Success 200 {object} models.Person "Данные человека"
// @Router /api/people/{id} [get]
func (pc *PersonController) GetPersonByID(c *gin.Context) {
	id := c.Param("id")
	var person models.Person
	if err := pc.DB.First(&person, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
		return
	}
	c.JSON(http.StatusOK, person)
}

// UpdatePerson godoc
// @Summary Обновить данные человека
// @Description Обновляет данные существующего человека.
// @Accept json
// @Produce json
// @Param id path int true "ID человека"
// @Param person body models.Person true "Обновленные данные человека"
// @Success 200 {object} models.Person "Данные успешно обновлены"
// @Router /api/people/{id} [put]
func (pc *PersonController) UpdatePerson(c *gin.Context) {
	id := c.Param("id")
	var person models.Person
	if err := pc.DB.First(&person, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
		return
	}

	var input models.Person
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pc.DB.Model(&person).Updates(input)
	c.JSON(http.StatusOK, person)
}

// DeletePerson godoc
// @Summary Удалить человека
// @Description Удаляет человека по его уникальному ID.
// @Produce json
// @Param id path int true "ID человека"
// @Success 200 {object} object "Человек успешно удален"
// @Router /api/people/{id} [delete]
func (pc *PersonController) DeletePerson(c *gin.Context) {
	id := c.Param("id")
	var person models.Person
	if err := pc.DB.First(&person, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
		return
	}

	pc.DB.Delete(&person)
	c.JSON(http.StatusOK, gin.H{"message": "Person deleted"})
}

func (pc *PersonController) addPerson(person *models.Person) error {
	age, err := pc.API.GetAge(person.Name)
	if err != nil {
		return err
	}
	person.Age = age

	gender, err := pc.API.GetGender(person.Name)
	if err != nil {
		return err
	}
	person.Gender = gender

	nationality, err := pc.API.GetNationality(person.Name)
	if err != nil {
		return err
	}
	person.Nationality = nationality

	return nil
}
