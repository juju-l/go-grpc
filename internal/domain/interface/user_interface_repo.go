package interfacepri

type UserRepoistory interface {
	GetCrudRepo()
	Add()
	Delete()
	Update()
	Select()
	GetQueryableByPage()
	ConvertPages()
}
