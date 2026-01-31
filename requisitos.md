# O que é?

Uma ferramenta CLI para tomar notas.
A ideia é facilitar o registro de informações das mais diversas naturezas via terminar, com poucos digitos.
Ao invés de ter que abrir um programa, um interface, mesmo que seja um vi ou nano, em uma linha consigo resolver meu problema de registrar coisas.

# Como vai funcionar?

Ainda estou pensando nisso.
Mas a ideia é ter algo como:

```bash
log in <book name> <note...>
```

A partir disso, podemos pensar em algumas regras:

- se o book não existir, criar automaticamente; se existir, adicionar uma nota nele
- cada nota é individual dentro do book
- é possível listar (de alguma forma 'bonita'):
	- todos os books
	- as notas de um book
	- todas as notas
- é possivel abrir/ler:
	- um book, exibindo as notas
	- selecionar notas especificas
- a exibição pode ser de diversos formatos, na minha cabeça tem 2 que seriam mais úteis:
	- foco no texto (cada nota preenche a tela - como no vim, por exemplo - e para alternar entre as notas pode-se navegar por elas)
	- foco na apresenção, onde as notas selecionadas são apresentadas na tela num esquema de 'cards' - aqui a ideia é poder apresentar sprint review dos meus registros através delas
- é possível editar uma nota
- é possível apagar uma nota? acho que não quero esse recurso, uma vez feita, está feita, mas ainda não estou certo disso

# Comandos?

Minha ideia é que os comandos sejam intuitivos e próximos à linguagem natural. Exemplos:

'''bash
# para logar com titulo e conteudo
log in tasks -t este é meu titulo -c este é meu conteudo 
'''

# Questões?

É realmente interessante ou é reescrever a roda?

---

# Resumo
# CLI de Notas — Resumo Consolidado

## Ideia central
Uma **ferramenta CLI para registro rápido de notas**, focada em **mínima fricção cognitiva**.

- Um comando
- Poucos caracteres
- Sem abrir editor
- Registro imediato > edição

---

## Conceitos

### Book
- Contexto semântico (ex.: `tasks`, `ideas`, `work`)
- Criado automaticamente
- Não é só uma pasta

### Nota
- Unidade individual
- Timestamp automático
- Opcionalmente título
- Imutável por padrão (editar só para correção)
- Não apagar (no máximo arquivar/marcar)

---

## Comandos principais

### Registro rápido
```bash
log in <book> "<nota>"
log in tasks "terminar POC do scraper"
```

Se o book não existir: cria
Se existir: adiciona nota

Registro com mais detalhes (opcional)
log in tasks -t "POC scraper" -c "terminar fluxo de retry e proxy"

Listar books
log books

Listar notas de um book
log ls tasks

Saída compacta:
[12] 2026-01-27 22:14  terminar POC do scraper

Listar todas as notas
log ls --all
log ls -a

Leitura / Visualização
Modo texto (foco no conteúdo)
log open tasks
- Uma nota por tela
- Navegação simples (estilo less / vim ou outro melhor)
- Busca /

Modo apresentação (cards)
- log show tasks --last 5 --cards
- Visual resumido
- Útil para review, retrospectiva e apresentação
- cores em cada coisa para diferenciar o que for relevante

Operações permitidas
"Append" nota:
log append tasks 12 "bloqueado aguardando resposta do time X"
Visualmente algo como:
[12] terminar POC do scraper
- 22:14 criado
- 23:02 bloqueado aguardando resposta do time X
- 10:11 retomado após liberação do proxy


Armazenamento
	Diretório local (~/.log)
	Um arquivo por book
	Formato simples (JSONL ou Markdown)
	Fácil de buscar, versionar e fazer backup


Linguagem? 
Go
