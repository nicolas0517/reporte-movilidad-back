package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type Movilidad struct {
	Id                       int                     `orm:"column(id);pk;auto"`
	FechaInicio              string                  `orm:"column(fecha_inicio)"`
	FechaFin                 string                  `orm:"column(fecha_fin)"`
	NombreActoAdministrativo string                  `orm:"column(nombre_acto_administrativo)"`
	EnlaceActoAdministrativo string                  `orm:"column(enlace_acto_administrativo)"`
	AreaConocimiento         string                  `orm:"column(area_conocimiento);null"`
	ConceptoParticipacion    string                  `orm:"column(concepto_participacion);null"`
	Pais                     int                     `orm:"column(pais)"`
	TipoMovilidad            *TipoMovilidad          `orm:"column(tipo_movilidad);rel(fk)"`
	ClasificacionDuracion    *ClasificacionDuracion  `orm:"column(clasificacion_duracion);rel(fk)"`
	CategoriaMovilidad       *CategoriaMovilidad     `orm:"column(categoria_movilidad);rel(fk)"`
	Institucion              int                     `orm:"column(institucion)"`
	Convenio                 *Convenio               `orm:"column(convenio);rel(fk);null"`
	Dependencia              int                     `orm:"column(dependencia);null"`
	TipoPersona              *TipoPersona            `orm:"column(tipo_persona);rel(fk)"`
	ClasificacionMovilidad   *ClasificacionMovilidad `orm:"column(clasificacion_movilidad);rel(fk)"`
	Persona                  int                     `orm:"column(persona)"`
}

func (t *Movilidad) TableName() string {
	return "movilidad"
}

func init() {
	orm.RegisterModel(new(Movilidad))
}

// AddMovilidad insert a new Movilidad into database and returns
// last inserted Id on success.
func AddMovilidad(m *Movilidad) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetMovilidadById retrieves Movilidad by Id. Returns error if
// Id doesn't exist
func GetMovilidadById(id int) (v *Movilidad, err error) {
	o := orm.NewOrm()
	v = &Movilidad{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllMovilidad retrieves all Movilidad matches certain condition. Returns empty list if
// no records exist
func GetAllMovilidad(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Movilidad))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Movilidad
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateMovilidad updates Movilidad by Id and returns error if
// the record to be updated doesn't exist
func UpdateMovilidadById(m *Movilidad) (err error) {
	o := orm.NewOrm()
	v := Movilidad{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteMovilidad deletes Movilidad by Id and returns error if
// the record to be deleted doesn't exist
func DeleteMovilidad(id int) (err error) {
	o := orm.NewOrm()
	v := Movilidad{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Movilidad{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
