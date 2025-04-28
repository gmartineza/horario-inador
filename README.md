# Horario-inador - Especificación de requerimientos de software

## 1. Introducción

### 1.1 Propósito

El propósito de este proyecto es crear una aplicación a la que se le pueda pasar un json con horarios de clases para generar una tabla con formato intuitivo, para luego crear una imagen en distintas relaciones de aspecto, pensada para ser usada como fondo de pantalla en smartphones.

### 1.2. Alcance

La aplicación será desarrollada en dos etapas. La primera instancia tendrá foco en la funcionalidad principal: scrapear el archivojson, generar la tabla de horarios, generar una imagen con dicha tabla. La segunda instancia se centrará en facilitar el acceso a la aplicación desde un smartphone [método a definir].

## 2. Descripción general

### 2.1 Perspectiva del producto

El generador de imágenes será una aplicación independiente que obtiene el horario de clases del usuario de un archivo JSON y generará una imagen personalizada pensada para ser usada como fondo de pantalla en el smartphone de dicho usuario.

### 2.2 Funciones del producto

1. Scraping de datos: la aplicación leerá el horario de clases del usuario de un archivo JSON.
2. Generación de imagen: la aplicación generará una imagen basada en los datos de los horarios, considerando la relación de aspecto usual de pantallas de smartphones. [a confirmar: También permitirá establecer una relación de aspecto personalizada].
3. Guardado de imagen: la aplicación permitirá guardar la imagen generada.

## 3. Requerimientos específicos

### 3.1 Requerimientos funcionales

1. Scraping de datos:
  - La aplicación deberá leer el horario de clases del usuario de un archivo JSON.
  - La aplicación deberá manejar cualquier error o dato no válido en el archivo JSON.
2. Generación de imagen:
  - La aplicación deberá generar una imagen con la relación de aspecto adecuada para pantallas de smartphone.
  - La imagen deberá mostrar el horario de clases del usuario de una manera clara y visualmente atractiva.
3. Guardado de imagen:
  - La aplicación deberá guardar la imagen generada en un archivo.

### 3.2 Requerimientos no funcionales

1. Usabilidad:
  - La aplicación deberá tener una interfaz de usuario simple e intuitiva.
2. Rendimiento:
  - La aplicación deberá generar la imagen rápidamente, con demora mínima para el usuario.

## 4. Información adicional

Este es un proyecto personal con fines de  experimentación y aprendizaje. El enfoque está en practicar principios de la ingeniería de software y explorar el lenguaje de programación Go.
