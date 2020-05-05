package models

import "time"

//User : Definicion de Collection User para DB
type User struct {
	UserName     string
	Password     string
	LastTime     time.Time
	HistoryLogin []Session
}

//Session : Instancias de Sesion
type Session struct {
	RemoteAddr  string
	Date        time.Time
	DavinciCode string
	Online      bool
}

//Contact : Referencias de un Contacto.
type Contact struct {
	FullName string `bson:"fullname"`
	Email    string `bson:"email"`
}

//DataInvention : Inventos con Data
type DataInvention struct {
	KeyValue    string    `bson:"_id"`
	ContentJSON string    `bson:"content"`
	Update      time.Time `bson:"update"`
}

//Repository :  Repositorio de inventos con contenido.
type Repository struct {
	InventionCode string          `bson:"_id"`
	Invention     string          `bson:"name"`
	KeyRef        string          `bson:"key"`
	LastTime      time.Time       `bson:"lastTime"`
	WareHouse     []DataInvention `bson:"data"`
}

//Project : Instancias de Sesion
type Project struct {
	Code          string        `bson:"_id"`
	Author        string        `bson:"author"`
	Name          string        `bson:"name"`
	Alias         string        `bson:"alias"`
	Resume        string        `bson:"resume"`
	Company       string        `bson:"company"`
	Date          time.Time     `bson:"date"`
	Administrator Contact       `bson:"administrator"`
	Avatar64      string        `bson:"avatar64"`
	Repository    []Repository  `bson:"repository"`
	Users         []UserProject `bson:"users"`
	DataStored    DataStored    `bson:"dataStored"`
	Epics         Epics         `bson:"epics"`
	UserStories   UserStories   `bson:"userStories"`
}

//UserProject : usuario de un proyecto
type UserProject struct {
	Code     string `bson:"_id"`
	UserName string `bson:"user"`
	Password string `bson:"password"`
}

//UserStories : Historias de usuarios del proyecto y tipos de epicos
type UserStories struct {
	Repository []UserStory       `bson:"repository"`
	Types      []UserStoriesType `bson:"types"`
}

//Epics : Epicos del proyecto y tipos de epicos
type Epics struct {
	Repository []Epic     `bson:"repository"`
	Types      []EpicType `bson:"types"`
}

//Invention : Inventos para crear contexto a historias de usuario.
type Invention struct {
	Code      string     `bson:"_id"`
	Name      string     `bson:"name"`
	KeyLabel  string     `bson:"key_label"`
	Icon      string     `bson:"icon"`
	KeyValue  string     `bson:"key_value"`
	Date      time.Time  `bson:"date"`
	Artifacts []Artifact `bson:"artifacts"`
}

//Artifact : Artefactos para crear inventos
type Artifact struct {
	Code     string  `bson:"_id"`
	Name     string  `bson:"name"`
	NickName string  `bson:"nickname"`
	TypeRef  TypeRef `bson:"type_ref"`
}

//TypeRef : Valores Basicos de Datos : text,number,date,bool,json
type TypeRef struct {
	Code        string `bson:"_id"`
	Name        string `bson:"name"`
	IsEssential bool   `bson:"is_essential"`
	IsList      bool   `bson:"is_list"`
}

//InventionVO : Representacion Front-END de inventos
type InventionVO struct {
	Code      string          `json:"code"`
	Name      string          `json:"name"`
	Artifacts []ArtifactVO    `json:"artifacts"`
	Title     string          `json:"title"`
	Subtitle  string          `json:"subtitle"`
	Slots     string          `json:"slots"`
	Icon      string          `json:"icon"`
	KeyLabel  string          `json:"keyLabel"`
	KeyValue  string          `json:"keyValue"`
	Data      []DataInvention `json:"WareHouse"`
}

//ArtifactVO : Representacion Front-END de artefactos
type ArtifactVO struct {
	Name string `json:"name"`
	Code string `json:"code"`

	TypeName string `json:"typeName"`
	NickName string `json:"nickName"`

	TypeRef string `json:"typeRef"`

	IsEssential bool `json:"isEssential"`
	IsList      bool `json:"isList"`

	IsText    bool `json:"isText,omitempty"`
	IsNumeric bool `json:"isNumeric,omitempty"`
	IsDate    bool `json:"isDate,omitempty"`
	IsBool    bool `json:"isBool,omitempty"`
	IsJSON    bool `json:"isJson,omitempty"`

	List []struct {
		Code string `json:"code"`
		Name string `json:"name"`
	} `json:"list,omitempty"`
}

//EpicType , Epicos
type EpicType struct {
	Code       string          `bson:"_id"`
	Name       string          `json:"name,omitempty" bson:"name"`
	Definition string          `json:"definition,omitempty" bson:"definition,omitempty"`
	Reference  string          `json:"reference,omitempty" bson:"reference,omitempty"`
	Attributes []AttributeEpic `json:"attributes" bson:"attributes"`
	Inventions []InventionEpic `json:"inventions" bson:"inventions"`
}

//AttributeEpic , atributos del epico
type AttributeEpic struct {
	Name         string `json:"name" bson:"name"`
	ArtifactType string `json:"artifactType" bson:"artifactType"`
	Value        string `json:"value" bson:"value"`
}

//InventionEpic , inventos para epico
type InventionEpic struct {
	Code  string `json:"code" bson:"code"`
	Name  string `json:"name" bson:"name"`
	Value string `json:"value,omitempty" bson:"value,omitempty"`
}

//EpicTypeRequest , Epicos
type EpicTypeRequest struct {
	Name        string          `json:"name"`
	ProjectCode string          `json:"projectCode"`
	Definition  string          `json:"definition"`
	Reference   string          `json:"reference"`
	Attributes  []AttributeEpic `json:"attributes"`
	Inventions  []InventionEpic `json:"inventions"`
}

//UserStoriesType . Tipos de Historias de Usuario
type UserStoriesType struct {
	Code          string         `json:"code" bson:"_id"`
	Title         string         `json:"title"  bson:"title"`
	PreConditions []PreCondition `json:"preConditions"  bson:"preconditions"`
	CodeFormat    []string       `json:"codeFormat"  bson:"codeformat"`
	Update        time.Time      `json:"update"  bson:"update"`
}

//ArtifactUS , referencias de artefactos para historias de usuario tipo.
type ArtifactUS struct {
	Name        string `json:"name"  bson:"name"`
	Code        string `json:"code" bson:"code"`
	TypeName    string `json:"typeName"  bson:"typename"`
	NickName    string `json:"nickName" bson:"nickname"`
	TypeRef     string `json:"typeRef"  bson:"typeref"`
	IsEssential bool   `json:"isEssential"  bson:"isessential"`
	IsList      bool   `json:"isList" bson:"islist"`
	IsText      bool   `json:"isText,omitempty" bson:"istext"`
	IsJSON      bool   `json:"isJson,omitempty" bson:"isjson"`
}

//Parent ,artefacto padre del cual depende el atributo
type Parent struct {
	ID           string `json:"id" bson:"ID"`
	ArtifactName string `json:"artifactName" bson:"ArtifactName"`
}

//PreCondition , Precondiciones de un historia de usuario.
type PreCondition struct {
	ID        string       `json:"id" bson:"id"`
	Code      string       `json:"code" bson:"code"`
	Artifacts []ArtifactUS `json:"artifacts" bson:"artifacts"`
	Name      string       `json:"name" bson:"name"`
	Label     string       `json:"label" bson:"label"`
	Parent    Parent       `json:"parent,omitempty" bson:"parent"`
	Value     string       `json:"value" bson:"value"`
}

//Fields , Atributos de la Historia de Usuario.
type Fields struct {
	Description    string `json:"description" bson:"description"`
	ExpectedResult string `json:"expectedResult" bson:"expectedResult"`
}

//UserStory , Historias de USuario.
type UserStory struct {
	ID               string         `json:"id" bson:"id"`
	Code             UsCode         `json:"code" bson:"code"`
	Type             string         `json:"type" bson:"type"`
	LastVersionIndex int            `json:"lastVersionIndex" bson:"lastVersionIndex"`
	Versions         []SubUserStory `json:"versions" bson:"versions"`
}

//SubUserStory , Historias de USuario.
type SubUserStory struct {
	Version       string         `json:"version" bson:"version"`
	Fields        Fields         `json:"fields" bson:"fields"`
	PreConditions []PreCondition `json:"preConditions" bson:"preConditions"`
	Steps         []Step         `json:"steps"  bson:"steps"`
	Date          time.Time      `json:"date" bson:"date"`
	Author        string         `json:"author" bson:"author"`
}

//UsCode historia de usuario codigo.
type UsCode struct {
	Value             string   `json:"value" bson:"value"`
	Format            []string `json:"format" bson:"format"`
	InventionSequence string   `json:"inventionSequence" bson:"inventionSequence"`
}

//PreConditionQuery , Query de busqueda de casos similares por precondiciones
type PreConditionQuery struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}

//Step ,  Pasos
type Step struct {
	Step   string `json:"step" bson:"step"`
	Result string `json:"result" bson:"result"`
	ID     int    `json:"id" bson:"id"`
}

//Epic : Epicos del proyecto
type Epic struct {
	ID           string       `json:"id" bson:"_id"`
	EpicTypeCode string       `json:"type" bson:"epictype"`
	Code         string       `json:"code" bson:"code"`
	Attributes   []Attributes `json:"attributes" bson:"attributes"`
	Author       string       `json:"author" bson:"author"`
	Activities   []Activities `json:"activities" bson:"activities"`
	Date         time.Time    `json:"date" bson:"date"`
	Save         bool         `json:"save" bson:"-"`
}

//Attributes , atributos del Epico
type Attributes struct {
	Name  string `json:"name" bson:"name"`
	Code  string `json:"code" bson:"code"`
	Value string `json:"value" bson:"value"`
}

//InventionsContext , filtro de actividades
type InventionsContext struct {
	Name  string   `json:"name" bson:"name"`
	Code  string   `json:"code" bson:"code"`
	Value []string `json:"value" bson:"value"`
}

//UserStoriesRef , historias de usuarios relacionadas
type UserStoriesRef struct {
	Code         string `json:"code" bson:"code"`
	Version      string `json:"version" bson:"version"`
	IndexVersion int    `json:"indexVersion" bson:"indexVersion"`
	ExecuteOrder int    `json:"executeOrder" bson:"executeOrder"`
}

//Activities , Actividades del EPICO
type Activities struct {
	Code           string              `json:"code" bson:"code"`
	Inventions     []InventionsContext `json:"inventions" bson:"inventionsContext"`
	UserStoriesRef []UserStoriesRef    `json:"userStoriesRef" bson:"userStoriesRef"`
}

//DataStored , Repositorio de Datos de Prueba
type DataStored struct {
	DataTypes  []DataType `json:"dataTypes" bson:"dataTypes"`
	Repository []TestData `json:"repository" bson:"repository"`
}

//DataType , es la definicion del tipo de dato a cargar.
type DataType struct {
	ID           string    `json:"id" bson:"_id"`
	PrincipalKey string    `json:"principalKey" bson:"principalKey"`
	Name         string    `json:"name" bson:"name"`
	Date         time.Time `json:"date" bson:"date"`
	Description  string    `json:"description" bson:"description"`
	Icon         string    `json:"icon64" bson:"icon64"`
}

//TestData , Almacen de Datos
type TestData struct {
	ID        string            `json:"id" bson:"_id"`
	DataType  DataType          `json:"dataType" bson:"dataType"`
	Atributes map[string]string `json:"atributes" bson:"atributes"`
	Date      time.Time         `json:"date" bson:"date"`
	Owner     string            `json:"owner" bson:"owner"`
}
