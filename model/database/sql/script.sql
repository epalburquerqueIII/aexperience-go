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

CREATE Table usuarios_roles (
 id integer PRIMARY KeY NOT NULL,
 nombre varchar(30)NOT NULL );

INSERT INTO `usuarios_roles` (`id`, `nombre`) VALUES (1, 'usuario');
INSERT INTO `usuarios_roles` (`id`, `nombre`) VALUES (0, 'administrador');

CREATE Table menu_parent (
    id integer PRIMARY KEY,
    titulo varchar(50) NOT NULL,
    tipo integer NOT NULL
);
-- Tipo de tabla menu_parent
-- 0-normal
-- 1-desplegable

CREATE Table menus (
    id integer AUTO_INCREMENT PRIMARY KEY,
    parent_id integer NOT NULL,
    orden integer NOT NULL,
    titulo varchar(50) NOT NULL,
    icono varchar(50) NOT NULL,
    url varchar(50) NOT NULL,
    handleFunc varchar(50) NOT NULL,
    FOREIGN KEY (parent_id) REFERENCES menu_parent(id)
    );
    CREATE Table menu_usuarios_roles (
    id integer AUTO_INCREMENT PRIMARY KEY,
    idMenu integer NOT NULL,
    idUsuario_roles integer NOT NULL,
    FOREIGN KEY (idMenu) REFERENCES menus(id),
    FOREIGN KEY (idUsuario_roles) REFERENCES usuarios_roles(id)
    );

    INSERT INTO `menu_parent` (`id`, `titulo`, `tipo`) VALUES (1, 'Estadísticas', 0);
    INSERT INTO `menu_parent` (`id`, `titulo`, `tipo`) VALUES (2, 'Gestión de datos', 1);
    INSERT INTO `menu_parent` (`id`, `titulo`, `tipo`) VALUES (3, 'Perfil', 1);
    INSERT INTO `menu_parent` (`id`, `titulo`, `tipo`) VALUES (4, 'Ajustes', 1);

    INSERT INTO `menus` (`parent_id`,`orden`, `titulo`, `icono`, `url`, `handleFunc`) VALUES (2, 1, 'Estadísticas', 'fas fa-fw fa-chart-bar', 'http://localhost:3000/estadisticas', 'http.HandleFunc("/estadisticas", controller.Estadisticas)' );
    INSERT INTO `menus` (`parent_id`,`orden`, `titulo`, `icono`, `url`, `handleFunc`) VALUES (3, 1, 'IVA', 'fas fa-fw fa-cog', 'http://localhost:3000/iva', 'http.HandleFunc("/iva", controller.Iva)' );
    INSERT INTO `menus` (`parent_id`,`orden`, `titulo`, `icono`, `url`, `handleFunc`) VALUES (4, 1, 'Iniciar sesión', 'fas fa-fw fa-user', 'http://localhost:3000/login', 'http.HandleFunc("/login", controller.Login)' );
    INSERT INTO `menus` (`parent_id`,`orden`, `titulo`, `icono`, `url`, `handleFunc`) VALUES (4, 2, 'Registro', 'fas fa-fw fa-user', 'http://localhost:3000/registro', 'http.HandleFunc("/registro", controller.Registro)' );
    INSERT INTO `menus` (`parent_id`,`orden`, `titulo`, `icono`, `url`, `handleFunc`) VALUES (4, 3, 'Recuperar Contraseña', 'fas fa-fw fa-user', 'http://localhost:3000/olvido-contrasena', 'http.HandleFunc("/olvido-contrasena", controller.Olvidocontrasena)' );
    INSERT INTO `menus` (`parent_id`,`orden`, `titulo`, `icono`, `url`, `handleFunc`) VALUES (1, 1, 'Usuarios', 'fas fa-fw fa-file-signature', 'http://localhost:3000/usuario', 'http.HandleFunc("/usuario", controller.Usuario)' );
    INSERT INTO `menus` (`parent_id`,`orden`, `titulo`, `icono`, `url`, `handleFunc`) VALUES (1, 2, 'Autorizados', 'fas fa-fw fa-file-signature', 'http://localhost:3000/autorizados', 'http.HandleFunc("/autorizado", controller.Autorizado)' );
    INSERT INTO `menus` (`parent_id`,`orden`, `titulo`, `icono`, `url`, `handleFunc`) VALUES (1, 3, 'Consumo Bonos', 'fas fa-fw fa-file-signature', 'http://localhost:3000/consumoBonos', 'http.HandleFunc("/consumoBonos", controller.ConsumoBonos)' );
    INSERT INTO `menus` (`parent_id`,`orden`, `titulo`, `icono`, `url`, `handleFunc`) VALUES (1, 4, 'Espacios', 'fas fa-fw fa-file-signature', 'http://localhost:3000/espacios', 'http.HandleFunc("/espacios", controller.Espacios)' );
    INSERT INTO `menus` (`parent_id`,`orden`, `titulo`, `icono`, `url`, `handleFunc`) VALUES (1, 5, 'Reservas', 'fas fa-fw fa-file-signature', 'http://localhost:3000/reservas', 'http.HandleFunc("/reservas", controller.Reservas)' );
    INSERT INTO `menus` (`parent_id`,`orden`, `titulo`, `icono`, `url`, `handleFunc`) VALUES (1, 6, 'Horarios', 'fas fa-fw fa-file-signature',
    'http://localhost:3000/horarios', 'http.HandleFunc("/horarios", controller.Horarios)' );
    INSERT INTO `menus` (`parent_id`,`orden`, `titulo`, `icono`, `url`, `handleFunc`) VALUES (1, 7, 'Tipos Pagos', 'fas fa-fw fa-file-signature', 'http://localhost:3000/tiposPagos', 'http.HandleFunc("/tiposPagos", controller.TiposPagos)' );
    INSERT INTO `menus` (`parent_id`,`orden`, `titulo`, `icono`, `url`, `handleFunc`) VALUES (1, 8, 'Tipos Eventos', 'fas fa-fw fa-file-signature', 'http://localhost:3000/tiposEventos', 'http.HandleFunc("/tiposEventos", controller.TiposEventos)' );
    INSERT INTO `menus` (`parent_id`,`orden`, `titulo`, `icono`, `url`, `handleFunc`) VALUES (1, 9, 'Bonos', 'fas fa-fw fa-file-signature', 'http://localhost:3000/bonos', 'http.HandleFunc("/bonos", controller.Bonos)' );
    INSERT INTO `menus` (`parent_id`,`orden`, `titulo`, `icono`, `url`, `handleFunc`) VALUES (1, 10, 'Pagos', 'fas fa-fw fa-file-signature', 'http://localhost:3000/pagos', 'http.HandleFunc("/pagos", controller.Pagos)' );
    INSERT INTO `menus` (`parent_id`,`orden`, `titulo`, `icono`, `url`, `handleFunc`) VALUES (1, 11, 'Usuarios_roles', 'fas fa-fw fa-file-signature', 'http://localhost:3000/usuarios_roles', 'http.HandleFunc("/usuarios_roles", controller.Usuarios_roles)' );
    INSERT INTO `menus` (`parent_id`,`orden`, `titulo`, `icono`, `url`, `handleFunc`) VALUES (1, 12, 'Menu_Usuarios_roles', 'fas fa-fw fa-file-signature', 'http://localhost:3000/menu_usuarios_roles', 'http.HandleFunc("/menu_usuarios_roles", controller.Menu_usuarios_roles)' );