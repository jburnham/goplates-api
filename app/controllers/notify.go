package controllers

import (
	"goplates-api/app/models"
	"strings"

	"github.com/revel/revel"
)

type Notify struct {
	Application
}

func (c Notify) Blocked(employeeId int) revel.Result {
	employee := c.loadEmployeeById(employeeId)
	return c.RenderJson(employee)
}

func (c Notify) Blocking(employeeId int) revel.Result {
	employee := c.loadEmployeeById(employeeId)
	return c.RenderJson(employee)
}

func (c Notify) loadByPlate(plate string) *models.Plate {
	h, err := c.Txn.Get(models.Plate{}, strings.ToUpper(plate))
	if err != nil {
		panic(err)
	}
	if h == nil {
		return nil
	}
	return h.(*models.Plate)
}

func (c Notify) loadEmployeeById(id int) *models.Employee {
	h, err := c.Txn.Get(models.Employee{}, id)
	if err != nil {
		panic(err)
	}
	if h == nil {
		return nil
	}
	return h.(*models.Employee)
}
