// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/nicolas0517/reporte-movilidad/controllers"

	"github.com/astaxie/beego"

	"github.com/astaxie/beego/plugins/cors"
)

func init() {
	beego.Debug("Filters init...")

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))

	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/tipo_movilidad",
			beego.NSInclude(
				&controllers.TipoMovilidadController{},
			),
		),

		beego.NSNamespace("/clasificacion_duracion",
			beego.NSInclude(
				&controllers.ClasificacionDuracionController{},
			),
		),

		beego.NSNamespace("/categoria_movilidad",
			beego.NSInclude(
				&controllers.CategoriaMovilidadController{},
			),
		),

		beego.NSNamespace("/tipo_presupuesto",
			beego.NSInclude(
				&controllers.TipoPresupuestoController{},
			),
		),

		beego.NSNamespace("/presupuesto",
			beego.NSInclude(
				&controllers.PresupuestoController{},
			),
		),

		beego.NSNamespace("/convenio",
			beego.NSInclude(
				&controllers.ConvenioController{},
			),
		),

		beego.NSNamespace("/tipo_persona",
			beego.NSInclude(
				&controllers.TipoPersonaController{},
			),
		),

		beego.NSNamespace("/clasificacion_movilidad",
			beego.NSInclude(
				&controllers.ClasificacionMovilidadController{},
			),
		),

		beego.NSNamespace("/movilidad",
			beego.NSInclude(
				&controllers.MovilidadController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
