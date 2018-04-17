CREATE TABLE "movilidad" (
"id" serial4 NOT NULL,
"fecha_inicio" varchar(255) NOT NULL,
"fecha_fin" varchar(255) NOT NULL,
"nombre_acto_administrativo" varchar(255) NOT NULL,
"enlace_acto_administrativo" varchar(255) NOT NULL,
"area_conocimiento" varchar(255),
"concepto_participacion" varchar(255),
"pais" int4 NOT NULL,
"tipo_movilidad" int4 NOT NULL,
"clasificacion_duracion" int4 NOT NULL,
"categoria_movilidad" int4 NOT NULL,
"institucion" int4 NOT NULL,
"convenio" int4,
"dependencia" int4,
"tipo_persona" int4 NOT NULL,
"clasificacion_movilidad" int4 NOT NULL,
"persona" int4 NOT NULL,
PRIMARY KEY ("id") 
)
WITHOUT OIDS;

CREATE TABLE "tipo_movilidad" (
"id" serial4 NOT NULL,
"nombre" varchar(255) NOT NULL,
PRIMARY KEY ("id") 
)
WITHOUT OIDS;

CREATE TABLE "clasificacion_duracion" (
"id" serial4 NOT NULL,
"nombre" varchar(255) NOT NULL,
"criterio" varchar(255) NOT NULL,
PRIMARY KEY ("id") 
)
WITHOUT OIDS;

CREATE TABLE "categoria_movilidad" (
"id" serial4 NOT NULL,
"nombre" varchar(255) NOT NULL,
"disponible" boolean,
"descripcion" varchar(255),
PRIMARY KEY ("id") 
)
WITHOUT OIDS;

CREATE TABLE "presupuesto" (
"id" serial4 NOT NULL,
"descripcion" varchar(255) NOT NULL,
"movilidad" int4 NOT NULL,
"tipo_presupuesto" int4 NOT NULL,
PRIMARY KEY ("id") 
)
WITHOUT OIDS;

CREATE TABLE "tipo_presupuesto" (
"id" serial4 NOT NULL,
"nombre" varchar(255) NOT NULL,
PRIMARY KEY ("id") 
)
WITHOUT OIDS;

CREATE TABLE "convenio" (
"id" serial4 NOT NULL,
"nombre" varchar(255) NOT NULL,
"disponible" boolean,
"descripcion" varchar(255),
PRIMARY KEY ("id") 
)
WITHOUT OIDS;

CREATE TABLE "tipo_persona" (
"id" serial4 NOT NULL,
"nombre" varchar(255) NOT NULL,
PRIMARY KEY ("id") 
)
WITHOUT OIDS;

CREATE TABLE "clasificacion_movilidad" (
"id" serial4 NOT NULL,
"nombre" varchar(255) NOT NULL,
PRIMARY KEY ("id") 
)
WITHOUT OIDS;


ALTER TABLE "movilidad" ADD CONSTRAINT "tipo_movilidad" FOREIGN KEY ("tipo_movilidad") REFERENCES "tipo_movilidad" ("id");
ALTER TABLE "movilidad" ADD CONSTRAINT "clasificacion_duracion" FOREIGN KEY ("clasificacion_duracion") REFERENCES "clasificacion_duracion" ("id");
ALTER TABLE "movilidad" ADD CONSTRAINT "categoria_movilidad" FOREIGN KEY ("categoria_movilidad") REFERENCES "categoria_movilidad" ("id");
ALTER TABLE "presupuesto" ADD CONSTRAINT "movilidad" FOREIGN KEY ("movilidad") REFERENCES "movilidad" ("id");
ALTER TABLE "presupuesto" ADD CONSTRAINT "tipo_presupuesto" FOREIGN KEY ("tipo_presupuesto") REFERENCES "tipo_presupuesto" ("id");
ALTER TABLE "movilidad" ADD CONSTRAINT "convenio" FOREIGN KEY ("convenio") REFERENCES "convenio" ("id");
ALTER TABLE "movilidad" ADD CONSTRAINT "tipo_persona" FOREIGN KEY ("tipo_persona") REFERENCES "tipo_persona" ("id");
ALTER TABLE "movilidad" ADD CONSTRAINT "clasificacion_movilidad" FOREIGN KEY ("clasificacion_movilidad") REFERENCES "clasificacion_movilidad" ("id");

