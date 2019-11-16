package main

import (
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func streamHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("vid-id")
	vl := VIDEO_DIR + vid
	video, err := os.Open(vl)
	defer video.Close()
	if err != nil {
		log.Printf(err.Error())
		sendErrorResponse(w, http.StatusInternalServerError, "Internal serve error")
		return
	}

	w.Header().Set("Content-Type", "video/mp4")
	http.ServeContent(w, r, "", time.Now(), video)
}

func uploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)
	if e := r.ParseMultipartForm(MAX_UPLOAD_SIZE); e != nil {
		sendErrorResponse(w, http.StatusBadRequest, "File is to big")
		return
	}
	// todo: 验证
	file, _, err := r.FormFile("file")
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, "InternalServerError")
		return
	}
	data, e := ioutil.ReadAll(file)
	if e != nil {
		log.Printf("Read file error: %v", e)
		sendErrorResponse(w, http.StatusInternalServerError, "InternalServerError")
		return
	}

	filename := p.ByName("vid-id")
	e = ioutil.WriteFile(VIDEO_DIR + filename, data, 0666)
	if e != nil {
		log.Printf(e.Error())
		sendErrorResponse(w, http.StatusInternalServerError, "InternalServerError")
		return
	}
	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "Uploaded successful")
}