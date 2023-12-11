package handler

import (
	"fmt"
	"log"
	"net/http"

	"forum/internal/models"
)

func (h *Handler) createPostVotePOST(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/postvote/create" {
		log.Printf("createPostVotePOST: not found %s\n", r.URL.Path)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound) // 404
		return
	}
	if r.Method != http.MethodPost {
		log.Printf("createPostVotePOST: method not allowed %s\n", r.Method)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed) // 405
		return
	}
	if err := r.ParseForm(); err != nil {
		log.Printf("createPostVotePOST:ParseForm:%s\n", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError) // 500
		return
	}

	vote, err := h.getVote(r.Form.Get("vote"))
	if err != nil {
		log.Printf("createPostVotePOST:getVote:%s\n", err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest) // 400
		return
	}
	postId, err := h.getPostIdFromForm(r.Form.Get("post_id"))
	if err != nil {
		log.Printf("createPostVotePOST:getPostIdFromForm:%s\n", err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest) // 400
		return
	}

	user := h.getUserFromContext(r)

	newVote := &models.PostVote{
		PostId: postId,
		UserId: user.Id,
		Vote:   vote,
	}

	err = h.service.CreatePostVote(newVote)
	if err != nil {
		log.Printf("createPostVotePOST:CreatePostVote:%s\n", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError) // 500
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/post/%d", postId), http.StatusSeeOther) // 303
}