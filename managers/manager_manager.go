// package manager

// import (
// 	model "crudecho/models"
// 	"crudecho/request"

// 	service "crudecho/services"
// 	"errors"
// 	"log"
// )

// func HandleCreate(data interface{}, usePostgres bool) error {
// 	var err error

// 	// Switch over the data to handle different types
// 	switch record := data.(type) {
// 	case *request.CreateSiteManagerRequest:
// 		// Creating a Site Manager
// 		siteManager := &model.SiteManagerDetails{
// 			Name:  record.Name,
// 			Email: record.Email,
// 		}

// 		// Decide whether to use PostgreSQL or MongoDB
// 		if usePostgres {
// 			_, err = service.PGServiceCreate(siteManager)
// 		} else {
// 			_, err = service.MongoServiceCreate(siteManager)
// 		}

// 	case *request.CreateWorkerRequest:
// 		// Creating a Worker (handle Worker data here)
// 		worker := &model.Worker{
// 			Name:   record.Name,
// 			Role:   record.Role,
// 			SiteID: record.SiteID,
// 		}

// 		// Decide whether to use PostgreSQL or MongoDB
// 		if usePostgres {
// 			_, err = service.PGServiceCreate(worker)
// 		} else {
// 			_, err = service.MongoServiceCreate(worker)
// 		}

// 	default:
// 		log.Printf("Unknown type: %T", data)
// 		return errors.New("invalid data type")
// 	}

// 	// Check if an error occurred during record creation
// 	if err != nil {
// 		log.Printf("Error creating record: %v", err)
// 		return err
// 	}

//		return nil
//	}
package manager

import (
	"errors"
	"log"

	model "crudecho/models"
	"crudecho/request"
	service "crudecho/services"
)

// HandleCreate handles the creation of Site Managers and Workers.
func HandleCreate(data interface{}, usePostgres bool) error {
	var err error
	switch record := data.(type) {
	case *request.CreateSiteManagerRequest:
		siteManager := &model.SiteManagerDetails{
			Name:  record.Name,
			Email: record.Email,
		}
		if usePostgres {
			_, err = service.PGServiceCreate(siteManager)
		} else {
			_, err = service.MongoServiceCreate(siteManager)
		}
	case *request.CreateWorkerRequest:
		worker := &model.Worker{
			Name:   record.Name,
			Role:   record.Role,
			SiteID: record.SiteID,
		}
		if usePostgres {
			_, err = service.PGServiceCreate(worker)
		} else {
			_, err = service.MongoServiceCreate(worker)
		}
	default:
		log.Printf("Unknown type: %T", data)
		return errors.New("invalid data type")
	}
	if err != nil {
		log.Printf("Error creating record: %v", err)
		return err
	}
	return nil
}

// HandleGetAll retrieves all Site Managers and Workers.

// HandleGetAll retrieves all Site Managers and Workers.
func HandleGetAll(usePostgres bool) ([]interface{}, error) {
	if usePostgres {
		return service.PGServiceGetAll()
	}
	return service.MongoServiceGetAll()
}

// HandleGetById retrieves a record by ID.

// HandleUpdate updates a record by ID.
// func HandleUpdate(id int, data interface{}, usePostgres bool) error {
// 	var err error
// 	switch record := data.(type) {
// 	case *request.CreateSiteManagerRequest:
// 		siteManager := &model.SiteManagerDetails{
// 			Name:  record.Name,
// 			Email: record.Email,
// 		}
// 		if usePostgres {
// 			err = service.PGServiceUpdate(id, siteManager)
// 		} else {
// 			err = service.MongoServiceUpdate(id, siteManager)
// 		}
// 	case *request.CreateWorkerRequest:
// 		worker := &model.Worker{
// 			Name:   record.Name,
// 			Role:   record.Role,
// 			SiteID: record.SiteID,
// 		}
// 		if usePostgres {
// 			err = service.PGServiceUpdate(id, worker)
// 		} else {
// 			err = service.MongoServiceUpdate(id, worker)
// 		}
// 	default:
// 		log.Printf("Unknown type: %T", data)
// 		return errors.New("invalid data type")
// 	}
// 	if err != nil {
// 		log.Printf("Error updating record: %v", err)
// 		return err
// 	}
// 	return nil
// }

// HandleDelete deletes a record by ID.

// HandleDelete deletes a record by ID.
func HandleDelete(id interface{}, usePostgres bool) error {
	if usePostgres {
		return service.PGServiceDelete(id.(int))
	}
	return service.MongoServiceDelete(id.(string))
}

// HandleGetById retrieves a record by ID.

// HandleUpdate updates a record by ID.
func HandleUpdate(id interface{}, data interface{}, usePostgres bool) error {
	var err error
	switch record := data.(type) {
	case *request.CreateSiteManagerRequest:
		siteManager := &model.SiteManagerDetails{
			Name:  record.Name,
			Email: record.Email,
		}
		if usePostgres {
			err = service.PGServiceUpdate(id.(int), siteManager)
		} else {
			err = service.MongoServiceUpdate(id.(string), siteManager)
		}
	case *request.CreateWorkerRequest:
		worker := &model.Worker{
			Name:   record.Name,
			Role:   record.Role,
			SiteID: record.SiteID,
		}
		if usePostgres {
			err = service.PGServiceUpdate(id.(int), worker)
		} else {
			err = service.MongoServiceUpdate(id.(string), worker)
		}
	default:
		log.Printf("Unknown type: %T", data)
		return errors.New("invalid data type")
	}
	if err != nil {
		log.Printf("Error updating record: %v", err)
		return err
	}
	return nil
}

// HandleDelete deletes a record by ID.
