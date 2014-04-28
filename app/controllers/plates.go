package controllers

import (
	"goplates-api/app/models"
	"strings"

	"github.com/revel/revel"
)

type Plates struct {
	Application
}

func (c Plates) Show(plate string) revel.Result {
	i := c.loadByPlate(plate)
	if i == nil {
		return c.RenderJson(struct{}{})
	}
	employee := c.loadEmployeeById(i.EmployeeID)
	return c.RenderJson(employee)
}

func (c Plates) Blocked(plate string) revel.Result {
	i := c.loadByPlate(plate)
	if i == nil {
		return c.RenderJson(struct{}{})
	}
	employee := c.loadEmployeeById(i.EmployeeID)
	return c.RenderJson(employee)
}

func (c Plates) Blocking(plate string) revel.Result {
	i := c.loadByPlate(plate)
	if i == nil {
		return c.RenderJson(struct{}{})
	}
	employee := c.loadEmployeeById(i.EmployeeID)
	return c.RenderJson(employee)
}

func (c Plates) loadByPlate(plate string) *models.Plate {
	h, err := c.Txn.Get(models.Plate{}, strings.ToUpper(plate))
	if err != nil {
		panic(err)
	}
	if h == nil {
		return nil
	}
	return h.(*models.Plate)
}

func (c Plates) loadEmployeeById(id int) *models.Employee {
	h, err := c.Txn.Get(models.Employee{}, id)
	if err != nil {
		panic(err)
	}
	if h == nil {
		return nil
	}
	return h.(*models.Employee)
}
