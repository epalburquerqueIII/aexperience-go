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
    newletter tinyint,
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

CREATE Table menus (
    id integer AUTO_INCREMENT PRIMARY KEY,
    parent_id integer NOT NULL,
    orden integer NOT NULL,
    titulo varchar(50) NOT NULL,
    icono varchar(50) NOT NULL,
    url varchar(50) NOT NULL,
    hanledFunc varchar(50) NOT NULL
    );

CREATE Table menu_usuarios_roles (
    id integer AUTO_INCREMENT PRIMARY KEY,
    idMenu integer NOT NULL,
    idUsuario integer NOT NULL,
    FOREIGN KEY (idMenu) REFERENCES menus(id),
    FOREIGN KEY (idUsuario) REFERENCES usuarios(id)
    );

