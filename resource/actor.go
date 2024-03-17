package resource

import (
	"fmt"
	"musical_wiki/config"
	"musical_wiki/models"
)

type ActorResource struct {
	Model models.Actor
}

type ActorSliceResource struct {
	ModelSlice []models.Actor
}

func (resource *ActorResource) ToMap() map[string]interface{} {
	actor := map[string]interface{}{
		"id":             resource.Model.Id,
		"name":           resource.Model.Name,
		"translatedName": resource.Model.TranslatedName,
		"nickName":       resource.Model.NickName,
		"nationality":    resource.Model.Nationality,
		"born":           resource.Model.Born,
		"avatar":         fmt.Sprint(config.Endpoint, "/", resource.Model.Avatar.ImageName),
		"content":        resource.Model.Content,
		"credits":        resource.Model.Credits,
		"socials":        resource.Model.Socials,
	}
	gallery := make([]map[string]interface{}, len(resource.Model.Gallery))
	for i, g := range resource.Model.Gallery {
		gallery[i] = map[string]interface{}{
			"id":  g.Id,
			"url": fmt.Sprint(config.Endpoint, "/", g.ImageName),
		}
	}
	actor["gallery"] = gallery
	return actor
}

func (resource *ActorSliceResource) ToSliceResource() []map[string]interface{} {
	result := make([]map[string]interface{}, len(resource.ModelSlice))
	for i, m := range resource.ModelSlice {
		result[i] = map[string]interface{}{
			"id":             m.Id,
			"name":           m.Name,
			"translatedName": m.TranslatedName,
		}
	}
	return result
}
