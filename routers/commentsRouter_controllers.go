package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:CategoriaMovilidadController"] = append(beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:CategoriaMovilidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:CategoriaMovilidadController"] = append(beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:CategoriaMovilidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:CategoriaMovilidadController"] = append(beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:CategoriaMovilidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:CategoriaMovilidadController"] = append(beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:CategoriaMovilidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:CategoriaMovilidadController"] = append(beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:CategoriaMovilidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:ClasificacionDuracionController"] = append(beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:ClasificacionDuracionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:ClasificacionDuracionController"] = append(beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:ClasificacionDuracionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:ClasificacionDuracionController"] = append(beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:ClasificacionDuracionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:ClasificacionDuracionController"] = append(beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:ClasificacionDuracionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:ClasificacionDuracionController"] = append(beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:ClasificacionDuracionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:ClasificacionMovilidadController"] = append(beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:ClasificacionMovilidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:ClasificacionMovilidadController"] = append(beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:ClasificacionMovilidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:ClasificacionMovilidadController"] = append(beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:ClasificacionMovilidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:ClasificacionMovilidadController"] = append(beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:ClasificacionMovilidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:ClasificacionMovilidadController"] = append(beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:ClasificacionMovilidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:ConvenioController"] = append(beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:ConvenioController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:ConvenioController"] = append(beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:ConvenioController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:ConvenioController"] = append(beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:ConvenioController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:ConvenioController"] = append(beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:ConvenioController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:ConvenioController"] = append(beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:ConvenioController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:MovilidadController"] = append(beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:MovilidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:MovilidadController"] = append(beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:MovilidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:MovilidadController"] = append(beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:MovilidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:MovilidadController"] = append(beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:MovilidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:MovilidadController"] = append(beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:MovilidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:PresupuestoController"] = append(beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:PresupuestoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:PresupuestoController"] = append(beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:PresupuestoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:PresupuestoController"] = append(beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:PresupuestoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:PresupuestoController"] = append(beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:PresupuestoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:PresupuestoController"] = append(beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:PresupuestoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:TipoMovilidadController"] = append(beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:TipoMovilidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:TipoMovilidadController"] = append(beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:TipoMovilidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:TipoMovilidadController"] = append(beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:TipoMovilidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:TipoMovilidadController"] = append(beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:TipoMovilidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:TipoMovilidadController"] = append(beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:TipoMovilidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:TipoPersonaController"] = append(beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:TipoPersonaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:TipoPersonaController"] = append(beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:TipoPersonaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:TipoPersonaController"] = append(beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:TipoPersonaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:TipoPersonaController"] = append(beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:TipoPersonaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:TipoPersonaController"] = append(beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:TipoPersonaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:TipoPresupuestoController"] = append(beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:TipoPresupuestoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:TipoPresupuestoController"] = append(beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:TipoPresupuestoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:TipoPresupuestoController"] = append(beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:TipoPresupuestoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:TipoPresupuestoController"] = append(beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:TipoPresupuestoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:TipoPresupuestoController"] = append(beego.GlobalControllerRouter["github.com/nicolas0517/reporte-movilidad/controllers:TipoPresupuestoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

}
