select * from movies_db.movies;

select m.id,m.title, a.first_name,a.last_name, a.favorite_movie_id 
from movies m
inner join actors a
on m.id = a.favorite_movie_id;


#LEFT

SELECT m.title as "Pelicula Favorita",concat( a.first_name," ",a.last_name) as "ACTOR"
from movies as m
left join actors a on m.id = a.favorite_movie_id
where a.favorite_movie_id is null;

#RIGHT

SELECT m.title as "Pelicula Favorita",concat( a.first_name," ",a.last_name) as "ACTOR"
from movies as m
right join actors a on m.id = a.favorite_movie_id
where a.favorite_movie_id is null;

## Otras consultas

select m.title,Sum(m.awards) as total_awards
from movies m
group by id having total_awards = 0;


# Seleccionar el nombre, el puesto y la localidad de los departamentos donde trabajan los vendedores.
select e.nombre, e.puesto,d.localidad
from empleado e 
inner join departamento d
on e.depto_nro = d.depto_nro;

# Visualizar los departamentos con más de cinco empleados.


# Mostrar el nombre, salario y nombre del departamento de los empleados que tengan el mismo puesto que ‘Mito Barchuk’.
select e.nombre, e.salario, d.nombre_depto 
from empleado e
inner join departamento d
on e.depto_nro = d.depto_nro where e.puesto in (select puesto from empleado where nombre = "Mito" and apellido = "Barchuk");
