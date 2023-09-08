LOAD CSV WITH HEADERS FROM "https://docs.google.com/spreadsheets/d/1gsZIAAhfu_yLMyeQ55TvDi7eLeZBdNcGPKHWwNBTM8g/edit?usp=sharing" AS row
MERGE (parrain:Node {name: toString(row.`ID parrain`)})
MERGE (fils:Node {name: toString(row.`ID fils`)})

FOREACH (ignored IN CASE toInteger(row.Direct) WHEN 1 THEN [1] ELSE [] END |
    MERGE (parrain)-[rel:DIRECT]->(fils)
    ON CREATE SET rel.type = 'direct'
    ON MATCH SET rel.type = 'direct'
)

FOREACH (ignored IN CASE toInteger(row.Direct) WHEN 0 THEN [1] ELSE [] END |
    MERGE (fils)-[rel:INDIRECT]->(parrain)
    ON CREATE SET rel.type = 'indirect'
    ON MATCH SET rel.type = 'indirect'
);
