package alanube

// --- Document Types (used in webhooks and responses) ---

// DocumentType represents the type identifier used in webhook payloads and API responses.
type DocumentType string

const (
	DocumentTypeFiscalInvoice      DocumentType = "fiscalInvoice"      // Factura de crédito fiscal (31)
	DocumentTypeInvoice            DocumentType = "invoice"            // Factura de consumo (32)
	DocumentTypeDebitNote          DocumentType = "debitNote"          // Nota débito (33)
	DocumentTypeCreditNote         DocumentType = "creditNote"         // Nota crédito (34)
	DocumentTypePurchase           DocumentType = "purchase"           // Compras (41)
	DocumentTypeMinorExpense       DocumentType = "minorExpense"       // Gastos menores (43)
	DocumentTypeSpecialRegime      DocumentType = "specialRegime"      // Regímenes especiales (44)
	DocumentTypeGubernamental      DocumentType = "gubernamental"      // Gubernamental (45)
	DocumentTypeExportSupport      DocumentType = "exportSupport"      // Exportación (46)
	DocumentTypePaymentAbroadSupport DocumentType = "paymentAbroadSupport" // Pagos al exterior (47)
)

// --- Document Statuses ---

// Status represents the processing status of a document.
type Status string

const (
	StatusRegistered      Status = "REGISTERED"       // Document was registered
	StatusToSend          Status = "TO_SEND"           // Registered and ready to be sent
	StatusFailed          Status = "FAILED"            // Sending failed (e.g., duplicate eNCF)
	StatusWaitingResponse Status = "WAITING_RESPONSE"  // Sent, awaiting DGII response
	StatusToNotify        Status = "TO_NOTIFY"         // Got DGII response, will notify via webhook
	StatusFinished        Status = "FINISHED"          // Processing completed
)

// --- Legal Statuses ---

// LegalStatus represents the legal status assigned by the DGII.
type LegalStatus string

const (
	LegalStatusRejected                LegalStatus = "REJECTED"                  // Document was rejected
	LegalStatusAccepted                LegalStatus = "ACCEPTED"                  // Document was accepted
	LegalStatusAcceptedWithObservations LegalStatus = "ACCEPTED_WITH_OBSERVATIONS" // Document was partially accepted
	LegalStatusInProcess               LegalStatus = "IN_PROCESS"                // Document is being processed
)

// --- Additional Tax Types (Tipos de Impuestos Adicionales) ---

// TaxType represents an additional tax type code.
type TaxType int

const (
	TaxTypePropinaLegal                TaxType = 1  // Propina Legal (10%)
	TaxTypeCDT                         TaxType = 2  // Contribución al Desarrollo de las Telecomunicaciones (2%)
	TaxTypeISCSeguros                  TaxType = 3  // ISC Servicios Seguros en general (16%)
	TaxTypeISCTelecomunicaciones       TaxType = 4  // ISC Servicios de Telecomunicaciones (10%)
	TaxTypePrimeraPlaca                TaxType = 5  // Impuesto sobre Primer Registro de Vehículos (17%)
	TaxTypeISCEspecificoCerveza        TaxType = 6  // ISC Específico - Cerveza (632.58)
	TaxTypeISCEspecificoVinosUva       TaxType = 7  // ISC Específico - Vinos de uva (632.58)
	TaxTypeISCEspecificoVermut         TaxType = 8  // ISC Específico - Vermut y demás vinos (632.58)
	TaxTypeISCEspecificoBebidas        TaxType = 9  // ISC Específico - Demás bebidas fermentadas (632.58)
	TaxTypeISCEspecificoAlcohol80      TaxType = 10 // ISC Específico - Alcohol Etílico >= 80% (632.58)
	TaxTypeISCEspecificoAlcoholMenos80 TaxType = 11 // ISC Específico - Alcohol Etílico < 80% (632.58)
	TaxTypeISCEspecificoAguardiente    TaxType = 12 // ISC Específico - Aguardientes de uva (632.58)
	TaxTypeISCEspecificoWhisky         TaxType = 13 // ISC Específico - Whisky (632.58)
	TaxTypeISCEspecificoRon            TaxType = 14 // ISC Específico - Ron y aguardientes de caña (632.58)
	TaxTypeISCEspecificoGin            TaxType = 15 // ISC Específico - Gin y Ginebra (632.58)
	TaxTypeISCEspecificoVodka          TaxType = 16 // ISC Específico - Vodka (632.58)
	TaxTypeISCEspecificoLicores        TaxType = 17 // ISC Específico - Licores (632.58)
	TaxTypeISCEspecificoDemasAlcohol   TaxType = 18 // ISC Específico - Los demás Bebidas y Alcoholes (632.58)
	TaxTypeISCEspecificoCigarrillos20  TaxType = 19 // ISC Específico - Cigarrillos cajetilla 20 uds (53.51)
	TaxTypeISCEspecificoDemasCig20     TaxType = 20 // ISC Específico - Los demás Cigarrillos 20 uds (53.51)
	TaxTypeISCEspecificoCigarrillos10  TaxType = 21 // ISC Específico - Cigarrillos 10 uds (26.75)
	TaxTypeISCEspecificoDemasCig10     TaxType = 22 // ISC Específico - Los demás Cigarrillos 10 uds (26.75)
	TaxTypeISCAdValoremCerveza         TaxType = 23 // ISC Ad Valorem - Cerveza (10%)
	TaxTypeISCAdValoremVinosUva        TaxType = 24 // ISC Ad Valorem - Vinos de uva (10%)
	TaxTypeISCAdValoremVermut          TaxType = 25 // ISC Ad Valorem - Vermut y demás vinos (10%)
	TaxTypeISCAdValoremBebidas         TaxType = 26 // ISC Ad Valorem - Demás bebidas fermentadas (10%)
	TaxTypeISCAdValoremAlcohol80       TaxType = 27 // ISC Ad Valorem - Alcohol Etílico >= 80% (10%)
	TaxTypeISCAdValoremAlcoholMenos80  TaxType = 28 // ISC Ad Valorem - Alcohol Etílico < 80% (10%)
	TaxTypeISCAdValoremAguardiente     TaxType = 29 // ISC Ad Valorem - Aguardientes de uva (10%)
	TaxTypeISCAdValoremWhisky          TaxType = 30 // ISC Ad Valorem - Whisky (10%)
	TaxTypeISCAdValoremRon             TaxType = 31 // ISC Ad Valorem - Ron y aguardientes de caña (10%)
	TaxTypeISCAdValoremGin             TaxType = 32 // ISC Ad Valorem - Gin y Ginebra (10%)
	TaxTypeISCAdValoremVodka           TaxType = 33 // ISC Ad Valorem - Vodka (10%)
	TaxTypeISCAdValoremLicores         TaxType = 34 // ISC Ad Valorem - Licores (10%)
	TaxTypeISCAdValoremDemasAlcohol    TaxType = 35 // ISC Ad Valorem - Los demás Bebidas y Alcoholes (10%)
	TaxTypeISCAdValoremCigarrillos20   TaxType = 36 // ISC Ad Valorem - Cigarrillos cajetilla 20 uds (20%)
	TaxTypeISCAdValoremDemasCig20      TaxType = 37 // ISC Ad Valorem - Los demás Cigarrillos 20 uds (20%)
	TaxTypeISCAdValoremCigarrillos10   TaxType = 38 // ISC Ad Valorem - Cigarrillos 10 uds (20%)
	TaxTypeISCAdValoremDemasCig10      TaxType = 39 // ISC Ad Valorem - Los demás Cigarrillos 10 uds (20%)
)

// --- Unit of Measure (Unidades de Medida) ---

// UnitMeasure represents a unit of measure code.
type UnitMeasure int

const (
	UnitBarril              UnitMeasure = 1  // BARR - Barril
	UnitBolsa               UnitMeasure = 2  // BOL - Bolsa
	UnitBote                UnitMeasure = 3  // BOT - Bote
	UnitBultos              UnitMeasure = 4  // BULTO - Bultos
	UnitBotella             UnitMeasure = 5  // BOTELLA - Botella
	UnitCaja                UnitMeasure = 6  // CAJ - Caja
	UnitCajetilla           UnitMeasure = 7  // CAJETILLA - Cajetilla
	UnitCentimetro          UnitMeasure = 8  // CM - Centímetro
	UnitCilindro            UnitMeasure = 9  // CIL - Cilindro
	UnitConjunto            UnitMeasure = 10 // CONJ - Conjunto
	UnitContenedor          UnitMeasure = 11 // CONT - Contenedor
	UnitDia                 UnitMeasure = 12 // DÍA - Día
	UnitDocena              UnitMeasure = 13 // DOC - Docena
	UnitFardo               UnitMeasure = 14 // FARD - Fardo
	UnitGalones             UnitMeasure = 15 // GL - Galones
	UnitGrado               UnitMeasure = 16 // GRAD - Grado
	UnitGramo               UnitMeasure = 17 // GR - Gramo
	UnitGranel              UnitMeasure = 18 // GRAN - Granel
	UnitHora                UnitMeasure = 19 // HOR - Hora
	UnitHuacal              UnitMeasure = 20 // HUAC - Huacal
	UnitKilogramo           UnitMeasure = 21 // KG - Kilogramo
	UnitKilovatioHora       UnitMeasure = 22 // kWh - Kilovatio Hora
	UnitLibra               UnitMeasure = 23 // LB - Libra
	UnitLitro               UnitMeasure = 24 // LITRO - Litro
	UnitLote                UnitMeasure = 25 // LOT - Lote
	UnitMetro               UnitMeasure = 26 // M - Metro
	UnitMetroCuadrado       UnitMeasure = 27 // M2 - Metro Cuadrado
	UnitMetroCubico         UnitMeasure = 28 // M3 - Metro Cúbico
	UnitMMBTU               UnitMeasure = 29 // MMBTU - Millones de Unidades Térmicas
	UnitMinuto              UnitMeasure = 30 // MIN - Minuto
	UnitPaquete             UnitMeasure = 31 // PAQ - Paquete
	UnitPar                 UnitMeasure = 32 // PAR - Par
	UnitPie                 UnitMeasure = 33 // PIE - Pie
	UnitPieza               UnitMeasure = 34 // PZA - Pieza
	UnitRollo               UnitMeasure = 35 // ROL - Rollo
	UnitSobre               UnitMeasure = 36 // SOBR - Sobre
	UnitSegundo             UnitMeasure = 37 // SEG - Segundo
	UnitTanque              UnitMeasure = 38 // TANQUE - Tanque
	UnitTonelada            UnitMeasure = 39 // TONE - Tonelada
	UnitTubo                UnitMeasure = 40 // TUB - Tubo
	UnitYarda               UnitMeasure = 41 // YD - Yarda
	UnitYardaCuadrada       UnitMeasure = 42 // YD2 - Yarda Cuadrada
	UnitUnidad              UnitMeasure = 43 // UND - Unidad
	UnitElemento            UnitMeasure = 44 // EA - Elemento
	UnitMillar              UnitMeasure = 45 // MILLAR - Millar
	UnitSaco                UnitMeasure = 46 // SAC - Saco
	UnitLata                UnitMeasure = 47 // LAT - Lata
	UnitDisplay             UnitMeasure = 48 // DIS - Display
	UnitBidon               UnitMeasure = 49 // BID - Bidón
	UnitRacion              UnitMeasure = 50 // RAC - Ración
	UnitQuintal             UnitMeasure = 51 // Q - Quintal
	UnitToneladasRegistro   UnitMeasure = 52 // GRT - Toneladas de registro bruto
	UnitPieCuadrado         UnitMeasure = 53 // P2 - Pie cuadrado
	UnitPasajero            UnitMeasure = 54 // PAX - Pasajero
	UnitPulgadas            UnitMeasure = 55 // PULG - Pulgadas
	UnitParqueoBarcos       UnitMeasure = 56 // STAY - Parqueo barcos en muelle
	UnitBandeja             UnitMeasure = 57 // BDJ - Bandeja
	UnitHectarea            UnitMeasure = 58 // HA - Hectárea
	UnitMililitro           UnitMeasure = 59 // ML - Mililitro
	UnitMiligramo           UnitMeasure = 60 // MG - Miligramo
	UnitOnzas               UnitMeasure = 61 // OZ - Onzas
	UnitOnzasTroy           UnitMeasure = 62 // OZT - Onzas Troy
)

// --- Currency Codes (Codificación Monedas) ---

// CurrencyCode represents an ISO currency code accepted by the API.
type CurrencyCode string

const (
	CurrencyBRL CurrencyCode = "BRL" // Real Brasileño
	CurrencyCAD CurrencyCode = "CAD" // Dólar Canadiense
	CurrencyCHF CurrencyCode = "CHF" // Franco Suizo
	CurrencyCHY CurrencyCode = "CHY" // Yuan Chino
	CurrencyXDR CurrencyCode = "XDR" // Derecho Especial de Giro
	CurrencyDKK CurrencyCode = "DKK" // Corona Danesa
	CurrencyEUR CurrencyCode = "EUR" // Euro
	CurrencyGBP CurrencyCode = "GBP" // Libra Esterlina
	CurrencyJPY CurrencyCode = "JPY" // Yen Japonés
	CurrencyNOK CurrencyCode = "NOK" // Corona Noruega
	CurrencySCP CurrencyCode = "SCP" // Libra Escocesa
	CurrencySEK CurrencyCode = "SEK" // Corona Sueca
	CurrencyUSD CurrencyCode = "USD" // Dólar Estadounidense
	CurrencyVEF CurrencyCode = "VEF" // Bolívar Fuerte Venezolano
	CurrencyHTG CurrencyCode = "HTG" // Gurda Haitiana
	CurrencyMXN CurrencyCode = "MXN" // Peso Mexicano
)
