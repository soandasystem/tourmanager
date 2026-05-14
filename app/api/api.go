package api

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"tourmanager/app/handlers"
	"tourmanager/config"
	"tourmanager/core/ports"
	"tourmanager/core/services"
	"tourmanager/infrastructure/b2"
	"tourmanager/infrastructure/postgres"
	"tourmanager/util"

	"github.com/antoniomarfa/hexatools/infrastructure/middleware"

	"github.com/antoniomarfa/hexatools/infrastructure"

	"github.com/gin-gonic/gin"
)

type api struct {
	config   config.Config
	services svs
}

type svs struct {
	company    ports.CompanyService
	colegios   ports.ColegiosService
	comunas    ports.ComunasService
	curso      ports.CursoService
	fmedica    ports.FmedicaService
	gatewaysc  ports.GatewayscService
	gateways   ports.GatewaysService
	ingreso    ports.IngresoService
	pagos      ports.PagosService
	permission ports.PermissionService
	programac  ports.ProgramacService
	programad  ports.ProgramadService
	quotes     ports.QuotesService
	region     ports.RegionService
	roles      ports.RolesService
	sale       ports.SaleService
	users      ports.UsersService
	voucher    ports.VoucherService
	airports   ports.AirportsService
	country    ports.CountryService
	programs   ports.ProgramsService
	upload     ports.UploadService
}

// New creates a new API
func New(ctx context.Context, cfg config.Config) (a api) {
	a.config = cfg

	var companyRepo ports.CompanyRepository
	var colegiosRepo ports.ColegiosRepository
	var comunasRepo ports.ComunasRepository
	var cursoRepo ports.CursoRepository
	var fmedicaRepo ports.FmedicaRepository
	var gatewayscRepo ports.GatewayscRepository
	var gatewaysRepo ports.GatewaysRepository
	var ingresoRepo ports.IngresoRepository
	var pagosRepo ports.PagosRepository
	var permissionRepo ports.PermissionRepository
	var programacRepo ports.ProgramacRepository
	var programadRepo ports.ProgramadRepository
	var quotesRepo ports.QuotesRepository
	var regionRepo ports.RegionRepository
	var rolesRepo ports.RolesRepository
	var saleRepo ports.SaleRepository
	var usersRepo ports.UsersRepository
	var voucherRepo ports.VoucherRepository
	var airportsRepo ports.AirportsRepository
	var countryRepo ports.CountryRepository
	var programsRepo ports.ProgramsRepository

	//abre la base de datos
	db, err := infrastructure.ConnectPostgresOrm(ctx, a.config.DSN)
	if err != nil {
		log.Fatal(err)
	}

	companyRepo = postgres.NewCompanyRepository(ctx, db)
	colegiosRepo = postgres.NewColegiosRepository(ctx, db)
	comunasRepo = postgres.NewComunasRepository(ctx, db)
	cursoRepo = postgres.NewCursoRepository(ctx, db)
	fmedicaRepo = postgres.NewFmedicaRepository(ctx, db)
	gatewayscRepo = postgres.NewGatewayscRepository(ctx, db)
	gatewaysRepo = postgres.NewGatewaysRepository(ctx, db)
	ingresoRepo = postgres.NewIngresoRepository(ctx, db)
	pagosRepo = postgres.NewPagosRepository(ctx, db)
	permissionRepo = postgres.NewPermissionRepository(ctx, db)
	programacRepo = postgres.NewProgramacRepository(ctx, db)
	programadRepo = postgres.NewProgramadRepository(ctx, db)
	quotesRepo = postgres.NewQuotesRepository(ctx, db)
	regionRepo = postgres.NewRegionRepository(ctx, db)
	rolesRepo = postgres.NewRolesRepository(ctx, db)
	saleRepo = postgres.NewSaleRepository(ctx, db)
	usersRepo = postgres.NewUsersRepository(ctx, db)
	voucherRepo = postgres.NewVoucherRepository(ctx, db)
	airportsRepo = postgres.NewAirportsRepository(ctx, db)
	countryRepo = postgres.NewCountryRepository(ctx, db)
	programsRepo = postgres.NewProgramsRepository(ctx, db)

	a.services.company = services.NewCompanyService(a.config, companyRepo)
	a.services.colegios = services.NewColegiosService(a.config, colegiosRepo)
	a.services.comunas = services.NewComunasService(a.config, comunasRepo)
	a.services.curso = services.NewCursoService(a.config, cursoRepo)
	a.services.fmedica = services.NewFmedicaService(a.config, fmedicaRepo)
	a.services.gatewaysc = services.NewGatewayscService(a.config, gatewayscRepo)
	a.services.gateways = services.NewGatewaysService(a.config, gatewaysRepo)
	a.services.ingreso = services.NewIngresoService(a.config, ingresoRepo)
	a.services.pagos = services.NewPagosService(a.config, pagosRepo)
	a.services.permission = services.NewPermissionService(a.config, permissionRepo)
	a.services.programac = services.NewProgramacService(a.config, programacRepo)
	a.services.programad = services.NewProgramadService(a.config, programadRepo)
	a.services.quotes = services.NewQuotesService(a.config, quotesRepo)
	a.services.region = services.NewRegionService(a.config, regionRepo)
	a.services.roles = services.NewRolesService(a.config, rolesRepo)
	a.services.sale = services.NewSaleService(a.config, saleRepo)
	a.services.users = services.NewUsersService(a.config, usersRepo)
	a.services.voucher = services.NewVoucherService(a.config, voucherRepo)
	a.services.airports = services.NewAirportsService(a.config, airportsRepo)
	a.services.country = services.NewCountryService(a.config, countryRepo)
	a.services.programs = services.NewProgramsService(a.config, programsRepo)

	// Inicializar B2 Storage y Upload Service
	if a.config.B2Endpoint != "" {
		b2Storage := b2.NewB2Storage(
			ctx,
			a.config.B2KeyID,
			a.config.B2Application,
			a.config.B2Bucket,
			a.config.B2Region,
			a.config.B2Endpoint,
		)
		a.services.upload = services.NewUploadService(a.config, b2Storage)
	}

	return a
}

// Run API
func (a *api) Run(ctx context.Context, cancel context.CancelFunc) func() error {
	return func() error {
		defer cancel()
		gin.SetMode(gin.ReleaseMode)
		router := gin.Default()

		// Configurar CORS globalmente
		router.Use(middleware.CORSMiddleware())

		//asigna el schema a trabajar
		router.Use(util.SchemaMiddleware())

		// Luego este que inyecta el contexto con timeout y schema:
		router.Use(util.ContextWithSchemaMiddleware(a.config.Timeout.Duration))

		handlers.SetHealthRoutes(ctx, a.config, router)
		handlers.SetCompanyRoutes(ctx, a.config, router, a.services.company)
		handlers.SetColegiosRoutes(ctx, a.config, router, a.services.colegios)
		handlers.SetComunasRoutes(ctx, a.config, router, a.services.comunas)
		handlers.SetCursoRoutes(ctx, a.config, router, a.services.curso)
		handlers.SetFmedicaRoutes(ctx, a.config, router, a.services.fmedica)
		handlers.SetGatewayscRoutes(ctx, a.config, router, a.services.gatewaysc)
		handlers.SetGatewaysRoutes(ctx, a.config, router, a.services.gateways)
		handlers.SetIngresoRoutes(ctx, a.config, router, a.services.ingreso)
		handlers.SetPagosRoutes(ctx, a.config, router, a.services.pagos)
		handlers.SetPermissionRoutes(ctx, a.config, router, a.services.permission)
		handlers.SetProgramacRoutes(ctx, a.config, router, a.services.programac)
		handlers.SetProgramadRoutes(ctx, a.config, router, a.services.programad)
		handlers.SetQuotesRoutes(ctx, a.config, router, a.services.quotes)
		handlers.SetRegionRoutes(ctx, a.config, router, a.services.region)
		handlers.SetRolesRoutes(ctx, a.config, router, a.services.roles)
		handlers.SetSaleRoutes(ctx, a.config, router, a.services.sale)
		handlers.SetUsersRoutes(ctx, a.config, router, a.services.users)
		handlers.SetVoucherRoutes(ctx, a.config, router, a.services.voucher)
		handlers.SetAirportsRoutes(ctx, a.config, router, a.services.airports)
		handlers.SetCountryRoutes(ctx, a.config, router, a.services.country)
		handlers.SetProgramsRoutes(ctx, a.config, router, a.services.programs)

		if a.services.upload != nil {
			handlers.SetUploadRoutes(ctx, a.config, router, a.services.upload)
		}

		log.Printf("Version: %s", a.config.Version)
		log.Printf("Environment: %s", a.config.Environment)
		log.Printf("Database: %s", a.config.Database)
		log.Printf("Listening on port %d", a.config.Port)

		server := &http.Server{
			Addr:    fmt.Sprintf(":%d", a.config.Port),
			Handler: router,
		}
		go shutdown(ctx, server)
		return server.ListenAndServe()
	}
}

func SchemaMiddleware() gin.HandlerFunc {
	panic("unimplemented")
}

func shutdown(ctx context.Context, server *http.Server) {
	<-ctx.Done()
	log.Printf("Shutting down API gracefully...")
	server.Shutdown(ctx)
}
