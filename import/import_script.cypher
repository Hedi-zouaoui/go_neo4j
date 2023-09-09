LOAD CSV WITH HEADERS 

FROM "https://docs.google.com/spreadsheets/d/e/2PACX-1vRsr421ImlqDv4JJf5NvjbRphlb_QMGZKAjqIbePdnNDjwLwMnT2yuXJmwCt7afXdjyS-CHCnbqYQls/pub?output=csv" AS csv fieldterminator ','
MERGE (parrain:Node {name: toString(csv.`ID parrain`)})
MERGE (fils:Node {name: toString(csv.`ID fils`)})

FOREACH (ignored IN CASE toInteger(csv.Direct) WHEN 1 THEN [1] ELSE [] END |
    MERGE (fils)-[rel:direct]->(parrain)
    ON CREATE SET rel.type = 'direct'
    ON MATCH SET rel.type = 'direct'
)

FOREACH (ignored IN CASE toInteger(csv.Direct) WHEN 0 THEN [1] ELSE [] END |
    MERGE (fils)-[rel:indirect]->(parrain)
    ON CREATE SET rel.type = 'indirect'
    ON MATCH SET rel.type = 'indirect'
);
