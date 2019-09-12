CREATE Table bonos (
    precio integer PRIMARY KEY,
    sesiones integer NOT NULL
    );
INSERT INTO `bonos` (`precio`, `sesiones`) VALUES (10, 12);
INSERT INTO `bonos` (`precio`, `sesiones`) VALUES (25, 25);
INSERT INTO `bonos` (`precio`, `sesiones`) VALUES (50, 70);

CREATE TABLE usuarios (
    id int AUTO_INCREMENT PRIMARY KEY,
    nombre varchar(50) NOT NULL,
    nif varchar(30) NOT NULL,
    email varchar(30) NOT NULL,
    tipo integer not null,
    telefono varchar(30) NOT NULL,
    sesionesBonos integer,
    newsletter tinyint,
    fechaBaja date
);
CREATE Table tiposPago (
    id integer AUTO_INCREMENT PRIMARY KEY,
    nombre varchar(30)
    );
CREATE Table tiposEvento (
    id integer AUTO_INCREMENT PRIMARY KEY,
    nombre varchar(30)
    );
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
    fechaPago date NOT NULL,
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
    numeroTarjeta varchar(50) NOT NULL,
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
CREATE Table menuUsuariosRoles (
    id integer AUTO_INCREMENT PRIMARY KEY,
    idMenu integer NOT NULL,
    idUsuarioRoles integer NOT NULL,
    FOREIGN KEY (idMenu) REFERENCES menus(id),
    FOREIGN KEY (idUsuarioRoles) REFERENCES usuariosRoles(id)
);

    INSERT INTO `menuParent` (`id`, `titulo`, `tipo`) VALUES (1, 'Estadísticas', 0);
    INSERT INTO `menuParent` (`id`, `titulo`, `tipo`) VALUES (2, 'Gestión de datos', 1);
    INSERT INTO `menuParent` (`id`, `titulo`, `tipo`) VALUES (3, 'Perfil', 1);
    INSERT INTO `menuParent` (`id`, `titulo`, `tipo`) VALUES (4, 'Ajustes', 1);

    INSERT INTO `menus` (`parentId`,`orden`, `titulo`, `icono`, `url`, `handleFunc`) VALUES (1, 1, 'Estadísticas', 'fas fa-fw fa-chart-bar', 'http://localhost:3000/estadisticas', 'http.HandleFunc("/estadisticas", controller.Estadisticas)' );

    INSERT INTO `menus` (`parentId`,`orden`, `titulo`, `icono`, `url`, `handleFunc`) VALUES (2, 1, 'Usuarios', 'fas fa-fw fa-file-signature', 'http://localhost:3000/usuario', 'http.HandleFunc("/usuario", controller.Usuario)' );
    INSERT INTO `menus` (`parentId`,`orden`, `titulo`, `icono`, `url`, `handleFunc`) VALUES (2, 2, 'Autorizados', 'fas fa-fw fa-file-signature', 'http://localhost:3000/autorizados', 'http.HandleFunc("/autorizado", controller.Autorizado)' );
    INSERT INTO `menus` (`parentId`,`orden`, `titulo`, `icono`, `url`, `handleFunc`) VALUES (2, 3, 'Consumo Bonos', 'fas fa-fw fa-file-signature', 'http://localhost:3000/consumoBonos', 'http.HandleFunc("/consumoBonos", controller.ConsumoBonos)' );
    INSERT INTO `menus` (`parentId`,`orden`, `titulo`, `icono`, `url`, `handleFunc`) VALUES (2, 4, 'Espacios', 'fas fa-fw fa-file-signature', 'http://localhost:3000/espacios', 'http.HandleFunc("/espacios", controller.Espacios)' );
    INSERT INTO `menus` (`parentId`,`orden`, `titulo`, `icono`, `url`, `handleFunc`) VALUES (2, 5, 'Reservas', 'fas fa-fw fa-file-signature', 'http://localhost:3000/reservas', 'http.HandleFunc("/reservas", controller.Reservas)' );
    INSERT INTO `menus` (`parentId`,`orden`, `titulo`, `icono`, `url`, `handleFunc`) VALUES (2, 6, 'Horarios', 'fas fa-fw fa-file-signature',
    'http://localhost:3000/horarios', 'http.HandleFunc("/horarios", controller.Horarios)' );
    INSERT INTO `menus` (`parentId`,`orden`, `titulo`, `icono`, `url`, `handleFunc`) VALUES (2, 7, 'Tipos Pagos', 'fas fa-fw fa-file-signature', 'http://localhost:3000/tiposPagos', 'http.HandleFunc("/tiposPagos", controller.TiposPagos)' );
    INSERT INTO `menus` (`parentId`,`orden`, `titulo`, `icono`, `url`, `handleFunc`) VALUES (2, 8, 'Tipos Eventos', 'fas fa-fw fa-file-signature', 'http://localhost:3000/tiposEventos', 'http.HandleFunc("/tiposEventos", controller.TiposEventos)' );
    INSERT INTO `menus` (`parentId`,`orden`, `titulo`, `icono`, `url`, `handleFunc`) VALUES (2, 9, 'Bonos', 'fas fa-fw fa-file-signature', 'http://localhost:3000/bonos', 'http.HandleFunc("/bonos", controller.Bonos)' );
    INSERT INTO `menus` (`parentId`,`orden`, `titulo`, `icono`, `url`, `handleFunc`) VALUES (2, 10, 'Pagos', 'fas fa-fw fa-file-signature', 'http://localhost:3000/pagos', 'http.HandleFunc("/pagos", controller.Pagos)' );
    INSERT INTO `menus` (`parentId`,`orden`, `titulo`, `icono`, `url`, `handleFunc`) VALUES (2, 11, 'UsuariosRoles', 'fas fa-fw fa-file-signature', 'http://localhost:3000/usuariosRoles', 'http.HandleFunc("/usuariosRoles", controller.UsuariosRoles)' );
    INSERT INTO `menus` (`parentId`,`orden`, `titulo`, `icono`, `url`, `handleFunc`) VALUES (2, 12, 'MenuUsuariosRoles', 'fas fa-fw fa-file-signature', 'http://localhost:3000/menuUsuariosRoles', 'http.HandleFunc("/menuUsuariosRoles", controller.MenuUsuariosRoles)' );

    INSERT INTO `menus` (`parentId`,`orden`, `titulo`, `icono`, `url`, `handleFunc`) VALUES (3, 1, 'Iniciar sesión', 'fas fa-fw fa-user', 'http://localhost:3000/login', 'http.HandleFunc("/login", controller.Login)' );
    INSERT INTO `menus` (`parentId`,`orden`, `titulo`, `icono`, `url`, `handleFunc`) VALUES (3, 2, 'Registro', 'fas fa-fw fa-user', 'http://localhost:3000/registro', 'http.HandleFunc("/registro", controller.Registro)' );
    INSERT INTO `menus` (`parentId`,`orden`, `titulo`, `icono`, `url`, `handleFunc`) VALUES (3, 3, 'Recuperar Contraseña', 'fas fa-fw fa-user', 'http://localhost:3000/olvido-contrasena', 'http.HandleFunc("/olvido-contrasena", controller.Olvidocontrasena)' );

    INSERT INTO `menus` (`parentId`,`orden`, `titulo`, `icono`, `url`, `handleFunc`) VALUES (4, 1, 'Ajustes', 'fas fa-fw fa-cog', 'http://localhost:3000/ajustes', 'http.HandleFunc("/ajustes", controller.Ajustes)' );

    