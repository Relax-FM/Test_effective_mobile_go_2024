package handler

import (
	tem2024 "github.com/Relax-FM/Test_effective_mobile_go_2024"
	"github.com/gin-gonic/gin"
)

type getAllMusicItemsResponse struct {
	Data []tem2024.MusicItem `json:"data"`
}

// @Summary Get all music
// @Tags music
// @Description Get all music in our library
// @ID get-all-music
// @Accept json
// @Produce json
// @Param sort_param query string false "Sorted by"
// @Param desc query boolean false "Desc/asc"
// @Param limit query int false "limit for pagination"
// @Success 200 {object} getAllMusicItemsResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/music [get]
func (h *Handler) getAllMusic(c *gin.Context) {

}

type getMusicTextResponse struct {
	Text []tem2024.MusicText `json:"text"`
}

// @Summary Get music text
// @Tags music
// @Description REturn music text with pagination
// @ID get-music-text
// @Accept json
// @Produce json
// @Param id path int true "Music ID"
// @Success 200 {object} getMusicTextResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/music/{id} [get]
func (h *Handler) getMusicText(c *gin.Context) {
	
}

// @Summary Delete music item
// @Tags music
// @Description Delete music item by id
// @ID delete-music-item
// @Accept json
// @Produce json
// @Param id path int true "Music ID"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/music/{id} [delete]
func (h *Handler) deleteMusicItem(c *gin.Context) {
	
}

// @Summary Update music
// @Tags music
// @Description Update music info by id
// @ID update-music
// @Accept json
// @Produce json
// @Param id path int true "Music ID"
// @Param input body tem2024.UpdateMusicInput
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/music/{id} [put]
func (h *Handler) updateMusicItem(c *gin.Context) {
	
}

// @Summary Create music item
// @Tags music
// @Description create todo list
// @ID create-list
// @Accept json
// @Produce json
// @Param input body tem2024.CreateMusicInput true "Music info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/music [post]
func (h *Handler) addMusicItem(c *gin.Context) {
	
}