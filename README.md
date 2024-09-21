-- Tabla Clientes
CREATE TABLE Clientes (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL,
    telefono VARCHAR(20)
);

-- Tabla Productos
CREATE TABLE Productos (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    precio DECIMAL(10, 2) NOT NULL,
    stock INT NOT NULL
);

-- Tabla Órdenes
CREATE TABLE Ordenes (
    id INT AUTO_INCREMENT PRIMARY KEY,
    cliente_id INT NOT NULL,
    fecha DATE NOT NULL,
    total DECIMAL(10, 2),
    FOREIGN KEY (cliente_id) REFERENCES Clientes(id)
);

-- Tabla Detalle de Órdenes
CREATE TABLE DetalleOrdenes (
    id INT AUTO_INCREMENT PRIMARY KEY,
    orden_id INT NOT NULL,
    producto_id INT NOT NULL,
    cantidad INT NOT NULL,
    precio_unitario DECIMAL(10, 2) NOT NULL,
    FOREIGN KEY (orden_id) REFERENCES Ordenes(id) ON DELETE CASCADE,
    FOREIGN KEY (producto_id) REFERENCES Productos(id)
);

endpoint: ("/orden") post
ejemplo de json:

{
  "cliente_id": 1,
  "fecha": "2024-09-20",
  "total": 150.75,
  "detalles": [
    {
      "producto_id": 1,
      "cantidad": 2,
      "precio_unitario": 50.25
    },
    {
      "producto_id": 2,
      "cantidad": 1,
      "precio_unitario": 50.50
    }
  ]
}

endpoint: ("/productos/mas-vendidos")
ejemplo de json:
{
    "fecha_inicio": "2024-01-01",
    "fecha_fin": "2024-09-30"
}
