## Desafío Base de Datos

En el sistema de ventas de Fantasy Products, se perdieron los datos de la base de datos, pero antes de que eso fallase, alguien pudo descargar unos archivos .json que hacen referencia a las tablas que se borraron.

Este es el DER del sistema:
![DER del Sistema](https://lh7-us.googleusercontent.com/AWEb1QWEzcMS2ojlm2jpMeKaEqW3Zs9rHB5Fk5aVyOp6srIVdoKkYv9EVll0lrMqHX2_r-wZQDyflAqJXEv4bbUgfVutW3t7J5uDjuZKMJ8gx0uG2KDudAFCxMFT4HeDuf9gN1v5JYe75XXGtL21lMs)

Encontramos una base de proyecto para comenzar dentro de la carpeta data que tiene 4 archivos, es decir, 4 tablas (sales.json, products.json, invoices.json, customers.json), cada uno de ellos tiene los registros de dicha tabla, junto con un script sql para levantar la base de datos.

A su vez, hay una estructura de api con algunos métodos ya implementados **(Create y ReadAll)**.

Se sabe que los campos y el orden de cada tabla son:

- **sales.json** 💸 id, product_id, invoice_id, quantity                                                
- **products.json** 🛒 id,description price                                                        
- **invoices.json** 🧾 id,datetime,customer_id, total                                                
- **customers.json** 👨‍💼 id,last_name,first_name,condition

## Objetivos

Para levantar la base de datos podemos utilizar el siguiente comando en una terminal en la base del proyecto:

**sudo mysql -u root -p -v < ./docs/db/mysql/database.sql**

Como se habrán dado cuenta, la tabla de invoices, perdió el dato de total, por lo tanto es necesario que podamos recalcular con los datos que dispone entre sales, invoices y products.

Realizar las siguientes tareas:

- Crear un app que permita cargar los datos del json al storage respectivo.
- Crear un método endpoint que permita actualizar los datos de invoices.

## Nuevas Consultas a realizar:

1) Montos totales redondeados a 2 decimales por **condition** del **customer**

|Condition|Total|
|---------|-----|
|Inactivo(0)|605929.10|
|Activo(1)|716792.33|

2) Top **5** de los **products** más vendidos y sus cantidades vendidas
Salida esperada

| Description             | Total |
|-------------------------|-------|
| Vinegar - Raspberry     | 660   |
| Flour - Corn, Fine      | 521   |
| Cookie - Oatmeal        | 467   |
| Pepper - Red Chili      | 439   |
| Chocolate - Milk Coating| 436   |

3) El Top 5 de los **customers activos** que gastaron la mayor cantidad de dinero                                                                
Salida esperada
| First Name | Last Name | Amount    |
|------------|-----------|-----------|
| Lannie     | Tortis    | 58513.55  |
| Jasen      | Crowcum   | 48291.03  |
| Elvina     | Ovell     | 43590.75  |
| Lazaro     | Anstis    | 40792.06  |
| Wilden     | Oaten     | 39786.79  |

- Realizar test unitarios sobre las nuevas funcionalidades a incorporar en los **storages** respectivos, utilizando el paquete **go-txdb** (aclaración: algunas consultas utilizan inner join)
- Crear los **handlers** y registrarlos en **endpoints**.
