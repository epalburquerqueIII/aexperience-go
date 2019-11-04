CREATE Table bonos (
    id integer NOT NULL PRIMARY KEY,
    precio float,
    sesiones integer NOT NULL
    );
INSERT INTO `bonos` (`id`, `precio`, `sesiones`) VALUES (0, 10, 12);
INSERT INTO `bonos` (`id`, `precio`, `sesiones`) VALUES (1, 25, 25);
INSERT INTO `bonos` (`id`, `precio`, `sesiones`) VALUES (2, 50, 70);
INSERT INTO `bonos` (`id`, `precio`, `sesiones`) VALUES (3, 3, 1);
INSERT INTO `bonos` (`id`, `precio`, `sesiones`) VALUES (4, 1.5, 1);


CREATE TABLE usuarios (
    id int AUTO_INCREMENT PRIMARY KEY,
    nombre varchar(50) NOT NULL,
    nif varchar(30) NOT NULL,
    email varchar(30) NOT NULL,
    fechaNacimiento date NOT NULL,
    idUsuarioRol integer not null,
    telefono varchar(30) NOT NULL,
    Password varchar(100) NOT NULL,
    sesionesBonos integer,
    newsletter tinyint,
    fechaBaja date
);

--Newsletter
-- 0 = no
-- 1 = si


CREATE Table tiposPago (
    id integer AUTO_INCREMENT PRIMARY KEY,
    nombre varchar(30)
    );
INSERT INTO `tiposPago` (`id`, `nombre`) VALUES (1, "Efectivo");
INSERT INTO `tiposPago` (`id`, `nombre`) VALUES (2, "Transferencia");


CREATE Table tiposEvento (
    id integer AUTO_INCREMENT PRIMARY KEY,
    nombre varchar(30)
    );

    -- tabla de espacios y eventos
CREATE Table espacios (
    id integer AUTO_INCREMENT PRIMARY KEY, 
    descripcion varchar(50), 
    estado tinyint NOT NULL, 
    modo tinyint NOT NULL, 
    precio double NOT NULL,
    idTipoevento integer, 
    fecha date NOT NULL, 
    aforo integer,
    numeroReservaslimite integer NOT NULL,
    FOREIGN KEY (idTipoevento) REFERENCES tiposevento(id) 
);
CREATE Table autorizados (
    id integer AUTO_INCREMENT PRIMARY KEY,
    idUsuario integer,
    nombreAutorizado varchar(30) NOT NULL,
    nif varchar(30) NOT NULL,
    FOREIGN KEY (idUsuario) REFERENCES usuarios(id)
    );
CREATE Table consumoBonos (
    id integer AUTO_INCREMENT PRIMARY KEY,
    fecha date NOT NULL,
    sesiones integer NOT NULL,
    idUsuario integer,
    idEspacio integer,
    idAutorizado integer,
    FOREIGN KEY (idUsuario) REFERENCES usuarios(id),
    FOREIGN KEY (idEspacio) REFERENCES espacios(id),
    FOREIGN KEY (idAutorizado) REFERENCES autorizados(id)
    );
CREATE Table reservas (
    id integer AUTO_INCREMENT PRIMARY KEY,
    fecha date NOT NULL,
    Sesiones integer,
    hora integer,
    idUsuario integer,
    idEspacio integer,
    idAutorizado integer,
    FOREIGN KEY (idUsuario) REFERENCES usuarios(id),
    FOREIGN KEY (idEspacio) REFERENCES espacios(id),
    FOREIGN KEY (idAutorizado) REFERENCES autorizados(id)
    );
CREATE Table horarios (
    id integer AUTO_INCREMENT PRIMARY KEY,
    idEspacio integer,
    descripcion varchar(50) NOT NULL,
    fechaInicio date NOT NULL,
    fechaFin date NOT NULL,
    hora integer NOT NULL,
    FOREIGN KEY (idEspacio) REFERENCES espacios(id)
    );
CREATE Table pagos (
    id integer AUTO_INCREMENT PRIMARY KEY,
    idReserva integer,
    fechaPago date NOT NULL,
    idTipopago integer,
    importe float,
    numeroTarjeta varchar(50) NOT NULL,
    FOREIGN KEY (idReserva) REFERENCES reservas(id),
    FOREIGN KEY (idTipopago) REFERENCES tiposPago(id)
    );

CREATE Table pagosPendientes (
    id integer AUTO_INCREMENT PRIMARY KEY,
    idReserva integer,
    fechaPago date NOT NULL,
    idTipopago integer,
    numeroTarjeta varchar(50) NOT NULL,
    importe float NOT NULL, 
    FOREIGN KEY (idReserva) REFERENCES reservas(id),
    FOREIGN KEY (idTipopago) REFERENCES tiposPago(id)
    );

CREATE Table usuariosRoles (
 id integer PRIMARY KeY NOT NULL,
 nombre varchar(30)NOT NULL );

INSERT INTO `usuariosRoles` (`id`, `nombre`) VALUES (1, 'usuario');
INSERT INTO `usuariosRoles` (`id`, `nombre`) VALUES (0, 'administrador');

CREATE Table menuParent (
    id integer PRIMARY KEY,
    titulo varchar(50) NOT NULL,
    tipo integer NOT NULL
);
-- Tipo de tabla menuParent
-- 0-normal
-- 1-desplegable

CREATE Table menus (
    id integer AUTO_INCREMENT PRIMARY KEY,
    parentId integer NOT NULL,
    orden integer NOT NULL,
    titulo varchar(50) NOT NULL,
    icono varchar(50) NOT NULL,
    url varchar(50) NOT NULL,
    handleFunc varchar(50) NOT NULL,
    FOREIGN KEY (parentId) REFERENCES menuParent(id)
);
CREATE TABLE tiponoticias (
    id integer auto_increment primary key,
	nombre varchar(50) not null);
  
INSERT INTO `tiponoticias`(`id`, `nombre`) VALUES (1,'Deportes');
INSERT INTO `tiponoticias`(`id`, `nombre`) VALUES (2,'Cultura');
INSERT INTO `tiponoticias`(`id`, `nombre`) VALUES (3,'Eventos');
INSERT INTO `tiponoticias`(`id`, `nombre`) VALUES (4,'Noticias');
INSERT INTO `tiponoticias`(`id`, `nombre`) VALUES (5,'Musica');
INSERT INTO `tiponoticias`(`id`, `nombre`) VALUES (6,'Actividades');
INSERT INTO `tiponoticias`(`id`, `nombre`) VALUES (7,'Ferias');
INSERT INTO `tiponoticias`(`id`, `nombre`) VALUES (8,'Naturaleza');
INSERT INTO `tiponoticias`(`id`, `nombre`) VALUES (9,'Fiestas Regionales');
INSERT INTO `tiponoticias`(`id`, `nombre`) VALUES (10,'Mancomunidad Lácara-Los Baldíos');

CREATE TABLE newsletter (
	id integer auto_increment primary key,
	email varchar(50) not null,
	idtiponoticias integer not null,
	FOREIGN key (idtiponoticias) references tiponoticias(id));

CREATE Table menuUsuariosRoles (
    id integer AUTO_INCREMENT PRIMARY KEY,
    idMenu integer NOT NULL,
    idUsuarioRoles integer NOT NULL,
    FOREIGN KEY (idMenu) REFERENCES menus(id),
    FOREIGN KEY (idUsuarioRoles) REFERENCES usuariosRoles(id)
);

    INSERT INTO `menuParent` (`id`, `titulo`, `tipo`) VALUES (1, 'Inicio', 0);
    INSERT INTO `menuParent` (`id`, `titulo`, `tipo`) VALUES (2, 'Estadísticas', 0);
    INSERT INTO `menuParent` (`id`, `titulo`, `tipo`) VALUES (3, 'Gestión de datos', 1);
    INSERT INTO `menuParent` (`id`, `titulo`, `tipo`) VALUES (4, 'Perfil', 1);
    INSERT INTO `menuParent` (`id`, `titulo`, `tipo`) VALUES (5, 'Ajustes', 1);

    INSERT INTO `menus` (`parentId`,`orden`, `titulo`, `icono`, `url`, `handleFunc`) VALUES (1, 1, 'Inicio', 'fas fa-fw fa-home', 'http://localhost:3000/index', 'http.HandleFunc("/", index)' );
    INSERT INTO `menus` (`parentId`,`orden`, `titulo`, `icono`, `url`, `handleFunc`) VALUES (2, 1, 'Estadísticas', 'fas fa-fw fa-chart-bar', 'http://localhost:3000/estadisticas', 'http.HandleFunc("/estadisticas", controller.Estadisticas)' );

    INSERT INTO `menus` (`parentId`,`orden`, `titulo`, `icono`, `url`, `handleFunc`) VALUES (3, 1, 'Usuarios', 'fas fa-fw fa-file-signature', 'http://localhost:3000/usuarios', 'http.HandleFunc("/usuarios", controller.Usuarios)' );
    INSERT INTO `menus` (`parentId`,`orden`, `titulo`, `icono`, `url`, `handleFunc`) VALUES (3, 2, 'Autorizados', 'fas fa-fw fa-file-signature', 'http://localhost:3000/autorizados', 'http.HandleFunc("/autorizados", controller.Autorizados)' );
    INSERT INTO `menus` (`parentId`,`orden`, `titulo`, `icono`, `url`, `handleFunc`) VALUES (3, 3, 'Consumo de Bonos', 'fas fa-fw fa-file-signature', 'http://localhost:3000/consumobonos', 'http.HandleFunc("/consumobonos", controller.ConsumoBonos)' );
    INSERT INTO `menus` (`parentId`,`orden`, `titulo`, `icono`, `url`, `handleFunc`) VALUES (3, 4, 'Espacios', 'fas fa-fw fa-file-signature', 'http://localhost:3000/espacios', 'http.HandleFunc("/espacios", controller.Espacios)' );
    INSERT INTO `menus` (`parentId`,`orden`, `titulo`, `icono`, `url`, `handleFunc`) VALUES (3, 5, 'Reservas', 'fas fa-fw fa-file-signature', 'http://localhost:3000/reservas', 'http.HandleFunc("/reservas", controller.Reservas)' );
    INSERT INTO `menus` (`parentId`,`orden`, `titulo`, `icono`, `url`, `handleFunc`) VALUES (3, 6, 'Horarios', 'fas fa-fw fa-file-signature', 'http://localhost:3000/horarios', 'http.HandleFunc("/horarios", controller.Horarios)' );
     
    INSERT INTO `menus` (`parentId`,`orden`, `titulo`, `icono`, `url`, `handleFunc`) VALUES (3, 8, 'Tipos de eventos', 'fas fa-fw fa-file-signature', 'http://localhost:3000/tiposeventos', 'http.HandleFunc("/tiposeventos", controller.Tiposeventos)' );
    INSERT INTO `menus` (`parentId`,`orden`, `titulo`, `icono`, `url`, `handleFunc`) VALUES (3, 9, 'Bonos', 'fas fa-fw fa-file-signature', 'http://localhost:3000/bonos', 'http.HandleFunc("/bonos", controller.Bonos)' );
    INSERT INTO `menus` (`parentId`,`orden`, `titulo`, `icono`, `url`, `handleFunc`) VALUES (3, 10, 'Pagos', 'fas fa-fw fa-file-signature', 'http://localhost:3000/pagos', 'http.HandleFunc("/pagos", controller.Pagos)' );
    INSERT INTO `menus` (`parentId`,`orden`, `titulo`, `icono`, `url`, `handleFunc`) VALUES (3, 11, 'Roles', 'fas fa-fw fa-file-signature', 'http://localhost:3000/usuariosroles', 'http.HandleFunc("/usuariosroles", controller.UsuariosRoles)' );
    INSERT INTO `menus` (`parentId`,`orden`, `titulo`, `icono`, `url`, `handleFunc`) VALUES (3, 12, 'Roles por menú', 'fas fa-fw fa-file-signature', 'http://localhost:3000/menuroles', 'http.HandleFunc("/menuroles", controller.MenuRoles)' );
    INSERT INTO `menus` (`parentId`,`orden`, `titulo`, `icono`, `url`, `handleFunc`) VALUES (3, 13, 'Newsletter', 'fas fa-fw fa-file-signature', 'http://localhost:3000/newsletter', 'http.HandleFunc("/newsletter", controller.Newsletter)' );
    INSERT INTO `menus` (`parentId`,`orden`, `titulo`, `icono`, `url`, `handleFunc`) VALUES (3, 14, 'Pagos Pendientes', 'fas fa-fw fa-file-signature', 'http://localhost:3000/pagospendientes', 'http.HandleFunc("/pagospendientes", controller.PagosPendientes)' );
    INSERT INTO `menus` (`parentId`,`orden`, `titulo`, `icono`, `url`, `handleFunc`) VALUES (3, 15, 'Menus', 'fas fa-fw fa-file-signature', 'http://localhost:3000/menus', 'http.HandleFunc("/menus", controller.Menus)' );
   
    INSERT INTO `menus` (`parentId`,`orden`, `titulo`, `icono`, `url`, `handleFunc`) VALUES (4, 1, 'Iniciar sesión', 'fas fa-fw fa-user', 'http://localhost:3000/login', 'http.HandleFunc("/login", controller.Login)' );
    INSERT INTO `menus` (`parentId`,`orden`, `titulo`, `icono`, `url`, `handleFunc`) VALUES (4, 2, 'Registro', 'fas fa-fw fa-user', 'http://localhost:3000/registro', 'http.HandleFunc("/registro", controller.Registro)' );
    INSERT INTO `menus` (`parentId`,`orden`, `titulo`, `icono`, `url`, `handleFunc`) VALUES (4, 3, 'Recuperar Contraseña', 'fas fa-fw fa-user', 'http://localhost:3000/recuperarcontrasena', 'http.HandleFunc("/recuperarcontrasena", controller.Recuperarcontrasena)' );

    INSERT INTO `menus` (`parentId`,`orden`, `titulo`, `icono`, `url`, `handleFunc`) VALUES (5, 1, 'Ajustes', 'fas fa-fw fa-cog', 'http://localhost:3000/ajustes', 'http.HandleFunc("/ajustes", controller.Ajustes)' );



