LOAD CSV WITH HEADERS FROM 'https://docs.google.com/spreadsheets/d/e/2PACX-1vR-ueVXdAYKaeB-gbnT--pic5awlwsOVdt7qD9ZpeMPZiggsqnNR_tJxJiqpmT6zmo7pYcHokzRBpY8/pub?output=csv' AS row
CREATE (:Node {name: row.name , IdAd : toInteger(row.id) });

LOAD CSV WITH HEADERS FROM "https://docs.google.com/spreadsheets/d/e/2PACX-1vSiaR8_EGoUNFYwRmSiJO5S7IQI7tgUlZYAZkBeyYTBBsEwxhC4H1C_KUudhzUn6b2i9qUauUB2iNDD/pub?output=csv" AS csv 
MATCH (parrain:Node {IdAd: toInteger(csv.`ID parrain`)})
MATCH (fils:Node {IdAd: toInteger(csv.`ID fils`)})

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