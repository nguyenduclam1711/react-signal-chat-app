package modules

import (
	"sync"

	"github.com/gofiber/fiber/v2"
)

type CreateModuleFunction func(app *fiber.App)

func CreateAllModules(app *fiber.App) {
	var wg sync.WaitGroup

	allCreateModuleFuncs := []CreateModuleFunction{
		CreateUserModule,
		CreateChatModule,
	}

	// increment the wait group counter
	wg.Add(len(allCreateModuleFuncs))

	for _, createModule := range allCreateModuleFuncs {
		// launch each module setup in a go routine
		go func(cm CreateModuleFunction) {
			// defer the counter when the go routine is complete
			defer wg.Done()
			cm(app)
		}(createModule)
	}

	// wait for all goroutines to complete
	wg.Wait()
}
