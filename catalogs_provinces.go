package alanube

// --- Provinces (Dominican Republic) ---

// Province represents a Dominican Republic province code used in DGII electronic invoicing.
type Province string

const (
	ProvinceDistritoNacional      Province = "010000" // DISTRITO NACIONAL
	ProvinceAzua                  Province = "020000" // PROVINCIA AZUA
	ProvinceBahoruco              Province = "030000" // PROVINCIA BAHORUCO
	ProvinceBarahona              Province = "040000" // PROVINCIA BARAHONA
	ProvinceDajabon               Province = "050000" // PROVINCIA DAJABÓN
	ProvinceDuarte                Province = "060000" // PROVINCIA DUARTE
	ProvinceEliasPina             Province = "070000" // PROVINCIA ELÍAS PIÑA
	ProvinceElSeibo               Province = "080000" // PROVINCIA EL SEIBO
	ProvinceEspaillat             Province = "090000" // PROVINCIA ESPAILLAT
	ProvinceIndependencia         Province = "100000" // PROVINCIA INDEPENDENCIA
	ProvinceLaAltagracia          Province = "110000" // PROVINCIA LA ALTAGRACIA
	ProvinceLaRomana              Province = "120000" // PROVINCIA LA ROMANA
	ProvinceLaVega                Province = "130000" // PROVINCIA LA VEGA
	ProvinceMariaTrindadSanchez   Province = "140000" // PROVINCIA MARÍA TRINIDAD SÁNCHEZ
	ProvinceMonteCristi           Province = "150000" // PROVINCIA MONTE CRISTI
	ProvincePedernales            Province = "160000" // PROVINCIA PEDERNALES
	ProvincePeravia               Province = "170000" // PROVINCIA PERAVIA
	ProvincePuertoPlata           Province = "180000" // PROVINCIA PUERTO PLATA
	ProvinceHermanasMirabal       Province = "190000" // PROVINCIA HERMANAS MIRABAL
	ProvinceSamana                Province = "200000" // PROVINCIA SAMANÁ
	ProvinceSanCristobal          Province = "210000" // PROVINCIA SAN CRISTÓBAL
	ProvinceSanJuan               Province = "220000" // PROVINCIA SAN JUAN
	ProvinceSanPedrodeMacoris     Province = "230000" // PROVINCIA SAN PEDRO DE MACORÍS
	ProvinceSanchezRamirez        Province = "240000" // PROVINCIA SÁNCHEZ RAMÍREZ
	ProvinceSantiago              Province = "250000" // PROVINCIA SANTIAGO
	ProvinceSantiagoRodriguez     Province = "260000" // PROVINCIA SANTIAGO RODRÍGUEZ
	ProvinceValverde              Province = "270000" // PROVINCIA VALVERDE
	ProvinceMonsenorNouel         Province = "280000" // PROVINCIA MONSEÑOR NOUEL
	ProvinceMontePlata            Province = "290000" // PROVINCIA MONTE PLATA
	ProvinceHatoMayor             Province = "300000" // PROVINCIA HATO MAYOR
	ProvinceSanJosedeOcoa         Province = "310000" // PROVINCIA SAN JOSÉ DE OCOA
	ProvinceSantoDomingo          Province = "320000" // PROVINCIA SANTO DOMINGO
)

// --- Municipalities (Dominican Republic) ---

// Municipality represents a Dominican Republic municipality code used in DGII electronic invoicing.
type Municipality string

const (
	// Distrito Nacional (01)
	MunicipalitySantoDomingoDeGuzman Municipality = "010100" // SANTO DOMINGO DE GUZMÁN

	// Provincia Azua (02)
	MunicipalityAzua              Municipality = "020100" // AZUA
	MunicipalityLasCharcas        Municipality = "020200" // LAS CHARCAS
	MunicipalityLasYayasDeViajama Municipality = "020300" // LAS YAYAS DE VIAJAMA
	MunicipalityPadreLasCasas     Municipality = "020400" // PADRE LAS CASAS
	MunicipalityPeralta           Municipality = "020500" // PERALTA
	MunicipalitySabanaYegua       Municipality = "020600" // SABANA YEGUA
	MunicipalityPuebloViejo       Municipality = "020700" // PUEBLO VIEJO
	MunicipalityTabaraArriba      Municipality = "020800" // TÁBARA ARRIBA
	MunicipalityGuayabal          Municipality = "020900" // GUAYABAL
	MunicipalityEstebanía         Municipality = "021000" // ESTEBANÍA

	// Provincia Bahoruco (03)
	MunicipalityNeiba        Municipality = "030001" // NEIBA
	MunicipalityGalvan       Municipality = "030200" // GALVÁN
	MunicipalityTamayo       Municipality = "030300" // TAMAYO
	MunicipalityVillaJaragua Municipality = "030400" // VILLA JARAGUA
	MunicipalityLosRios      Municipality = "030500" // LOS RÍOS

	// Provincia Barahona (04)
	MunicipalityBarahona   Municipality = "040100" // BARAHONA
	MunicipalityCabral     Municipality = "040200" // CABRAL
	MunicipalityEnriquillo Municipality = "040300" // ENRIQUILLO
	MunicipalityParaiso    Municipality = "040400" // PARAÍSO
	MunicipalityVicenteNoble Municipality = "040500" // VICENTE NOBLE
	MunicipalityElPenon    Municipality = "040600" // EL PEÑÓN
	MunicipalityLaCienaga  Municipality = "040700" // LA CIÉNAGA
	MunicipalityFundacion  Municipality = "040800" // FUNDACIÓN
	MunicipalityLasSalinas Municipality = "040900" // LAS SALINAS
	MunicipalityPolo       Municipality = "041000" // POLO
	MunicipalityJaquimeyes Municipality = "041100" // JAQUIMEYES

	// Provincia Dajabón (05)
	MunicipalityDajabon        Municipality = "050100" // DAJABÓN
	MunicipalityLomaDeCabrera  Municipality = "050200" // LOMA DE CABRERA
	MunicipalityPartido        Municipality = "050300" // PARTIDO
	MunicipalityRestauracion   Municipality = "050400" // RESTAURACIÓN
	MunicipalityElPino         Municipality = "050500" // EL PINO

	// Provincia Duarte (06)
	MunicipalitySanFranciscoDeMacoris Municipality = "060100" // SAN FRANCISCO DE MACORÍS
	MunicipalityArenoso               Municipality = "060200" // ARENOSO
	MunicipalityCastillo              Municipality = "060300" // CASTILLO
	MunicipalityPimentel              Municipality = "060400" // PIMENTEL
	MunicipalityVillaRiva             Municipality = "060500" // VILLA RIVA
	MunicipalityLasGuaranas           Municipality = "060600" // LAS GUÁRANAS
	MunicipalityEugenioMariaDeHostos  Municipality = "060700" // EUGENIO MARÍA DE HOSTOS

	// Provincia Elías Piña (07)
	MunicipalityComendador   Municipality = "070100" // COMENDADOR
	MunicipalityBanica       Municipality = "070200" // BÁNICA
	MunicipalityElLlano      Municipality = "070300" // EL LLANO
	MunicipalityHondoValle   Municipality = "070400" // HONDO VALLE
	MunicipalityPedroSantana Municipality = "070500" // PEDRO SANTANA
	MunicipalityJuanSantiago Municipality = "070600" // JUAN SANTIAGO

	// Provincia El Seibo (08)
	MunicipalityElSeibo Municipality = "080100" // EL SEIBO
	MunicipalityMiches Municipality = "080200" // MICHES

	// Provincia Espaillat (09)
	MunicipalityMoca             Municipality = "090100" // MOCA
	MunicipalityCayetanoGermosen Municipality = "090200" // CAYETANO GERMOSÉN
	MunicipalityGasparHernandez  Municipality = "090300" // GASPAR HERNÁNDEZ
	MunicipalityJamaoAlNorte     Municipality = "090400" // JAMAO AL NORTE

	// Provincia Independencia (10)
	MunicipalityJimani         Municipality = "100100" // JIMANÍ
	MunicipalityDuverge        Municipality = "100200" // DUVERGÉ
	MunicipalityLaDescubierta  Municipality = "100300" // LA DESCUBIERTA
	MunicipalityPostrerRio     Municipality = "100400" // POSTRER RÍO
	MunicipalityCristobal      Municipality = "100500" // CRISTÓBAL
	MunicipalityMella          Municipality = "100600" // MELLA

	// Provincia La Altagracia (11)
	MunicipalityHiguey            Municipality = "110100" // HIGÜEY
	MunicipalitySanRafaelDelYuma  Municipality = "110200" // SAN RAFAEL DEL YUMA

	// Provincia La Romana (12)
	MunicipalityLaRomana     Municipality = "120100" // LA ROMANA
	MunicipalityGuaymate     Municipality = "120200" // GUAYMATE
	MunicipalityVillaHermosa Municipality = "120300" // VILLA HERMOSA

	// Provincia La Vega (13)
	MunicipalityLaVega    Municipality = "130100" // LA VEGA
	MunicipalityConstanza Municipality = "130200" // CONSTANZA
	MunicipalityJarabacoa Municipality = "130300" // JARABACOA
	MunicipalityJimaAbajo Municipality = "130400" // JIMA ABAJO

	// Provincia María Trinidad Sánchez (14)
	MunicipalityNagua      Municipality = "140100" // NAGUA
	MunicipalityCabrera    Municipality = "140200" // CABRERA
	MunicipalityElFactor   Municipality = "140300" // EL FACTOR
	MunicipalityRioSanJuan Municipality = "140400" // RÍO SAN JUAN

	// Provincia Monte Cristi (15)
	MunicipalityMonteCristi           Municipality = "150100" // MONTE CRISTI
	MunicipalityCastanuelas           Municipality = "150200" // CASTAÑUELAS
	MunicipalityGuayubin              Municipality = "150300" // GUAYUBÍN
	MunicipalityLasMatasDeSantaCruz   Municipality = "150400" // LAS MATAS DE SANTA CRUZ
	MunicipalityPepilloSalcedo        Municipality = "150500" // PEPILLO SALCEDO
	MunicipalityVillaVasquez          Municipality = "150600" // VILLA VÁSQUEZ

	// Provincia Pedernales (16)
	MunicipalityPedernales Municipality = "160100" // PEDERNALES
	MunicipalityOviedo     Municipality = "160200" // OVIEDO

	// Provincia Peravia (17)
	MunicipalityBani     Municipality = "170100" // BANÍ
	MunicipalityNizao    Municipality = "170200" // NIZAO
	MunicipalityMatanzas Municipality = "170300" // MATANZAS

	// Provincia Puerto Plata (18)
	MunicipalityPuertoPlata     Municipality = "180100" // PUERTO PLATA
	MunicipalityAltamira        Municipality = "180200" // ALTAMIRA
	MunicipalityGuananico       Municipality = "180300" // GUANANICO
	MunicipalityImbert          Municipality = "180400" // IMBERT
	MunicipalityLosHidalgos     Municipality = "180500" // LOS HIDALGOS
	MunicipalityLuperon         Municipality = "180600" // LUPERÓN
	MunicipalitySosua           Municipality = "180700" // SOSÚA
	MunicipalityVillaIsabela    Municipality = "180800" // VILLA ISABELA
	MunicipalityVillaMontellano Municipality = "180900" // VILLA MONTELLANO

	// Provincia Hermanas Mirabal (19)
	MunicipalitySalcedo    Municipality = "190100" // SALCEDO
	MunicipalityTenares    Municipality = "190200" // TENARES
	MunicipalityVillaTapia Municipality = "190300" // VILLA TAPIA

	// Provincia Samaná (20)
	MunicipalitySamana      Municipality = "200100" // SAMANÁ
	MunicipalitySanchez     Municipality = "200200" // SÁNCHEZ
	MunicipalityLasTerrenas Municipality = "200300" // LAS TERRENAS

	// Provincia San Cristóbal (21)
	MunicipalitySanCristobal           Municipality = "210100" // SAN CRISTÓBAL
	MunicipalitySabanaGrandeDePalenque Municipality = "210200" // SABANA GRANDE DE PALENQUE
	MunicipalityBajosDeHaina           Municipality = "210300" // BAJOS DE HAINA
	MunicipalityCambitaGarabitos       Municipality = "210400" // CAMBITA GARABITOS
	MunicipalityVillaAltagracia        Municipality = "210500" // VILLA ALTAGRACIA
	MunicipalityYaguate                Municipality = "210600" // YAGUATE
	MunicipalitySanGregorioDeNigua     Municipality = "210700" // SAN GREGORIO DE NIGUA
	MunicipalityLosCacaos              Municipality = "210800" // LOS CACAOS

	// Provincia San Juan (22)
	MunicipalitySanJuan          Municipality = "220100" // SAN JUAN
	MunicipalityBohechio         Municipality = "220200" // BOHECHÍO
	MunicipalityElCercado        Municipality = "220300" // EL CERCADO
	MunicipalityJuanDeHerrera    Municipality = "220400" // JUAN DE HERRERA
	MunicipalityLasMatasDeFarfan Municipality = "220500" // LAS MATAS DE FARFÁN
	MunicipalityVallejuelo       Municipality = "220600" // VALLEJUELO

	// Provincia San Pedro de Macorís (23)
	MunicipalitySanPedroDeMacoris Municipality = "230100" // SAN PEDRO DE MACORÍS
	MunicipalityLosLlanos         Municipality = "230200" // LOS LLANOS
	MunicipalityRamonSantana      Municipality = "230300" // RAMÓN SANTANA
	MunicipalityConsuelo          Municipality = "230400" // CONSUELO
	MunicipalityQuisqueya         Municipality = "230500" // QUISQUEYA
	MunicipalityGuayacanes        Municipality = "230600" // GUAYACANES

	// Provincia Sánchez Ramírez (24)
	MunicipalityCotui   Municipality = "240100" // COTUÍ
	MunicipalityCevicos Municipality = "240200" // CEVICOS
	MunicipalityFantino Municipality = "240300" // FANTINO
	MunicipalityLaMata  Municipality = "240400" // LA MATA

	// Provincia Santiago (25)
	MunicipalitySantiago           Municipality = "250100" // SANTIAGO
	MunicipalityBisono             Municipality = "250200" // BISONÓ
	MunicipalityJanico             Municipality = "250300" // JÁNICO
	MunicipalityLiceyAlMedio       Municipality = "250400" // LICEY AL MEDIO
	MunicipalitySanJoseDeLasMatas  Municipality = "250500" // SAN JOSÉ DE LAS MATAS
	MunicipalityTamboril           Municipality = "250600" // TAMBORIL
	MunicipalityVillaGonzalez      Municipality = "250700" // VILLA GONZÁLEZ
	MunicipalityPunal              Municipality = "250800" // PUÑAL
	MunicipalitySabanaIglesia      Municipality = "250900" // SABANA IGLESIA
	MunicipalityBaitoa             Municipality = "251000" // BAITOA

	// Provincia Santiago Rodríguez (26)
	MunicipalitySanIgnacioDeSabaneta Municipality = "260100" // SAN IGNACIO DE SABANETA
	MunicipalityVillaLosAlmacigos   Municipality = "260200" // VILLA LOS ALMÁCIGOS
	MunicipalityMoncion             Municipality = "260300" // MONCIÓN

	// Provincia Valverde (27)
	MunicipalityMao          Municipality = "270100" // MAO
	MunicipalityEsperanza    Municipality = "270200" // ESPERANZA
	MunicipalityLagunaSalada Municipality = "270300" // LAGUNA SALADA

	// Provincia Monseñor Nouel (28)
	MunicipalityBonao       Municipality = "280100" // BONAO
	MunicipalityMaimon      Municipality = "280200" // MAIMÓN
	MunicipalityPiedraBlanca Municipality = "280300" // PIEDRA BLANCA

	// Provincia Monte Plata (29)
	MunicipalityMontePlata          Municipality = "290100" // MONTE PLATA
	MunicipalityBayaguana           Municipality = "290200" // BAYAGUANA
	MunicipalitySabanaGrandeDeBoya  Municipality = "290300" // SABANA GRANDE DE BOYÁ
	MunicipalityYamasa              Municipality = "290400" // YAMASÁ
	MunicipalityPeralvillo          Municipality = "290500" // PERALVILLO

	// Provincia Hato Mayor (30)
	MunicipalityHatoMayor       Municipality = "300100" // HATO MAYOR
	MunicipalitySabanaDeLaMar   Municipality = "300200" // SABANA DE LA MAR
	MunicipalityElValle         Municipality = "300300" // EL VALLE

	// Provincia San José de Ocoa (31)
	MunicipalitySanJoseDeOcoa Municipality = "310100" // SAN JOSÉ DE OCOA
	MunicipalitySabanaLarga   Municipality = "310200" // SABANA LARGA
	MunicipalityRanchoArriba  Municipality = "310300" // RANCHO ARRIBA

	// Provincia Santo Domingo (32)
	MunicipalitySantoDomingoEste     Municipality = "320100" // SANTO DOMINGO ESTE
	MunicipalitySantoDomingoOeste    Municipality = "320200" // SANTO DOMINGO OESTE
	MunicipalitySantoDomingoNorte    Municipality = "320300" // SANTO DOMINGO NORTE
	MunicipalityBocaChica            Municipality = "320400" // BOCA CHICA
	MunicipalitySanAntonioDeGuerra   Municipality = "320500" // SAN ANTONIO DE GUERRA
	MunicipalityLosAlcarrizos        Municipality = "320600" // LOS ALCARRIZOS
	MunicipalityPedroBrand           Municipality = "320700" // PEDRO BRAND
)
