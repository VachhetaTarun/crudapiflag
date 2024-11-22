// package controller

// import (
// 	manager "crudecho/managers"
// 	"crudecho/request"
// 	"encoding/json"
// 	"io/ioutil"

// 	"log"
// 	"net/http"

// 	"github.com/labstack/echo/v4"
// )

// // Create handles the creation of either a Site Manager or a Worker.
// func Create(c echo.Context) error {
// 	usePostgres := c.QueryParam("use_postgres") == "true"
// 	log.Printf("usePostgres: %v", usePostgres) // Debug log

// 	// Read raw body content
// 	body, err := ioutil.ReadAll(c.Request().Body)
// 	if err != nil {
// 		log.Printf("Error reading request body: %v", err)
// 		return c.JSON(http.StatusBadRequest, "Invalid data")
// 	}
// 	log.Printf("Raw request body: %s", body)

// 	// Attempt to bind the incoming data as a SiteManager
// 	var siteManagerRequest request.CreateSiteManagerRequest
// 	if err := json.Unmarshal(body, &siteManagerRequest); err == nil {
// 		if siteManagerRequest.Name != "" && siteManagerRequest.Email != "" {
// 			log.Printf("Site Manager payload: %+v", siteManagerRequest)

// 			// Create Site Manager (PostgreSQL/MongoDB based on usePostgres)
// 			if err := manager.HandleCreate(&siteManagerRequest, usePostgres); err != nil {
// 				log.Printf("Error in HandleCreate: %v", err)
// 				return c.JSON(http.StatusInternalServerError, "Failed to create Site Manager")
// 			}
// 			return c.JSON(http.StatusOK, "Site Manager created successfully")
// 		}
// 	} else {
// 		log.Printf("Error unmarshalling Site Manager request: %v", err)
// 	}

// 	// If binding fails or is incomplete for SiteManager, try Worker
// 	var workerRequest request.CreateWorkerRequest
// 	if err := json.Unmarshal(body, &workerRequest); err == nil {
// 		if workerRequest.Name != "" && workerRequest.Role != "" && workerRequest.SiteID != 0 {
// 			log.Printf("Worker payload: %+v", workerRequest)

// 			// Create Worker (PostgreSQL/MongoDB based on usePostgres)
// 			if err := manager.HandleCreate(&workerRequest, usePostgres); err != nil {
// 				log.Printf("Error in HandleCreate: %v", err)
// 				return c.JSON(http.StatusInternalServerError, "Failed to create Worker")
// 			}
// 			return c.JSON(http.StatusOK, "Worker created successfully")
// 		}
// 	} else {
// 		log.Printf("Error unmarshalling Worker request: %v", err)
// 	}

// 	// If binding for both fails, return an invalid data error
// 	return c.JSON(http.StatusBadRequest, "Invalid data")
// }

package controller

import (
	manager "crudecho/managers"
	"crudecho/request"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Create handles the creation of either a Site Manager or a Worker.
func Create(c echo.Context) error {
	usePostgres := c.QueryParam("use_postgres") == "true"
	log.Printf("usePostgres: %v", usePostgres) // Debug log

	// Read raw body content
	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Printf("Error reading request body: %v", err)
		return c.JSON(http.StatusBadRequest, "Invalid data")
	}
	log.Printf("Raw request body: %s", body)

	// Attempt to bind the incoming data as a SiteManager
	var siteManagerRequest request.CreateSiteManagerRequest
	if err := json.Unmarshal(body, &siteManagerRequest); err == nil {
		if siteManagerRequest.Name != "" && siteManagerRequest.Email != "" {
			log.Printf("Site Manager payload: %+v", siteManagerRequest)

			// Create Site Manager (PostgreSQL/MongoDB based on usePostgres)
			if err := manager.HandleCreate(&siteManagerRequest, usePostgres); err != nil {
				log.Printf("Error in HandleCreate: %v", err)
				return c.JSON(http.StatusInternalServerError, "Failed to create Site Manager")
			}
			return c.JSON(http.StatusOK, "Site Manager created successfully")
		}
	} else {
		log.Printf("Error unmarshalling Site Manager request: %v", err)
	}

	// If binding fails or is incomplete for SiteManager, try Worker
	var workerRequest request.CreateWorkerRequest
	if err := json.Unmarshal(body, &workerRequest); err == nil {
		if workerRequest.Name != "" && workerRequest.Role != "" && workerRequest.SiteID != 0 {
			log.Printf("Worker payload: %+v", workerRequest)

			// Create Worker (PostgreSQL/MongoDB based on usePostgres)
			if err := manager.HandleCreate(&workerRequest, usePostgres); err != nil {
				log.Printf("Error in HandleCreate: %v", err)
				return c.JSON(http.StatusInternalServerError, "Failed to create Worker")
			}
			return c.JSON(http.StatusOK, "Worker created successfully")
		}
	} else {
		log.Printf("Error unmarshalling Worker request: %v", err)
	}

	// If binding for both fails, return an invalid data error
	return c.JSON(http.StatusBadRequest, "Invalid data")
}

// GetAll fetches all Site Managers and Workers.
func GetAll(c echo.Context) error {
	usePostgres := c.QueryParam("use_postgres") == "true"
	records, err := manager.HandleGetAll(usePostgres)
	if err != nil {
		log.Printf("Error in HandleGetAll: %v", err)
		return c.JSON(http.StatusInternalServerError, "Failed to fetch records")
	}
	return c.JSON(http.StatusOK, records)
}

// GetById fetches a Site Manager or Worker by ID.

// Update updates a Site Manager or Worker by ID.
// func Update(c echo.Context) error {
// 	usePostgres := c.QueryParam("use_postgres") == "true"
// 	id := c.Param("id")
// 	intId, err := strconv.Atoi(id)
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, "Invalid ID")
// 	}

// 	// Read raw body content
// 	body, err := ioutil.ReadAll(c.Request().Body)
// 	if err != nil {
// 		log.Printf("Error reading request body: %v", err)
// 		return c.JSON(http.StatusBadRequest, "Invalid data")
// 	}
// 	log.Printf("Raw request body: %s", body)

// 	// Attempt to bind the incoming data as a SiteManager
// 	var siteManagerRequest request.CreateSiteManagerRequest
// 	if err := json.Unmarshal(body, &siteManagerRequest); err == nil {
// 		if siteManagerRequest.Name != "" && siteManagerRequest.Email != "" {
// 			log.Printf("Site Manager payload: %+v", siteManagerRequest)

// 			// Update Site Manager (PostgreSQL/MongoDB based on usePostgres)
// 			if err := manager.HandleUpdate(intId, &siteManagerRequest, usePostgres); err != nil {
// 				log.Printf("Error in HandleUpdate: %v", err)
// 				return c.JSON(http.StatusInternalServerError, "Failed to update Site Manager")
// 			}
// 			return c.JSON(http.StatusOK, "Site Manager updated successfully")
// 		}
// 	} else {
// 		log.Printf("Error unmarshalling Site Manager request: %v", err)
// 	}

// 	// If binding fails or is incomplete for SiteManager, try Worker
// 	var workerRequest request.CreateWorkerRequest
// 	if err := json.Unmarshal(body, &workerRequest); err == nil {
// 		if workerRequest.Name != "" && workerRequest.Role != "" && workerRequest.SiteID != 0 {
// 			log.Printf("Worker payload: %+v", workerRequest)

// 			// Update Worker (PostgreSQL/MongoDB based on usePostgres)
// 			if err := manager.HandleUpdate(intId, &workerRequest, usePostgres); err != nil {
// 				log.Printf("Error in HandleUpdate: %v", err)
// 				return c.JSON(http.StatusInternalServerError, "Failed to update Worker")
// 			}
// 			return c.JSON(http.StatusOK, "Worker updated successfully")
// 		}
// 	} else {
// 		log.Printf("Error unmarshalling Worker request: %v", err)
// 	}

// 	// If binding for both fails, return an invalid data error
// 	return c.JSON(http.StatusBadRequest, "Invalid data")
// }

// Delete deletes a Site Manager or Worker by ID.

// Delete deletes a Site Manager or Worker by ID.
// Delete deletes a Site Manager or Worker by ID.
func Delete(c echo.Context) error {
	usePostgres := c.QueryParam("use_postgres") == "true"
	id := c.Param("id")

	if usePostgres {
		intId, err := strconv.Atoi(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "Invalid ID")
		}
		if err := manager.HandleDelete(intId, usePostgres); err != nil {
			log.Printf("Error in HandleDelete: %v", err)
			return c.JSON(http.StatusInternalServerError, "Failed to delete record")
		}
	} else {
		if err := manager.HandleDelete(id, usePostgres); err != nil {
			log.Printf("Error in HandleDelete: %v", err)
			return c.JSON(http.StatusInternalServerError, "Failed to delete record")
		}
	}

	return c.JSON(http.StatusOK, "Record deleted successfully")
}

// Update updates a Site Manager or Worker by ID.
func Update(c echo.Context) error {
	usePostgres := c.QueryParam("use_postgres") == "true"
	id := c.Param("id")

	// Read raw body content
	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Printf("Error reading request body: %v", err)
		return c.JSON(http.StatusBadRequest, "Invalid data")
	}
	log.Printf("Raw request body: %s", body)

	if usePostgres {
		intId, err := strconv.Atoi(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "Invalid ID")
		}

		// Attempt to bind the incoming data as a SiteManager
		var siteManagerRequest request.CreateSiteManagerRequest
		if err := json.Unmarshal(body, &siteManagerRequest); err == nil {
			if siteManagerRequest.Name != "" && siteManagerRequest.Email != "" {
				log.Printf("Site Manager payload: %+v", siteManagerRequest)

				// Update Site Manager (PostgreSQL)
				if err := manager.HandleUpdate(intId, &siteManagerRequest, usePostgres); err != nil {
					log.Printf("Error in HandleUpdate: %v", err)
					return c.JSON(http.StatusInternalServerError, "Failed to update Site Manager")
				}
				return c.JSON(http.StatusOK, "Site Manager updated successfully")
			}
		} else {
			log.Printf("Error unmarshalling Site Manager request: %v", err)
		}

		// If binding fails or is incomplete for SiteManager, try Worker
		var workerRequest request.CreateWorkerRequest
		if err := json.Unmarshal(body, &workerRequest); err == nil {
			if workerRequest.Name != "" && workerRequest.Role != "" && workerRequest.SiteID != 0 {
				log.Printf("Worker payload: %+v", workerRequest)

				// Update Worker (PostgreSQL)
				if err := manager.HandleUpdate(intId, &workerRequest, usePostgres); err != nil {
					log.Printf("Error in HandleUpdate: %v", err)
					return c.JSON(http.StatusInternalServerError, "Failed to update Worker")
				}
				return c.JSON(http.StatusOK, "Worker updated successfully")
			}
		} else {
			log.Printf("Error unmarshalling Worker request: %v", err)
		}
	} else {
		// Attempt to bind the incoming data as a SiteManager
		var siteManagerRequest request.CreateSiteManagerRequest
		if err := json.Unmarshal(body, &siteManagerRequest); err == nil {
			if siteManagerRequest.Name != "" && siteManagerRequest.Email != "" {
				log.Printf("Site Manager payload: %+v", siteManagerRequest)

				// Update Site Manager (MongoDB)
				if err := manager.HandleUpdate(id, &siteManagerRequest, usePostgres); err != nil {
					log.Printf("Error in HandleUpdate: %v", err)
					return c.JSON(http.StatusInternalServerError, "Failed to update Site Manager")
				}
				return c.JSON(http.StatusOK, "Site Manager updated successfully")
			}
		} else {
			log.Printf("Error unmarshalling Site Manager request: %v", err)
		}

		// If binding fails or is incomplete for SiteManager, try Worker
		var workerRequest request.CreateWorkerRequest
		if err := json.Unmarshal(body, &workerRequest); err == nil {
			if workerRequest.Name != "" && workerRequest.Role != "" && workerRequest.SiteID != 0 {
				log.Printf("Worker payload: %+v", workerRequest)

				// Update Worker (MongoDB)
				if err := manager.HandleUpdate(id, &workerRequest, usePostgres); err != nil {
					log.Printf("Error in HandleUpdate: %v", err)
					return c.JSON(http.StatusInternalServerError, "Failed to update Worker")
				}
				return c.JSON(http.StatusOK, "Worker updated successfully")
			}
		} else {
			log.Printf("Error unmarshalling Worker request: %v", err)
		}
	}

	// If binding for both fails, return an invalid data error
	return c.JSON(http.StatusBadRequest, "Invalid data")
}
