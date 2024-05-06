package repository

type createRepoFunc func()

func CreateAllRepositories() {
	allCreateRepositoriesFuncs := []createRepoFunc{
		NewUserRepository,
		NewUserCredentialRepository,
	}

	for _, function := range allCreateRepositoriesFuncs {
		go function()
	}
}
