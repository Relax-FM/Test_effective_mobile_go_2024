package handler

import (
	"fmt"
	"net/http"
	"strconv"

	tem2024 "github.com/Relax-FM/Test_effective_mobile_go_2024"
	"github.com/gin-gonic/gin"
)

type getAllPagesMusicItemsResponse struct {
	Data []tem2024.PageMusicItemsResponse `json:"data"`
}

// @Summary 			Get all music
// @Tags 				music
// @Description 		Get all music in our library
// @ID 					get-all-music
// @Accept 				json
// @Produce 			json
// @Param 	sort_by 	query 	string 		false "Sorted by"				default(id)
// @Param 	desc 		query 	boolean 	false "Desc/asc"				default(true)
// @Param 	limit 		query 	int 		false "limit for pagination"	default(50)
// @Success 200 		{object} 			getAllPagesMusicItemsResponse
// @Failure 400,404 	{object} 			errorResponse
// @Failure 500 		{object} 			errorResponse
// @Failure default 	{object} 			errorResponse
// @Router 	/api/music 	[get]
func (h *Handler) getAllMusic(c *gin.Context) {
	var inputParams tem2024.QueryParams
	if err := c.Bind(&inputParams); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	allMusic, err := h.services.MusicItem.GetAllMusic(inputParams)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllPagesMusicItemsResponse{
		Data: allMusic,
	})
}

type getMusicTextResponse struct {
	Text []tem2024.CoupletMusicText `json:"text"`
}

// @Summary 			Get music text
// @Tags 				music
// @Description 		Return music text with pagination
// @ID 					get-music-text
// @Accept 				json
// @Produce 			json
// @Param 	id 			path 		int 	true 	"Music ID"
// @Success 200 		{object} 	getMusicTextResponse
// @Failure 400,404 	{object} 	errorResponse
// @Failure 500 		{object} 	errorResponse
// @Failure default 	{object} 	errorResponse
// @Router 	/api/music/{id} 		[get]
func (h *Handler) getMusicText(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	couplets, err := h.services.MusicItem.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getMusicTextResponse{
		Text: couplets,
	})
}

// @Summary 			Delete music item
// @Tags 				music
// @Description 		Delete music item by id
// @ID 					delete-music-item
// @Accept 				json
// @Produce 			json
// @Param 	id 	path 	int 	true 	"Music ID"
// @Success 200 		{object} 	statusResponse
// @Failure 400,404 	{object} 	errorResponse
// @Failure 500 		{object} 	errorResponse
// @Failure default 	{object} 	errorResponse
// @Router 	/api/music/{id} 		[delete]
func (h *Handler) deleteMusicItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.MusicItem.Delete(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: fmt.Sprintf("Element with id %d deleted successfully", id),
	})
}

// @Summary 			Update music
// @Tags 				music
// @Description 		Update music info by id
// @ID 					update-music
// @Accept 				json
// @Produce 			json
// @Param 	id 		path 	int 						true 	"Music ID"
// @Param 	input 	body 	tem2024.UpdateMusicInput 	true 	"Update info"
// @Success 200 	{object} 	statusResponse
// @Failure 400,404 {object} 	errorResponse
// @Failure 500 	{object} 	errorResponse
// @Failure default {object} 	errorResponse
// @Router 	/api/music/{id} 	[put]
func (h *Handler) updateMusicItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input tem2024.UpdateMusicInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.MusicItem.Update(id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: fmt.Sprintf("Element with id %d updated successfully", id),
	})
}

// @Summary 			Create music item
// @Tags 				music
// @Description 		create todo list
// @ID 					create-list
// @Accept 				json
// @Produce 			json
// @Param 	input 	body 		tem2024.CreateMusicInput 	true 	"creating music info"
// @Success 200 	{object} 	statusResponse
// @Failure 400,404 {object} 	errorResponse
// @Failure 500 	{object} 	errorResponse
// @Failure default {object} 	errorResponse
// @Router 		/api/music 		[post]
func (h *Handler) addMusicItem(c *gin.Context) {
	var input tem2024.CreateMusicInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.MusicItem.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: fmt.Sprintf("Element with id %d added successfully", id),
	})
}