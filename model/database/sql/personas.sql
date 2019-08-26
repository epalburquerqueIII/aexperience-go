create table provincia(
    id integer primary key,
    nombre varchar(40))
create table personas(
    id integer primary key,
    nombre varchar(50),
    direccion varchar(100),
    poblacion varchar(50),
    provinciaid integer,
    telefono varchar(50),
    email varchar(50),
    foreign key (provinciaid) references provincias(id)
)


En sistemas operativos, un hilo (del inglés thread), hebra (del inglés fiber), proceso ligero o subproceso es una secuencia de tareas encadenadas muy pequeña que puede ser ejecutada por un sistema operativo. 15112018epa
8 8:30 59
115

1. Despliegue de componentes
- Modelos de despliegue:
? Diseño sin repositorio:
- Diseño y ejecución sin despliegue
- Ejemplos: UML
? Diseño con repositorio sólo para el depósito de componentes:
- Tipos de contenedores
- Ejemplos: EJBs, .NET, CCM, Servicios web
? Despliegue con repositorio:
- Composición y depósito de componentes
- Ejemplo: JavaBean
? Diseño con repositorio:
- Tipos de conectores
- Ejemplos: Koala
2. Selección de componentes
- Tipos:
? Componentes comerciales:
- Sin posibilidad de modificaciones (COTS)
- Con posibilidad de adaptaciones (MOTS)
? Componentes de fuente abierta
? Ventajas e inconvenientes
- Métodos de personalización de componentes:
? Parametrización
? Uso de extensiones (plugins)
- Criterios de selección de componentes reutilizables:
? Adaptabilidad
? Auditabilidad
? Estandarización
? Características de concurrencia
? Rendimiento
? Consumo de recursos
? Seguridad
? Características de mantenimiento y actualización
- Proceso de selección de componentes:
? Evaluación de componentes según requisitos
? Diseño y codificación (código de enlace):
- Enlace de componentes con otros sistemas
- Integración
- Configuración
? Diseño de pruebas
? Detección de fallos
? Mantenimiento y gestión de configuraciones
? Actualización de componentes
? Métodos de selección de uso común:
- CAP (COTS Acquisition Process)
- RUP (Rational Unified Process)
3. Control de calidad de componentes
- Métodos de evaluación de calidad de componentes. Estándares de calidad
- Categorías y métricas de evaluación
- Proceso de validación y medición de calidad:
? Pruebas de conformidad a requisitos funcionales
? Pruebas de integración con otros sistemas
? Pruebas de aspectos no funcionales:
- Rendimiento
- Seguridad
- Integración
- Documentación de componentes
- Descripción funcional
- Descripción de aspectos no funcionales
- Descripción del proceso de instalación y despliegue:
? Descripción del empaquetamiento (packaging)
? Requisitos de implantación
? Parametrización y ajuste