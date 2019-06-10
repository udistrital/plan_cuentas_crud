package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_crud/controllers:ApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_crud/controllers:ApropiacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_crud/controllers:ApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_crud/controllers:ApropiacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_crud/controllers:ApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_crud/controllers:ApropiacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_crud/controllers:ApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_crud/controllers:ApropiacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_crud/controllers:ApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_crud/controllers:ApropiacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_crud/controllers:ApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_crud/controllers:ApropiacionController"],
        beego.ControllerComments{
            Method: "AprobarPresupuesto",
            Router: `/AprobacionAsignacionInicial/`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_crud/controllers:Date"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_crud/controllers:Date"],
        beego.ControllerComments{
            Method: "FechaActual",
            Router: `/FechaActual/:formato`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_crud/controllers:RamaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_crud/controllers:RamaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_crud/controllers:RamaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_crud/controllers:RamaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_crud/controllers:RamaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_crud/controllers:RamaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_crud/controllers:RamaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_crud/controllers:RamaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_crud/controllers:RamaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_crud/controllers:RamaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_crud/controllers:RamaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_crud/controllers:RamaController"],
        beego.ControllerComments{
            Method: "DeleteRubroRelation",
            Router: `/DeleteRubroRelation/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_crud/controllers:RubroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_crud/controllers:RubroController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_crud/controllers:RubroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_crud/controllers:RubroController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_crud/controllers:RubroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_crud/controllers:RubroController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_crud/controllers:RubroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_crud/controllers:RubroController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_crud/controllers:RubroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_crud/controllers:RubroController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_crud/controllers:RubroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_crud/controllers:RubroController"],
        beego.ControllerComments{
            Method: "GetArbolMigracion",
            Router: `/GetArbolMigracion/:a`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
