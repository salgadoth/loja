create table public.produtos(
	id serial primary key,
	nome varchar,
	descricao varchar,
	preco decimal,
	quantidade integer
)

insert into public.produtos(nome, descricao, preco, quantidade) values 
	('Camiseta', 'Preta', 19.00, 10),
	('Fone', 'Surround sound', 90.00, 5);
	