package render

import (
	"bytes"
	"html/template"
	"log"
	"github.com/ElmerMenjivar1999/bookings/pkg/config"
	//"myapp/pkg/handlers"
	"github.com/ElmerMenjivar1999/bookings/pkg/models"
	"net/http"
	"path/filepath"
)

// var functions = template.FuncMap{

// }
var app *config.AppConfig

//NewTemplates sets the config for he template package
func NewTemplates(a *config.AppConfig){
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData{
	
	
	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string,td *models.TemplateData){
	var tc map[string]*template.Template
	if app.UseCache {
		//get the template cache from the app config
		tc = app.TemplateCache
		log.Println("Obteniendo el template del cache")
	}else{
		
		tc,_ = CreateTemplateCache()
	}
	
	//Create a template cache
	// tc,err := CreateTemplateCache()
	// if err != nil{
	// 	log.Fatal(err)
	// }
	//get requested template from cache
	t,ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}
	buf := new(bytes.Buffer)
	td = AddDefaultData(td)

	_ = t.Execute(buf,td)
	
	//render the template
	_,err := buf.WriteTo(w)
	if err != nil{
		log.Println(err)
	}
	// parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl, "./templates/base.html")
	// err := parsedTemplate.Execute(w,nil)
	// if err != nil {
	// 	fmt.Println("error parsing template",err)
	// 	return
	// }

}

//more complex template cache
func CreateTemplateCache() (map[string]*template.Template,error){
	//myCache := make(map[string]*template.Template)
	myCache := map[string]*template.Template{}

	//get all of the files named *.page.html from ./templates
	pages,err := filepath.Glob("./templates/*.html")
	
	if err != nil{
		return myCache,err
	}

	//range through all files ending with *.html
	for _, page := range pages{
		name := filepath.Base(page)
		
		ts,err := template.New(name).ParseFiles(page)
		
		if err != nil{
			return myCache,err
		}
		matches,err := filepath.Glob("./templates/base.html")
		
		if err != nil{
			return myCache,err
		}
		if len(matches) > 0{
			ts,err = ts.ParseGlob("./templates/base.html")
			if err != nil{
				return myCache,err
			}

		}
		myCache[name] = ts
		
	}
	return myCache,nil
}





//simple template cache
// var tc = make(map[string]*template.Template)


// func RenderTemplate(w http.ResponseWriter, t string){
// 	var tmpl *template.Template
// 	var err error
// 	//check to see if we already have the template in our cache
// 	_, inMap := tc[t]
// 	if !inMap{
// 		//need to create the template
// 		log.Println("Creating template and adding to cache")
// 		err = createTemplateCache(t)
// 		if err != nil{
// 			log.Println(err)
// 		}
// 	} else{
// 		//we have the template in our cache
// 		log.Println("Using cache template")
// 	}
// 	tmpl = tc[t]

// 	err = tmpl.Execute(w,nil)
// 	if err != nil{
// 		log.Println(err)
// 	}

// }

// func createTemplateCache(t string) error {
// 	templates := []string{
// 		fmt.Sprintf("./templates/%s",t),
// 		"./templates/base.html",
// 	}

// 	//parse the template
// 	tmpl,err := template.ParseFiles(templates...)
// 	if err != nil {
// 		return err
// 	}

// 	//add template to cache
// 	tc[t] = tmpl
// 	return nil
//}
