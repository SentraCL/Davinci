package controllers

import (
	models "../models"
)

var sesionModel = models.SessionModel{}
var typeRefModel = models.TypeRefModel{}
var artifactModel = models.ArtifactModel{}
var inventionModel = models.InventionModel{}
var projectModel = models.ProjectModel{}

type EpicExportRequest struct {
	ID         string   `json:"id"`
	FormatDate string   `json:"formatDate"`
	Name       string   `json:"name"`
	Type       string   `json:"type"`
	Title      string   `json:"title"`
	EpicForm   EpicForm `json:"epicForm"`
}

type AttributesEpicExport struct {
	Name         string `json:"name"`
	ArtifactType string `json:"artifactType"`
	Value        string `json:"value"`
}

type Inventions struct {
	Code  string `json:"code"`
	Name  string `json:"name"`
	Value string `json:"value,omitempty"`
}

type UserStoriesRef struct {
	Code         string `json:"code"`
	Version      string `json:"version"`
	IndexVersion int    `json:"indexVersion"`
	ExecuteOrder int    `json:"executeOrder"`
}

type Activities struct {
	Code           string           `json:"code"`
	Inventions     []Inventions     `json:"inventions"`
	UserStoriesRef []UserStoriesRef `json:"userStoriesRef"`
}

type EpicForm struct {
	Reference  string       `json:"reference"`
	Attributes []AttributesEpicExport `json:"attributes"`
	Inventions []Inventions `json:"inventions"`
	Activities []Activities `json:"activities"`
}
