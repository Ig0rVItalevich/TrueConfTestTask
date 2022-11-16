package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
	refactoring "refactoring/models"
	"strconv"
)

func (h *Handler) getAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.services.GetAll()
	if err != nil {
		RenderErrResponse(w, r, http.StatusInternalServerError, err)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, users)
}

func (h *Handler) createUser(w http.ResponseWriter, r *http.Request) {
	var input refactoring.User
	if err := render.Bind(r, &input); err != nil {
		RenderErrResponse(w, r, http.StatusBadRequest, err)
		return
	}

	id, err := h.services.Create(input)
	if err != nil {
		RenderErrResponse(w, r, http.StatusInternalServerError, err)
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, map[string]interface{}{
		"user_id": id,
	})
}

func (h *Handler) getUserByid(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		RenderErrResponse(w, r, http.StatusBadRequest, err)
		return
	}

	user, err := h.services.GetById(int(idInt))
	if err != nil {
		RenderErrResponse(w, r, http.StatusInternalServerError, err)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, user)
}

func (h *Handler) updateUser(w http.ResponseWriter, r *http.Request) {
	var input refactoring.UpdateUser

	if err := render.Bind(r, &input); err != nil {
		RenderErrResponse(w, r, http.StatusBadRequest, err)
		return
	}

	id := chi.URLParam(r, "id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		RenderErrResponse(w, r, http.StatusBadRequest, err)
		return
	}

	if err := h.services.Update(int(idInt), input); err != nil {
		RenderErrResponse(w, r, http.StatusInternalServerError, err)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) deleteUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		RenderErrResponse(w, r, http.StatusBadRequest, err)
		return
	}

	if err := h.services.Delete(int(idInt)); err != nil {
		RenderErrResponse(w, r, http.StatusInternalServerError, err)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, map[string]interface{}{
		"status": "ok",
	})
}
