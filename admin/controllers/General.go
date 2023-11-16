package controllers

import (
	"blog/admin/helpers"
	"blog/admin/models"
	"fmt"
	"html/template"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/julienschmidt/httprouter"
)

type General struct{}

func (general General) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !helpers.CheckUser(w, r) {
		return
	}

	view, err := template.New("index").Funcs(template.FuncMap{}).ParseFiles(helpers.Include("setting/list")...)
	if err != nil {
		fmt.Println(err)
		return
	}

	data := make(map[string]interface{})
	data["Alert"] = helpers.GetAlert(w, r)
	view.ExecuteTemplate(w, "index", data)
}

func (general General) Add(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if !helpers.CheckUser(w, r) {
		return
	}

	if r.Method == http.MethodPost {
		err := r.ParseMultipartForm(10 << 20)
		if err != nil {
			fmt.Println(err)
			return
		}

		siteTitle := r.FormValue("name")
		footerTitle := r.FormValue("footer")

		logoFile, logoHandler, logoErr := r.FormFile("light")
		darkLogoFile, darkLogoHandler, darkLogoErr := r.FormFile("dark")
		faviconFile, faviconHandler, faviconErr := r.FormFile("favicon")

		if logoErr != nil || darkLogoErr != nil || faviconErr != nil {
			fmt.Println("Error uploading files:", logoErr, darkLogoErr, faviconErr)
			return
		}

		defer logoFile.Close()
		defer darkLogoFile.Close()
		defer faviconFile.Close()

		uploadsDir := "uploads/general"
		os.MkdirAll(uploadsDir, os.ModePerm)

		logoPath := filepath.Join(uploadsDir, logoHandler.Filename)
		darkLogoPath := filepath.Join(uploadsDir, darkLogoHandler.Filename)
		faviconPath := filepath.Join(uploadsDir, faviconHandler.Filename)

		saveFile(logoPath, logoFile)
		saveFile(darkLogoPath, darkLogoFile)
		saveFile(faviconPath, faviconFile)

		newSetting := models.General_Setting{
			SiteTitle:   siteTitle,
			FooterTitle: footerTitle,
			Logo:        logoPath,
			DarkLogo:    darkLogoPath,
			Favicon:     faviconPath,
		}
		newSetting.Add()

		http.Redirect(w, r, "/admin/setting", http.StatusSeeOther)
		return
	}

	view, err := template.New("add").Funcs(template.FuncMap{}).ParseFiles(helpers.Include("setting/add")...)
	if err != nil {
		fmt.Println(err)
		return
	}

	data := make(map[string]interface{})
	data["Alert"] = helpers.GetAlert(w, r)
	view.ExecuteTemplate(w, "add", data)
}

func saveFile(filePath string, file multipart.File) error {
	dir := filepath.Dir(filePath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return fmt.Errorf("Error creating directory: %v", err)
		}
	}

	out, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("Error creating file: %v", err)
	}
	defer out.Close()

	// Copy the file
	_, err = io.Copy(out, file)
	if err != nil {
		return fmt.Errorf("Error copying file: %v", err)
	}

	fmt.Printf("File %s uploaded successfully\n", filePath)
	return nil
}

func (general General) MyPage(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	latestSetting, err := models.GetLatestGeneralSetting()
	if err != nil {
		fmt.Println(err)
		return
	}

	tmpl, err := template.New("mypage").ParseFiles(`/admin/list`)
	if err != nil {
		fmt.Println(err)
		return
	}

	data := map[string]interface{}{
		"SiteTitle":   latestSetting.SiteTitle,
		"FooterTitle": latestSetting.FooterTitle,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		fmt.Println(err)
		return
	}
}
